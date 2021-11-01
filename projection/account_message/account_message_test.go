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
					"FromAddress:MsgSend",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"ToAddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"ToAddress:MsgSend",
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
						MessageType:     "MsgSend",
						Data: &event_usecase.MsgSend{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgSendCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgSend",
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
					"InputAddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"InputAddress:MsgMultiSend",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"OutputAddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"OutputAddress:MsgMultiSend",
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
						MessageType:     "MsgMultiSend",
						Data: &event_usecase.MsgMultiSend{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgMultiSendCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgMultiSend",
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
					[]string{"InputAddress", "OutputAddress"},
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
					"DelegatorAddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"DelegatorAddress:MsgSetWithdrawAddress",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"WithdrawAddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"WithdrawAddress:MsgSetWithdrawAddress",
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
						MessageType:     "MsgSetWithdrawAddress",
						Data: &event_usecase.MsgSetWithdrawAddress{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgSetWithdrawAddressCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgSetWithdrawAddress",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							MsgSetWithdrawAddressParams: usecase_model.MsgSetWithdrawAddressParams{
								DelegatorAddress: "DelegatorAddress",
								WithdrawAddress:  "WithdrawAddress",
							},
						},
					},
					[]string{"DelegatorAddress", "WithdrawAddress"},
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
					"DelegatorAddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"DelegatorAddress:MsgWithdrawDelegatorReward",
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
						MessageType:     "MsgWithdrawDelegatorReward",
						Data: &event_usecase.MsgWithdrawDelegatorReward{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgWithdrawDelegatorRewardCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgWithdrawDelegatorReward",
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
					[]string{"DelegatorAddress"},
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
					"RecipientAddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"RecipientAddress:MsgWithdrawValidatorCommission",
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
						MessageType:     "MsgWithdrawValidatorCommission",
						Data: &event_usecase.MsgWithdrawValidatorCommission{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgWithdrawValidatorCommissionCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgWithdrawValidatorCommission",
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
					[]string{"RecipientAddress"},
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
					"Depositor:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Depositor:MsgFundCommunityPool",
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
						MessageType:     "MsgFundCommunityPool",
						Data: &event_usecase.MsgFundCommunityPool{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgFundCommunityPoolCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgFundCommunityPool",
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
					[]string{"Depositor"},
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
					"ProposerAddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"ProposerAddress:MsgSubmitParamUpdateProposal",
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
						MessageType:     "MsgSubmitParamUpdateProposal",
						Data: &event_usecase.MsgSubmitParamChangeProposal{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgSubmitParamUpdateProposalCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgSubmitParamUpdateProposal",
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
					[]string{"ProposerAddress"},
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
					"ProposerAddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"ProposerAddress:MsgSubmitCommunityPoolSpendProposal",
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
						MessageType:     "MsgSubmitCommunityPoolSpendProposal",
						Data: &event_usecase.MsgSubmitCommunityPoolSpendProposal{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgSubmitCommunityPoolSpendProposalCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgSubmitCommunityPoolSpendProposal",
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
					[]string{"ProposerAddress"},
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
					"ProposerAddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"ProposerAddress:MsgSubmitCancelSoftwareUpgradeProposal",
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
						MessageType:     "MsgSubmitCancelSoftwareUpgradeProposal",
						Data: &event_usecase.MsgSubmitCancelSoftwareUpgradeProposal{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgSubmitCancelSoftwareUpgradeProposalCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgSubmitCancelSoftwareUpgradeProposal",
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
					[]string{"ProposerAddress"},
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
					"Depositor:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Depositor:MsgDeposit",
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
						MessageType:     "MsgDeposit",
						Data: &event_usecase.MsgDeposit{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgDepositCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgDeposit",
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
					[]string{"Depositor"},
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
					"Voter:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Voter:MsgVote",
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
						MessageType:     "MsgVote",
						Data: &event_usecase.MsgVote{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgVoteCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgVote",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							ProposalId: "ProposalId",
							Voter:      "Voter",
							Option:     "Option",
						},
					},
					[]string{"Voter"},
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
					"DelegatorAddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"DelegatorAddress:MsgCreateValidator",
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
						MessageType:     "MsgCreateValidator",
						Data: &event_usecase.MsgCreateValidator{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgCreateValidatorCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgCreateValidator",
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
					[]string{"DelegatorAddress"},
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
					"accountaddressprefix1fmprm0sjy6lz9llv7rltn0v2azzwcwzvgd99ce:MsgEditValidator",
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
						MessageType:     "MsgEditValidator",
						Data: &event_usecase.MsgEditValidator{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgEditValidatorCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgEditValidator",
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
					"DelegatorAddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"DelegatorAddress:MsgDelegate",
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
						MessageType:     "MsgDelegate",
						Data: &event_usecase.MsgDelegate{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgDelegateCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgDelegate",
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
					[]string{"DelegatorAddress"},
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
					"DelegatorAddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"DelegatorAddress:MsgUndelegate",
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
						MessageType:     "MsgUndelegate",
						Data: &event_usecase.MsgUndelegate{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgUndelegateCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgUndelegate",
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
					[]string{"DelegatorAddress"},
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
					"DelegatorAddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"DelegatorAddress:MsgBeginRedelegate",
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
						MessageType:     "MsgBeginRedelegate",
						Data: &event_usecase.MsgBeginRedelegate{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgBeginRedelegateCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgBeginRedelegate",
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
					[]string{"DelegatorAddress"},
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
					"accountaddressprefix1fmprm0sjy6lz9llv7rltn0v2azzwcwzvgd99ce:MsgUnjail",
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
						MessageType:     "MsgUnjail",
						Data: &event_usecase.MsgUnjail{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgUnjailCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgUnjail",
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
					"Sender:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Sender:MsgIssueDenom",
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
						MessageType:     "MsgIssueDenom",
						Data: &event_usecase.MsgNFTIssueDenom{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgIssueDenomCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgIssueDenom",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							DenomId:   "DenomId",
							DenomName: "DenomName",
							Schema:    "Schema",
							Sender:    "Sender",
						},
					},
					[]string{"Sender"},
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
					"Sender:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Sender:MsgMintNFT",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Recipient:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Recipient:MsgMintNFT",
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
						MessageType:     "MsgMintNFT",
						Data: &event_usecase.MsgNFTMintNFT{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgMintNFTCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgMintNFT",
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
					[]string{"Sender", "Recipient"},
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
					"Sender:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Sender:MsgTransferNFT",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Recipient:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Recipient:MsgTransferNFT",
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
						MessageType:     "MsgTransferNFT",
						Data: &event_usecase.MsgNFTTransferNFT{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgTransferNFTCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgTransferNFT",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							DenomId:   "DenomId",
							TokenId:   "TokenId",
							Sender:    "Sender",
							Recipient: "Recipient",
						},
					},
					[]string{"Sender", "Recipient"},
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
					"Sender:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Sender:MsgEditNFT",
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
						MessageType:     "MsgEditNFT",
						Data: &event_usecase.MsgNFTEditNFT{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgEditNFTCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgEditNFT",
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
					[]string{"Sender"},
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
					"Sender:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Sender:MsgBurnNFT",
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
						MessageType:     "MsgBurnNFT",
						Data: &event_usecase.MsgNFTBurnNFT{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgBurnNFTCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgBurnNFT",
								MsgTxHash: "TxHash",
								MsgIndex:  0,
							},
							DenomId: "DenomId",
							TokenId: "TokenId",
							Sender:  "Sender",
						},
					},
					[]string{"Sender"},
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
					"Signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Signer:MsgCreateClient",
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
						MessageType:     "MsgCreateClient",
						Data: &event_usecase.MsgIBCCreateClient{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgCreateClientCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgCreateClient",
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
					[]string{"Signer"},
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
					"Signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Signer:MsgUpdateClient",
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
						MessageType:     "MsgUpdateClient",
						Data: &event_usecase.MsgIBCUpdateClient{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgUpdateClientCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgUpdateClient",
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
					[]string{"Signer"},
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
					"Signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Signer:MsgConnectionOpenInit",
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
						MessageType:     "MsgConnectionOpenInit",
						Data: &event_usecase.MsgIBCConnectionOpenInit{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgConnectionOpenInitCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgConnectionOpenInit",
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
					[]string{"Signer"},
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
					"Signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Signer:MsgConnectionOpenAck",
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
						MessageType:     "MsgConnectionOpenAck",
						Data: &event_usecase.MsgIBCConnectionOpenAck{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgConnectionOpenAckCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgConnectionOpenAck",
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
					[]string{"Signer"},
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
					"Signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Signer:MsgConnectionOpenTry",
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
						MessageType:     "MsgConnectionOpenTry",
						Data: &event_usecase.MsgIBCConnectionOpenTry{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgConnectionOpenTryCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgConnectionOpenTry",
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
					[]string{"Signer"},
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
					"Signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Signer:MsgConnectionOpenConfirm",
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
						MessageType:     "MsgConnectionOpenConfirm",
						Data: &event_usecase.MsgIBCConnectionOpenConfirm{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgConnectionOpenConfirmCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgConnectionOpenConfirm",
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
					[]string{"Signer"},
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
					"Signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Signer:MsgChannelOpenInit",
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
						MessageType:     "MsgChannelOpenInit",
						Data: &event_usecase.MsgIBCChannelOpenInit{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgChannelOpenInitCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgChannelOpenInit",
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
					[]string{"Signer"},
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
					"Signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Signer:MsgChannelOpenAck",
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
						MessageType:     "MsgChannelOpenAck",
						Data: &event_usecase.MsgIBCChannelOpenAck{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgChannelOpenAckCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgChannelOpenAck",
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
					[]string{"Signer"},
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
					"Signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Signer:MsgChannelOpenTry",
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
						MessageType:     "MsgChannelOpenTry",
						Data: &event_usecase.MsgIBCChannelOpenTry{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgChannelOpenTryCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgChannelOpenTry",
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
					[]string{"Signer"},
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
					"Signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Signer:MsgChannelOpenConfirm",
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
						MessageType:     "MsgChannelOpenConfirm",
						Data: &event_usecase.MsgIBCChannelOpenConfirm{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgChannelOpenConfirmCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgChannelOpenConfirm",
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
					[]string{"Signer"},
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
					"Signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Signer:MsgAcknowledgement",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Sender:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Sender:MsgAcknowledgement",
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
						MessageType:     "MsgAcknowledgement",
						Data: &event_usecase.MsgIBCAcknowledgement{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgAcknowledgementCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgAcknowledgement",
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
					[]string{"Signer", "Sender"},
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
					"Signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Signer:MsgRecvPacket",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Receiver:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Receiver:MsgRecvPacket",
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
						MessageType:     "MsgRecvPacket",
						Data: &event_usecase.MsgIBCRecvPacket{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgRecvPacketCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgRecvPacket",
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
					[]string{"Signer", "Receiver"},
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
					"Sender:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Sender:MsgTransfer",
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
						MessageType:     "MsgTransfer",
						Data: &event_usecase.MsgIBCTransferTransfer{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgTransferCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgTransfer",
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
					[]string{"Sender"},
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
					"Signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Signer:MsgTimeout",
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
						MessageType:     "MsgTimeout",
						Data: &event_usecase.MsgIBCTimeout{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgTimeoutCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgTimeout",
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
					[]string{"Signer"},
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
					"Signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Signer:MsgTimeoutOnClose",
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
						MessageType:     "MsgTimeoutOnClose",
						Data: &event_usecase.MsgIBCTimeoutOnClose{
							MsgBase: event_usecase.MsgBase{
								Base: entity_event.Base{
									EventName:    "MsgTimeoutOnCloseCreated",
									EventVersion: 1,
									BlockHeight:  1,
									EventUUID:    "TESTUUID",
								},
								MsgName:   "MsgTimeoutOnClose",
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
					[]string{"Signer"},
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
					"Signer:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Signer:MsgChannelCloseInit",
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
						MessageType:     "MsgChannelCloseInit",
						Data:            typedEvent,
					},
					[]string{"Signer"},
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
					"Granter:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Granter:MsgGrant",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Grantee:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Grantee:MsgGrant",
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
						MessageType:     "MsgGrant",
						Data:            typedEvent,
					},
					[]string{"Granter", "Grantee"},
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
					"Granter:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Granter:MsgRevoke",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Grantee:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Grantee:MsgRevoke",
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
						MessageType:     "MsgRevoke",
						Data:            typedEvent,
					},
					[]string{"Granter", "Grantee"},
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
					"Grantee:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Grantee:MsgExec",
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
						MessageType:     "MsgExec",
						Data:            typedEvent,
					},
					[]string{"Grantee"},
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
					"Granter:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Granter:MsgGrantAllowance",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Grantee:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Grantee:MsgGrantAllowance",
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
						MessageType:     "MsgGrantAllowance",
						Data:            typedEvent,
					},
					[]string{"Granter", "Grantee"},
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
					"Granter:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Granter:MsgRevokeAllowance",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Grantee:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"Grantee:MsgRevokeAllowance",
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
						MessageType:     "MsgRevokeAllowance",
						Data:            typedEvent,
					},
					[]string{"Granter", "Grantee"},
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
					"FromAddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"FromAddress:MsgCreateVestingAccount",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"ToAddress:-",
					int64(1),
				).Return(nil)

				mockAccountMessagesTotalView.On(
					"Increment",
					"ToAddress:MsgCreateVestingAccount",
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
						MessageType:     "MsgCreateVestingAccount",
						Data:            typedEvent,
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

		projection := NewAccountMessageProjection(mockRDbConn)
		err := projection.HandleEvents(1, tc.Events)
		assert.NoError(t, err)

		for _, m := range mocks {
			m.AssertExpectations(t)
		}

		fmt.Println(tc.Name, "Passed")
	}
}
