package transaction_test

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/stretchr/testify/assert"
	testify_mock "github.com/stretchr/testify/mock"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/projection/transaction"
	transaction_view "github.com/crypto-com/chain-indexing/projection/transaction/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	model_usecase "github.com/crypto-com/chain-indexing/usecase/model"
)

func NewTransactionProjection(rdbConn rdb.Conn) *transaction.Transaction {
	return transaction.NewTransaction(
		nil,
		rdbConn,
	)
}

func NewMockRDbConn() *test.MockRDbConn {
	mock := test.NewMockRDbConn()
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

func TestTransaction_HandleEvents(t *testing.T) {
	testCases := []struct {
		Name     string
		Events   []event_entity.Event
		MockFunc func(events []event_entity.Event) []*testify_mock.Mock
	}{
		{
			Name: "Handle TransactionCreated",
			Events: []event_entity.Event{
				&event_usecase.BlockCreated{
					Base: event_entity.NewBase(event_entity.BaseParams{
						Name:        event_usecase.BLOCK_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					Block: &model_usecase.Block{
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
				&event_usecase.TransactionCreated{
					Base: event_entity.Base{
						EventName:    "EventName",
						EventVersion: 0,
						BlockHeight:  1,
						EventUUID:    "EventUUID",
					},
					TxHash:   "TxHash",
					Index:    0,
					Code:     0,
					Log:      "Log",
					MsgCount: 1,
					Signers: []model_usecase.TransactionSigner{
						{
							Address: "SignerAddress",
							TransactionSignerInfo: model_usecase.TransactionSignerInfo{
								Type:       "SignerType",
								IsMultiSig: true,
								Pubkeys: []string{
									"pubkey1", "pubkey2", "pubkey3",
								},
								MaybeThreshold:  primptr.Int(2),
								AccountSequence: 0,
							},
						},
					},
					Senders:       []model_usecase.TransactionSigner{},
					Fee:           coin.Coins{coin.MustNewCoinFromString("basetcro", "1")},
					FeePayer:      "FeePayer",
					FeeGranter:    "FeeGranter",
					GasWanted:     200,
					GasUsed:       100,
					Memo:          "Memo",
					TimeoutHeight: 10,
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
			MockFunc: func(events []event_entity.Event) (mocks []*testify_mock.Mock) {

				msgEvent, _ := events[2].(event_usecase.MsgEvent)

				mockTransactionsView := transaction_view.NewMockTransactionsView(nil).(*transaction_view.MockTransactionsView)
				mocks = append(mocks, &mockTransactionsView.Mock)

				transaction.NewTransactions = func(_ *rdb.Handle) transaction_view.BlockTransactions {
					return mockTransactionsView
				}

				mockTransactionsView.On(
					"InsertAll",
					[]transaction_view.TransactionRow{
						{
							BlockHeight:   1,
							BlockTime:     utctime.UTCTime{},
							BlockHash:     "Hash",
							Hash:          "TxHash",
							Index:         0,
							Success:       true,
							Code:          0,
							Log:           "Log",
							Fee:           coin.Coins{coin.MustNewCoinFromString("basetcro", "1")},
							FeePayer:      "FeePayer",
							FeeGranter:    "FeeGranter",
							GasWanted:     200,
							GasUsed:       100,
							Memo:          "Memo",
							TimeoutHeight: 10,
							Signers: []transaction_view.TransactionRowSigner{
								{
									Type:       "SignerType",
									IsMultiSig: true,
									Pubkeys: []string{
										"pubkey1", "pubkey2", "pubkey3",
									},
									MaybeThreshold:  primptr.Int(2),
									AccountSequence: 0,
									Address:         "SignerAddress",
								},
							},
							Messages: []transaction_view.TransactionRowMessage{
								{
									Type:    "MsgSend",
									Content: msgEvent,
								},
							},
						},
					},
				).Return(nil)

				mockTransactionsTotalView := transaction_view.NewMockTransactionsTotalView(nil).(*transaction_view.MockTransactionsTotalView)
				mocks = append(mocks, &mockTransactionsTotalView.Mock)

				transaction.NewTransactionsTotal = func(_ *rdb.Handle) transaction_view.TransactionsTotal {
					return mockTransactionsTotalView
				}

				mockTransactionsTotalView.On(
					"Increment",
					"-",
					int64(1),
				).Return(nil)

				mockTransactionsTotalView.On(
					"Set",
					"1",
					int64(1),
				).Return(nil)

				return mocks
			},
		},
		{
			Name: "Handle TransactionFailed",
			Events: []event_entity.Event{
				&event_usecase.BlockCreated{
					Base: event_entity.NewBase(event_entity.BaseParams{
						Name:        event_usecase.BLOCK_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					Block: &model_usecase.Block{
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
				&event_usecase.TransactionFailed{
					Base: event_entity.Base{
						EventName:    "EventName",
						EventVersion: 0,
						BlockHeight:  1,
						EventUUID:    "EventUUID",
					},
					TxHash:   "TxHash",
					Index:    0,
					Code:     0,
					Log:      "Log",
					MsgCount: 1,
					Signers: []model_usecase.TransactionSigner{
						{
							Address: "SignerAddress",
							TransactionSignerInfo: model_usecase.TransactionSignerInfo{
								Type:       "SignerType",
								IsMultiSig: true,
								Pubkeys: []string{
									"pubkey1", "pubkey2", "pubkey3",
								},
								MaybeThreshold:  primptr.Int(2),
								AccountSequence: 0,
							},
						},
					},
					Senders:       []model_usecase.TransactionSigner{},
					Fee:           coin.Coins{coin.MustNewCoinFromString("basetcro", "1")},
					FeePayer:      "FeePayer",
					FeeGranter:    "FeeGranter",
					GasWanted:     200,
					GasUsed:       100,
					Memo:          "Memo",
					TimeoutHeight: 10,
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
			MockFunc: func(events []event_entity.Event) (mocks []*testify_mock.Mock) {

				msgEvent, _ := events[2].(event_usecase.MsgEvent)

				mockTransactionsView := transaction_view.NewMockTransactionsView(nil).(*transaction_view.MockTransactionsView)
				mocks = append(mocks, &mockTransactionsView.Mock)

				transaction.NewTransactions = func(_ *rdb.Handle) transaction_view.BlockTransactions {
					return mockTransactionsView
				}

				mockTransactionsView.On(
					"InsertAll",
					[]transaction_view.TransactionRow{
						{
							BlockHeight:   1,
							BlockTime:     utctime.UTCTime{},
							BlockHash:     "Hash",
							Hash:          "TxHash",
							Index:         0,
							Success:       false,
							Code:          0,
							Log:           "Log",
							Fee:           coin.Coins{coin.MustNewCoinFromString("basetcro", "1")},
							FeePayer:      "FeePayer",
							FeeGranter:    "FeeGranter",
							GasWanted:     200,
							GasUsed:       100,
							Memo:          "Memo",
							TimeoutHeight: 10,
							Signers: []transaction_view.TransactionRowSigner{
								{
									Type:       "SignerType",
									IsMultiSig: true,
									Pubkeys: []string{
										"pubkey1", "pubkey2", "pubkey3",
									},
									MaybeThreshold:  primptr.Int(2),
									AccountSequence: 0,
									Address:         "SignerAddress",
								},
							},
							Messages: []transaction_view.TransactionRowMessage{
								{
									Type:    "MsgSend",
									Content: msgEvent,
								},
							},
						},
					},
				).Return(nil)

				mockTransactionsTotalView := transaction_view.NewMockTransactionsTotalView(nil).(*transaction_view.MockTransactionsTotalView)
				mocks = append(mocks, &mockTransactionsTotalView.Mock)

				transaction.NewTransactionsTotal = func(_ *rdb.Handle) transaction_view.TransactionsTotal {
					return mockTransactionsTotalView
				}

				mockTransactionsTotalView.On(
					"Increment",
					"-",
					int64(1),
				).Return(nil)

				mockTransactionsTotalView.On(
					"Set",
					"1",
					int64(1),
				).Return(nil)

				return mocks
			},
		},
	}

	for _, tc := range testCases {
		mockRDbConn := NewMockRDbConn()
		mockTx := NewMockRDbTx()
		mockRDbConn.On("Begin").Return(mockTx, nil)
		mocks := tc.MockFunc(tc.Events)

		// We are not interested in testing the below function in unit test
		transaction.UpdateLastHandledEventHeight = func(_ *transaction.Transaction, _ *rdb.Handle, _ int64) error {
			return nil
		}

		projection := NewTransactionProjection(mockRDbConn)
		err := projection.HandleEvents(1, tc.Events)
		assert.NoError(t, err)

		for _, m := range mocks {
			m.AssertExpectations(t)
		}

		fmt.Println(tc.Name, "Passed")
	}
}
