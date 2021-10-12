package ibc_channel_message_test

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	testify_mock "github.com/stretchr/testify/mock"

	pagination2 "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/internal/primptr"
	"github.com/crypto-com/chain-indexing/internal/utctime"
	"github.com/crypto-com/chain-indexing/projection/ibc_channel_message"
	"github.com/crypto-com/chain-indexing/projection/ibc_channel_message/view"
	usecase_event "github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

func NewIBCChannelMessageProjection(rdbConn rdb.Conn) *ibc_channel_message.IBCChannelMessage {

	return ibc_channel_message.NewIBCChannelMessage(
		nil,
		rdbConn,
	)
}

func NewMockRDbConn() *test.MockRDbConn {
	mock := test.NewMockRDBbConn()
	mock.On("ToHandle").Return(&rdb.Handle{
		Runner:      mock,
		TypeConv:    &pg.PgxTypeConv{},
		StmtBuilder: sq.StatementBuilderType{},
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

type MockIBCChannelMessageView struct {
	testify_mock.Mock
}

func (ibcChannelMessagesView *MockIBCChannelMessageView) Insert(ibcChannelMessage *view.IBCChannelMessageRow) error {
	mockArgs := ibcChannelMessagesView.Called(ibcChannelMessage)
	return mockArgs.Error(0)
}

func (ibcChannelMessagesView *MockIBCChannelMessageView) ListByChannelID(
	channelID string,
	order view.IBCChannelMessagesListOrder,
	filter view.IBCChannelMessagesListFilter,
	pagination *pagination2.Pagination,
) (
	[]view.IBCChannelMessageRow,
	*pagination2.PaginationResult,
	error,
) {
	mockArgs := ibcChannelMessagesView.Called(channelID, order, filter, pagination)
	messages, _ := mockArgs.Get(0).([]view.IBCChannelMessageRow)
	paginationResult, _ := mockArgs.Get(1).(*pagination2.PaginationResult)
	return messages, paginationResult, mockArgs.Error(3)
}

type MockIBCChannelMessageTotalView struct {
	testify_mock.Mock
}

func (totalView *MockIBCChannelMessageTotalView) Increment(identity string, total int64) error {
	mockArgs := totalView.Called(identity, total)
	return mockArgs.Error(0)
}

func (totalView *MockIBCChannelMessageTotalView) SumBy(identities []string) (int64, error) {
	mockArgs := totalView.Called(identities)
	total, _ := mockArgs.Get(0).(int64)
	return total, mockArgs.Error(1)
}

func TestIBCChannelMessage_HandleEvents(t *testing.T) {
	testCases := []struct {
		Name     string
		Events   []entity_event.Event
		MockFunc func(mock *test.MockRDbConn, events []entity_event.Event) []*testify_mock.Mock
	}{
		{
			Name: "HandleMsgIBCChannelOpenInit",
			Events: []entity_event.Event{
				&usecase_event.MsgIBCChannelOpenInit{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_IBC_CHANNEL_OPEN_INIT,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgChannelOpenInitParams{
						ChannelID: "ChannelID",
						RawMsgChannelOpenInit: ibc_model.RawMsgChannelOpenInit{
							PortID: "PortID",
							Channel: ibc_model.Channel{
								Ordering: "Ordering",
								Counterparty: ibc_model.ChannelCounterparty{
									PortID:    "CounterpartyPortID",
									ChannelID: "CounterpartyChannelID",
								},
							},
							Signer: "Signer",
						},
					},
				},
			},
			MockFunc: func(mockCoon *test.MockRDbConn, events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockTx := NewMockRDbTx()
				mocks = append(mocks, &mockTx.Mock)
				mockCoon.On("Begin").Return(mockTx, nil)

				typedEvent := events[0].(*usecase_event.MsgIBCChannelOpenInit)

				mockIBCChannelMessageView := &MockIBCChannelMessageView{}
				mocks = append(mocks, &mockIBCChannelMessageView.Mock)
				mockIBCChannelMessageView.
					On("Insert", &view.IBCChannelMessageRow{
						ChannelID:       "ChannelID",
						BlockHeight:     int64(1),
						BlockTime:       utctime.UTCTime{},
						TransactionHash: typedEvent.TxHash(),
						MaybeRelayer:    primptr.String(typedEvent.Params.Signer),
						MessageType:     typedEvent.MsgName,
						Message:         typedEvent,
					}).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessagesView = func(handle *rdb.Handle) view.IBCChannelMessagesI {
					return mockIBCChannelMessageView
				}

				mockIBCChannelMessageTotalView := &MockIBCChannelMessageTotalView{}
				mocks = append(mocks, &mockIBCChannelMessageTotalView.Mock)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:-", "ChannelID"), int64(1)).
					Return(nil)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:%s", "ChannelID", typedEvent.MsgName), int64(1)).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessagesTotalView = func(handle *rdb.Handle) view.IBCChannelMessagesTotalI {
					return mockIBCChannelMessageTotalView
				}

				return mocks
			},
		},
	}

	for _, tc := range testCases {
		mockRDbConn := NewMockRDbConn()
		mocks := tc.MockFunc(mockRDbConn, tc.Events)
		mocks = append(mocks, &mockRDbConn.Mock)

		updateHeightCalled, actualInputHeight := fakeUpdateProjectionLastHandledEventHeight()

		projection := NewIBCChannelMessageProjection(mockRDbConn)
		err := projection.HandleEvents(1, tc.Events)
		assert.NoError(t, err)

		assert.True(t, *updateHeightCalled)
		assert.EqualValues(t, int64(1), *actualInputHeight)

		for _, m := range mocks {
			m.AssertExpectations(t)
		}

		fmt.Println(tc.Name, "Passed")
	}
}

func fakeUpdateProjectionLastHandledEventHeight() (*bool, *int64) {
	var called bool
	var inputHeight int64
	ibc_channel_message.UpdateProjectionLastHandledEventHeight = func(_ *ibc_channel_message.IBCChannelMessage, rdbHandle *rdb.Handle, height int64) error {
		called = true
		inputHeight = height
		return nil
	}
	return &called, &inputHeight
}
