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
	testify_mock "github.com/stretchr/testify/mock"
)

const DefaultMaxActiveBlocksPeriodLimit = 100000

func NewValidatorProjection(rdbConn rdb.Conn) *validator.Validator {
	return validator.NewValidator(
		test_logger.NewFakeLogger(),
		rdbConn,
		"tcrocncl",
		nil,
		&validator.Config{
			validator.AttentionStatusRules{
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
			DefaultMaxActiveBlocksPeriodLimit,
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
						OperatorAddress:         "tcrovaloper1v76r7u4uyr3ewdks8cqmuw7ca4lejvc8unag2m",
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
						MaybeAfterBlockTime:  &timestamp,
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
									"commissionRate":    nil,
									"validatorAddress":  "crcvaloper12luku6uxehhak02py4rcz65zu0swh7wj6ulrlg",
									"minSelfDelegation": nil,
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
									"commissionRate":    nil,
									"validatorAddress":  "crcvaloper12luku6uxehhak02py4rcz65zu0swh7wj6ulrlg",
									"minSelfDelegation": nil,
								},
							},
						},
					}, nil, nil)

				mockValidatorsView.On(
					"Update",
					&view.ValidatorRow{
						MaybeId:                 primptr.Int64(3),
						OperatorAddress:         "tcrovaloper1v76r7u4uyr3ewdks8cqmuw7ca4lejvc8unag2m",
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
					testify_mock.Anything,
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
					testify_mock.Anything,
					testify_mock.Anything,
				).Return(
					[]validator_view.ValidatorRow{}, nil)

				mockValidatorBlockCommitmentsTotalView := validator_view.NewMockValidatorBlockCommitmentsTotal(nil).(*validator_view.MockValidatorBlockCommitmentsTotalView)
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
					[]view.ValidatorRow(nil),
					[]view.ValidatorRow(nil),
					int64(1),
					int64(100000),
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
						OperatorAddress:         "tcrovaloper1v76r7u4uyr3ewdks8cqmuw7ca4lejvc8unag2m",
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
						MaybeAfterBlockTime:  &timestamp,
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
						OperatorAddress:         "tcrovaloper1v76r7u4uyr3ewdks8cqmuw7ca4lejvc8unag2m",
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
					testify_mock.Anything,
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
					testify_mock.Anything,
					testify_mock.Anything,
				).Return(
					[]validator_view.ValidatorRow{}, nil)

				mockValidatorBlockCommitmentsTotalView := validator_view.NewMockValidatorBlockCommitmentsTotal(nil).(*validator_view.MockValidatorBlockCommitmentsTotalView)
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
					[]view.ValidatorRow(nil),
					[]view.ValidatorRow(nil),
					int64(1),
					int64(100000),
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
						OperatorAddress:         "tcrovaloper1v76r7u4uyr3ewdks8cqmuw7ca4lejvc8unag2m",
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
						MaybeAfterBlockTime:  &timestamp,
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
						OperatorAddress:         "tcrovaloper1v76r7u4uyr3ewdks8cqmuw7ca4lejvc8unag2m",
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
					testify_mock.Anything,
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
					testify_mock.Anything,
					testify_mock.Anything,
				).Return(
					[]validator_view.ValidatorRow{}, nil)

				mockValidatorBlockCommitmentsTotalView := validator_view.NewMockValidatorBlockCommitmentsTotal(nil).(*validator_view.MockValidatorBlockCommitmentsTotalView)
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
					[]view.ValidatorRow(nil),
					[]view.ValidatorRow(nil),
					int64(1),
					int64(100000),
				).Return(nil)

				return mocks
			},
		},
		{
			Name: "Handle MsgEditValidator status should be Bonded",
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
						Website:         "newWebsite",
						SecurityContact: "[do-not-modify]",
						Details:         "[do-not-modify]",
					},
					ValidatorAddress:       "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
					MaybeCommissionRate:    primptr.String("0.1"),
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
						OperatorAddress:         "tcrovaloper1v76r7u4uyr3ewdks8cqmuw7ca4lejvc8unag2m",
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
						MaybeAfterBlockTime:  &timestamp,
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
										"website":         "newWebsite",
										"identity":        "[do-not-modify]",
										"securityContact": "[do-not-modify]",
									},
									"commissionRate":    nil,
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
										"website":         "newWebsite",
										"identity":        "[do-not-modify]",
										"securityContact": "[do-not-modify]",
									},
									"commissionRate":    nil,
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
						OperatorAddress:         "tcrovaloper1v76r7u4uyr3ewdks8cqmuw7ca4lejvc8unag2m",
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
						Website:                 "newWebsite",
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
					testify_mock.Anything,
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
					testify_mock.Anything,
					testify_mock.Anything,
				).Return(
					[]validator_view.ValidatorRow{}, nil)

				mockValidatorBlockCommitmentsTotalView := validator_view.NewMockValidatorBlockCommitmentsTotal(nil).(*validator_view.MockValidatorBlockCommitmentsTotalView)
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
					[]view.ValidatorRow(nil),
					[]view.ValidatorRow(nil),
					int64(1),
					int64(100000),
				).Return(nil)

				return mocks
			},
		},
		{
			Name: "Handle validators active blocks up time",
			Events: []event_entity.Event{
				&event_usecase.BlockCreated{
					Base: event_entity.NewBase(event_entity.BaseParams{
						Name:        event_usecase.BLOCK_CREATED,
						Version:     1,
						BlockHeight: 11,
					}),
					Block: &model_usecase.Block{
						Height:          11,
						Hash:            "Hash",
						Time:            utctime.FromUnixNano(1665902976672205000),
						AppHash:         "AppHash",
						ProposerAddress: "ProposerAddress",
						Txs:             nil,
						Signatures: []model.BlockSignature{
							{
								BlockIdFlag:      2,
								ValidatorAddress: "F9E6FFB9B536956201AA138224FD888D03775AB4",
								Timestamp:        utctime.FromUnixNano(int64(1000000)),
								Signature:        "ZW2pUcKFN/oPQCmdCouchXmgpPyd/Ddo45dhHEMwsBeHTBuSJh15zUMmfl5FZsPHeKC8citFvOm/52bgl5XHCw==",
							},
							{
								BlockIdFlag:      2,
								ValidatorAddress: "031E3891DDB94FC7C7C132B7CD9736738110C889",
								Timestamp:        utctime.FromUnixNano(int64(2000000)),
								Signature:        "uhWDC9NDT86FbRVGbOM2lGY8sVkWU51JJ9F8gPwTfK0ebcui1R34oM+jhPKdStn/4sq4qDgzbsN66cQ5kl8NAw==",
							},
						},
					},
				},
			},
			MockFunc: func(events []event_entity.Event) (mocks []*testify_mock.Mock) {
				mockValidatorsView := validator_view.NewMockValidatorsView(nil).(*validator_view.MockValidatorsView)
				mocks = append(mocks, &mockValidatorsView.Mock)

				validator.NewValidators = func(_ *rdb.Handle) validator_view.Validators {
					return mockValidatorsView
				}

				mockValidatorsView.On(
					"ListAll",
					view.ValidatorsListFilter{
						MaybeStatuses: nil,
					},
					view.ValidatorsListOrder{MaybePower: nil},
				).Return(
					[]validator_view.ValidatorRow{
						{
							MaybeId:                 nil,
							OperatorAddress:         "OperatorAddress1",
							ConsensusNodeAddress:    "ConsensusNodeAddress1",
							InitialDelegatorAddress: "InitialDelegatorAddress1",
							TendermintPubkey:        "TendermintPubkey1",
							TendermintAddress:       "F9E6FFB9B536956201AA138224FD888D03775AB4",
							Status:                  "Bonded",
							Jailed:                  false,
							JoinedAtBlockHeight:     1,
							Power:                   "",
							Moniker:                 "",
							Identity:                "",
							Website:                 "",
							SecurityContact:         "",
							Details:                 "",
							CommissionRate:          "",
							CommissionMaxRate:       "",
							CommissionMaxChangeRate: "",
							MinSelfDelegation:       "",
							TotalSignedBlock:        0,
							TotalActiveBlock:        0,
							ImpreciseUpTime:         big.NewFloat(0),
							VotedGovProposal:        nil,
							RecentActiveBlocks:      []int64{},
							TotalRecentActiveBlocks: 0,
						},
						{
							MaybeId:                 nil,
							OperatorAddress:         "OperatorAddress2",
							ConsensusNodeAddress:    "ConsensusNodeAddress2",
							InitialDelegatorAddress: "InitialDelegatorAddress2",
							TendermintPubkey:        "TendermintPubkey2",
							TendermintAddress:       "031E3891DDB94FC7C7C132B7CD9736738110C889",
							Status:                  "Bonded",
							Jailed:                  false,
							JoinedAtBlockHeight:     1,
							Power:                   "",
							Moniker:                 "",
							Identity:                "",
							Website:                 "",
							SecurityContact:         "",
							Details:                 "",
							CommissionRate:          "",
							CommissionMaxRate:       "",
							CommissionMaxChangeRate: "",
							MinSelfDelegation:       "",
							TotalSignedBlock:        0,
							TotalActiveBlock:        0,
							ImpreciseUpTime:         big.NewFloat(0),
							VotedGovProposal:        nil,
							RecentActiveBlocks:      []int64{},
							TotalRecentActiveBlocks: 0,
						},
						{
							MaybeId:                 nil,
							OperatorAddress:         "OperatorAddress3",
							ConsensusNodeAddress:    "ConsensusNodeAddress3",
							InitialDelegatorAddress: "InitialDelegatorAddress3",
							TendermintPubkey:        "TendermintPubkey3",
							TendermintAddress:       "TendermintAddress3",
							Status:                  "Bonded",
							Jailed:                  false,
							JoinedAtBlockHeight:     2,
							Power:                   "",
							Moniker:                 "",
							Identity:                "",
							Website:                 "",
							SecurityContact:         "",
							Details:                 "",
							CommissionRate:          "",
							CommissionMaxRate:       "",
							CommissionMaxChangeRate: "",
							MinSelfDelegation:       "",
							TotalSignedBlock:        0,
							TotalActiveBlock:        0,
							ImpreciseUpTime:         big.NewFloat(0),
							VotedGovProposal:        nil,
							RecentActiveBlocks:      []int64{},
							TotalRecentActiveBlocks: 0,
						},
						{
							MaybeId:                 nil,
							OperatorAddress:         "OperatorAddress4",
							ConsensusNodeAddress:    "ConsensusNodeAddress4",
							InitialDelegatorAddress: "InitialDelegatorAddress4",
							TendermintPubkey:        "TendermintPubkey4",
							TendermintAddress:       "TendermintAddress4",
							Status:                  "Bonded",
							Jailed:                  false,
							JoinedAtBlockHeight:     1,
							Power:                   "",
							Moniker:                 "",
							Identity:                "",
							Website:                 "",
							SecurityContact:         "",
							Details:                 "",
							CommissionRate:          "",
							CommissionMaxRate:       "",
							CommissionMaxChangeRate: "",
							MinSelfDelegation:       "",
							TotalSignedBlock:        0,
							TotalActiveBlock:        0,
							ImpreciseUpTime:         big.NewFloat(0),
							VotedGovProposal:        nil,
							RecentActiveBlocks:      []int64{},
							TotalRecentActiveBlocks: 0,
						},
						{
							MaybeId:                 nil,
							OperatorAddress:         "OperatorAddress5",
							ConsensusNodeAddress:    "ConsensusNodeAddress5",
							InitialDelegatorAddress: "InitialDelegatorAddress5",
							TendermintPubkey:        "TendermintPubkey5",
							TendermintAddress:       "TendermintAddress5",
							Status:                  "Unbonded",
							Jailed:                  false,
							JoinedAtBlockHeight:     1,
							Power:                   "",
							Moniker:                 "",
							Identity:                "",
							Website:                 "",
							SecurityContact:         "",
							Details:                 "",
							CommissionRate:          "",
							CommissionMaxRate:       "",
							CommissionMaxChangeRate: "",
							MinSelfDelegation:       "",
							TotalSignedBlock:        0,
							TotalActiveBlock:        0,
							ImpreciseUpTime:         big.NewFloat(0),
							VotedGovProposal:        nil,
							RecentActiveBlocks:      []int64{},
							TotalRecentActiveBlocks: 0,
						},
					}, nil)

				mockValidatorBlockCommitmentsView := validator_view.NewMockValidatorBlockCommitmentsView(nil).(*validator_view.MockValidatorBlockCommitmentsView)
				mocks = append(mocks, &mockValidatorBlockCommitmentsView.Mock)

				validator.NewValidatorBlockCommitments = func(_ *rdb.Handle) validator_view.ValidatorBlockCommitments {
					return mockValidatorBlockCommitmentsView
				}

				mockValidatorBlockCommitmentsTotalView := validator_view.NewMockValidatorBlockCommitmentsTotal(nil).(*validator_view.MockValidatorBlockCommitmentsTotalView)
				mocks = append(mocks, &mockValidatorBlockCommitmentsTotalView.Mock)

				validator.NewValidatorBlockCommitmentsTotal = func(_ *rdb.Handle) validator_view.ValidatorBlockCommitmentsTotal {
					return mockValidatorBlockCommitmentsTotalView
				}

				mockValidatorBlockCommitmentsView.On(
					"InsertAll",
					[]view.ValidatorBlockCommitmentRow{
						{
							MaybeId:              nil,
							ConsensusNodeAddress: "ConsensusNodeAddress1",
							BlockHeight:          1,
							IsProposer:           false,
							Signature:            "ZW2pUcKFN/oPQCmdCouchXmgpPyd/Ddo45dhHEMwsBeHTBuSJh15zUMmfl5FZsPHeKC8citFvOm/52bgl5XHCw==",
							Timestamp:            utctime.FromUnixNano(int64(1000000)),
						},
						{
							MaybeId:              nil,
							ConsensusNodeAddress: "ConsensusNodeAddress2",
							BlockHeight:          1,
							IsProposer:           false,
							Signature:            "uhWDC9NDT86FbRVGbOM2lGY8sVkWU51JJ9F8gPwTfK0ebcui1R34oM+jhPKdStn/4sq4qDgzbsN66cQ5kl8NAw==",
							Timestamp:            utctime.FromUnixNano(int64(2000000)),
						},
					},
				).Return(nil).Once()

				mockValidatorBlockCommitmentsTotalView.On(
					"Set",
					"1:-",
					int64(2),
				).Return(nil)

				mockValidatorBlockCommitmentsTotalView.On(
					"IncrementAll",
					[]string{"-:ConsensusNodeAddress1", "-:ConsensusNodeAddress2"},
					int64(1),
				).Return(nil)

				mockValidatorBlockCommitmentsTotalView.On(
					"IncrementAll",
					[]string{"1:ConsensusNodeAddress1", "1:ConsensusNodeAddress2"},
					int64(1),
				).Return(nil)

				mockValidatorBlockCommitmentsTotalView.On(
					"Increment",
					"-:-",
					int64(2),
				).Return(nil)

				mockValidatorsView.On(
					"UpdateAllValidatorUpTime",
					[]view.ValidatorRow{
						{
							MaybeId:                 nil,
							OperatorAddress:         "OperatorAddress1",
							ConsensusNodeAddress:    "ConsensusNodeAddress1",
							InitialDelegatorAddress: "InitialDelegatorAddress1",
							TendermintPubkey:        "TendermintPubkey1",
							TendermintAddress:       "F9E6FFB9B536956201AA138224FD888D03775AB4",
							Status:                  "Bonded",
							Jailed:                  false,
							JoinedAtBlockHeight:     1,
							Power:                   "",
							Moniker:                 "",
							Identity:                "",
							Website:                 "",
							SecurityContact:         "",
							Details:                 "",
							CommissionRate:          "",
							CommissionMaxRate:       "",
							CommissionMaxChangeRate: "",
							MinSelfDelegation:       "",
							TotalSignedBlock:        1,
							TotalActiveBlock:        1,
							ImpreciseUpTime:         big.NewFloat(1),
							VotedGovProposal:        nil,
							RecentActiveBlocks:      []int64{},
							TotalRecentActiveBlocks: 0,
						},
						{
							MaybeId:                 nil,
							OperatorAddress:         "OperatorAddress2",
							ConsensusNodeAddress:    "ConsensusNodeAddress2",
							InitialDelegatorAddress: "InitialDelegatorAddress2",
							TendermintPubkey:        "TendermintPubkey2",
							TendermintAddress:       "031E3891DDB94FC7C7C132B7CD9736738110C889",
							Status:                  "Bonded",
							Jailed:                  false,
							JoinedAtBlockHeight:     1,
							Power:                   "",
							Moniker:                 "",
							Identity:                "",
							Website:                 "",
							SecurityContact:         "",
							Details:                 "",
							CommissionRate:          "",
							CommissionMaxRate:       "",
							CommissionMaxChangeRate: "",
							MinSelfDelegation:       "",
							TotalSignedBlock:        1,
							TotalActiveBlock:        1,
							ImpreciseUpTime:         big.NewFloat(1),
							VotedGovProposal:        nil,
							RecentActiveBlocks:      []int64{},
							TotalRecentActiveBlocks: 0,
						},
						{
							MaybeId:                 nil,
							OperatorAddress:         "OperatorAddress3",
							ConsensusNodeAddress:    "ConsensusNodeAddress3",
							InitialDelegatorAddress: "InitialDelegatorAddress3",
							TendermintPubkey:        "TendermintPubkey3",
							TendermintAddress:       "TendermintAddress3",
							Status:                  "Bonded",
							Jailed:                  false,
							JoinedAtBlockHeight:     2,
							Power:                   "",
							Moniker:                 "",
							Identity:                "",
							Website:                 "",
							SecurityContact:         "",
							Details:                 "",
							CommissionRate:          "",
							CommissionMaxRate:       "",
							CommissionMaxChangeRate: "",
							MinSelfDelegation:       "",
							TotalSignedBlock:        1,
							TotalActiveBlock:        1,
							ImpreciseUpTime:         big.NewFloat(1),
							VotedGovProposal:        nil,
							RecentActiveBlocks:      []int64{},
							TotalRecentActiveBlocks: 0,
						},
					},
					[]view.ValidatorRow{
						{
							MaybeId:                 nil,
							OperatorAddress:         "OperatorAddress4",
							ConsensusNodeAddress:    "ConsensusNodeAddress4",
							InitialDelegatorAddress: "InitialDelegatorAddress4",
							TendermintPubkey:        "TendermintPubkey4",
							TendermintAddress:       "TendermintAddress4",
							Status:                  "Bonded",
							Jailed:                  false,
							JoinedAtBlockHeight:     1,
							Power:                   "",
							Moniker:                 "",
							Identity:                "",
							Website:                 "",
							SecurityContact:         "",
							Details:                 "",
							CommissionRate:          "",
							CommissionMaxRate:       "",
							CommissionMaxChangeRate: "",
							MinSelfDelegation:       "",
							TotalSignedBlock:        0,
							TotalActiveBlock:        1,
							ImpreciseUpTime:         big.NewFloat(0),
							VotedGovProposal:        nil,
							RecentActiveBlocks:      []int64{},
							TotalRecentActiveBlocks: 0,
						},
					},
					int64(11),
					int64(100000),
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
