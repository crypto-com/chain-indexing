package account_raw_event_test

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	entity_event "github.com/crypto-com/chain-indexing/entity/event"

	"github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/projection/account_message"
	"github.com/crypto-com/chain-indexing/projection/account_raw_event"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	usecase_model "github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/stretchr/testify/assert"
	testify_mock "github.com/stretchr/testify/mock"
)

func NewAccountRawEventProjection(rdbConn rdb.Conn) *account_raw_event.AccountRawEvent {
	return account_raw_event.NewAccountRawEvent(
		nil,
		rdbConn,
		"tcro",
		nil,
	)
}

func NewMockRDbConn() *test.MockRDbConn {
	mock := test.NewMockRDbConn()
	mock.On("ToHandle").Return(&rdb.Handle{
		Runner:   mock,
		TypeConv: &pg.PgxTypeConv{},
		StmtBuilder: &rdb.StatementBuilder{
			StatementBuilderType: sq.StatementBuilderType{},
			PlaceholderFormat:    nil,
		},
	})

	return mock
}

func NewMockRDbTx() *test.MockRDbTx {
	mockTx := &test.MockRDbTx{}
	mockTx.On("ToHandle").Return(&rdb.Handle{
		Runner:      mockTx,
		TypeConv:    &pg.PgxTypeConv{},
		StmtBuilder: pg.PostgresStmtBuilder,
	}).Maybe()
	mockTx.On("Rollback").Return(nil).Maybe()
	mockTx.On("Commit").Return(nil).Maybe()

	return mockTx
}
func TestAccountMessage_HandleEvents(t *testing.T) {
	testCases := []struct {
		Name     string
		Events   []entity_event.Event
		MockFunc func(events []entity_event.Event) []*testify_mock.Mock
	}{
		{
			Name: "HandleMsgSend",
			Events: []entity_event.Event{
				&event_usecase.BlockCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					Block: &usecase_model.Block{
						Height:          1,
						Hash:            "Hash",
						Time:            utctime.UTCTime{},
						AppHash:         "AppHash",
						ProposerAddress: "ProposerAddress",
						Txs:             nil,
						Signatures:      nil,
						Evidences:       nil,
					},
				},
				&event_usecase.MsgSend{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_SEND,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					FromAddress: "FromAddress",
					ToAddress:   "ToAddress",
					Amount: coin.Coins{
						coin.Coin{
							Denom:  "Denom",
							Amount: coin.NewInt(100),
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockAccountMessagesTotalView := account_message_view.NewMockAccountMessagesTotalView(nil).(*account_message_view.MockAccountMessagesTotalView)
				mocks = append(mocks, &mockAccountMessagesTotalView.Mock)

				account_message.NewAccountMessagesTotal = func(_ *rdb.Handle) account_message_view.AccountMessagesTotal {
					return mockAccountMessagesTotalView
				}

				mockAccountMessagesTotalView.On(
					"Increment",
					"FromAddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"FromAddress:/cosmos.bank.v1beta1.MsgSend",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"ToAddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"ToAddress:/cosmos.bank.v1beta1.MsgSend",
					int64(1),
				).Return(nil)

				mockAccountMessagesView := account_message_view.NewMockAccountMessagesView(nil).(*account_message_view.MockAccountMessagesView)
				mocks = append(mocks, &mockAccountMessagesView.Mock)

				account_message.NewAccountMessages = func(_ *rdb.Handle) account_message_view.AccountMessages {
					return mockAccountMessagesView
				}

				mockAccountMessagesView.On(
					"Insert",
					&account_message_view.AccountMessageRow{
						MaybeAccount:    (*string)(nil),
						BlockHeight:     1,
						BlockHash:       "Hash",
						BlockTime:       utctime.UTCTime{},
						TransactionHash: "TxHash",
						Success:         true,
						MessageIndex:    0,
						MessageType:     "/cosmos.bank.v1beta1.MsgSend",
						Data: &event_usecase.MsgSend{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/cosmos.bank.v1beta1.MsgSend.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/cosmos.bank.v1beta1.MsgSend",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							FromAddress: "FromAddress",
							ToAddress:   "ToAddress",
							Amount: coin.Coins{
								{
									Denom:  "Denom",
									Amount: coin.NewInt(100),
								},
							},
						},
					},
					[]string{"FromAddress", "ToAddress"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
	}
	for _, tc := range testCases {
		mockRDbConn := NewMockRDbConn()
		mockTx := NewMockRDbTx()
		mockRDbConn.On("Begin").Return(mockTx, nil)
		mocks := tc.MockFunc(tc.Events)

		projection := NewAccountRawEventProjection(mockRDbConn)
		err := projection.HandleEvents(1, tc.Events)
		assert.NoError(t, err)

		for _, m := range mocks {
			m.AssertExpectations(t)
		}

		fmt.Println(tc.Name, "Passed")
	}
}
