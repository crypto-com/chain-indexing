package account_message_test

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/external/json"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/projection/account_message"
	account_message_view "github.com/crypto-com/chain-indexing/projection/account_message/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	usecase_model "github.com/crypto-com/chain-indexing/usecase/model"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	"github.com/stretchr/testify/assert"
	testify_mock "github.com/stretchr/testify/mock"
)

func NewAccountMessageProjection(rdbConn rdb.Conn) *account_message.AccountMessage {
	return account_message.NewAccountMessage(
		nil,
		rdbConn,
		"accountaddressprefix",
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
					"fromaddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"fromaddress:/cosmos.bank.v1beta1.MsgSend",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"toaddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"toaddress:/cosmos.bank.v1beta1.MsgSend",
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
					[]string{"fromaddress", "toaddress"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgMultiSend",
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
				&event_usecase.MsgMultiSend{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_MULTI_SEND,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Inputs: []usecase_model.MsgMultiSendInput{
						{
							Address: "InputAddress",
							Amount: coin.Coins{
								coin.Coin{
									Denom:  "Denom",
									Amount: coin.NewInt(100),
								},
							},
						},
					},
					Outputs: []usecase_model.MsgMultiSendOutput{
						{
							Address: "OutputAddress",
							Amount: coin.Coins{
								coin.Coin{
									Denom:  "Denom",
									Amount: coin.NewInt(100),
								},
							},
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
					"inputaddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"inputaddress:/cosmos.bank.v1beta1.MsgMultiSend",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"outputaddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"outputaddress:/cosmos.bank.v1beta1.MsgMultiSend",
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
						MessageType:     "/cosmos.bank.v1beta1.MsgMultiSend",
						Data: &event_usecase.MsgMultiSend{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/cosmos.bank.v1beta1.MsgMultiSend.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/cosmos.bank.v1beta1.MsgMultiSend",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Inputs: []usecase_model.MsgMultiSendInput{
								{
									Address: "InputAddress",
									Amount: coin.Coins{
										coin.Coin{
											Denom:  "Denom",
											Amount: coin.NewInt(100),
										},
									},
								},
							},
							Outputs: []usecase_model.MsgMultiSendOutput{
								{
									Address: "OutputAddress",
									Amount: coin.Coins{
										coin.Coin{
											Denom:  "Denom",
											Amount: coin.NewInt(100),
										},
									},
								},
							},
						},
					},
					[]string{"inputaddress", "outputaddress"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgSetWithdrawAddress",
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
				&event_usecase.MsgSetWithdrawAddress{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_SET_WITHDRAW_ADDRESS,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					MsgSetWithdrawAddressParams: usecase_model.MsgSetWithdrawAddressParams{
						DelegatorAddress: "DelegatorAddress",
						WithdrawAddress:  "WithdrawAddress",
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
					"delegatoraddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"delegatoraddress:/cosmos.distribution.v1beta1.MsgSetWithdrawAddress",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"withdrawaddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"withdrawaddress:/cosmos.distribution.v1beta1.MsgSetWithdrawAddress",
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
						MessageType:     "/cosmos.distribution.v1beta1.MsgSetWithdrawAddress",
						Data: &event_usecase.MsgSetWithdrawAddress{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/cosmos.distribution.v1beta1.MsgSetWithdrawAddress.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/cosmos.distribution.v1beta1.MsgSetWithdrawAddress",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							MsgSetWithdrawAddressParams: usecase_model.MsgSetWithdrawAddressParams{
								DelegatorAddress: "DelegatorAddress",
								WithdrawAddress:  "WithdrawAddress",
							},
						},
					},
					[]string{"delegatoraddress", "withdrawaddress"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgWithdrawDelegatorReward",
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
				&event_usecase.MsgWithdrawDelegatorReward{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_WITHDRAW_DELEGATOR_REWARD,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					MsgWithdrawDelegatorRewardParams: usecase_model.MsgWithdrawDelegatorRewardParams{
						DelegatorAddress: "DelegatorAddress",
						ValidatorAddress: "ValidatorAddress",
						RecipientAddress: "RecipientAddress",
						Amount: coin.Coins{
							coin.Coin{
								Denom:  "Denom",
								Amount: coin.NewInt(100),
							},
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
					"delegatoraddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"delegatoraddress:/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
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
						MessageType:     "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
						Data: &event_usecase.MsgWithdrawDelegatorReward{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							MsgWithdrawDelegatorRewardParams: usecase_model.MsgWithdrawDelegatorRewardParams{
								DelegatorAddress: "DelegatorAddress",
								ValidatorAddress: "ValidatorAddress",
								RecipientAddress: "RecipientAddress",
								Amount: coin.Coins{
									coin.Coin{
										Denom:  "Denom",
										Amount: coin.NewInt(100),
									},
								},
							},
						},
					},
					[]string{"delegatoraddress"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgWithdrawValidatorCommission",
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
				&event_usecase.MsgWithdrawValidatorCommission{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_WITHDRAW_VALIDATOR_COMMISSION,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					MsgWithdrawValidatorCommissionParams: usecase_model.MsgWithdrawValidatorCommissionParams{
						ValidatorAddress: "ValidatorAddress",
						RecipientAddress: "RecipientAddress",
						Amount: coin.Coins{
							coin.Coin{
								Denom:  "Denom",
								Amount: coin.NewInt(100),
							},
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
					"recipientaddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"recipientaddress:/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission",
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
						MessageType:     "/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission",
						Data: &event_usecase.MsgWithdrawValidatorCommission{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							MsgWithdrawValidatorCommissionParams: usecase_model.MsgWithdrawValidatorCommissionParams{
								ValidatorAddress: "ValidatorAddress",
								RecipientAddress: "RecipientAddress",
								Amount: coin.Coins{
									coin.Coin{
										Denom:  "Denom",
										Amount: coin.NewInt(100),
									},
								},
							},
						},
					},
					[]string{"recipientaddress"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgFundCommunityPool",
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
				&event_usecase.MsgFundCommunityPool{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_FUND_COMMUNITY_POOL,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Depositor: "Depositor",
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
					"depositor:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"depositor:/cosmos.distribution.v1beta1.MsgFundCommunityPool",
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
						MessageType:     "/cosmos.distribution.v1beta1.MsgFundCommunityPool",
						Data: &event_usecase.MsgFundCommunityPool{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/cosmos.distribution.v1beta1.MsgFundCommunityPool.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/cosmos.distribution.v1beta1.MsgFundCommunityPool",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Depositor: "Depositor",
							Amount: coin.Coins{
								coin.Coin{
									Denom:  "Denom",
									Amount: coin.NewInt(100),
								},
							},
						},
					},
					[]string{"depositor"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgSubmitParamChangeProposal",
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
				&event_usecase.MsgSubmitParamChangeProposal{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_SUBMIT_PARAM_CHANGE_PROPOSAL,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					MsgSubmitParamChangeProposalParams: usecase_model.MsgSubmitParamChangeProposalParams{
						MaybeProposalId: nil,
						Content: usecase_model.MsgSubmitParamChangeProposalContent{
							Type:        "Type",
							Title:       "Title",
							Description: "Description",
							Changes: []usecase_model.MsgSubmitParamChangeProposalChange{
								{
									Subspace: "Subspace",
									Key:      "Key",
									Value:    nil,
								},
							},
						},
						ProposerAddress: "ProposerAddress",
						InitialDeposit: coin.Coins{
							{
								Denom:  "Denom",
								Amount: coin.NewInt(100),
							},
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
					"proposeraddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"proposeraddress:/cosmos.params.v1beta1.ParameterChangeProposal",
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
						MessageType:     "/cosmos.params.v1beta1.ParameterChangeProposal",
						Data: &event_usecase.MsgSubmitParamChangeProposal{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/cosmos.params.v1beta1.ParameterChangeProposal.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/cosmos.params.v1beta1.ParameterChangeProposal",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							MsgSubmitParamChangeProposalParams: usecase_model.MsgSubmitParamChangeProposalParams{
								MaybeProposalId: nil,
								Content: usecase_model.MsgSubmitParamChangeProposalContent{
									Type:        "Type",
									Title:       "Title",
									Description: "Description",
									Changes: []usecase_model.MsgSubmitParamChangeProposalChange{
										{
											Subspace: "Subspace",
											Key:      "Key",
											Value:    nil,
										},
									},
								},
								ProposerAddress: "ProposerAddress",
								InitialDeposit: coin.Coins{
									{
										Denom:  "Denom",
										Amount: coin.NewInt(100),
									},
								},
							},
						},
					},
					[]string{"proposeraddress"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgSubmitCommunityPoolSpendProposal",
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
				&event_usecase.MsgSubmitCommunityPoolSpendProposal{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					MsgSubmitCommunityPoolSpendProposalParams: usecase_model.MsgSubmitCommunityPoolSpendProposalParams{
						MaybeProposalId: nil,
						Content: usecase_model.MsgSubmitCommunityPoolSpendProposalContent{
							Type:             "Type",
							Title:            "Title",
							Description:      "Description",
							RecipientAddress: "RecipientAddress",
							Amount: coin.Coins{
								{
									Denom:  "Denom",
									Amount: coin.NewInt(100),
								},
							},
						},
						ProposerAddress: "ProposerAddress",
						InitialDeposit: coin.Coins{
							{
								Denom:  "Denom",
								Amount: coin.NewInt(100),
							},
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
					"proposeraddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"proposeraddress:/cosmos.distribution.v1beta1.CommunityPoolSpendProposal",
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
						MessageType:     "/cosmos.distribution.v1beta1.CommunityPoolSpendProposal",
						Data: &event_usecase.MsgSubmitCommunityPoolSpendProposal{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/cosmos.distribution.v1beta1.CommunityPoolSpendProposal.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/cosmos.distribution.v1beta1.CommunityPoolSpendProposal",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							MsgSubmitCommunityPoolSpendProposalParams: usecase_model.MsgSubmitCommunityPoolSpendProposalParams{
								MaybeProposalId: nil,
								Content: usecase_model.MsgSubmitCommunityPoolSpendProposalContent{
									Type:             "Type",
									Title:            "Title",
									Description:      "Description",
									RecipientAddress: "RecipientAddress",
									Amount: coin.Coins{
										{
											Denom:  "Denom",
											Amount: coin.NewInt(100),
										},
									},
								},
								ProposerAddress: "ProposerAddress",
								InitialDeposit: coin.Coins{
									{
										Denom:  "Denom",
										Amount: coin.NewInt(100),
									},
								},
							},
						},
					},
					[]string{"proposeraddress"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgSubmitCancelSoftwareUpgradeProposal",
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
				&event_usecase.MsgSubmitCancelSoftwareUpgradeProposal{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					MsgSubmitCancelSoftwareUpgradeProposalParams: usecase_model.MsgSubmitCancelSoftwareUpgradeProposalParams{
						MaybeProposalId: nil,
						Content: usecase_model.MsgSubmitCancelSoftwareUpgradeProposalContent{
							Type:        "Type",
							Title:       "Title",
							Description: "Description",
						},
						ProposerAddress: "ProposerAddress",
						InitialDeposit: coin.Coins{
							{
								Denom:  "Denom",
								Amount: coin.NewInt(100),
							},
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
					"proposeraddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"proposeraddress:/cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal",
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
						MessageType:     "/cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal",
						Data: &event_usecase.MsgSubmitCancelSoftwareUpgradeProposal{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							MsgSubmitCancelSoftwareUpgradeProposalParams: usecase_model.MsgSubmitCancelSoftwareUpgradeProposalParams{
								MaybeProposalId: nil,
								Content: usecase_model.MsgSubmitCancelSoftwareUpgradeProposalContent{
									Type:        "Type",
									Title:       "Title",
									Description: "Description",
								},
								ProposerAddress: "ProposerAddress",
								InitialDeposit: coin.Coins{
									{
										Denom:  "Denom",
										Amount: coin.NewInt(100),
									},
								},
							},
						},
					},
					[]string{"proposeraddress"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgDeposit",
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
				&event_usecase.MsgDeposit{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_DEPOSIT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					ProposalId: "ProposalId",
					Depositor:  "Depositor",
					Amount: coin.Coins{
						{
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
					"depositor:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"depositor:/cosmos.gov.v1beta1.MsgDeposit",
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
						MessageType:     "/cosmos.gov.v1beta1.MsgDeposit",
						Data: &event_usecase.MsgDeposit{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/cosmos.gov.v1beta1.MsgDeposit.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/cosmos.gov.v1beta1.MsgDeposit",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							ProposalId: "ProposalId",
							Depositor:  "Depositor",
							Amount: coin.Coins{
								{
									Denom:  "Denom",
									Amount: coin.NewInt(100),
								},
							},
						},
					},
					[]string{"depositor"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgVote",
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
				&event_usecase.MsgVote{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_VOTE,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					ProposalId: "ProposalId",
					Voter:      "Voter",
					Option:     "Option",
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
					"voter:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"voter:/cosmos.gov.v1beta1.MsgVote",
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
						MessageType:     "/cosmos.gov.v1beta1.MsgVote",
						Data: &event_usecase.MsgVote{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/cosmos.gov.v1beta1.MsgVote.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/cosmos.gov.v1beta1.MsgVote",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							ProposalId: "ProposalId",
							Voter:      "Voter",
							Option:     "Option",
						},
					},
					[]string{"voter"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgCreateValidator",
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
				&event_usecase.MsgCreateValidator{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_CREATE_VALIDATOR,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Description: usecase_model.ValidatorDescription{
						Moniker:         "Moniker",
						Identity:        "Identity",
						Website:         "Website",
						SecurityContact: "SecurityContact",
						Details:         "Details",
					},
					CommissionRates: usecase_model.ValidatorCommission{
						Rate:          "Rate",
						MaxRate:       "MaxRate",
						MaxChangeRate: "MaxChangeRate",
					},
					MinSelfDelegation: "MinSelfDelegation",
					DelegatorAddress:  "DelegatorAddress",
					ValidatorAddress:  "ValidatorAddress",
					TendermintPubkey:  "TendermintPubkey",
					Amount: coin.Coin{
						Denom:  "Denom",
						Amount: coin.NewInt(100),
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
					"delegatoraddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"delegatoraddress:/cosmos.staking.v1beta1.MsgCreateValidator",
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
						MessageType:     "/cosmos.staking.v1beta1.MsgCreateValidator",
						Data: &event_usecase.MsgCreateValidator{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/cosmos.staking.v1beta1.MsgCreateValidator.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/cosmos.staking.v1beta1.MsgCreateValidator",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Description: usecase_model.ValidatorDescription{
								Moniker:         "Moniker",
								Identity:        "Identity",
								Website:         "Website",
								SecurityContact: "SecurityContact",
								Details:         "Details",
							},
							CommissionRates: usecase_model.ValidatorCommission{
								Rate:          "Rate",
								MaxRate:       "MaxRate",
								MaxChangeRate: "MaxChangeRate",
							},
							MinSelfDelegation: "MinSelfDelegation",
							DelegatorAddress:  "DelegatorAddress",
							ValidatorAddress:  "ValidatorAddress",
							TendermintPubkey:  "TendermintPubkey",
							Amount: coin.Coin{
								Denom:  "Denom",
								Amount: coin.NewInt(100),
							},
						},
					},
					[]string{"delegatoraddress"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgEditValidator",
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
				&event_usecase.MsgEditValidator{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_EDIT_VALIDATOR,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Description: usecase_model.ValidatorDescription{
						Moniker:         "Moniker",
						Identity:        "Identity",
						Website:         "Website",
						SecurityContact: "SecurityContact",
						Details:         "Details",
					},
					ValidatorAddress:       "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
					MaybeCommissionRate:    nil,
					MaybeMinSelfDelegation: nil,
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
					"accountaddressprefix1fmprm0sjy6lz9llv7rltn0v2azzwcwzvgd99ce:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"accountaddressprefix1fmprm0sjy6lz9llv7rltn0v2azzwcwzvgd99ce:/cosmos.staking.v1beta1.MsgEditValidator",
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
						MessageType:     "/cosmos.staking.v1beta1.MsgEditValidator",
						Data: &event_usecase.MsgEditValidator{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/cosmos.staking.v1beta1.MsgEditValidator.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/cosmos.staking.v1beta1.MsgEditValidator",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Description: usecase_model.ValidatorDescription{
								Moniker:         "Moniker",
								Identity:        "Identity",
								Website:         "Website",
								SecurityContact: "SecurityContact",
								Details:         "Details",
							},
							ValidatorAddress:       "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
							MaybeCommissionRate:    nil,
							MaybeMinSelfDelegation: nil,
						},
					},
					[]string{"accountaddressprefix1fmprm0sjy6lz9llv7rltn0v2azzwcwzvgd99ce"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgDelegate",
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
				&event_usecase.MsgDelegate{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_DELEGATE,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Amount: coin.Coin{
						Denom:  "AmountDenom",
						Amount: coin.NewInt(100),
					},
					AutoClaimedRewards: coin.Coin{
						Denom:  "AutoClaimedRewardsDenom",
						Amount: coin.NewInt(1000),
					},
					DelegatorAddress: "DelegatorAddress",
					ValidatorAddress: "ValidatorAddress",
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
					"delegatoraddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"delegatoraddress:/cosmos.staking.v1beta1.MsgDelegate",
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
						MessageType:     "/cosmos.staking.v1beta1.MsgDelegate",
						Data: &event_usecase.MsgDelegate{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/cosmos.staking.v1beta1.MsgDelegate.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/cosmos.staking.v1beta1.MsgDelegate",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Amount: coin.Coin{
								Denom:  "AmountDenom",
								Amount: coin.NewInt(100),
							},
							AutoClaimedRewards: coin.Coin{
								Denom:  "AutoClaimedRewardsDenom",
								Amount: coin.NewInt(1000),
							},
							DelegatorAddress: "DelegatorAddress",
							ValidatorAddress: "ValidatorAddress",
						},
					},
					[]string{"delegatoraddress"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgUndelegate",
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
				&event_usecase.MsgUndelegate{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_UNDELEGATE,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Amount: coin.Coin{
						Denom:  "AmountDenom",
						Amount: coin.NewInt(100),
					},
					AutoClaimedRewards: coin.Coin{
						Denom:  "AutoClaimedRewardsDenom",
						Amount: coin.NewInt(1000),
					},
					DelegatorAddress:      "DelegatorAddress",
					ValidatorAddress:      "ValidatorAddress",
					MaybeUnbondCompleteAt: nil,
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
					"delegatoraddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"delegatoraddress:/cosmos.staking.v1beta1.MsgUndelegate",
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
						MessageType:     "/cosmos.staking.v1beta1.MsgUndelegate",
						Data: &event_usecase.MsgUndelegate{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/cosmos.staking.v1beta1.MsgUndelegate.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/cosmos.staking.v1beta1.MsgUndelegate",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Amount: coin.Coin{
								Denom:  "AmountDenom",
								Amount: coin.NewInt(100),
							},
							AutoClaimedRewards: coin.Coin{
								Denom:  "AutoClaimedRewardsDenom",
								Amount: coin.NewInt(1000),
							},
							DelegatorAddress:      "DelegatorAddress",
							ValidatorAddress:      "ValidatorAddress",
							MaybeUnbondCompleteAt: nil,
						},
					},
					[]string{"delegatoraddress"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgBeginRedelegate",
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
				&event_usecase.MsgBeginRedelegate{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_BEGIN_REDELEGATE,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Amount: coin.Coin{
						Denom:  "AmountDenom",
						Amount: coin.NewInt(100),
					},
					AutoClaimedRewards: coin.Coin{
						Denom:  "AutoClaimedRewardsDenom",
						Amount: coin.NewInt(1000),
					},
					DelegatorAddress:    "DelegatorAddress",
					ValidatorSrcAddress: "ValidatorSrcAddress",
					ValidatorDstAddress: "ValidatorDstAddress",
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
					"delegatoraddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"delegatoraddress:/cosmos.staking.v1beta1.MsgBeginRedelegate",
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
						MessageType:     "/cosmos.staking.v1beta1.MsgBeginRedelegate",
						Data: &event_usecase.MsgBeginRedelegate{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/cosmos.staking.v1beta1.MsgBeginRedelegate.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/cosmos.staking.v1beta1.MsgBeginRedelegate",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Amount: coin.Coin{
								Denom:  "AmountDenom",
								Amount: coin.NewInt(100),
							},
							AutoClaimedRewards: coin.Coin{
								Denom:  "AutoClaimedRewardsDenom",
								Amount: coin.NewInt(1000),
							},
							DelegatorAddress:    "DelegatorAddress",
							ValidatorSrcAddress: "ValidatorSrcAddress",
							ValidatorDstAddress: "ValidatorDstAddress",
						},
					},
					[]string{"delegatoraddress"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgUnjail",
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
				&event_usecase.MsgUnjail{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_UNJAIL,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					ValidatorAddr: "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
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
					"accountaddressprefix1fmprm0sjy6lz9llv7rltn0v2azzwcwzvgd99ce:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"accountaddressprefix1fmprm0sjy6lz9llv7rltn0v2azzwcwzvgd99ce:/cosmos.slashing.v1beta1.MsgUnjail",
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
						MessageType:     "/cosmos.slashing.v1beta1.MsgUnjail",
						Data: &event_usecase.MsgUnjail{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/cosmos.slashing.v1beta1.MsgUnjail.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/cosmos.slashing.v1beta1.MsgUnjail",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							ValidatorAddr: "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
						},
					},
					[]string{"accountaddressprefix1fmprm0sjy6lz9llv7rltn0v2azzwcwzvgd99ce"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgNFTIssueDenom",
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
				&event_usecase.MsgNFTIssueDenom{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_NFT_ISSUE_DENOM,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					DenomId:   "DenomId",
					DenomName: "DenomName",
					Schema:    "Schema",
					Sender:    "Sender",
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
					"sender:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"sender:/chainmain.nft.v1.MsgIssueDenom",
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
						MessageType:     "/chainmain.nft.v1.MsgIssueDenom",
						Data: &event_usecase.MsgNFTIssueDenom{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/chainmain.nft.v1.MsgIssueDenom.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/chainmain.nft.v1.MsgIssueDenom",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							DenomId:   "DenomId",
							DenomName: "DenomName",
							Schema:    "Schema",
							Sender:    "Sender",
						},
					},
					[]string{"sender"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgNFTMintNFT",
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
				&event_usecase.MsgNFTMintNFT{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_NFT_MINT_NFT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					DenomId:   "DenomId",
					TokenId:   "TokenId",
					TokenName: "TokenName",
					URI:       "URI",
					Data:      "Data",
					Sender:    "Sender",
					Recipient: "Recipient",
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
					"sender:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"sender:/chainmain.nft.v1.MsgMintNFT",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"recipient:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"recipient:/chainmain.nft.v1.MsgMintNFT",
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
						MessageType:     "/chainmain.nft.v1.MsgMintNFT",
						Data: &event_usecase.MsgNFTMintNFT{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/chainmain.nft.v1.MsgMintNFT.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/chainmain.nft.v1.MsgMintNFT",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							DenomId:   "DenomId",
							TokenId:   "TokenId",
							TokenName: "TokenName",
							URI:       "URI",
							Data:      "Data",
							Sender:    "Sender",
							Recipient: "Recipient",
						},
					},
					[]string{"sender", "recipient"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgNFTTransferNFT",
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
				&event_usecase.MsgNFTTransferNFT{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_NFT_TRANSFER_NFT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					DenomId:   "DenomId",
					TokenId:   "TokenId",
					Sender:    "Sender",
					Recipient: "Recipient",
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
					"sender:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"sender:/chainmain.nft.v1.MsgTransferNFT",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"recipient:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"recipient:/chainmain.nft.v1.MsgTransferNFT",
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
						MessageType:     "/chainmain.nft.v1.MsgTransferNFT",
						Data: &event_usecase.MsgNFTTransferNFT{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/chainmain.nft.v1.MsgTransferNFT.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/chainmain.nft.v1.MsgTransferNFT",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							DenomId:   "DenomId",
							TokenId:   "TokenId",
							Sender:    "Sender",
							Recipient: "Recipient",
						},
					},
					[]string{"sender", "recipient"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgNFTEditNFT",
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
				&event_usecase.MsgNFTEditNFT{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_NFT_EDIT_NFT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					DenomId:   "DenomId",
					TokenId:   "TokenId",
					TokenName: "TokenName",
					URI:       "URI",
					Data:      "Data",
					Sender:    "Sender",
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
					"sender:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"sender:/chainmain.nft.v1.MsgEditNFT",
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
						MessageType:     "/chainmain.nft.v1.MsgEditNFT",
						Data: &event_usecase.MsgNFTEditNFT{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/chainmain.nft.v1.MsgEditNFT.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/chainmain.nft.v1.MsgEditNFT",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							DenomId:   "DenomId",
							TokenId:   "TokenId",
							TokenName: "TokenName",
							URI:       "URI",
							Data:      "Data",
							Sender:    "Sender",
						},
					},
					[]string{"sender"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgNFTBurnNFT",
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
				&event_usecase.MsgNFTBurnNFT{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_NFT_BURN_NFT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					DenomId: "DenomId",
					TokenId: "TokenId",
					Sender:  "Sender",
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
					"sender:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"sender:/chainmain.nft.v1.MsgBurnNFT",
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
						MessageType:     "/chainmain.nft.v1.MsgBurnNFT",
						Data: &event_usecase.MsgNFTBurnNFT{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/chainmain.nft.v1.MsgBurnNFT.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/chainmain.nft.v1.MsgBurnNFT",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							DenomId: "DenomId",
							TokenId: "TokenId",
							Sender:  "Sender",
						},
					},
					[]string{"sender"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCCreateClient",
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
				&event_usecase.MsgIBCCreateClient{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_CREATE_CLIENT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgCreateClientParams{
						Signer:                     "Signer",
						ClientID:                   "ClientID",
						ClientType:                 "ClientType",
						MaybeTendermintLightClient: &ibc_model.TendermintLightClient{},
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
					"signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"signer:/ibc.core.client.v1.MsgCreateClient",
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
						MessageType:     "/ibc.core.client.v1.MsgCreateClient",
						Data: &event_usecase.MsgIBCCreateClient{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/ibc.core.client.v1.MsgCreateClient.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/ibc.core.client.v1.MsgCreateClient",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Params: ibc_model.MsgCreateClientParams{
								Signer:                     "Signer",
								ClientID:                   "ClientID",
								ClientType:                 "ClientType",
								MaybeTendermintLightClient: &ibc_model.TendermintLightClient{},
							},
						},
					},
					[]string{"signer"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCUpdateClient",
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
				&event_usecase.MsgIBCUpdateClient{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_UPDATE_CLIENT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgUpdateClientParams{
						Signer:                           "Signer",
						ClientID:                         "ClientID",
						ClientType:                       "ClientType",
						MaybeTendermintLightClientUpdate: &ibc_model.TendermintLightClientUpdate{},
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
					"signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"signer:/ibc.core.client.v1.MsgUpdateClient",
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
						MessageType:     "/ibc.core.client.v1.MsgUpdateClient",
						Data: &event_usecase.MsgIBCUpdateClient{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/ibc.core.client.v1.MsgUpdateClient.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/ibc.core.client.v1.MsgUpdateClient",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Params: ibc_model.MsgUpdateClientParams{
								Signer:                           "Signer",
								ClientID:                         "ClientID",
								ClientType:                       "ClientType",
								MaybeTendermintLightClientUpdate: &ibc_model.TendermintLightClientUpdate{},
							},
						},
					},
					[]string{"signer"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCConnectionOpenInit",
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
				&event_usecase.MsgIBCConnectionOpenInit{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_CONNECTION_OPEN_INIT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgConnectionOpenInitParams{
						RawMsgConnectionOpenInit: ibc_model.RawMsgConnectionOpenInit{
							Signer: "Signer",
						},
						ConnectionID: "ConnectionID",
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
					"signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"signer:/ibc.core.connection.v1.MsgConnectionOpenInit",
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
						MessageType:     "/ibc.core.connection.v1.MsgConnectionOpenInit",
						Data: &event_usecase.MsgIBCConnectionOpenInit{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/ibc.core.connection.v1.MsgConnectionOpenInit.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/ibc.core.connection.v1.MsgConnectionOpenInit",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Params: ibc_model.MsgConnectionOpenInitParams{
								RawMsgConnectionOpenInit: ibc_model.RawMsgConnectionOpenInit{
									Signer: "Signer",
								},
								ConnectionID: "ConnectionID",
							},
						},
					},
					[]string{"signer"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCConnectionOpenAck",
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
				&event_usecase.MsgIBCConnectionOpenAck{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_CONNECTION_OPEN_ACK,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgConnectionOpenAckParams{
						MsgConnectionOpenAckBaseParams: ibc_model.MsgConnectionOpenAckBaseParams{
							Signer: "Signer",
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
					"signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"signer:/ibc.core.connection.v1.MsgConnectionOpenAck",
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
						MessageType:     "/ibc.core.connection.v1.MsgConnectionOpenAck",
						Data: &event_usecase.MsgIBCConnectionOpenAck{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/ibc.core.connection.v1.MsgConnectionOpenAck.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/ibc.core.connection.v1.MsgConnectionOpenAck",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Params: ibc_model.MsgConnectionOpenAckParams{
								MsgConnectionOpenAckBaseParams: ibc_model.MsgConnectionOpenAckBaseParams{
									Signer: "Signer",
								},
							},
						},
					},
					[]string{"signer"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCConnectionOpenTry",
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
				&event_usecase.MsgIBCConnectionOpenTry{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_CONNECTION_OPEN_TRY,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgConnectionOpenTryParams{
						MsgConnectionOpenTryBaseParams: ibc_model.MsgConnectionOpenTryBaseParams{
							Signer: "Signer",
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
					"signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"signer:/ibc.core.connection.v1.MsgConnectionOpenTry",
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
						MessageType:     "/ibc.core.connection.v1.MsgConnectionOpenTry",
						Data: &event_usecase.MsgIBCConnectionOpenTry{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/ibc.core.connection.v1.MsgConnectionOpenTry.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/ibc.core.connection.v1.MsgConnectionOpenTry",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Params: ibc_model.MsgConnectionOpenTryParams{
								MsgConnectionOpenTryBaseParams: ibc_model.MsgConnectionOpenTryBaseParams{
									Signer: "Signer",
								},
							},
						},
					},
					[]string{"signer"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCConnectionOpenConfirm",
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
				&event_usecase.MsgIBCConnectionOpenConfirm{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_CONNECTION_OPEN_CONFIRM,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgConnectionOpenConfirmParams{
						RawMsgConnectionOpenConfirm: ibc_model.RawMsgConnectionOpenConfirm{
							Signer: "Signer",
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
					"signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"signer:/ibc.core.connection.v1.MsgConnectionOpenConfirm",
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
						MessageType:     "/ibc.core.connection.v1.MsgConnectionOpenConfirm",
						Data: &event_usecase.MsgIBCConnectionOpenConfirm{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/ibc.core.connection.v1.MsgConnectionOpenConfirm.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/ibc.core.connection.v1.MsgConnectionOpenConfirm",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Params: ibc_model.MsgConnectionOpenConfirmParams{
								RawMsgConnectionOpenConfirm: ibc_model.RawMsgConnectionOpenConfirm{
									Signer: "Signer",
								},
							},
						},
					},
					[]string{"signer"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCChannelOpenInit",
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
				&event_usecase.MsgIBCChannelOpenInit{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_CHANNEL_OPEN_INIT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgChannelOpenInitParams{
						RawMsgChannelOpenInit: ibc_model.RawMsgChannelOpenInit{
							Signer: "Signer",
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
					"signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"signer:/ibc.core.channel.v1.MsgChannelOpenInit",
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
						MessageType:     "/ibc.core.channel.v1.MsgChannelOpenInit",
						Data: &event_usecase.MsgIBCChannelOpenInit{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/ibc.core.channel.v1.MsgChannelOpenInit.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/ibc.core.channel.v1.MsgChannelOpenInit",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Params: ibc_model.MsgChannelOpenInitParams{
								RawMsgChannelOpenInit: ibc_model.RawMsgChannelOpenInit{
									Signer: "Signer",
								},
							},
						},
					},
					[]string{"signer"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCChannelOpenAck",
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
				&event_usecase.MsgIBCChannelOpenAck{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_CHANNEL_OPEN_ACK,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgChannelOpenAckParams{
						RawMsgChannelOpenAck: ibc_model.RawMsgChannelOpenAck{
							Signer: "Signer",
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
					"signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"signer:/ibc.core.channel.v1.MsgChannelOpenAck",
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
						MessageType:     "/ibc.core.channel.v1.MsgChannelOpenAck",
						Data: &event_usecase.MsgIBCChannelOpenAck{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/ibc.core.channel.v1.MsgChannelOpenAck.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/ibc.core.channel.v1.MsgChannelOpenAck",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Params: ibc_model.MsgChannelOpenAckParams{
								RawMsgChannelOpenAck: ibc_model.RawMsgChannelOpenAck{
									Signer: "Signer",
								},
							},
						},
					},
					[]string{"signer"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCChannelOpenTry",
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
				&event_usecase.MsgIBCChannelOpenTry{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_CHANNEL_OPEN_TRY,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgChannelOpenTryParams{
						RawMsgChannelOpenTry: ibc_model.RawMsgChannelOpenTry{
							Signer: "Signer",
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
					"signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"signer:/ibc.core.channel.v1.MsgChannelOpenTry",
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
						MessageType:     "/ibc.core.channel.v1.MsgChannelOpenTry",
						Data: &event_usecase.MsgIBCChannelOpenTry{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/ibc.core.channel.v1.MsgChannelOpenTry.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/ibc.core.channel.v1.MsgChannelOpenTry",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Params: ibc_model.MsgChannelOpenTryParams{
								RawMsgChannelOpenTry: ibc_model.RawMsgChannelOpenTry{
									Signer: "Signer",
								},
							},
						},
					},
					[]string{"signer"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCChannelOpenConfirm",
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
				&event_usecase.MsgIBCChannelOpenConfirm{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_CHANNEL_OPEN_CONFIRM,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgChannelOpenConfirmParams{
						RawMsgChannelOpenConfirm: ibc_model.RawMsgChannelOpenConfirm{
							Signer: "Signer",
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
					"signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"signer:/ibc.core.channel.v1.MsgChannelOpenConfirm",
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
						MessageType:     "/ibc.core.channel.v1.MsgChannelOpenConfirm",
						Data: &event_usecase.MsgIBCChannelOpenConfirm{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/ibc.core.channel.v1.MsgChannelOpenConfirm.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/ibc.core.channel.v1.MsgChannelOpenConfirm",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Params: ibc_model.MsgChannelOpenConfirmParams{
								RawMsgChannelOpenConfirm: ibc_model.RawMsgChannelOpenConfirm{
									Signer: "Signer",
								},
							},
						},
					},
					[]string{"signer"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCAcknowledgement",
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
				&event_usecase.MsgIBCAcknowledgement{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_ACKNOWLEDGEMENT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgAcknowledgementParams{
						MaybeFungibleTokenPacketData: &ibc_model.MsgAcknowledgementFungibleTokenPacketData{
							FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
								Sender:   "Sender",
								Receiver: "Receiver",
								Denom:    "Denom",
								Amount:   json.NewNumericStringFromUint64(100),
							},
							Success: true,
						},
						RawMsgAcknowledgement: ibc_model.RawMsgAcknowledgement{
							Signer: "Signer",
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
					"signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"signer:/ibc.core.channel.v1.MsgAcknowledgement",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"sender:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"sender:/ibc.core.channel.v1.MsgAcknowledgement",
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
						MessageType:     "/ibc.core.channel.v1.MsgAcknowledgement",
						Data: &event_usecase.MsgIBCAcknowledgement{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/ibc.core.channel.v1.MsgAcknowledgement.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/ibc.core.channel.v1.MsgAcknowledgement",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Params: ibc_model.MsgAcknowledgementParams{
								MaybeFungibleTokenPacketData: &ibc_model.MsgAcknowledgementFungibleTokenPacketData{
									FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
										Sender:   "Sender",
										Receiver: "Receiver",
										Denom:    "Denom",
										Amount:   json.NewNumericStringFromUint64(100),
									},
									Success: true,
								},
								RawMsgAcknowledgement: ibc_model.RawMsgAcknowledgement{
									Signer: "Signer",
								},
							},
						},
					},
					[]string{"signer", "sender"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCRecvPacket",
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
				&event_usecase.MsgIBCRecvPacket{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_RECV_PACKET,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgRecvPacketParams{
						MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
							FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
								Sender:   "Sender",
								Receiver: "Receiver",
								Denom:    "Denom",
								Amount:   json.NewNumericStringFromUint64(100),
							},
							Success:                true,
							MaybeDenominationTrace: nil,
						},
						RawMsgRecvPacket: ibc_model.RawMsgRecvPacket{
							Signer: "Signer",
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
					"signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"signer:/ibc.core.channel.v1.MsgRecvPacket",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"receiver:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"receiver:/ibc.core.channel.v1.MsgRecvPacket",
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
						MessageType:     "/ibc.core.channel.v1.MsgRecvPacket",
						Data: &event_usecase.MsgIBCRecvPacket{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/ibc.core.channel.v1.MsgRecvPacket.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/ibc.core.channel.v1.MsgRecvPacket",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Params: ibc_model.MsgRecvPacketParams{
								MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
									FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
										Sender:   "Sender",
										Receiver: "Receiver",
										Denom:    "Denom",
										Amount:   json.NewNumericStringFromUint64(100),
									},
									Success:                true,
									MaybeDenominationTrace: nil,
								},
								RawMsgRecvPacket: ibc_model.RawMsgRecvPacket{
									Signer: "Signer",
								},
							},
						},
					},
					[]string{"signer", "receiver"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCTransferTransfer",
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
				&event_usecase.MsgIBCTransferTransfer{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_TRANSFER_TRANSFER,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgTransferParams{
						RawMsgTransfer: ibc_model.RawMsgTransfer{
							Sender: "Sender",
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
					"sender:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"sender:/ibc.applications.transfer.v1.MsgTransfer",
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
						MessageType:     "/ibc.applications.transfer.v1.MsgTransfer",
						Data: &event_usecase.MsgIBCTransferTransfer{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/ibc.applications.transfer.v1.MsgTransfer.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/ibc.applications.transfer.v1.MsgTransfer",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Params: ibc_model.MsgTransferParams{
								RawMsgTransfer: ibc_model.RawMsgTransfer{
									Sender: "Sender",
								},
							},
						},
					},
					[]string{"sender"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCTimeout",
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
				&event_usecase.MsgIBCTimeout{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_TIMEOUT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgTimeoutParams{
						RawMsgTimeout: ibc_model.RawMsgTimeout{
							Signer: "Signer",
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
					"signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"signer:/ibc.core.channel.v1.MsgTimeout",
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
						MessageType:     "/ibc.core.channel.v1.MsgTimeout",
						Data: &event_usecase.MsgIBCTimeout{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/ibc.core.channel.v1.MsgTimeout.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/ibc.core.channel.v1.MsgTimeout",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Params: ibc_model.MsgTimeoutParams{
								RawMsgTimeout: ibc_model.RawMsgTimeout{
									Signer: "Signer",
								},
							},
						},
					},
					[]string{"signer"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCTimeoutOnClose",
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
				&event_usecase.MsgIBCTimeoutOnClose{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_TIMEOUT_ON_CLOSE,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgTimeoutOnCloseParams{
						RawMsgTimeoutOnClose: ibc_model.RawMsgTimeoutOnClose{
							Signer: "Signer",
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
					"signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"signer:/ibc.core.channel.v1.MsgTimeoutOnClose",
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
						MessageType:     "/ibc.core.channel.v1.MsgTimeoutOnClose",
						Data: &event_usecase.MsgIBCTimeoutOnClose{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "/ibc.core.channel.v1.MsgTimeoutOnClose.Created",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "/ibc.core.channel.v1.MsgTimeoutOnClose",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							Params: ibc_model.MsgTimeoutOnCloseParams{
								RawMsgTimeoutOnClose: ibc_model.RawMsgTimeoutOnClose{
									Signer: "Signer",
								},
							},
						},
					},
					[]string{"signer"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCChannelCloseInit",
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
				&event_usecase.MsgIBCChannelCloseInit{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_CHANNEL_CLOSE_INIT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgChannelCloseInitParams{
						RawMsgChannelCloseInit: ibc_model.RawMsgChannelCloseInit{
							Signer: "Signer",
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[1].(*event_usecase.MsgIBCChannelCloseInit)

				mockAccountMessagesTotalView := account_message_view.NewMockAccountMessagesTotalView(nil).(*account_message_view.MockAccountMessagesTotalView)
				mocks = append(mocks, &mockAccountMessagesTotalView.Mock)

				account_message.NewAccountMessagesTotal = func(_ *rdb.Handle) account_message_view.AccountMessagesTotal {
					return mockAccountMessagesTotalView
				}

				mockAccountMessagesTotalView.On(
					"Increment",
					"signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"signer:/ibc.core.channel.v1.MsgChannelCloseInit",
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
						MessageType:     "/ibc.core.channel.v1.MsgChannelCloseInit",
						Data:            typedEvent,
					},
					[]string{"signer"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgGrant",
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
				&event_usecase.MsgGrant{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_GRANT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: usecase_model.MsgGrantParams{
						MaybeSendGrant: &model.RawMsgSendGrant{
							Granter: "Granter",
							Grantee: "Grantee",
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[1].(*event_usecase.MsgGrant)

				mockAccountMessagesTotalView := account_message_view.NewMockAccountMessagesTotalView(nil).(*account_message_view.MockAccountMessagesTotalView)
				mocks = append(mocks, &mockAccountMessagesTotalView.Mock)

				account_message.NewAccountMessagesTotal = func(_ *rdb.Handle) account_message_view.AccountMessagesTotal {
					return mockAccountMessagesTotalView
				}

				mockAccountMessagesTotalView.On(
					"Increment",
					"granter:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"granter:/cosmos.authz.v1beta1.MsgGrant",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"grantee:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"grantee:/cosmos.authz.v1beta1.MsgGrant",
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
						MessageType:     "/cosmos.authz.v1beta1.MsgGrant",
						Data:            typedEvent,
					},
					[]string{"granter", "grantee"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgRevoke",
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
				&event_usecase.MsgRevoke{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_REVOKE,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: usecase_model.MsgRevokeParams{
						RawMsgRevoke: usecase_model.RawMsgRevoke{
							Granter: "Granter",
							Grantee: "Grantee",
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[1].(*event_usecase.MsgRevoke)

				mockAccountMessagesTotalView := account_message_view.NewMockAccountMessagesTotalView(nil).(*account_message_view.MockAccountMessagesTotalView)
				mocks = append(mocks, &mockAccountMessagesTotalView.Mock)

				account_message.NewAccountMessagesTotal = func(_ *rdb.Handle) account_message_view.AccountMessagesTotal {
					return mockAccountMessagesTotalView
				}

				mockAccountMessagesTotalView.On(
					"Increment",
					"granter:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"granter:/cosmos.authz.v1beta1.MsgRevoke",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"grantee:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"grantee:/cosmos.authz.v1beta1.MsgRevoke",
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
						MessageType:     "/cosmos.authz.v1beta1.MsgRevoke",
						Data:            typedEvent,
					},
					[]string{"granter", "grantee"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgExec",
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
				&event_usecase.MsgExec{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_EXEC,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: usecase_model.MsgExecParams{
						RawMsgExec: usecase_model.RawMsgExec{
							Grantee: "Grantee",
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[1].(*event_usecase.MsgExec)

				mockAccountMessagesTotalView := account_message_view.NewMockAccountMessagesTotalView(nil).(*account_message_view.MockAccountMessagesTotalView)
				mocks = append(mocks, &mockAccountMessagesTotalView.Mock)

				account_message.NewAccountMessagesTotal = func(_ *rdb.Handle) account_message_view.AccountMessagesTotal {
					return mockAccountMessagesTotalView
				}

				mockAccountMessagesTotalView.On(
					"Increment",
					"grantee:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"grantee:/cosmos.authz.v1beta1.MsgExec",
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
						MessageType:     "/cosmos.authz.v1beta1.MsgExec",
						Data:            typedEvent,
					},
					[]string{"grantee"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgGrantAllowance",
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
				&event_usecase.MsgGrantAllowance{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_GRANT_ALLOWANCE,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: usecase_model.MsgGrantAllowanceParams{
						MaybeBasicAllowance: &usecase_model.RawMsgGrantBasicAllowance{
							Granter: "Granter",
							Grantee: "Grantee",
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[1].(*event_usecase.MsgGrantAllowance)

				mockAccountMessagesTotalView := account_message_view.NewMockAccountMessagesTotalView(nil).(*account_message_view.MockAccountMessagesTotalView)
				mocks = append(mocks, &mockAccountMessagesTotalView.Mock)

				account_message.NewAccountMessagesTotal = func(_ *rdb.Handle) account_message_view.AccountMessagesTotal {
					return mockAccountMessagesTotalView
				}

				mockAccountMessagesTotalView.On(
					"Increment",
					"granter:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"granter:/cosmos.feegrant.v1beta1.MsgGrantAllowance",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"grantee:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"grantee:/cosmos.feegrant.v1beta1.MsgGrantAllowance",
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
						MessageType:     "/cosmos.feegrant.v1beta1.MsgGrantAllowance",
						Data:            typedEvent,
					},
					[]string{"granter", "grantee"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgRevokeAllowance",
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
				&event_usecase.MsgRevokeAllowance{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_REVOKE_ALLOWANCE,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: usecase_model.MsgRevokeAllowanceParams{
						RawMsgRevokeAllowance: usecase_model.RawMsgRevokeAllowance{
							Granter: "Granter",
							Grantee: "Grantee",
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[1].(*event_usecase.MsgRevokeAllowance)

				mockAccountMessagesTotalView := account_message_view.NewMockAccountMessagesTotalView(nil).(*account_message_view.MockAccountMessagesTotalView)
				mocks = append(mocks, &mockAccountMessagesTotalView.Mock)

				account_message.NewAccountMessagesTotal = func(_ *rdb.Handle) account_message_view.AccountMessagesTotal {
					return mockAccountMessagesTotalView
				}

				mockAccountMessagesTotalView.On(
					"Increment",
					"granter:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"granter:/cosmos.feegrant.v1beta1.MsgRevokeAllowance",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"grantee:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"grantee:/cosmos.feegrant.v1beta1.MsgRevokeAllowance",
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
						MessageType:     "/cosmos.feegrant.v1beta1.MsgRevokeAllowance",
						Data:            typedEvent,
					},
					[]string{"granter", "grantee"},
				).Return(nil)

				account_message.UpdateLastHandledEventHeight = func(_ *account_message.AccountMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgCreateVestingAccount",
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
				&event_usecase.MsgCreateVestingAccount{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_CREATE_VESTING_ACCOUNT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: usecase_model.MsgCreateVestingAccountParams{
						RawMsgCreateVestingAccount: usecase_model.RawMsgCreateVestingAccount{
							FromAddress: "FromAddress",
							ToAddress:   "ToAddress",
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[1].(*event_usecase.MsgCreateVestingAccount)

				mockAccountMessagesTotalView := account_message_view.NewMockAccountMessagesTotalView(nil).(*account_message_view.MockAccountMessagesTotalView)
				mocks = append(mocks, &mockAccountMessagesTotalView.Mock)

				account_message.NewAccountMessagesTotal = func(_ *rdb.Handle) account_message_view.AccountMessagesTotal {
					return mockAccountMessagesTotalView
				}

				mockAccountMessagesTotalView.On(
					"Increment",
					"fromaddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"fromaddress:/cosmos.vesting.v1beta1.MsgCreateVestingAccount",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"toaddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"toaddress:/cosmos.vesting.v1beta1.MsgCreateVestingAccount",
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
						MessageType:     "/cosmos.vesting.v1beta1.MsgCreateVestingAccount",
						Data:            typedEvent,
					},
					[]string{"fromaddress", "toaddress"},
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

		projection := NewAccountMessageProjection(mockRDbConn)
		err := projection.HandleEvents(1, tc.Events)
		assert.NoError(t, err)

		for _, m := range mocks {
			m.AssertExpectations(t)
		}

		fmt.Println(tc.Name, "Passed")
	}
}
