package validator_test

import (
	"fmt"
	"math/big"
	"testing"

	sq "github.com/Masterminds/squirrel"
	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	test_logger "github.com/crypto-com/chain-indexing/external/logger/test"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/projection/validator"
	"github.com/crypto-com/chain-indexing/projection/validator/view"
	validator_view "github.com/crypto-com/chain-indexing/projection/validator/view"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	model_usecase "github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	testify_mock "github.com/stretchr/testify/mock"
)

func NewValidatorProjection(rdbConn rdb.Conn) *validator.Validator {
	return validator.NewValidator(
		test_logger.NewFakeLogger(),
		rdbConn,
		"tcrocncl",
		nil,
		&validator.Config{
			validator.Rules{
				validator.MaxCommissionRateChange{
					Enable:    true,
					MaxChange: float64(0.1),
				},
				validator.MaxEditQuota{
					Enable: true,
					Quota: map[string]int{
						"moniker":        2,
						"commissionRate": 2,
					},
					Interval: "24h",
				},
			},
		},
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

func TestValidator_HandleEvents(t *testing.T) {
	testCases := []struct {
		Name     string
		Events   []event_entity.Event
		MockFunc func(events []event_entity.Event) []*testify_mock.Mock
	}{
		{
			Name: "Handle MsgEditValidator moniker changed more than twice",
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
						Time:            utctime.FromUnixNano(1665902976672205000),
						AppHash:         "AppHash",
						ProposerAddress: "ProposerAddress",
						Txs:             nil,
						Signatures:      nil,
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
					Description: model.ValidatorDescription{
						Moniker:         "newMoniker",
						Identity:        "[do-not-modify]",
						Website:         "[do-not-modify]",
						SecurityContact: "[do-not-modify]",
						Details:         "[do-not-modify]",
					},
					ValidatorAddress:       "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
					MaybeCommissionRate:    nil,
					MaybeMinSelfDelegation: nil,
				},
			},
			MockFunc: func(events []event_entity.Event) (mocks []*testify_mock.Mock) {
				mockValidatorsView := validator_view.NewMockValidatorsView(nil).(*validator_view.MockValidatorsView)
				mocks = append(mocks, &mockValidatorsView.Mock)

				validator.NewValidators = func(_ *rdb.Handle) validator_view.Validators {
					return mockValidatorsView
				}

				mockValidatorsView.On(
					"FindBy",
					validator_view.ValidatorIdentity{
						MaybeOperatorAddress: primptr.String("tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus"),
					},
				).Return(
					&validator_view.ValidatorRow{
						MaybeId:                 primptr.Int64(3),
						OperatorAddress:         "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
						ConsensusNodeAddress:    "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
						InitialDelegatorAddress: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						TendermintPubkey:        "Kpox5fS2po0sJUHmzllExuJ4uZ5nm0bbCp6UQKESsnE=",
						TendermintAddress:       "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						Status:                  "Bonded",
						Jailed:                  false,
						JoinedAtBlockHeight:     1,
						Power:                   "",
						Moniker:                 "[do-not-modify]",
						Identity:                "[do-not-modify]",
						Website:                 "[do-not-modify]",
						SecurityContact:         "[do-not-modify]",
						Details:                 "[do-not-modify]",
						CommissionRate:          "0.1",
						CommissionMaxRate:       "0.2",
						CommissionMaxChangeRate: "0.01",
						MinSelfDelegation:       "1",
						TotalSignedBlock:        1,
						TotalActiveBlock:        1,
						ImpreciseUpTime:         big.NewFloat(0.1),
						VotedGovProposal:        big.NewInt(1),
					}, nil)

				mockValidatorActivitiesView := validator_view.NewMockValidatorActivitiesView(nil).(*validator_view.MockValidatorActivitiesView)
				mocks = append(mocks, &mockValidatorActivitiesView.Mock)

				validator.NewValidatorActivities = func(_ *rdb.Handle) validator_view.ValidatorActivities {
					return mockValidatorActivitiesView
				}

				timestamp := utctime.FromUnixNano(1665816576672205000)

				mockValidatorActivitiesView.On(
					"List",
					validator_view.ValidatorActivitiesListFilter{
						AfterBlockTime:       &timestamp,
						MaybeOperatorAddress: primptr.String("tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus"),
					},
					validator_view.ValidatorActivitiesListOrder{},
					&pagination_interface.Pagination{},
				).Return(
					[]validator_view.ValidatorActivityRow{
						{

							BlockHeight:          1,
							BlockTime:            utctime.FromUnixNano(166590297672205000),
							BlockHash:            "Hash",
							MaybeTransactionHash: primptr.String("TxHash"),
							OperatorAddress:      "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
							Success:              true,
							Data: validator_view.ValidatorActivityRowData{
								Type: "/cosmos.staking.v1beta1.MsgEditValidator",
								Content: map[string]interface{}{
									"name":    "/cosmos.staking.v1beta1.MsgEditValidator.Created",
									"height":  37333,
									"txHash":  "BD5D2BE22166DB174B93AA42689AFC737DE42365F79912E80B64890C127A28DA",
									"msgName": "/cosmos.staking.v1beta1.MsgEditValidator",
									"description": map[string]interface{}{
										"details":         "[do-not-modify]",
										"moniker":         "newMoniker",
										"website":         "[do-not-modify]",
										"identity":        "[do-not-modify]",
										"securityContact": "[do-not-modify]",
									},
									"commissionRate":    "null",
									"validatorAddress":  "crcvaloper12luku6uxehhak02py4rcz65zu0swh7wj6ulrlg",
									"minSelfDelegation": "null",
								},
							},
						},
						{

							BlockHeight:          1,
							BlockTime:            utctime.FromUnixNano(166590297672205000),
							BlockHash:            "Hash",
							MaybeTransactionHash: primptr.String("TxHash"),
							OperatorAddress:      "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
							Success:              true,
							Data: validator_view.ValidatorActivityRowData{
								Type: "/cosmos.staking.v1beta1.MsgEditValidator",
								Content: map[string]interface{}{
									"name":    "/cosmos.staking.v1beta1.MsgEditValidator.Created",
									"height":  37333,
									"txHash":  "BD5D2BE22166DB174B93AA42689AFC737DE42365F79912E80B64890C127A28DA",
									"msgName": "/cosmos.staking.v1beta1.MsgEditValidator",
									"description": map[string]interface{}{
										"details":         "[do-not-modify]",
										"moniker":         "newMoniker",
										"website":         "[do-not-modify]",
										"identity":        "[do-not-modify]",
										"securityContact": "[do-not-modify]",
									},
									"commissionRate":    "null",
									"validatorAddress":  "crcvaloper12luku6uxehhak02py4rcz65zu0swh7wj6ulrlg",
									"minSelfDelegation": "null",
								},
							},
						},
					}, nil, nil)

				mockValidatorsView.On(
					"Update",
					&view.ValidatorRow{
						MaybeId:                 primptr.Int64(3),
						OperatorAddress:         "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
						ConsensusNodeAddress:    "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
						InitialDelegatorAddress: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						TendermintPubkey:        "Kpox5fS2po0sJUHmzllExuJ4uZ5nm0bbCp6UQKESsnE=",
						TendermintAddress:       "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						Status:                  "Attention",
						Jailed:                  false,
						JoinedAtBlockHeight:     1,
						Power:                   "",
						Moniker:                 "newMoniker",
						Identity:                "[do-not-modify]",
						Website:                 "[do-not-modify]",
						SecurityContact:         "[do-not-modify]",
						Details:                 "[do-not-modify]",
						CommissionRate:          "0.1",
						CommissionMaxRate:       "0.2",
						CommissionMaxChangeRate: "0.01",
						MinSelfDelegation:       "1",
						TotalSignedBlock:        1,
						TotalActiveBlock:        1,
						ImpreciseUpTime:         big.NewFloat(0.1),
						VotedGovProposal:        big.NewInt(1),
					},
				).Return(nil)

				mockValidatorActivitiesView.On(
					"InsertAll",
					mock.Anything,
				).Return(nil)

				mockValidatorActivitiesTotalView := validator_view.NewMockValidatorActivitiesTotalView(nil).(*validator_view.MockValidatorActivitiesTotalView)
				mocks = append(mocks, &mockValidatorActivitiesTotalView.Mock)

				validator.NewValidatorActivitiesTotal = func(_ *rdb.Handle) validator_view.ValidatorActivitiesTotal {
					return mockValidatorActivitiesTotalView
				}

				mockValidatorActivitiesTotalView.On(
					"Increment",
					"-",
					int64(1),
				).Return(nil)

				mockValidatorActivitiesTotalView.On(
					"Increment",
					"tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
					int64(1),
				).Return(nil)

				mockValidatorActivitiesTotalView.On(
					"Increment",
					"tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus:/cosmos.staking.v1beta1.MsgEditValidator.Created",
					int64(1),
				).Return(nil)

				mockValidatorActivitiesTotalView.On(
					"Increment",
					"-:/cosmos.staking.v1beta1.MsgEditValidator.Created",
					int64(1),
				).Return(nil)

				mockValidatorsView.On(
					"ListAll",
					mock.Anything,
					mock.Anything,
				).Return(
					[]validator_view.ValidatorRow{}, nil)

				mockValidatorBlockCommitmentsTotalView := validator_view.NewMockValidatorActivitiesTotalView(nil).(*validator_view.MockValidatorActivitiesTotalView)
				mocks = append(mocks, &mockValidatorBlockCommitmentsTotalView.Mock)

				validator.NewValidatorBlockCommitmentsTotal = func(_ *rdb.Handle) validator_view.ValidatorBlockCommitmentsTotal {
					return mockValidatorBlockCommitmentsTotalView
				}

				mockValidatorBlockCommitmentsTotalView.On(
					"Set",
					"1:-",
					int64(0),
				).Return(nil)

				mockValidatorBlockCommitmentsTotalView.On(
					"IncrementAll",
					[]string{},
					int64(1),
				).Return(nil)

				mockValidatorBlockCommitmentsTotalView.On(
					"Increment",
					"-:-",
					int64(0),
				).Return(nil)

				mockValidatorsView.On(
					"UpdateAllValidatorUpTime",
					mock.Anything,
				).Return(nil)

				return mocks
			},
		},
		{
			Name: "Handle MsgEditValidator commission changed more than twice",
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
						Time:            utctime.FromUnixNano(1665902976672205000),
						AppHash:         "AppHash",
						ProposerAddress: "ProposerAddress",
						Txs:             nil,
						Signatures:      nil,
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
					Description: model.ValidatorDescription{
						Moniker:         "[do-not-modify]",
						Identity:        "[do-not-modify]",
						Website:         "[do-not-modify]",
						SecurityContact: "[do-not-modify]",
						Details:         "[do-not-modify]",
					},
					ValidatorAddress:       "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
					MaybeCommissionRate:    primptr.String("0.2"),
					MaybeMinSelfDelegation: nil,
				},
			},
			MockFunc: func(events []event_entity.Event) (mocks []*testify_mock.Mock) {
				mockValidatorsView := validator_view.NewMockValidatorsView(nil).(*validator_view.MockValidatorsView)
				mocks = append(mocks, &mockValidatorsView.Mock)

				validator.NewValidators = func(_ *rdb.Handle) validator_view.Validators {
					return mockValidatorsView
				}

				mockValidatorsView.On(
					"FindBy",
					validator_view.ValidatorIdentity{
						MaybeOperatorAddress: primptr.String("tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus"),
					},
				).Return(
					&validator_view.ValidatorRow{
						MaybeId:                 primptr.Int64(3),
						OperatorAddress:         "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
						ConsensusNodeAddress:    "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
						InitialDelegatorAddress: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						TendermintPubkey:        "Kpox5fS2po0sJUHmzllExuJ4uZ5nm0bbCp6UQKESsnE=",
						TendermintAddress:       "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						Status:                  "Bonded",
						Jailed:                  false,
						JoinedAtBlockHeight:     1,
						Power:                   "",
						Moniker:                 "[do-not-modify]",
						Identity:                "[do-not-modify]",
						Website:                 "[do-not-modify]",
						SecurityContact:         "[do-not-modify]",
						Details:                 "[do-not-modify]",
						CommissionRate:          "0.1",
						CommissionMaxRate:       "0.2",
						CommissionMaxChangeRate: "0.01",
						MinSelfDelegation:       "1",
						TotalSignedBlock:        1,
						TotalActiveBlock:        1,
						ImpreciseUpTime:         big.NewFloat(0.1),
						VotedGovProposal:        big.NewInt(1),
					}, nil)

				mockValidatorActivitiesView := validator_view.NewMockValidatorActivitiesView(nil).(*validator_view.MockValidatorActivitiesView)
				mocks = append(mocks, &mockValidatorActivitiesView.Mock)

				validator.NewValidatorActivities = func(_ *rdb.Handle) validator_view.ValidatorActivities {
					return mockValidatorActivitiesView
				}

				timestamp := utctime.FromUnixNano(1665816576672205000)

				mockValidatorActivitiesView.On(
					"List",
					validator_view.ValidatorActivitiesListFilter{
						AfterBlockTime:       &timestamp,
						MaybeOperatorAddress: primptr.String("tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus"),
					},
					validator_view.ValidatorActivitiesListOrder{},
					&pagination_interface.Pagination{},
				).Return(
					[]validator_view.ValidatorActivityRow{
						{

							BlockHeight:          1,
							BlockTime:            utctime.FromUnixNano(166590297672205000),
							BlockHash:            "Hash",
							MaybeTransactionHash: primptr.String("TxHash"),
							OperatorAddress:      "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
							Success:              true,
							Data: validator_view.ValidatorActivityRowData{
								Type: "/cosmos.staking.v1beta1.MsgEditValidator",
								Content: map[string]interface{}{
									"name":    "/cosmos.staking.v1beta1.MsgEditValidator.Created",
									"height":  37333,
									"txHash":  "BD5D2BE22166DB174B93AA42689AFC737DE42365F79912E80B64890C127A28DA",
									"msgName": "/cosmos.staking.v1beta1.MsgEditValidator",
									"description": map[string]interface{}{
										"details":         "[do-not-modify]",
										"moniker":         "[do-not-modify]",
										"website":         "[do-not-modify]",
										"identity":        "[do-not-modify]",
										"securityContact": "[do-not-modify]",
									},
									"commissionRate":    "0.1",
									"validatorAddress":  "crcvaloper12luku6uxehhak02py4rcz65zu0swh7wj6ulrlg",
									"minSelfDelegation": "null",
								},
							},
						},
						{

							BlockHeight:          1,
							BlockTime:            utctime.FromUnixNano(166590297672205000),
							BlockHash:            "Hash",
							MaybeTransactionHash: primptr.String("TxHash"),
							OperatorAddress:      "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
							Success:              true,
							Data: validator_view.ValidatorActivityRowData{
								Type: "/cosmos.staking.v1beta1.MsgEditValidator",
								Content: map[string]interface{}{
									"name":    "/cosmos.staking.v1beta1.MsgEditValidator.Created",
									"height":  37333,
									"txHash":  "BD5D2BE22166DB174B93AA42689AFC737DE42365F79912E80B64890C127A28DA",
									"msgName": "/cosmos.staking.v1beta1.MsgEditValidator",
									"description": map[string]interface{}{
										"details":         "[do-not-modify]",
										"moniker":         "[do-not-modify]",
										"website":         "[do-not-modify]",
										"identity":        "[do-not-modify]",
										"securityContact": "[do-not-modify]",
									},
									"commissionRate":    "0.1",
									"validatorAddress":  "crcvaloper12luku6uxehhak02py4rcz65zu0swh7wj6ulrlg",
									"minSelfDelegation": "null",
								},
							},
						},
					}, nil, nil)

				mockValidatorsView.On(
					"Update",
					&view.ValidatorRow{
						MaybeId:                 primptr.Int64(3),
						OperatorAddress:         "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
						ConsensusNodeAddress:    "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
						InitialDelegatorAddress: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						TendermintPubkey:        "Kpox5fS2po0sJUHmzllExuJ4uZ5nm0bbCp6UQKESsnE=",
						TendermintAddress:       "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						Status:                  "Attention",
						Jailed:                  false,
						JoinedAtBlockHeight:     1,
						Power:                   "",
						Moniker:                 "[do-not-modify]",
						Identity:                "[do-not-modify]",
						Website:                 "[do-not-modify]",
						SecurityContact:         "[do-not-modify]",
						Details:                 "[do-not-modify]",
						CommissionRate:          "0.2",
						CommissionMaxRate:       "0.2",
						CommissionMaxChangeRate: "0.01",
						MinSelfDelegation:       "1",
						TotalSignedBlock:        1,
						TotalActiveBlock:        1,
						ImpreciseUpTime:         big.NewFloat(0.1),
						VotedGovProposal:        big.NewInt(1),
					},
				).Return(nil)

				mockValidatorActivitiesView.On(
					"InsertAll",
					mock.Anything,
				).Return(nil)

				mockValidatorActivitiesTotalView := validator_view.NewMockValidatorActivitiesTotalView(nil).(*validator_view.MockValidatorActivitiesTotalView)
				mocks = append(mocks, &mockValidatorActivitiesTotalView.Mock)

				validator.NewValidatorActivitiesTotal = func(_ *rdb.Handle) validator_view.ValidatorActivitiesTotal {
					return mockValidatorActivitiesTotalView
				}

				mockValidatorActivitiesTotalView.On(
					"Increment",
					"-",
					int64(1),
				).Return(nil)

				mockValidatorActivitiesTotalView.On(
					"Increment",
					"tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
					int64(1),
				).Return(nil)

				mockValidatorActivitiesTotalView.On(
					"Increment",
					"tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus:/cosmos.staking.v1beta1.MsgEditValidator.Created",
					int64(1),
				).Return(nil)

				mockValidatorActivitiesTotalView.On(
					"Increment",
					"-:/cosmos.staking.v1beta1.MsgEditValidator.Created",
					int64(1),
				).Return(nil)

				mockValidatorsView.On(
					"ListAll",
					mock.Anything,
					mock.Anything,
				).Return(
					[]validator_view.ValidatorRow{}, nil)

				mockValidatorBlockCommitmentsTotalView := validator_view.NewMockValidatorActivitiesTotalView(nil).(*validator_view.MockValidatorActivitiesTotalView)
				mocks = append(mocks, &mockValidatorBlockCommitmentsTotalView.Mock)

				validator.NewValidatorBlockCommitmentsTotal = func(_ *rdb.Handle) validator_view.ValidatorBlockCommitmentsTotal {
					return mockValidatorBlockCommitmentsTotalView
				}

				mockValidatorBlockCommitmentsTotalView.On(
					"Set",
					"1:-",
					int64(0),
				).Return(nil)

				mockValidatorBlockCommitmentsTotalView.On(
					"IncrementAll",
					[]string{},
					int64(1),
				).Return(nil)

				mockValidatorBlockCommitmentsTotalView.On(
					"Increment",
					"-:-",
					int64(0),
				).Return(nil)

				mockValidatorsView.On(
					"UpdateAllValidatorUpTime",
					mock.Anything,
				).Return(nil)

				return mocks
			},
		},
		{
			Name: "Handle MsgEditValidator changed with extreme commission",
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
						Time:            utctime.FromUnixNano(1665902976672205000),
						AppHash:         "AppHash",
						ProposerAddress: "ProposerAddress",
						Txs:             nil,
						Signatures:      nil,
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
					Description: model.ValidatorDescription{
						Moniker:         "[do-not-modify]",
						Identity:        "[do-not-modify]",
						Website:         "[do-not-modify]",
						SecurityContact: "[do-not-modify]",
						Details:         "[do-not-modify]",
					},
					ValidatorAddress:       "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
					MaybeCommissionRate:    primptr.String("0.21"),
					MaybeMinSelfDelegation: nil,
				},
			},
			MockFunc: func(events []event_entity.Event) (mocks []*testify_mock.Mock) {
				mockValidatorsView := validator_view.NewMockValidatorsView(nil).(*validator_view.MockValidatorsView)
				mocks = append(mocks, &mockValidatorsView.Mock)

				validator.NewValidators = func(_ *rdb.Handle) validator_view.Validators {
					return mockValidatorsView
				}

				mockValidatorsView.On(
					"FindBy",
					validator_view.ValidatorIdentity{
						MaybeOperatorAddress: primptr.String("tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus"),
					},
				).Return(
					&validator_view.ValidatorRow{
						MaybeId:                 primptr.Int64(3),
						OperatorAddress:         "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
						ConsensusNodeAddress:    "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
						InitialDelegatorAddress: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						TendermintPubkey:        "Kpox5fS2po0sJUHmzllExuJ4uZ5nm0bbCp6UQKESsnE=",
						TendermintAddress:       "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						Status:                  "Bonded",
						Jailed:                  false,
						JoinedAtBlockHeight:     1,
						Power:                   "",
						Moniker:                 "[do-not-modify]",
						Identity:                "[do-not-modify]",
						Website:                 "[do-not-modify]",
						SecurityContact:         "[do-not-modify]",
						Details:                 "[do-not-modify]",
						CommissionRate:          "0.1",
						CommissionMaxRate:       "0.2",
						CommissionMaxChangeRate: "0.01",
						MinSelfDelegation:       "1",
						TotalSignedBlock:        1,
						TotalActiveBlock:        1,
						ImpreciseUpTime:         big.NewFloat(0.1),
						VotedGovProposal:        big.NewInt(1),
					}, nil)

				mockValidatorActivitiesView := validator_view.NewMockValidatorActivitiesView(nil).(*validator_view.MockValidatorActivitiesView)
				mocks = append(mocks, &mockValidatorActivitiesView.Mock)

				validator.NewValidatorActivities = func(_ *rdb.Handle) validator_view.ValidatorActivities {
					return mockValidatorActivitiesView
				}

				timestamp := utctime.FromUnixNano(1665816576672205000)

				mockValidatorActivitiesView.On(
					"List",
					validator_view.ValidatorActivitiesListFilter{
						AfterBlockTime:       &timestamp,
						MaybeOperatorAddress: primptr.String("tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus"),
					},
					validator_view.ValidatorActivitiesListOrder{},
					&pagination_interface.Pagination{},
				).Return(
					[]validator_view.ValidatorActivityRow{}, nil, nil)

				mockValidatorsView.On(
					"Update",
					&view.ValidatorRow{
						MaybeId:                 primptr.Int64(3),
						OperatorAddress:         "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
						ConsensusNodeAddress:    "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
						InitialDelegatorAddress: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						TendermintPubkey:        "Kpox5fS2po0sJUHmzllExuJ4uZ5nm0bbCp6UQKESsnE=",
						TendermintAddress:       "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						Status:                  "Attention",
						Jailed:                  false,
						JoinedAtBlockHeight:     1,
						Power:                   "",
						Moniker:                 "[do-not-modify]",
						Identity:                "[do-not-modify]",
						Website:                 "[do-not-modify]",
						SecurityContact:         "[do-not-modify]",
						Details:                 "[do-not-modify]",
						CommissionRate:          "0.21",
						CommissionMaxRate:       "0.2",
						CommissionMaxChangeRate: "0.01",
						MinSelfDelegation:       "1",
						TotalSignedBlock:        1,
						TotalActiveBlock:        1,
						ImpreciseUpTime:         big.NewFloat(0.1),
						VotedGovProposal:        big.NewInt(1),
					},
				).Return(nil)

				mockValidatorActivitiesView.On(
					"InsertAll",
					mock.Anything,
				).Return(nil)

				mockValidatorActivitiesTotalView := validator_view.NewMockValidatorActivitiesTotalView(nil).(*validator_view.MockValidatorActivitiesTotalView)
				mocks = append(mocks, &mockValidatorActivitiesTotalView.Mock)

				validator.NewValidatorActivitiesTotal = func(_ *rdb.Handle) validator_view.ValidatorActivitiesTotal {
					return mockValidatorActivitiesTotalView
				}

				mockValidatorActivitiesTotalView.On(
					"Increment",
					"-",
					int64(1),
				).Return(nil)

				mockValidatorActivitiesTotalView.On(
					"Increment",
					"tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
					int64(1),
				).Return(nil)

				mockValidatorActivitiesTotalView.On(
					"Increment",
					"tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus:/cosmos.staking.v1beta1.MsgEditValidator.Created",
					int64(1),
				).Return(nil)

				mockValidatorActivitiesTotalView.On(
					"Increment",
					"-:/cosmos.staking.v1beta1.MsgEditValidator.Created",
					int64(1),
				).Return(nil)

				mockValidatorsView.On(
					"ListAll",
					mock.Anything,
					mock.Anything,
				).Return(
					[]validator_view.ValidatorRow{}, nil)

				mockValidatorBlockCommitmentsTotalView := validator_view.NewMockValidatorActivitiesTotalView(nil).(*validator_view.MockValidatorActivitiesTotalView)
				mocks = append(mocks, &mockValidatorBlockCommitmentsTotalView.Mock)

				validator.NewValidatorBlockCommitmentsTotal = func(_ *rdb.Handle) validator_view.ValidatorBlockCommitmentsTotal {
					return mockValidatorBlockCommitmentsTotalView
				}

				mockValidatorBlockCommitmentsTotalView.On(
					"Set",
					"1:-",
					int64(0),
				).Return(nil)

				mockValidatorBlockCommitmentsTotalView.On(
					"IncrementAll",
					[]string{},
					int64(1),
				).Return(nil)

				mockValidatorBlockCommitmentsTotalView.On(
					"Increment",
					"-:-",
					int64(0),
				).Return(nil)

				mockValidatorsView.On(
					"UpdateAllValidatorUpTime",
					mock.Anything,
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
		validator.UpdateLastHandledEventHeight = func(_ *validator.Validator, _ *rdb.Handle, _ int64) error {
			return nil
		}

		projection := NewValidatorProjection(mockRDbConn)
		err := projection.HandleEvents(1, tc.Events)
		assert.NoError(t, err)

		for _, m := range mocks {
			m.AssertExpectations(t)
		}

		fmt.Println(tc.Name, "Passed")
	}
}
