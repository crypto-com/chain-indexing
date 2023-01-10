package proposal_test

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	testify_mock "github.com/stretchr/testify/mock"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbparambase"
	rdbparambase_types "github.com/crypto-com/chain-indexing/appinterface/projection/rdbparambase/types"
	rdbparambase_view "github.com/crypto-com/chain-indexing/appinterface/projection/rdbparambase/view"
	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbvalidatorbase"
	rdbvalidatorbase_view "github.com/crypto-com/chain-indexing/appinterface/projection/rdbvalidatorbase/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/projection/proposal"
	"github.com/crypto-com/chain-indexing/projection/proposal/types"
	"github.com/crypto-com/chain-indexing/projection/proposal/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	usecase_event "github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	v1_model "github.com/crypto-com/chain-indexing/usecase/model/v1"
)

func NewProposalProjection(rdbConn rdb.Conn) *proposal.Proposal {

	return proposal.NewProposal(
		nil,
		rdbConn,
		"",
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
	mockTx.On("ToHandle").Return(nil).Maybe()
	mockTx.On("Rollback").Return(nil).Maybe()
	mockTx.On("Commit").Return(nil).Maybe()

	return mockTx
}

func TestProposal_HandleEvents(t *testing.T) {
	testCases := []struct {
		Name     string
		Events   []entity_event.Event
		MockFunc func(events []entity_event.Event) []*testify_mock.Mock
	}{
		{
			Name: "HandleMsgSubmitTextProposal",
			Events: []entity_event.Event{
				&usecase_event.MsgSubmitTextProposal{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_SUBMIT_TEXT_PROPOSAL,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					MsgSubmitTextProposalParams: model.MsgSubmitTextProposalParams{
						MaybeProposalId: primptr.String("MaybeProposalId"),
						Content: model.MsgSubmitTextProposalContent{
							Title:       "Title",
							Description: "Description",
							Type:        "Type",
						},
						ProposerAddress: "ProposerAddress",
						InitialDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {

				mockValidatorsView := &rdbvalidatorbase_view.MockValidatorsView{}
				mocks = append(mocks, &mockValidatorsView.Mock)
				mockValidatorsView.
					On("FindLastBy", rdbvalidatorbase_view.ValidatorIdentity{
						MaybeInititalDelegatorAddress: primptr.String("ProposerAddress"),
					}).
					Return(&rdbvalidatorbase_view.ValidatorRow{
						OperatorAddress: "ProposerOperatorAddress",
					}, nil)

				proposal.ValidatorBaseGetView = func(
					_ *rdbvalidatorbase.Base,
					_ *rdb.Handle,
				) rdbvalidatorbase_view.Validators {
					return mockValidatorsView
				}

				mockParamsView := &rdbparambase_view.MockParamsView{}
				mocks = append(mocks, &mockParamsView.Mock)
				mockParamsView.
					On("FindDurationBy", rdbparambase_types.ParamAccessor{
						Module: "gov",
						Key:    "max_deposit_period",
					}).
					Return(time.Duration(1), nil)

				proposal.ParamBaseGetView = func(
					_ *rdbparambase.Base,
					_ *rdb.Handle,
				) rdbparambase_view.Params {
					return mockParamsView
				}

				mockProposalsView := &view.MockProposalsView{}
				mocks = append(mocks, &mockProposalsView.Mock)
				mockProposalsView.
					On("Insert", &view.ProposalRow{
						ProposalId:                   "MaybeProposalId",
						Title:                        "Title",
						Description:                  "Description",
						Type:                         "Type",
						Status:                       "DEPOSIT_PERIOD",
						ProposerAddress:              "ProposerAddress",
						MaybeProposerOperatorAddress: primptr.String("ProposerOperatorAddress"),
						Data:                         nil,
						InitialDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
						TotalDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
						TotalVote:                 big.NewInt(0),
						TransactionHash:           "TxHash",
						SubmitBlockHeight:         1,
						SubmitTime:                utctime.UTCTime{},
						DepositEndTime:            utctime.UTCTime{}.Add(time.Duration(1)),
						MaybeVotingStartTime:      nil,
						MaybeVotingEndTime:        nil,
						MaybeVotingEndBlockHeight: nil,
					}).
					Return(nil)

				proposal.NewProposals = func(
					_ *rdb.Handle,
				) view.Proposals {
					return mockProposalsView
				}

				mockDepositorsView := &view.MockDepositorsView{}
				mocks = append(mocks, &mockDepositorsView.Mock)
				mockDepositorsView.
					On("FindByProposalIdAndTxHash", "MaybeProposalId", "TxHash").
					Return(nil, nil)

				mockDepositorsView.
					On("Insert", &view.DepositorRow{
						ProposalId:                    "MaybeProposalId",
						DepositorAddress:              "ProposerAddress",
						MaybeDepositorOperatorAddress: primptr.String("ProposerOperatorAddress"),
						TransactionHash:               "TxHash",
						DepositAtBlockHeight:          1,
						DepositAtBlockTime:            utctime.UTCTime{},
						Amount: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
					}).
					Return(nil)

				proposal.NewDepositors = func(
					_ *rdb.Handle,
				) view.Depositors {
					return mockDepositorsView
				}

				mockDepositorsTotalView := &view.MockDepositorsTotalView{}
				mocks = append(mocks, &mockDepositorsTotalView.Mock)
				mockDepositorsTotalView.
					On("Increment", "MaybeProposalId", int64(1)).
					Return(nil)

				proposal.NewDepositorsTotal = func(handle *rdb.Handle) view.DepositorsTotal {
					return mockDepositorsTotalView
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgSubmitParamChangeProposal",
			Events: []entity_event.Event{
				&usecase_event.MsgSubmitParamChangeProposal{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_SUBMIT_PARAM_CHANGE_PROPOSAL,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					MsgSubmitParamChangeProposalParams: model.MsgSubmitParamChangeProposalParams{
						MaybeProposalId: primptr.String("MaybeProposalId"),
						Content: model.MsgSubmitParamChangeProposalContent{
							Title:       "Title",
							Description: "Description",
							Type:        "Type",
							Changes: []model.MsgSubmitParamChangeProposalChange{
								{
									Subspace: "Subspace",
									Key:      "Key",
									Value:    json.RawMessage([]byte("Value")),
								},
							},
						},
						ProposerAddress: "ProposerAddress",
						InitialDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {

				mockValidatorsView := &rdbvalidatorbase_view.MockValidatorsView{}
				mocks = append(mocks, &mockValidatorsView.Mock)
				mockValidatorsView.
					On("FindLastBy", rdbvalidatorbase_view.ValidatorIdentity{
						MaybeInititalDelegatorAddress: primptr.String("ProposerAddress"),
					}).
					Return(&rdbvalidatorbase_view.ValidatorRow{
						OperatorAddress: "ProposerOperatorAddress",
					}, nil)

				proposal.ValidatorBaseGetView = func(
					_ *rdbvalidatorbase.Base,
					_ *rdb.Handle,
				) rdbvalidatorbase_view.Validators {
					return mockValidatorsView
				}

				mockParamsView := &rdbparambase_view.MockParamsView{}
				mocks = append(mocks, &mockParamsView.Mock)
				mockParamsView.
					On("FindDurationBy", rdbparambase_types.ParamAccessor{
						Module: "gov",
						Key:    "max_deposit_period",
					}).
					Return(time.Duration(1), nil)

				proposal.ParamBaseGetView = func(
					_ *rdbparambase.Base,
					_ *rdb.Handle,
				) rdbparambase_view.Params {
					return mockParamsView
				}

				mockProposalsView := &view.MockProposalsView{}
				mocks = append(mocks, &mockProposalsView.Mock)
				mockProposalsView.
					On("Insert", &view.ProposalRow{
						ProposalId:                   "MaybeProposalId",
						Title:                        "Title",
						Description:                  "Description",
						Type:                         "Type",
						Status:                       "DEPOSIT_PERIOD",
						ProposerAddress:              "ProposerAddress",
						MaybeProposerOperatorAddress: primptr.String("ProposerOperatorAddress"),
						Data: []model.MsgSubmitParamChangeProposalChange{
							{
								Subspace: "Subspace",
								Key:      "Key",
								Value:    json.RawMessage([]byte("Value")),
							},
						},
						InitialDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
						TotalDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
						TotalVote:                 big.NewInt(0),
						TransactionHash:           "TxHash",
						SubmitBlockHeight:         1,
						SubmitTime:                utctime.UTCTime{},
						DepositEndTime:            utctime.UTCTime{}.Add(time.Duration(1)),
						MaybeVotingStartTime:      nil,
						MaybeVotingEndTime:        nil,
						MaybeVotingEndBlockHeight: nil,
					}).
					Return(nil)

				proposal.NewProposals = func(
					_ *rdb.Handle,
				) view.Proposals {
					return mockProposalsView
				}

				mockDepositorsView := &view.MockDepositorsView{}
				mocks = append(mocks, &mockDepositorsView.Mock)
				mockDepositorsView.
					On("FindByProposalIdAndTxHash", "MaybeProposalId", "TxHash").
					Return(nil, nil)

				mockDepositorsView.
					On("Insert", &view.DepositorRow{
						ProposalId:                    "MaybeProposalId",
						DepositorAddress:              "ProposerAddress",
						MaybeDepositorOperatorAddress: primptr.String("ProposerOperatorAddress"),
						TransactionHash:               "TxHash",
						DepositAtBlockHeight:          1,
						DepositAtBlockTime:            utctime.UTCTime{},
						Amount: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
					}).
					Return(nil)

				proposal.NewDepositors = func(
					_ *rdb.Handle,
				) view.Depositors {
					return mockDepositorsView
				}

				mockDepositorsTotalView := &view.MockDepositorsTotalView{}
				mocks = append(mocks, &mockDepositorsTotalView.Mock)
				mockDepositorsTotalView.
					On("Increment", "MaybeProposalId", int64(1)).
					Return(nil)

				proposal.NewDepositorsTotal = func(handle *rdb.Handle) view.DepositorsTotal {
					return mockDepositorsTotalView
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgSubmitCommunityPoolSpendProposal",
			Events: []entity_event.Event{
				&usecase_event.MsgSubmitCommunityPoolSpendProposal{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					MsgSubmitCommunityPoolSpendProposalParams: model.MsgSubmitCommunityPoolSpendProposalParams{
						MaybeProposalId: primptr.String("MaybeProposalId"),
						Content: model.MsgSubmitCommunityPoolSpendProposalContent{
							Title:            "Title",
							Description:      "Description",
							Type:             "Type",
							RecipientAddress: "RecipientAddress",
							Amount: []coin.Coin{
								coin.NewInt64Coin("DENOM", 1000),
							},
						},
						ProposerAddress: "ProposerAddress",
						InitialDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {

				mockValidatorsView := &rdbvalidatorbase_view.MockValidatorsView{}
				mocks = append(mocks, &mockValidatorsView.Mock)
				mockValidatorsView.
					On("FindLastBy", rdbvalidatorbase_view.ValidatorIdentity{
						MaybeInititalDelegatorAddress: primptr.String("ProposerAddress"),
					}).
					Return(&rdbvalidatorbase_view.ValidatorRow{
						OperatorAddress: "ProposerOperatorAddress",
					}, nil)

				proposal.ValidatorBaseGetView = func(
					_ *rdbvalidatorbase.Base,
					_ *rdb.Handle,
				) rdbvalidatorbase_view.Validators {
					return mockValidatorsView
				}

				mockParamsView := &rdbparambase_view.MockParamsView{}
				mocks = append(mocks, &mockParamsView.Mock)
				mockParamsView.
					On("FindDurationBy", rdbparambase_types.ParamAccessor{
						Module: "gov",
						Key:    "max_deposit_period",
					}).
					Return(time.Duration(1), nil)

				proposal.ParamBaseGetView = func(
					_ *rdbparambase.Base,
					_ *rdb.Handle,
				) rdbparambase_view.Params {
					return mockParamsView
				}

				mockProposalsView := &view.MockProposalsView{}
				mocks = append(mocks, &mockProposalsView.Mock)
				mockProposalsView.
					On("Insert", &view.ProposalRow{
						ProposalId:                   "MaybeProposalId",
						Title:                        "Title",
						Description:                  "Description",
						Type:                         "Type",
						Status:                       "DEPOSIT_PERIOD",
						ProposerAddress:              "ProposerAddress",
						MaybeProposerOperatorAddress: primptr.String("ProposerOperatorAddress"),
						Data: types.CommunityPoolSpendData{
							RecipientAddress: "RecipientAddress",
							Amount: []coin.Coin{
								coin.NewInt64Coin("DENOM", 1000),
							},
						},
						InitialDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
						TotalDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
						TotalVote:                 big.NewInt(0),
						TransactionHash:           "TxHash",
						SubmitBlockHeight:         1,
						SubmitTime:                utctime.UTCTime{},
						DepositEndTime:            utctime.UTCTime{}.Add(time.Duration(1)),
						MaybeVotingStartTime:      nil,
						MaybeVotingEndTime:        nil,
						MaybeVotingEndBlockHeight: nil,
					}).
					Return(nil)

				proposal.NewProposals = func(
					_ *rdb.Handle,
				) view.Proposals {
					return mockProposalsView
				}

				mockDepositorsView := &view.MockDepositorsView{}
				mocks = append(mocks, &mockDepositorsView.Mock)
				mockDepositorsView.
					On("FindByProposalIdAndTxHash", "MaybeProposalId", "TxHash").
					Return(nil, nil)

				mockDepositorsView.
					On("Insert", &view.DepositorRow{
						ProposalId:                    "MaybeProposalId",
						DepositorAddress:              "ProposerAddress",
						MaybeDepositorOperatorAddress: primptr.String("ProposerOperatorAddress"),
						TransactionHash:               "TxHash",
						DepositAtBlockHeight:          1,
						DepositAtBlockTime:            utctime.UTCTime{},
						Amount: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
					}).
					Return(nil)

				proposal.NewDepositors = func(
					_ *rdb.Handle,
				) view.Depositors {
					return mockDepositorsView
				}

				mockDepositorsTotalView := &view.MockDepositorsTotalView{}
				mocks = append(mocks, &mockDepositorsTotalView.Mock)
				mockDepositorsTotalView.
					On("Increment", "MaybeProposalId", int64(1)).
					Return(nil)

				proposal.NewDepositorsTotal = func(handle *rdb.Handle) view.DepositorsTotal {
					return mockDepositorsTotalView
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgSubmitSoftwareUpgradeProposal",
			Events: []entity_event.Event{
				&usecase_event.MsgSubmitSoftwareUpgradeProposal{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					MsgSubmitSoftwareUpgradeProposalParams: model.MsgSubmitSoftwareUpgradeProposalParams{
						MaybeProposalId: primptr.String("MaybeProposalId"),
						Content: model.MsgSubmitSoftwareUpgradeProposalContent{
							Title:       "Title",
							Description: "Description",
							Type:        "Type",
							Plan: model.MsgSubmitSoftwareUpgradeProposalPlan{
								Name:   "Name",
								Time:   utctime.UTCTime{},
								Height: 1,
								Info:   "Info",
							},
						},
						ProposerAddress: "ProposerAddress",
						InitialDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {

				mockValidatorsView := &rdbvalidatorbase_view.MockValidatorsView{}
				mocks = append(mocks, &mockValidatorsView.Mock)
				mockValidatorsView.
					On("FindLastBy", rdbvalidatorbase_view.ValidatorIdentity{
						MaybeInititalDelegatorAddress: primptr.String("ProposerAddress"),
					}).
					Return(&rdbvalidatorbase_view.ValidatorRow{
						OperatorAddress: "ProposerOperatorAddress",
					}, nil)

				proposal.ValidatorBaseGetView = func(
					_ *rdbvalidatorbase.Base,
					_ *rdb.Handle,
				) rdbvalidatorbase_view.Validators {
					return mockValidatorsView
				}

				mockParamsView := &rdbparambase_view.MockParamsView{}
				mocks = append(mocks, &mockParamsView.Mock)
				mockParamsView.
					On("FindDurationBy", rdbparambase_types.ParamAccessor{
						Module: "gov",
						Key:    "max_deposit_period",
					}).
					Return(time.Duration(1), nil)

				proposal.ParamBaseGetView = func(
					_ *rdbparambase.Base,
					_ *rdb.Handle,
				) rdbparambase_view.Params {
					return mockParamsView
				}

				mockProposalsView := &view.MockProposalsView{}
				mocks = append(mocks, &mockProposalsView.Mock)
				mockProposalsView.
					On("Insert", &view.ProposalRow{
						ProposalId:                   "MaybeProposalId",
						Title:                        "Title",
						Description:                  "Description",
						Type:                         "Type",
						Status:                       "DEPOSIT_PERIOD",
						ProposerAddress:              "ProposerAddress",
						MaybeProposerOperatorAddress: primptr.String("ProposerOperatorAddress"),
						Data: model.MsgSubmitSoftwareUpgradeProposalPlan{
							Name:   "Name",
							Time:   utctime.UTCTime{},
							Height: 1,
							Info:   "Info",
						},
						InitialDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
						TotalDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
						TotalVote:                 big.NewInt(0),
						TransactionHash:           "TxHash",
						SubmitBlockHeight:         1,
						SubmitTime:                utctime.UTCTime{},
						DepositEndTime:            utctime.UTCTime{}.Add(time.Duration(1)),
						MaybeVotingStartTime:      nil,
						MaybeVotingEndTime:        nil,
						MaybeVotingEndBlockHeight: nil,
					}).
					Return(nil)

				proposal.NewProposals = func(
					_ *rdb.Handle,
				) view.Proposals {
					return mockProposalsView
				}

				mockDepositorsView := &view.MockDepositorsView{}
				mocks = append(mocks, &mockDepositorsView.Mock)
				mockDepositorsView.
					On("FindByProposalIdAndTxHash", "MaybeProposalId", "TxHash").
					Return(nil, nil)

				mockDepositorsView.
					On("Insert", &view.DepositorRow{
						ProposalId:                    "MaybeProposalId",
						DepositorAddress:              "ProposerAddress",
						MaybeDepositorOperatorAddress: primptr.String("ProposerOperatorAddress"),
						TransactionHash:               "TxHash",
						DepositAtBlockHeight:          1,
						DepositAtBlockTime:            utctime.UTCTime{},
						Amount: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
					}).
					Return(nil)

				proposal.NewDepositors = func(
					_ *rdb.Handle,
				) view.Depositors {
					return mockDepositorsView
				}

				mockDepositorsTotalView := &view.MockDepositorsTotalView{}
				mocks = append(mocks, &mockDepositorsTotalView.Mock)
				mockDepositorsTotalView.
					On("Increment", "MaybeProposalId", int64(1)).
					Return(nil)

				proposal.NewDepositorsTotal = func(handle *rdb.Handle) view.DepositorsTotal {
					return mockDepositorsTotalView
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgSubmitCancelSoftwareUpgradeProposal",
			Events: []entity_event.Event{
				&usecase_event.MsgSubmitCancelSoftwareUpgradeProposal{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					MsgSubmitCancelSoftwareUpgradeProposalParams: model.MsgSubmitCancelSoftwareUpgradeProposalParams{
						MaybeProposalId: primptr.String("MaybeProposalId"),
						Content: model.MsgSubmitCancelSoftwareUpgradeProposalContent{
							Title:       "Title",
							Description: "Description",
							Type:        "Type",
						},
						ProposerAddress: "ProposerAddress",
						InitialDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {

				mockValidatorsView := &rdbvalidatorbase_view.MockValidatorsView{}
				mocks = append(mocks, &mockValidatorsView.Mock)
				mockValidatorsView.
					On("FindLastBy", rdbvalidatorbase_view.ValidatorIdentity{
						MaybeInititalDelegatorAddress: primptr.String("ProposerAddress"),
					}).
					Return(&rdbvalidatorbase_view.ValidatorRow{
						OperatorAddress: "ProposerOperatorAddress",
					}, nil)

				proposal.ValidatorBaseGetView = func(
					_ *rdbvalidatorbase.Base,
					_ *rdb.Handle,
				) rdbvalidatorbase_view.Validators {
					return mockValidatorsView
				}

				mockParamsView := &rdbparambase_view.MockParamsView{}
				mocks = append(mocks, &mockParamsView.Mock)
				mockParamsView.
					On("FindDurationBy", rdbparambase_types.ParamAccessor{
						Module: "gov",
						Key:    "max_deposit_period",
					}).
					Return(time.Duration(1), nil)

				proposal.ParamBaseGetView = func(
					_ *rdbparambase.Base,
					_ *rdb.Handle,
				) rdbparambase_view.Params {
					return mockParamsView
				}

				mockProposalsView := &view.MockProposalsView{}
				mocks = append(mocks, &mockProposalsView.Mock)
				mockProposalsView.
					On("Insert", &view.ProposalRow{
						ProposalId:                   "MaybeProposalId",
						Title:                        "Title",
						Description:                  "Description",
						Type:                         "Type",
						Status:                       "DEPOSIT_PERIOD",
						ProposerAddress:              "ProposerAddress",
						MaybeProposerOperatorAddress: primptr.String("ProposerOperatorAddress"),
						Data:                         nil,
						InitialDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
						TotalDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
						TotalVote:                 big.NewInt(0),
						TransactionHash:           "TxHash",
						SubmitBlockHeight:         1,
						SubmitTime:                utctime.UTCTime{},
						DepositEndTime:            utctime.UTCTime{}.Add(time.Duration(1)),
						MaybeVotingStartTime:      nil,
						MaybeVotingEndTime:        nil,
						MaybeVotingEndBlockHeight: nil,
					}).
					Return(nil)

				proposal.NewProposals = func(
					_ *rdb.Handle,
				) view.Proposals {
					return mockProposalsView
				}

				mockDepositorsView := &view.MockDepositorsView{}
				mocks = append(mocks, &mockDepositorsView.Mock)
				mockDepositorsView.
					On("FindByProposalIdAndTxHash", "MaybeProposalId", "TxHash").
					Return(nil, nil)

				mockDepositorsView.
					On("Insert", &view.DepositorRow{
						ProposalId:                    "MaybeProposalId",
						DepositorAddress:              "ProposerAddress",
						MaybeDepositorOperatorAddress: primptr.String("ProposerOperatorAddress"),
						TransactionHash:               "TxHash",
						DepositAtBlockHeight:          1,
						DepositAtBlockTime:            utctime.UTCTime{},
						Amount: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
					}).
					Return(nil)

				proposal.NewDepositors = func(
					_ *rdb.Handle,
				) view.Depositors {
					return mockDepositorsView
				}

				mockDepositorsTotalView := &view.MockDepositorsTotalView{}
				mocks = append(mocks, &mockDepositorsTotalView.Mock)
				mockDepositorsTotalView.
					On("Increment", "MaybeProposalId", int64(1)).
					Return(nil)

				proposal.NewDepositorsTotal = func(handle *rdb.Handle) view.DepositorsTotal {
					return mockDepositorsTotalView
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgSoftwareUpgrade",
			Events: []entity_event.Event{
				&usecase_event.MsgSoftwareUpgrade{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_SOFTWARE_UPGRADE,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					MsgSoftwareUpgradeParams: v1_model.MsgSoftwareUpgradeParams{
						MaybeProposalId: primptr.String("MaybeProposalId"),
						Authority:       "Authority",
						Proposer:        "Proposer",

						Plan: v1_model.MsgSoftwareUpgradePlan{
							Name:   "Name",
							Time:   utctime.FromUnixNano(int64(1000)),
							Height: int64(1000),
							Info:   "Info",
						},
						InitialDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
						Metadata: "Metadata",
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {

				mockValidatorsView := &rdbvalidatorbase_view.MockValidatorsView{}
				mocks = append(mocks, &mockValidatorsView.Mock)
				mockValidatorsView.
					On("FindLastBy", rdbvalidatorbase_view.ValidatorIdentity{
						MaybeInititalDelegatorAddress: primptr.String("Proposer"),
					}).
					Return(&rdbvalidatorbase_view.ValidatorRow{
						OperatorAddress: "ProposerOperatorAddress",
					}, nil)

				proposal.ValidatorBaseGetView = func(
					_ *rdbvalidatorbase.Base,
					_ *rdb.Handle,
				) rdbvalidatorbase_view.Validators {
					return mockValidatorsView
				}

				mockParamsView := &rdbparambase_view.MockParamsView{}
				mocks = append(mocks, &mockParamsView.Mock)
				mockParamsView.
					On("FindDurationBy", rdbparambase_types.ParamAccessor{
						Module: "gov",
						Key:    "max_deposit_period",
					}).
					Return(time.Duration(1), nil)

				proposal.ParamBaseGetView = func(
					_ *rdbparambase.Base,
					_ *rdb.Handle,
				) rdbparambase_view.Params {
					return mockParamsView
				}

				mockProposalsView := &view.MockProposalsView{}
				mocks = append(mocks, &mockProposalsView.Mock)
				mockProposalsView.
					On("Insert", &view.ProposalRow{
						ProposalId:                   "MaybeProposalId",
						Title:                        "",
						Description:                  "",
						Type:                         "/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade",
						Status:                       "DEPOSIT_PERIOD",
						ProposerAddress:              "Proposer",
						MaybeProposerOperatorAddress: primptr.String("ProposerOperatorAddress"),
						Data: types.MsgSoftwareUpgradeData{
							Type:      "/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade",
							Authority: "Authority",
							Plan: v1_model.MsgSoftwareUpgradePlan{
								Name:   "Name",
								Time:   utctime.FromUnixNano(int64(1000)),
								Height: 1000,
								Info:   "Info",
							},
						},
						InitialDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
						TotalDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
						TotalVote:                 big.NewInt(0),
						TransactionHash:           "TxHash",
						SubmitBlockHeight:         1,
						SubmitTime:                utctime.UTCTime{},
						DepositEndTime:            utctime.UTCTime{}.Add(time.Duration(1)),
						MaybeVotingStartTime:      nil,
						MaybeVotingEndTime:        nil,
						MaybeVotingEndBlockHeight: nil,
						Metadata:                  "Metadata",
					}).
					Return(nil)

				proposal.NewProposals = func(
					_ *rdb.Handle,
				) view.Proposals {
					return mockProposalsView
				}

				mockDepositorsView := &view.MockDepositorsView{}
				mocks = append(mocks, &mockDepositorsView.Mock)
				mockDepositorsView.
					On("FindByProposalIdAndTxHash", "MaybeProposalId", "TxHash").
					Return(nil, nil)

				mockDepositorsView.
					On("Insert", &view.DepositorRow{
						ProposalId:                    "MaybeProposalId",
						DepositorAddress:              "Proposer",
						MaybeDepositorOperatorAddress: primptr.String("ProposerOperatorAddress"),
						TransactionHash:               "TxHash",
						DepositAtBlockHeight:          1,
						DepositAtBlockTime:            utctime.UTCTime{},
						Amount: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
					}).
					Return(nil)

				proposal.NewDepositors = func(
					_ *rdb.Handle,
				) view.Depositors {
					return mockDepositorsView
				}

				mockDepositorsTotalView := &view.MockDepositorsTotalView{}
				mocks = append(mocks, &mockDepositorsTotalView.Mock)
				mockDepositorsTotalView.
					On("Increment", "MaybeProposalId", int64(1)).
					Return(nil)

				proposal.NewDepositorsTotal = func(handle *rdb.Handle) view.DepositorsTotal {
					return mockDepositorsTotalView
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgCancelUpgrade",
			Events: []entity_event.Event{
				&usecase_event.MsgCancelUpgrade{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_CANCEL_UPGRADE,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					MsgCancelUpgradeParams: v1_model.MsgCancelUpgradeParams{
						MaybeProposalId: primptr.String("MaybeProposalId"),
						Authority:       "Authority",
						Proposer:        "Proposer",
						InitialDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
						Metadata: "Metadata",
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {

				mockValidatorsView := &rdbvalidatorbase_view.MockValidatorsView{}
				mocks = append(mocks, &mockValidatorsView.Mock)
				mockValidatorsView.
					On("FindLastBy", rdbvalidatorbase_view.ValidatorIdentity{
						MaybeInititalDelegatorAddress: primptr.String("Proposer"),
					}).
					Return(&rdbvalidatorbase_view.ValidatorRow{
						OperatorAddress: "ProposerOperatorAddress",
					}, nil)

				proposal.ValidatorBaseGetView = func(
					_ *rdbvalidatorbase.Base,
					_ *rdb.Handle,
				) rdbvalidatorbase_view.Validators {
					return mockValidatorsView
				}

				mockParamsView := &rdbparambase_view.MockParamsView{}
				mocks = append(mocks, &mockParamsView.Mock)
				mockParamsView.
					On("FindDurationBy", rdbparambase_types.ParamAccessor{
						Module: "gov",
						Key:    "max_deposit_period",
					}).
					Return(time.Duration(1), nil)

				proposal.ParamBaseGetView = func(
					_ *rdbparambase.Base,
					_ *rdb.Handle,
				) rdbparambase_view.Params {
					return mockParamsView
				}

				mockProposalsView := &view.MockProposalsView{}
				mocks = append(mocks, &mockProposalsView.Mock)
				mockProposalsView.
					On("Insert", &view.ProposalRow{
						ProposalId:                   "MaybeProposalId",
						Title:                        "",
						Description:                  "",
						Type:                         "/cosmos.upgrade.v1beta1.MsgCancelUpgrade",
						Status:                       "DEPOSIT_PERIOD",
						ProposerAddress:              "Proposer",
						MaybeProposerOperatorAddress: primptr.String("ProposerOperatorAddress"),
						Data: types.MsgCancelUpgradeData{
							Type:      "/cosmos.upgrade.v1beta1.MsgCancelUpgrade",
							Authority: "Authority",
						},
						InitialDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
						TotalDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
						TotalVote:                 big.NewInt(0),
						TransactionHash:           "TxHash",
						SubmitBlockHeight:         1,
						SubmitTime:                utctime.UTCTime{},
						DepositEndTime:            utctime.UTCTime{}.Add(time.Duration(1)),
						MaybeVotingStartTime:      nil,
						MaybeVotingEndTime:        nil,
						MaybeVotingEndBlockHeight: nil,
						Metadata:                  "Metadata",
					}).
					Return(nil)

				proposal.NewProposals = func(
					_ *rdb.Handle,
				) view.Proposals {
					return mockProposalsView
				}

				mockDepositorsView := &view.MockDepositorsView{}
				mocks = append(mocks, &mockDepositorsView.Mock)
				mockDepositorsView.
					On("FindByProposalIdAndTxHash", "MaybeProposalId", "TxHash").
					Return(nil, nil)

				mockDepositorsView.
					On("Insert", &view.DepositorRow{
						ProposalId:                    "MaybeProposalId",
						DepositorAddress:              "Proposer",
						MaybeDepositorOperatorAddress: primptr.String("ProposerOperatorAddress"),
						TransactionHash:               "TxHash",
						DepositAtBlockHeight:          1,
						DepositAtBlockTime:            utctime.UTCTime{},
						Amount: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
					}).
					Return(nil)

				proposal.NewDepositors = func(
					_ *rdb.Handle,
				) view.Depositors {
					return mockDepositorsView
				}

				mockDepositorsTotalView := &view.MockDepositorsTotalView{}
				mocks = append(mocks, &mockDepositorsTotalView.Mock)
				mockDepositorsTotalView.
					On("Increment", "MaybeProposalId", int64(1)).
					Return(nil)

				proposal.NewDepositorsTotal = func(handle *rdb.Handle) view.DepositorsTotal {
					return mockDepositorsTotalView
				}

				return mocks
			},
		},
		{
			Name: "HandleProposalVotingPeriodStarted",
			Events: []entity_event.Event{
				&usecase_event.ProposalVotingPeriodStarted{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        usecase_event.PROPOSAL_VOTING_PERIOD_STARTED,
						Version:     1,
						BlockHeight: 1,
					}),
					ProposalId: "ProposalId",
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {

				mockProposalsView := &view.MockProposalsView{}
				mocks = append(mocks, &mockProposalsView.Mock)
				mockProposalsView.
					On("FindById", "ProposalId").
					Return([]view.ProposalWithMonikerRow{
						{
							ProposalRow: view.ProposalRow{
								ProposalId:           "ProposalId",
								Status:               "Status",
								MaybeVotingStartTime: nil,
								MaybeVotingEndTime:   nil,
							},
						},
					}, nil)

				proposal.NewProposals = func(
					_ *rdb.Handle,
				) view.Proposals {
					return mockProposalsView
				}

				mockParamsView := &rdbparambase_view.MockParamsView{}
				mocks = append(mocks, &mockParamsView.Mock)
				mockParamsView.
					On("FindDurationBy", rdbparambase_types.ParamAccessor{
						Module: "gov",
						Key:    "voting_period",
					}).
					Return(time.Duration(1), nil)

				proposal.ParamBaseGetView = func(
					_ *rdbparambase.Base,
					_ *rdb.Handle,
				) rdbparambase_view.Params {
					return mockParamsView
				}

				maybeVotingEndTime := utctime.UTCTime{}.Add(time.Duration(1))
				mockProposalsView.
					On("Update", &view.ProposalRow{
						ProposalId:           "ProposalId",
						Status:               "VOTING_PERIOD",
						MaybeVotingStartTime: &utctime.UTCTime{},
						MaybeVotingEndTime:   &maybeVotingEndTime,
					}).
					Return(nil)

				return mocks
			},
		},
		{
			Name: "HandleProposalInactived",
			Events: []entity_event.Event{
				&usecase_event.ProposalInactived{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        usecase_event.PROPOSAL_INACTIVED,
						Version:     1,
						BlockHeight: 1,
					}),
					ProposalId: "ProposalId",
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {

				mockProposalsView := &view.MockProposalsView{}
				mocks = append(mocks, &mockProposalsView.Mock)
				mockProposalsView.
					On("FindById", "ProposalId").
					Return([]view.ProposalWithMonikerRow{
						{
							ProposalRow: view.ProposalRow{
								ProposalId: "ProposalId",
								Status:     "Status",
							},
						},
					}, nil)
				mockProposalsView.
					On("Update", &view.ProposalRow{
						ProposalId: "ProposalId",
						Status:     "INACTIVE",
					}).
					Return(nil)

				proposal.NewProposals = func(
					_ *rdb.Handle,
				) view.Proposals {
					return mockProposalsView
				}

				return mocks
			},
		},
		{
			Name: "HandleProposalEnded proposal_passed",
			Events: []entity_event.Event{
				&usecase_event.ProposalEnded{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        usecase_event.PROPOSAL_ENDED,
						Version:     1,
						BlockHeight: 1,
					}),
					ProposalId: "ProposalId",
					Result:     "proposal_passed",
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {

				mockProposalsView := &view.MockProposalsView{}
				mocks = append(mocks, &mockProposalsView.Mock)
				mockProposalsView.
					On("FindById", "ProposalId").
					Return([]view.ProposalWithMonikerRow{
						{
							ProposalRow: view.ProposalRow{
								ProposalId: "ProposalId",
								Status:     "Status",
							},
						},
					}, nil)
				mockProposalsView.
					On("Update", &view.ProposalRow{
						ProposalId:                "ProposalId",
						Status:                    "PASSED",
						MaybeVotingEndBlockHeight: primptr.Int64(1),
					}).
					Return(nil)

				proposal.NewProposals = func(
					_ *rdb.Handle,
				) view.Proposals {
					return mockProposalsView
				}

				return mocks
			},
		},
		{
			Name: "HandleProposalEnded proposal_failed",
			Events: []entity_event.Event{
				&usecase_event.ProposalEnded{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        usecase_event.PROPOSAL_ENDED,
						Version:     1,
						BlockHeight: 1,
					}),
					ProposalId: "ProposalId",
					Result:     "proposal_failed",
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {

				mockProposalsView := &view.MockProposalsView{}
				mocks = append(mocks, &mockProposalsView.Mock)
				mockProposalsView.
					On("FindById", "ProposalId").
					Return([]view.ProposalWithMonikerRow{
						{
							ProposalRow: view.ProposalRow{
								ProposalId: "ProposalId",
								Status:     "Status",
							},
						},
					}, nil)
				mockProposalsView.
					On("Update", &view.ProposalRow{
						ProposalId:                "ProposalId",
						Status:                    "FAILED",
						MaybeVotingEndBlockHeight: primptr.Int64(1),
					}).
					Return(nil)

				proposal.NewProposals = func(
					_ *rdb.Handle,
				) view.Proposals {
					return mockProposalsView
				}

				return mocks
			},
		},
		{
			Name: "HandleProposalEnded proposal_rejected",
			Events: []entity_event.Event{
				&usecase_event.ProposalEnded{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        usecase_event.PROPOSAL_ENDED,
						Version:     1,
						BlockHeight: 1,
					}),
					ProposalId: "ProposalId",
					Result:     "proposal_rejected",
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {

				mockProposalsView := &view.MockProposalsView{}
				mocks = append(mocks, &mockProposalsView.Mock)
				mockProposalsView.
					On("FindById", "ProposalId").
					Return([]view.ProposalWithMonikerRow{
						{
							ProposalRow: view.ProposalRow{
								ProposalId: "ProposalId",
								Status:     "Status",
							},
						},
					}, nil)
				mockProposalsView.
					On("Update", &view.ProposalRow{
						ProposalId:                "ProposalId",
						Status:                    "REJECTED",
						MaybeVotingEndBlockHeight: primptr.Int64(1),
					}).
					Return(nil)

				proposal.NewProposals = func(
					_ *rdb.Handle,
				) view.Proposals {
					return mockProposalsView
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgDeposit",
			Events: []entity_event.Event{
				&usecase_event.MsgDeposit{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_DEPOSIT,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					ProposalId: "ProposalId",
					Depositor:  "Depositor",
					Amount: []coin.Coin{
						coin.NewInt64Coin("DENOM", 100),
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {

				mockProposalsView := &view.MockProposalsView{}
				mocks = append(mocks, &mockProposalsView.Mock)
				mockProposalsView.
					On("FindById", "ProposalId").
					Return([]view.ProposalWithMonikerRow{
						{

							ProposalRow: view.ProposalRow{
								ProposalId: "ProposalId",
								TotalDeposit: []coin.Coin{
									coin.NewInt64Coin("DENOM", 200),
								},
							},
						},
					}, nil)
				mockProposalsView.
					On("Update", &view.ProposalRow{
						ProposalId: "ProposalId",
						TotalDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 300),
						},
					}).
					Return(nil)

				proposal.NewProposals = func(
					_ *rdb.Handle,
				) view.Proposals {
					return mockProposalsView
				}

				mockValidatorsView := &rdbvalidatorbase_view.MockValidatorsView{}
				mocks = append(mocks, &mockValidatorsView.Mock)
				mockValidatorsView.
					On("FindLastBy", rdbvalidatorbase_view.ValidatorIdentity{
						MaybeInititalDelegatorAddress: primptr.String("Depositor"),
					}).
					Return(&rdbvalidatorbase_view.ValidatorRow{
						OperatorAddress: "DepositorOperatorAddress",
					}, nil)

				proposal.ValidatorBaseGetView = func(
					_ *rdbvalidatorbase.Base,
					_ *rdb.Handle,
				) rdbvalidatorbase_view.Validators {
					return mockValidatorsView
				}

				mockDepositorsView := &view.MockDepositorsView{}
				mocks = append(mocks, &mockDepositorsView.Mock)
				mockDepositorsView.
					On("Insert", &view.DepositorRow{
						ProposalId:                    "ProposalId",
						DepositorAddress:              "Depositor",
						MaybeDepositorOperatorAddress: primptr.String("DepositorOperatorAddress"),
						TransactionHash:               "TxHash",
						DepositAtBlockHeight:          1,
						DepositAtBlockTime:            utctime.UTCTime{},
						Amount: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
					}).
					Return(nil)

				proposal.NewDepositors = func(
					_ *rdb.Handle,
				) view.Depositors {
					return mockDepositorsView
				}

				mockDepositorsTotalView := &view.MockDepositorsTotalView{}
				mocks = append(mocks, &mockDepositorsTotalView.Mock)
				mockDepositorsTotalView.
					On("Increment", "ProposalId", int64(1)).
					Return(nil)

				proposal.NewDepositorsTotal = func(handle *rdb.Handle) view.DepositorsTotal {
					return mockDepositorsTotalView
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgDepositV1",
			Events: []entity_event.Event{
				&usecase_event.MsgDepositV1{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_DEPOSIT_V1,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					ProposalId: "ProposalId",
					Depositor:  "Depositor",
					Amount: []coin.Coin{
						coin.NewInt64Coin("DENOM", 100),
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {

				mockProposalsView := &view.MockProposalsView{}
				mocks = append(mocks, &mockProposalsView.Mock)
				mockProposalsView.
					On("FindById", "ProposalId").
					Return([]view.ProposalWithMonikerRow{
						{

							ProposalRow: view.ProposalRow{
								ProposalId: "ProposalId",
								TotalDeposit: []coin.Coin{
									coin.NewInt64Coin("DENOM", 200),
								},
							},
						},
					}, nil)
				mockProposalsView.
					On("Update", &view.ProposalRow{
						ProposalId: "ProposalId",
						TotalDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 300),
						},
					}).
					Return(nil)

				proposal.NewProposals = func(
					_ *rdb.Handle,
				) view.Proposals {
					return mockProposalsView
				}

				mockValidatorsView := &rdbvalidatorbase_view.MockValidatorsView{}
				mocks = append(mocks, &mockValidatorsView.Mock)
				mockValidatorsView.
					On("FindLastBy", rdbvalidatorbase_view.ValidatorIdentity{
						MaybeInititalDelegatorAddress: primptr.String("Depositor"),
					}).
					Return(&rdbvalidatorbase_view.ValidatorRow{
						OperatorAddress: "DepositorOperatorAddress",
					}, nil)

				proposal.ValidatorBaseGetView = func(
					_ *rdbvalidatorbase.Base,
					_ *rdb.Handle,
				) rdbvalidatorbase_view.Validators {
					return mockValidatorsView
				}

				mockDepositorsView := &view.MockDepositorsView{}
				mocks = append(mocks, &mockDepositorsView.Mock)
				mockDepositorsView.
					On("Insert", &view.DepositorRow{
						ProposalId:                    "ProposalId",
						DepositorAddress:              "Depositor",
						MaybeDepositorOperatorAddress: primptr.String("DepositorOperatorAddress"),
						TransactionHash:               "TxHash",
						DepositAtBlockHeight:          1,
						DepositAtBlockTime:            utctime.UTCTime{},
						Amount: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
					}).
					Return(nil)

				proposal.NewDepositors = func(
					_ *rdb.Handle,
				) view.Depositors {
					return mockDepositorsView
				}

				mockDepositorsTotalView := &view.MockDepositorsTotalView{}
				mocks = append(mocks, &mockDepositorsTotalView.Mock)
				mockDepositorsTotalView.
					On("Increment", "ProposalId", int64(1)).
					Return(nil)

				proposal.NewDepositorsTotal = func(handle *rdb.Handle) view.DepositorsTotal {
					return mockDepositorsTotalView
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgDepositV1 deposit to multi-msgs proposal",
			Events: []entity_event.Event{
				&usecase_event.MsgDepositV1{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_DEPOSIT_V1,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					ProposalId: "ProposalId",
					Depositor:  "Depositor",
					Amount: []coin.Coin{
						coin.NewInt64Coin("DENOM", 100),
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {

				mockProposalsView := &view.MockProposalsView{}
				mocks = append(mocks, &mockProposalsView.Mock)
				mockProposalsView.
					On("FindById", "ProposalId").
					Return([]view.ProposalWithMonikerRow{
						{

							ProposalRow: view.ProposalRow{
								MaybeId:    primptr.Int64(1),
								ProposalId: "ProposalId",
								Type:       "Type1",
								TotalDeposit: []coin.Coin{
									coin.NewInt64Coin("DENOM", 200),
								},
							},
						},
						{

							ProposalRow: view.ProposalRow{
								MaybeId:    primptr.Int64(2),
								ProposalId: "ProposalId",
								Type:       "Type2",
								TotalDeposit: []coin.Coin{
									coin.NewInt64Coin("DENOM", 200),
								},
							},
						},
					}, nil)
				mockProposalsView.
					On("Update", &view.ProposalRow{
						MaybeId:    primptr.Int64(1),
						ProposalId: "ProposalId",
						Type:       "Type1",
						TotalDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 300),
						},
					}).
					Return(nil)
				mockProposalsView.
					On("Update", &view.ProposalRow{
						MaybeId:    primptr.Int64(2),
						ProposalId: "ProposalId",
						Type:       "Type2",
						TotalDeposit: []coin.Coin{
							coin.NewInt64Coin("DENOM", 300),
						},
					}).
					Return(nil)

				proposal.NewProposals = func(
					_ *rdb.Handle,
				) view.Proposals {
					return mockProposalsView
				}

				mockValidatorsView := &rdbvalidatorbase_view.MockValidatorsView{}
				mocks = append(mocks, &mockValidatorsView.Mock)
				mockValidatorsView.
					On("FindLastBy", rdbvalidatorbase_view.ValidatorIdentity{
						MaybeInititalDelegatorAddress: primptr.String("Depositor"),
					}).
					Return(&rdbvalidatorbase_view.ValidatorRow{
						OperatorAddress: "DepositorOperatorAddress",
					}, nil)

				proposal.ValidatorBaseGetView = func(
					_ *rdbvalidatorbase.Base,
					_ *rdb.Handle,
				) rdbvalidatorbase_view.Validators {
					return mockValidatorsView
				}

				mockDepositorsView := &view.MockDepositorsView{}
				mocks = append(mocks, &mockDepositorsView.Mock)
				mockDepositorsView.
					On("Insert", &view.DepositorRow{
						ProposalId:                    "ProposalId",
						DepositorAddress:              "Depositor",
						MaybeDepositorOperatorAddress: primptr.String("DepositorOperatorAddress"),
						TransactionHash:               "TxHash",
						DepositAtBlockHeight:          1,
						DepositAtBlockTime:            utctime.UTCTime{},
						Amount: []coin.Coin{
							coin.NewInt64Coin("DENOM", 100),
						},
					}).
					Return(nil)

				proposal.NewDepositors = func(
					_ *rdb.Handle,
				) view.Depositors {
					return mockDepositorsView
				}

				mockDepositorsTotalView := &view.MockDepositorsTotalView{}
				mocks = append(mocks, &mockDepositorsTotalView.Mock)
				mockDepositorsTotalView.
					On("Increment", "ProposalId", int64(1)).
					Return(nil)

				proposal.NewDepositorsTotal = func(handle *rdb.Handle) view.DepositorsTotal {
					return mockDepositorsTotalView
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgVote voter record does not exist",
			Events: []entity_event.Event{
				&usecase_event.MsgVote{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_VOTE,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
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

				mockValidatorsView := &rdbvalidatorbase_view.MockValidatorsView{}
				mocks = append(mocks, &mockValidatorsView.Mock)
				mockValidatorsView.
					On("FindLastBy", rdbvalidatorbase_view.ValidatorIdentity{
						MaybeInititalDelegatorAddress: primptr.String("Voter"),
					}).
					Return(&rdbvalidatorbase_view.ValidatorRow{
						OperatorAddress: "VoterOperatorAddress",
					}, nil)

				proposal.ValidatorBaseGetView = func(
					_ *rdbvalidatorbase.Base,
					_ *rdb.Handle,
				) rdbvalidatorbase_view.Validators {
					return mockValidatorsView
				}

				mockVotesView := &view.MockVotesView{}
				mocks = append(mocks, &mockVotesView.Mock)
				mockVotesView.
					On("FindByProposalIdVoter", "ProposalId", "Voter").
					Return(&[]view.VoteWithMonikerRow{}, rdb.ErrNoRows)

				mockVotesView.
					On("Insert", &view.VoteRow{
						ProposalId:                "ProposalId",
						VoterAddress:              "Voter",
						MaybeVoterOperatorAddress: primptr.String("VoterOperatorAddress"),
						TransactionHash:           "TxHash",
						VoteAtBlockHeight:         1,
						VoteAtBlockTime:           utctime.UTCTime{},
						Answer:                    "Option",
						Histories:                 make([]view.VoteHistory, 0),
						Weight:                    "1.000000000000000000",
					}).
					Return(nil, rdb.ErrNoRows)

				proposal.NewVotes = func(
					_ *rdb.Handle,
				) view.Votes {
					return mockVotesView
				}

				mockProposalsView := &view.MockProposalsView{}
				mocks = append(mocks, &mockProposalsView.Mock)
				mockProposalsView.
					On("FindById", "ProposalId").
					Return([]view.ProposalWithMonikerRow{
						{
							ProposalRow: view.ProposalRow{
								MaybeId:    primptr.Int64(1),
								ProposalId: "ProposalId",
								TotalVote:  big.NewInt(1),
							},
						},
					}, nil)
				mockProposalsView.
					On("Update", &view.ProposalRow{
						MaybeId:    primptr.Int64(1),
						ProposalId: "ProposalId",
						TotalVote:  big.NewInt(2),
					}).
					Return(nil)

				proposal.NewProposals = func(
					_ *rdb.Handle,
				) view.Proposals {
					return mockProposalsView
				}

				mockVotesTotalView := &view.MockVotesTotalView{}
				mocks = append(mocks, &mockVotesTotalView.Mock)
				mockVotesTotalView.
					On("Increment", "ProposalId", int64(1)).
					Return(nil)

				proposal.NewVotesTotal = func(
					_ *rdb.Handle,
				) view.VotesTotal {
					return mockVotesTotalView
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgVote voter record exists",
			Events: []entity_event.Event{
				&usecase_event.MsgVote{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_VOTE,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
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

				mockValidatorsView := &rdbvalidatorbase_view.MockValidatorsView{}
				mocks = append(mocks, &mockValidatorsView.Mock)
				mockValidatorsView.
					On("FindLastBy", rdbvalidatorbase_view.ValidatorIdentity{
						MaybeInititalDelegatorAddress: primptr.String("Voter"),
					}).
					Return(&rdbvalidatorbase_view.ValidatorRow{
						OperatorAddress: "VoterOperatorAddress",
					}, nil)

				proposal.ValidatorBaseGetView = func(
					_ *rdbvalidatorbase.Base,
					_ *rdb.Handle,
				) rdbvalidatorbase_view.Validators {
					return mockValidatorsView
				}

				mockVotesView := &view.MockVotesView{}
				mocks = append(mocks, &mockVotesView.Mock)
				mockVotesView.
					On("FindByProposalIdVoter", "ProposalId", "Voter").
					Return([]view.VoteWithMonikerRow{
						{
							VoteRow: view.VoteRow{
								ProposalId:                "ProposalId",
								VoterAddress:              "Voter",
								TransactionHash:           "PreviousTransactionHash",
								MaybeVoterOperatorAddress: primptr.String("VoterOperatorAddress"),
								VoteAtBlockHeight:         0,
								VoteAtBlockTime:           utctime.FromUnixNano(-1),
								Answer:                    "PreviousAnswer",
								Histories:                 make([]view.VoteHistory, 0),
								Metadata:                  "PreviousMetadata",
								Weight:                    "1.000000000000000000",
							},
						},
					}, nil)

				mockVotesView.
					On("DeleteByProposalIdVoter", "ProposalId", "Voter").
					Return(int64(1), nil)

				mockVotesView.
					On("Insert", &view.VoteRow{
						ProposalId:                "ProposalId",
						VoterAddress:              "Voter",
						MaybeVoterOperatorAddress: primptr.String("VoterOperatorAddress"),
						TransactionHash:           "TxHash",
						VoteAtBlockHeight:         1,
						VoteAtBlockTime:           utctime.UTCTime{},
						Answer:                    "Option",
						Metadata:                  "",
						Weight:                    "1.000000000000000000",
						Histories: []view.VoteHistory{
							{
								TransactionHash:   "PreviousTransactionHash",
								VoteAtBlockHeight: 0,
								VoteAtBlockTime:   utctime.FromUnixNano(-1),
								Answer:            "PreviousAnswer",
								Metadata:          "PreviousMetadata",
								Weight:            "1.000000000000000000",
							},
						},
					}).
					Return(nil, rdb.ErrNoRows)

				proposal.NewVotes = func(
					_ *rdb.Handle,
				) view.Votes {
					return mockVotesView
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgVote vote for multi-msgs proposal",
			Events: []entity_event.Event{
				&usecase_event.MsgVote{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_VOTE,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
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

				mockValidatorsView := &rdbvalidatorbase_view.MockValidatorsView{}
				mocks = append(mocks, &mockValidatorsView.Mock)
				mockValidatorsView.
					On("FindLastBy", rdbvalidatorbase_view.ValidatorIdentity{
						MaybeInititalDelegatorAddress: primptr.String("Voter"),
					}).
					Return(&rdbvalidatorbase_view.ValidatorRow{
						OperatorAddress: "VoterOperatorAddress",
					}, nil)

				proposal.ValidatorBaseGetView = func(
					_ *rdbvalidatorbase.Base,
					_ *rdb.Handle,
				) rdbvalidatorbase_view.Validators {
					return mockValidatorsView
				}

				mockVotesView := &view.MockVotesView{}
				mocks = append(mocks, &mockVotesView.Mock)
				mockVotesView.
					On("FindByProposalIdVoter", "ProposalId", "Voter").
					Return(&[]view.VoteWithMonikerRow{}, rdb.ErrNoRows)

				mockVotesView.
					On("Insert", &view.VoteRow{
						ProposalId:                "ProposalId",
						VoterAddress:              "Voter",
						MaybeVoterOperatorAddress: primptr.String("VoterOperatorAddress"),
						TransactionHash:           "TxHash",
						VoteAtBlockHeight:         1,
						VoteAtBlockTime:           utctime.UTCTime{},
						Answer:                    "Option",
						Histories:                 make([]view.VoteHistory, 0),
						Weight:                    "1.000000000000000000",
					}).
					Return(nil, rdb.ErrNoRows)

				proposal.NewVotes = func(
					_ *rdb.Handle,
				) view.Votes {
					return mockVotesView
				}

				mockProposalsView := &view.MockProposalsView{}
				mocks = append(mocks, &mockProposalsView.Mock)
				mockProposalsView.
					On("FindById", "ProposalId").
					Return([]view.ProposalWithMonikerRow{
						{
							ProposalRow: view.ProposalRow{
								MaybeId:    primptr.Int64(1),
								ProposalId: "ProposalId",
								TotalVote:  big.NewInt(1),
							},
						},
						{
							ProposalRow: view.ProposalRow{
								MaybeId:    primptr.Int64(2),
								ProposalId: "ProposalId",
								TotalVote:  big.NewInt(1),
							},
						},
					}, nil)
				mockProposalsView.
					On("Update", &view.ProposalRow{
						MaybeId:    primptr.Int64(1),
						ProposalId: "ProposalId",
						TotalVote:  big.NewInt(2),
					}).
					Return(nil)
				mockProposalsView.
					On("Update", &view.ProposalRow{
						MaybeId:    primptr.Int64(2),
						ProposalId: "ProposalId",
						TotalVote:  big.NewInt(2),
					}).
					Return(nil)

				proposal.NewProposals = func(
					_ *rdb.Handle,
				) view.Proposals {
					return mockProposalsView
				}

				mockVotesTotalView := &view.MockVotesTotalView{}
				mocks = append(mocks, &mockVotesTotalView.Mock)
				mockVotesTotalView.
					On("Increment", "ProposalId", int64(1)).
					Return(nil)

				proposal.NewVotesTotal = func(
					_ *rdb.Handle,
				) view.VotesTotal {
					return mockVotesTotalView
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgVoteV1 voter record does not exist",
			Events: []entity_event.Event{
				&usecase_event.MsgVoteV1{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_VOTE_V1,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					ProposalId: "ProposalId",
					Voter:      "Voter",
					Option:     "Option",
					Metadata:   "Metadata",
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {

				mockValidatorsView := &rdbvalidatorbase_view.MockValidatorsView{}
				mocks = append(mocks, &mockValidatorsView.Mock)
				mockValidatorsView.
					On("FindLastBy", rdbvalidatorbase_view.ValidatorIdentity{
						MaybeInititalDelegatorAddress: primptr.String("Voter"),
					}).
					Return(&rdbvalidatorbase_view.ValidatorRow{
						OperatorAddress: "VoterOperatorAddress",
					}, nil)

				proposal.ValidatorBaseGetView = func(
					_ *rdbvalidatorbase.Base,
					_ *rdb.Handle,
				) rdbvalidatorbase_view.Validators {
					return mockValidatorsView
				}

				mockVotesView := &view.MockVotesView{}
				mocks = append(mocks, &mockVotesView.Mock)
				mockVotesView.
					On("FindByProposalIdVoter", "ProposalId", "Voter").
					Return(&[]view.VoteWithMonikerRow{}, rdb.ErrNoRows)

				mockVotesView.
					On("Insert", &view.VoteRow{
						ProposalId:                "ProposalId",
						VoterAddress:              "Voter",
						MaybeVoterOperatorAddress: primptr.String("VoterOperatorAddress"),
						TransactionHash:           "TxHash",
						VoteAtBlockHeight:         1,
						VoteAtBlockTime:           utctime.UTCTime{},
						Answer:                    "Option",
						Histories:                 make([]view.VoteHistory, 0),
						Weight:                    "1.000000000000000000",
						Metadata:                  "Metadata",
					}).
					Return(nil, rdb.ErrNoRows)

				proposal.NewVotes = func(
					_ *rdb.Handle,
				) view.Votes {
					return mockVotesView
				}

				mockProposalsView := &view.MockProposalsView{}
				mocks = append(mocks, &mockProposalsView.Mock)
				mockProposalsView.
					On("FindById", "ProposalId").
					Return([]view.ProposalWithMonikerRow{
						{
							ProposalRow: view.ProposalRow{
								MaybeId:    primptr.Int64(1),
								ProposalId: "ProposalId",
								TotalVote:  big.NewInt(1),
							},
						},
					}, nil)
				mockProposalsView.
					On("Update", &view.ProposalRow{
						MaybeId:    primptr.Int64(1),
						ProposalId: "ProposalId",
						TotalVote:  big.NewInt(2),
					}).
					Return(nil)

				proposal.NewProposals = func(
					_ *rdb.Handle,
				) view.Proposals {
					return mockProposalsView
				}

				mockVotesTotalView := &view.MockVotesTotalView{}
				mocks = append(mocks, &mockVotesTotalView.Mock)
				mockVotesTotalView.
					On("Increment", "ProposalId", int64(1)).
					Return(nil)

				proposal.NewVotesTotal = func(
					_ *rdb.Handle,
				) view.VotesTotal {
					return mockVotesTotalView
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgVoteV1 voter multi-records exist",
			Events: []entity_event.Event{
				&usecase_event.MsgVoteV1{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_VOTE_V1,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					ProposalId: "ProposalId",
					Voter:      "Voter",
					Option:     "Option",
					Metadata:   "Metadata",
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {

				mockValidatorsView := &rdbvalidatorbase_view.MockValidatorsView{}
				mocks = append(mocks, &mockValidatorsView.Mock)
				mockValidatorsView.
					On("FindLastBy", rdbvalidatorbase_view.ValidatorIdentity{
						MaybeInititalDelegatorAddress: primptr.String("Voter"),
					}).
					Return(&rdbvalidatorbase_view.ValidatorRow{
						OperatorAddress: "VoterOperatorAddress",
					}, nil)

				proposal.ValidatorBaseGetView = func(
					_ *rdbvalidatorbase.Base,
					_ *rdb.Handle,
				) rdbvalidatorbase_view.Validators {
					return mockValidatorsView
				}

				mockVotesView := &view.MockVotesView{}
				mocks = append(mocks, &mockVotesView.Mock)
				mockVotesView.
					On("FindByProposalIdVoter", "ProposalId", "Voter").
					Return([]view.VoteWithMonikerRow{
						{
							VoteRow: view.VoteRow{
								ProposalId:                "ProposalId",
								VoterAddress:              "Voter",
								TransactionHash:           "PreviousTransactionHash",
								MaybeVoterOperatorAddress: primptr.String("VoterOperatorAddress"),
								VoteAtBlockHeight:         0,
								VoteAtBlockTime:           utctime.FromUnixNano(-1),
								Answer:                    "PreviousAnswer1",
								Histories:                 make([]view.VoteHistory, 0),
								Metadata:                  "PreviousMetadata",
								Weight:                    "0.800000000000000000",
							},
						},
						{
							VoteRow: view.VoteRow{
								ProposalId:                "ProposalId",
								VoterAddress:              "Voter",
								TransactionHash:           "PreviousTransactionHash",
								MaybeVoterOperatorAddress: primptr.String("VoterOperatorAddress"),
								VoteAtBlockHeight:         0,
								VoteAtBlockTime:           utctime.FromUnixNano(-1),
								Answer:                    "PreviousAnswer2",
								Histories:                 make([]view.VoteHistory, 0),
								Metadata:                  "PreviousMetadata",
								Weight:                    "0.200000000000000000",
							},
						},
					}, nil)

				mockVotesView.
					On("DeleteByProposalIdVoter", "ProposalId", "Voter").
					Return(int64(2), nil)

				mockVotesView.
					On("Insert", &view.VoteRow{
						ProposalId:                "ProposalId",
						VoterAddress:              "Voter",
						MaybeVoterOperatorAddress: primptr.String("VoterOperatorAddress"),
						TransactionHash:           "TxHash",
						VoteAtBlockHeight:         1,
						VoteAtBlockTime:           utctime.UTCTime{},
						Answer:                    "Option",
						Histories: []view.VoteHistory{
							{
								TransactionHash:   "PreviousTransactionHash",
								VoteAtBlockHeight: 0,
								VoteAtBlockTime:   utctime.FromUnixNano(-1),
								Answer:            "PreviousAnswer1",
								Metadata:          "PreviousMetadata",
								Weight:            "0.800000000000000000",
							},
							{
								TransactionHash:   "PreviousTransactionHash",
								VoteAtBlockHeight: 0,
								VoteAtBlockTime:   utctime.FromUnixNano(-1),
								Answer:            "PreviousAnswer2",
								Metadata:          "PreviousMetadata",
								Weight:            "0.200000000000000000",
							},
						},
						Metadata: "Metadata",
						Weight:   "1.000000000000000000",
					}).
					Return(nil, rdb.ErrNoRows)

				proposal.NewVotes = func(
					_ *rdb.Handle,
				) view.Votes {
					return mockVotesView
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgVoteV1 voter record does not exist",
			Events: []entity_event.Event{
				&usecase_event.MsgVoteV1{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_VOTE_V1,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					ProposalId: "ProposalId",
					Voter:      "Voter",
					Option:     "Option",
					Metadata:   "Metadata",
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {

				mockValidatorsView := &rdbvalidatorbase_view.MockValidatorsView{}
				mocks = append(mocks, &mockValidatorsView.Mock)
				mockValidatorsView.
					On("FindLastBy", rdbvalidatorbase_view.ValidatorIdentity{
						MaybeInititalDelegatorAddress: primptr.String("Voter"),
					}).
					Return(&rdbvalidatorbase_view.ValidatorRow{
						OperatorAddress: "VoterOperatorAddress",
					}, nil)

				proposal.ValidatorBaseGetView = func(
					_ *rdbvalidatorbase.Base,
					_ *rdb.Handle,
				) rdbvalidatorbase_view.Validators {
					return mockValidatorsView
				}

				mockVotesView := &view.MockVotesView{}
				mocks = append(mocks, &mockVotesView.Mock)
				mockVotesView.
					On("FindByProposalIdVoter", "ProposalId", "Voter").
					Return(&[]view.VoteWithMonikerRow{}, rdb.ErrNoRows)

				mockVotesView.
					On("Insert", &view.VoteRow{
						ProposalId:                "ProposalId",
						VoterAddress:              "Voter",
						MaybeVoterOperatorAddress: primptr.String("VoterOperatorAddress"),
						TransactionHash:           "TxHash",
						VoteAtBlockHeight:         1,
						VoteAtBlockTime:           utctime.UTCTime{},
						Answer:                    "Option",
						Histories:                 make([]view.VoteHistory, 0),
						Weight:                    "1.000000000000000000",
						Metadata:                  "Metadata",
					}).
					Return(nil, rdb.ErrNoRows)

				proposal.NewVotes = func(
					_ *rdb.Handle,
				) view.Votes {
					return mockVotesView
				}

				mockProposalsView := &view.MockProposalsView{}
				mocks = append(mocks, &mockProposalsView.Mock)
				mockProposalsView.
					On("FindById", "ProposalId").
					Return([]view.ProposalWithMonikerRow{
						{
							ProposalRow: view.ProposalRow{
								MaybeId:    primptr.Int64(1),
								ProposalId: "ProposalId",
								TotalVote:  big.NewInt(1),
							},
						},
						{
							ProposalRow: view.ProposalRow{
								MaybeId:    primptr.Int64(2),
								ProposalId: "ProposalId",
								TotalVote:  big.NewInt(1),
							},
						},
					}, nil)
				mockProposalsView.
					On("Update", &view.ProposalRow{
						MaybeId:    primptr.Int64(1),
						ProposalId: "ProposalId",
						TotalVote:  big.NewInt(2),
					}).
					Return(nil)
				mockProposalsView.
					On("Update", &view.ProposalRow{
						MaybeId:    primptr.Int64(2),
						ProposalId: "ProposalId",
						TotalVote:  big.NewInt(2),
					}).
					Return(nil)

				proposal.NewProposals = func(
					_ *rdb.Handle,
				) view.Proposals {
					return mockProposalsView
				}

				mockVotesTotalView := &view.MockVotesTotalView{}
				mocks = append(mocks, &mockVotesTotalView.Mock)
				mockVotesTotalView.
					On("Increment", "ProposalId", int64(1)).
					Return(nil)

				proposal.NewVotesTotal = func(
					_ *rdb.Handle,
				) view.VotesTotal {
					return mockVotesTotalView
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgVoteWeightedV1 voter record does not exist",
			Events: []entity_event.Event{
				&usecase_event.MsgVoteWeightedV1{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_VOTE_WEIGHTED_V1,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					ProposalId: "ProposalId",
					Voter:      "Voter",
					VoteOptions: []v1_model.VoteOption{
						{
							Option: "Option1",
							Weight: "0.7",
						},
						{
							Option: "Option2",
							Weight: "0.2",
						},
						{
							Option: "Option3",
							Weight: "0.1",
						},
					},
					Metadata: "Metadata",
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {

				mockValidatorsView := &rdbvalidatorbase_view.MockValidatorsView{}
				mocks = append(mocks, &mockValidatorsView.Mock)
				mockValidatorsView.
					On("FindLastBy", rdbvalidatorbase_view.ValidatorIdentity{
						MaybeInititalDelegatorAddress: primptr.String("Voter"),
					}).
					Return(&rdbvalidatorbase_view.ValidatorRow{
						OperatorAddress: "VoterOperatorAddress",
					}, nil)

				proposal.ValidatorBaseGetView = func(
					_ *rdbvalidatorbase.Base,
					_ *rdb.Handle,
				) rdbvalidatorbase_view.Validators {
					return mockValidatorsView
				}

				mockVotesView := &view.MockVotesView{}
				mocks = append(mocks, &mockVotesView.Mock)
				mockVotesView.
					On("FindByProposalIdVoter", "ProposalId", "Voter").
					Return(&[]view.VoteWithMonikerRow{}, rdb.ErrNoRows)

				mockVotesView.
					On("Insert", &view.VoteRow{
						ProposalId:                "ProposalId",
						VoterAddress:              "Voter",
						MaybeVoterOperatorAddress: primptr.String("VoterOperatorAddress"),
						TransactionHash:           "TxHash",
						VoteAtBlockHeight:         1,
						VoteAtBlockTime:           utctime.UTCTime{},
						Answer:                    "Option1",
						Histories:                 make([]view.VoteHistory, 0),
						Weight:                    "0.7",
						Metadata:                  "Metadata",
					}).
					Return(nil, rdb.ErrNoRows)
				mockVotesView.
					On("Insert", &view.VoteRow{
						ProposalId:                "ProposalId",
						VoterAddress:              "Voter",
						MaybeVoterOperatorAddress: primptr.String("VoterOperatorAddress"),
						TransactionHash:           "TxHash",
						VoteAtBlockHeight:         1,
						VoteAtBlockTime:           utctime.UTCTime{},
						Answer:                    "Option2",
						Histories:                 make([]view.VoteHistory, 0),
						Weight:                    "0.2",
						Metadata:                  "Metadata",
					}).
					Return(nil, rdb.ErrNoRows)
				mockVotesView.
					On("Insert", &view.VoteRow{
						ProposalId:                "ProposalId",
						VoterAddress:              "Voter",
						MaybeVoterOperatorAddress: primptr.String("VoterOperatorAddress"),
						TransactionHash:           "TxHash",
						VoteAtBlockHeight:         1,
						VoteAtBlockTime:           utctime.UTCTime{},
						Answer:                    "Option3",
						Histories:                 make([]view.VoteHistory, 0),
						Weight:                    "0.1",
						Metadata:                  "Metadata",
					}).
					Return(nil, rdb.ErrNoRows)

				proposal.NewVotes = func(
					_ *rdb.Handle,
				) view.Votes {
					return mockVotesView
				}

				mockProposalsView := &view.MockProposalsView{}
				mocks = append(mocks, &mockProposalsView.Mock)
				mockProposalsView.
					On("FindById", "ProposalId").
					Return([]view.ProposalWithMonikerRow{
						{
							ProposalRow: view.ProposalRow{
								MaybeId:    primptr.Int64(1),
								ProposalId: "ProposalId",
								TotalVote:  big.NewInt(1),
							},
						},
					}, nil)
				mockProposalsView.
					On("Update", &view.ProposalRow{
						MaybeId:    primptr.Int64(1),
						ProposalId: "ProposalId",
						TotalVote:  big.NewInt(2),
					}).
					Return(nil)

				proposal.NewProposals = func(
					_ *rdb.Handle,
				) view.Proposals {
					return mockProposalsView
				}

				mockVotesTotalView := &view.MockVotesTotalView{}
				mocks = append(mocks, &mockVotesTotalView.Mock)
				mockVotesTotalView.
					On("Increment", "ProposalId", int64(1)).
					Return(nil)

				proposal.NewVotesTotal = func(
					_ *rdb.Handle,
				) view.VotesTotal {
					return mockVotesTotalView
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgVoteWeightedV1 voter multi-records exist",
			Events: []entity_event.Event{
				&usecase_event.MsgVoteWeightedV1{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_VOTE_WEIGHTED_V1,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					ProposalId: "ProposalId",
					Voter:      "Voter",
					VoteOptions: []v1_model.VoteOption{
						{
							Option: "Option1",
							Weight: "0.7",
						},
						{
							Option: "Option2",
							Weight: "0.3",
						},
					},
					Metadata: "Metadata",
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {

				mockValidatorsView := &rdbvalidatorbase_view.MockValidatorsView{}
				mocks = append(mocks, &mockValidatorsView.Mock)
				mockValidatorsView.
					On("FindLastBy", rdbvalidatorbase_view.ValidatorIdentity{
						MaybeInititalDelegatorAddress: primptr.String("Voter"),
					}).
					Return(&rdbvalidatorbase_view.ValidatorRow{
						OperatorAddress: "VoterOperatorAddress",
					}, nil)

				proposal.ValidatorBaseGetView = func(
					_ *rdbvalidatorbase.Base,
					_ *rdb.Handle,
				) rdbvalidatorbase_view.Validators {
					return mockValidatorsView
				}

				mockVotesView := &view.MockVotesView{}
				mocks = append(mocks, &mockVotesView.Mock)
				mockVotesView.
					On("FindByProposalIdVoter", "ProposalId", "Voter").
					Return([]view.VoteWithMonikerRow{
						{
							VoteRow: view.VoteRow{
								ProposalId:                "ProposalId",
								VoterAddress:              "Voter",
								TransactionHash:           "PreviousTransactionHash",
								MaybeVoterOperatorAddress: primptr.String("VoterOperatorAddress"),
								VoteAtBlockHeight:         0,
								VoteAtBlockTime:           utctime.FromUnixNano(-1),
								Answer:                    "PreviousAnswer1",
								Histories:                 make([]view.VoteHistory, 0),
								Metadata:                  "PreviousMetadata",
								Weight:                    "0.800000000000000000",
							},
						},
						{
							VoteRow: view.VoteRow{
								ProposalId:                "ProposalId",
								VoterAddress:              "Voter",
								TransactionHash:           "PreviousTransactionHash",
								MaybeVoterOperatorAddress: primptr.String("VoterOperatorAddress"),
								VoteAtBlockHeight:         0,
								VoteAtBlockTime:           utctime.FromUnixNano(-1),
								Answer:                    "PreviousAnswer2",
								Histories:                 make([]view.VoteHistory, 0),
								Metadata:                  "PreviousMetadata",
								Weight:                    "0.200000000000000000",
							},
						},
					}, nil)

				mockVotesView.
					On("DeleteByProposalIdVoter", "ProposalId", "Voter").
					Return(int64(2), nil)

				mockVotesView.
					On("Insert", &view.VoteRow{
						ProposalId:                "ProposalId",
						VoterAddress:              "Voter",
						MaybeVoterOperatorAddress: primptr.String("VoterOperatorAddress"),
						TransactionHash:           "TxHash",
						VoteAtBlockHeight:         1,
						VoteAtBlockTime:           utctime.UTCTime{},
						Answer:                    "Option1",
						Histories: []view.VoteHistory{
							{
								TransactionHash:   "PreviousTransactionHash",
								VoteAtBlockHeight: 0,
								VoteAtBlockTime:   utctime.FromUnixNano(-1),
								Answer:            "PreviousAnswer1",
								Metadata:          "PreviousMetadata",
								Weight:            "0.800000000000000000",
							},
							{
								TransactionHash:   "PreviousTransactionHash",
								VoteAtBlockHeight: 0,
								VoteAtBlockTime:   utctime.FromUnixNano(-1),
								Answer:            "PreviousAnswer2",
								Metadata:          "PreviousMetadata",
								Weight:            "0.200000000000000000",
							},
						},
						Metadata: "Metadata",
						Weight:   "0.7",
					}).
					Return(nil, rdb.ErrNoRows)
				mockVotesView.
					On("Insert", &view.VoteRow{
						ProposalId:                "ProposalId",
						VoterAddress:              "Voter",
						MaybeVoterOperatorAddress: primptr.String("VoterOperatorAddress"),
						TransactionHash:           "TxHash",
						VoteAtBlockHeight:         1,
						VoteAtBlockTime:           utctime.UTCTime{},
						Answer:                    "Option2",
						Histories: []view.VoteHistory{
							{
								TransactionHash:   "PreviousTransactionHash",
								VoteAtBlockHeight: 0,
								VoteAtBlockTime:   utctime.FromUnixNano(-1),
								Answer:            "PreviousAnswer1",
								Metadata:          "PreviousMetadata",
								Weight:            "0.800000000000000000",
							},
							{
								TransactionHash:   "PreviousTransactionHash",
								VoteAtBlockHeight: 0,
								VoteAtBlockTime:   utctime.FromUnixNano(-1),
								Answer:            "PreviousAnswer2",
								Metadata:          "PreviousMetadata",
								Weight:            "0.200000000000000000",
							},
						},
						Metadata: "Metadata",
						Weight:   "0.3",
					}).
					Return(nil, rdb.ErrNoRows)

				proposal.NewVotes = func(
					_ *rdb.Handle,
				) view.Votes {
					return mockVotesView
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgVoteWeightedV1 vote for multi-msgs proposal",
			Events: []entity_event.Event{
				&usecase_event.MsgVoteWeightedV1{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_VOTE_WEIGHTED_V1,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					ProposalId: "ProposalId",
					Voter:      "Voter",
					VoteOptions: []v1_model.VoteOption{
						{
							Option: "Option1",
							Weight: "0.7",
						},
						{
							Option: "Option2",
							Weight: "0.2",
						},
						{
							Option: "Option3",
							Weight: "0.1",
						},
					},
					Metadata: "Metadata",
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {

				mockValidatorsView := &rdbvalidatorbase_view.MockValidatorsView{}
				mocks = append(mocks, &mockValidatorsView.Mock)
				mockValidatorsView.
					On("FindLastBy", rdbvalidatorbase_view.ValidatorIdentity{
						MaybeInititalDelegatorAddress: primptr.String("Voter"),
					}).
					Return(&rdbvalidatorbase_view.ValidatorRow{
						OperatorAddress: "VoterOperatorAddress",
					}, nil)

				proposal.ValidatorBaseGetView = func(
					_ *rdbvalidatorbase.Base,
					_ *rdb.Handle,
				) rdbvalidatorbase_view.Validators {
					return mockValidatorsView
				}

				mockVotesView := &view.MockVotesView{}
				mocks = append(mocks, &mockVotesView.Mock)
				mockVotesView.
					On("FindByProposalIdVoter", "ProposalId", "Voter").
					Return(&[]view.VoteWithMonikerRow{}, rdb.ErrNoRows)

				mockVotesView.
					On("Insert", &view.VoteRow{
						ProposalId:                "ProposalId",
						VoterAddress:              "Voter",
						MaybeVoterOperatorAddress: primptr.String("VoterOperatorAddress"),
						TransactionHash:           "TxHash",
						VoteAtBlockHeight:         1,
						VoteAtBlockTime:           utctime.UTCTime{},
						Answer:                    "Option1",
						Histories:                 make([]view.VoteHistory, 0),
						Weight:                    "0.7",
						Metadata:                  "Metadata",
					}).
					Return(nil, rdb.ErrNoRows)
				mockVotesView.
					On("Insert", &view.VoteRow{
						ProposalId:                "ProposalId",
						VoterAddress:              "Voter",
						MaybeVoterOperatorAddress: primptr.String("VoterOperatorAddress"),
						TransactionHash:           "TxHash",
						VoteAtBlockHeight:         1,
						VoteAtBlockTime:           utctime.UTCTime{},
						Answer:                    "Option2",
						Histories:                 make([]view.VoteHistory, 0),
						Weight:                    "0.2",
						Metadata:                  "Metadata",
					}).
					Return(nil, rdb.ErrNoRows)
				mockVotesView.
					On("Insert", &view.VoteRow{
						ProposalId:                "ProposalId",
						VoterAddress:              "Voter",
						MaybeVoterOperatorAddress: primptr.String("VoterOperatorAddress"),
						TransactionHash:           "TxHash",
						VoteAtBlockHeight:         1,
						VoteAtBlockTime:           utctime.UTCTime{},
						Answer:                    "Option3",
						Histories:                 make([]view.VoteHistory, 0),
						Weight:                    "0.1",
						Metadata:                  "Metadata",
					}).
					Return(nil, rdb.ErrNoRows)

				proposal.NewVotes = func(
					_ *rdb.Handle,
				) view.Votes {
					return mockVotesView
				}

				mockProposalsView := &view.MockProposalsView{}
				mocks = append(mocks, &mockProposalsView.Mock)
				mockProposalsView.
					On("FindById", "ProposalId").
					Return([]view.ProposalWithMonikerRow{
						{
							ProposalRow: view.ProposalRow{
								MaybeId:    primptr.Int64(1),
								ProposalId: "ProposalId",
								Type:       "Type1",
								TotalVote:  big.NewInt(1),
							},
						},
						{
							ProposalRow: view.ProposalRow{
								MaybeId:    primptr.Int64(2),
								ProposalId: "ProposalId",
								Type:       "Type2",
								TotalVote:  big.NewInt(1),
							},
						},
					}, nil)
				mockProposalsView.
					On("Update", &view.ProposalRow{
						MaybeId:    primptr.Int64(1),
						ProposalId: "ProposalId",
						Type:       "Type1",
						TotalVote:  big.NewInt(2),
					}).
					Return(nil)
				mockProposalsView.
					On("Update", &view.ProposalRow{
						MaybeId:    primptr.Int64(2),
						ProposalId: "ProposalId",
						Type:       "Type2",
						TotalVote:  big.NewInt(2),
					}).
					Return(nil)

				proposal.NewProposals = func(
					_ *rdb.Handle,
				) view.Proposals {
					return mockProposalsView
				}

				mockVotesTotalView := &view.MockVotesTotalView{}
				mocks = append(mocks, &mockVotesTotalView.Mock)
				mockVotesTotalView.
					On("Increment", "ProposalId", int64(1)).
					Return(nil)

				proposal.NewVotesTotal = func(
					_ *rdb.Handle,
				) view.VotesTotal {
					return mockVotesTotalView
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
		mocks = append(mocks, &mockRDbConn.Mock)
		mocks = append(mocks, &mockTx.Mock)

		// Replace some function implementations.
		// We are not interested to test them in unit test.
		proposal.ParamBaseHandleEvents = func(
			_ *rdbparambase.Base,
			_ *rdb.Handle,
			_ logger.Logger,
			events []entity_event.Event,
		) error {
			return nil
		}
		proposal.ValidatorBaseHandleEvents = func(
			_ *rdbvalidatorbase.Base,
			_ *rdb.Handle,
			_ logger.Logger,
			events []entity_event.Event,
		) error {
			return nil
		}
		proposal.UpdateLastHandledEventHeight = func(
			_ *proposal.Proposal,
			_ *rdb.Handle,
			_ int64,
		) error {
			return nil
		}

		projection := NewProposalProjection(mockRDbConn)
		err := projection.HandleEvents(1, tc.Events)
		assert.NoError(t, err)

		for _, m := range mocks {
			m.AssertExpectations(t)
		}

		fmt.Println(tc.Name, "Passed")
	}
}
