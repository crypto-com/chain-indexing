package raw_transaction_test

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/projection/raw_transaction"
	raw_transaction_view "github.com/crypto-com/chain-indexing/projection/raw_transaction/view"
	"github.com/stretchr/testify/assert"
	testify_mock "github.com/stretchr/testify/mock"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	model_usecase "github.com/crypto-com/chain-indexing/usecase/model"
)

func NewRawTransactionProjection(rdbConn rdb.Conn) *raw_transaction.RawTransaction {
	return raw_transaction.NewRawTransaction(
		nil,
		rdbConn,
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

func TestRawTransaction_HandleEvents(t *testing.T) {
	testCases := []struct {
		Name     string
		Events   []event_entity.Event
		MockFunc func(events []event_entity.Event) []*testify_mock.Mock
	}{
		{
			Name: "Handle RawTransactionCreated",
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
				&event_usecase.RawTransactionCreated{
					Base: event_entity.Base{
						EventName:    "EventName",
						EventVersion: 0,
						BlockHeight:  1,
						EventUUID:    "EventUUID",
					},
					TxHash: "TxHash",
					Index:  0,
					Code:   0,
					Log:    "Log",
					Signers: []model_usecase.TransactionSigner{
						{
							MaybeKeyInfo: &model_usecase.TransactionSignerKeyInfo{
								Type:       "SignerType",
								IsMultiSig: true,
								Pubkeys: []string{
									"pubkey1", "pubkey2", "pubkey3",
								},
								MaybeThreshold: primptr.Int(2),
							},
							Address:         "SignerAddress",
							AccountSequence: 0,
						},
					},
					Fee:           coin.Coins{coin.MustNewCoinFromString("basetcro", "1")},
					FeePayer:      "FeePayer",
					FeeGranter:    "FeeGranter",
					GasWanted:     200,
					GasUsed:       100,
					Memo:          "Memo",
					TimeoutHeight: 10,
					Messages: []map[string]interface{}{
						{
							"@type": "/cosmos.bank.v1beta1.MsgSend",
							"amount": []map[string]interface{}{
								{"denom": "basetcro", "amount": "1000000000000"},
							},
							"to_address":   "tcro1k4nz2yalsuzuzxexcxuughp33cfqls7e6utarx",
							"from_address": "tcro1s4ggq2zuzvwg5k8vnx2xfwtdm4cz6wtnfl4xx7",
						},
					},
				},
			},
			MockFunc: func(events []event_entity.Event) (mocks []*testify_mock.Mock) {

				mockRawTransactionsView := raw_transaction_view.NewMockRawTransactionsView(nil).(*raw_transaction_view.MockRawTransactionsView)
				mocks = append(mocks, &mockRawTransactionsView.Mock)

				raw_transaction.NewRawTransactions = func(_ *rdb.Handle) raw_transaction_view.BlockRawTransactions {
					return mockRawTransactionsView
				}

				mockRawTransactionsView.On(
					"InsertAll",
					[]raw_transaction_view.RawTransactionRow{
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
							Signers: []raw_transaction_view.RawTransactionRowSigner{
								{
									MaybeKeyInfo: &raw_transaction_view.RawTransactionRowSignerKeyInfo{
										Type:       "SignerType",
										IsMultiSig: true,
										Pubkeys: []string{
											"pubkey1", "pubkey2", "pubkey3",
										},
										MaybeThreshold: primptr.Int(2),
									},
									AccountSequence: 0,
									Address:         "SignerAddress",
								},
							},
							Messages: []map[string]interface{}{
								{
									"@type": "/cosmos.bank.v1beta1.MsgSend",
									"amount": []map[string]interface{}{
										{"denom": "basetcro", "amount": "1000000000000"},
									},
									"to_address":   "tcro1k4nz2yalsuzuzxexcxuughp33cfqls7e6utarx",
									"from_address": "tcro1s4ggq2zuzvwg5k8vnx2xfwtdm4cz6wtnfl4xx7",
								},
							},
						},
					},
				).Return(nil)

				mockRawTransactionsTotalView := raw_transaction_view.NewMockRawTransactionsTotalView(nil).(*raw_transaction_view.MockRawTransactionsTotalView)
				mocks = append(mocks, &mockRawTransactionsTotalView.Mock)

				raw_transaction.NewRawTransactionsTotal = func(_ *rdb.Handle) raw_transaction_view.RawTransactionsTotal {
					return mockRawTransactionsTotalView
				}

				mockRawTransactionsTotalView.On(
					"Increment",
					"-",
					int64(1),
				).Return(nil)

				mockRawTransactionsTotalView.On(
					"Set",
					"1",
					int64(1),
				).Return(nil)

				return mocks
			},
		},
		{
			Name: "Handle RawTransactionFailed",
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
				&event_usecase.RawTransactionFailed{
					Base: event_entity.Base{
						EventName:    "EventName",
						EventVersion: 0,
						BlockHeight:  1,
						EventUUID:    "EventUUID",
					},
					TxHash: "TxHash",
					Index:  0,
					Code:   0,
					Log:    "Log",
					Signers: []model_usecase.TransactionSigner{
						{
							MaybeKeyInfo: &model_usecase.TransactionSignerKeyInfo{
								Type:       "SignerType",
								IsMultiSig: true,
								Pubkeys: []string{
									"pubkey1", "pubkey2", "pubkey3",
								},
								MaybeThreshold: primptr.Int(2),
							},
							Address:         "SignerAddress",
							AccountSequence: 0,
						},
					},
					Fee:           coin.Coins{coin.MustNewCoinFromString("basetcro", "1")},
					FeePayer:      "FeePayer",
					FeeGranter:    "FeeGranter",
					GasWanted:     200,
					GasUsed:       100,
					Memo:          "Memo",
					TimeoutHeight: 10,
					Messages: []map[string]interface{}{
						{
							"@type": "/cosmos.bank.v1beta1.MsgSend",
							"amount": []map[string]interface{}{
								{"denom": "basetcro", "amount": "1000000000000"},
							},
							"to_address":   "tcro1k4nz2yalsuzuzxexcxuughp33cfqls7e6utarx",
							"from_address": "tcro1s4ggq2zuzvwg5k8vnx2xfwtdm4cz6wtnfl4xx7",
						},
					},
				},
			},
			MockFunc: func(events []event_entity.Event) (mocks []*testify_mock.Mock) {

				mockRawTransactionsView := raw_transaction_view.NewMockRawTransactionsView(nil).(*raw_transaction_view.MockRawTransactionsView)
				mocks = append(mocks, &mockRawTransactionsView.Mock)

				raw_transaction.NewRawTransactions = func(_ *rdb.Handle) raw_transaction_view.BlockRawTransactions {
					return mockRawTransactionsView
				}

				mockRawTransactionsView.On(
					"InsertAll",
					[]raw_transaction_view.RawTransactionRow{
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
							Signers: []raw_transaction_view.RawTransactionRowSigner{
								{
									MaybeKeyInfo: &raw_transaction_view.RawTransactionRowSignerKeyInfo{
										Type:       "SignerType",
										IsMultiSig: true,
										Pubkeys: []string{
											"pubkey1", "pubkey2", "pubkey3",
										},
										MaybeThreshold: primptr.Int(2),
									},
									AccountSequence: 0,
									Address:         "SignerAddress",
								},
							},
							Messages: []map[string]interface{}{
								{
									"@type": "/cosmos.bank.v1beta1.MsgSend",
									"amount": []map[string]interface{}{
										{"denom": "basetcro", "amount": "1000000000000"},
									},
									"to_address":   "tcro1k4nz2yalsuzuzxexcxuughp33cfqls7e6utarx",
									"from_address": "tcro1s4ggq2zuzvwg5k8vnx2xfwtdm4cz6wtnfl4xx7",
								},
							},
						},
					},
				).Return(nil)

				mockRawTransactionsTotalView := raw_transaction_view.NewMockRawTransactionsTotalView(nil).(*raw_transaction_view.MockRawTransactionsTotalView)
				mocks = append(mocks, &mockRawTransactionsTotalView.Mock)

				raw_transaction.NewRawTransactionsTotal = func(_ *rdb.Handle) raw_transaction_view.RawTransactionsTotal {
					return mockRawTransactionsTotalView
				}

				mockRawTransactionsTotalView.On(
					"Increment",
					"-",
					int64(1),
				).Return(nil)

				mockRawTransactionsTotalView.On(
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
		raw_transaction.UpdateLastHandledEventHeight = func(_ *raw_transaction.RawTransaction, _ *rdb.Handle, _ int64) error {
			return nil
		}

		projection := NewRawTransactionProjection(mockRDbConn)
		err := projection.HandleEvents(1, tc.Events)
		assert.NoError(t, err)

		for _, m := range mocks {
			m.AssertExpectations(t)
		}

		fmt.Println(tc.Name, "Passed")
	}
}
