package validator_test

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"

	"github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	test_logger "github.com/crypto-com/chain-indexing/external/logger/test"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/projection/validator"
	validator_view "github.com/crypto-com/chain-indexing/projection/validator/view"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	model_usecase "github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/stretchr/testify/assert"
	testify_mock "github.com/stretchr/testify/mock"
)

func NewValidatorProjection(rdbConn rdb.Conn) *validator.Validator {
	return validator.NewValidator(
		test_logger.NewFakeLogger(),
		rdbConn,
		"crocnclcons",
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
func TestValidators_HandleEvents(t *testing.T) {
	testCases := []struct {
		Name     string
		Events   []event_entity.Event
		MockFunc func(events []event_entity.Event) []*testify_mock.Mock
	}{
		{
			Name: "Handle MsgEditValidator",
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
				fmt.Println("===> s1")

				mockValidatorsView.On("FindBy", validator_view.ValidatorIdentity{
					MaybeOperatorAddress: primptr.String("tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus"),
				}).Return(
					validator_view.ValidatorRow{
						// MaybeId:                 primptr.Int64(3),
						// OperatorAddress:         "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
						// ConsensusNodeAddress:    "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
						// InitialDelegatorAddress: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						// TendermintPubkey:        "Kpox5fS2po0sJUHmzllExuJ4uZ5nm0bbCp6UQKESsnE=",
						// TendermintAddress:       "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						// Status:                  "Attention",
						// Jailed:                  false,
						// JoinedAtBlockHeight:     1,
						// Power:                   "",
						// Moniker:                 "mymonicker",
						// Identity:                "myidentity",
						// Website:                 "mywebsite",
						// SecurityContact:         "mysecuritycontact",
						// Details:                 "mydetails",
						// CommissionRate:          "0.100000000000000000",
						// CommissionMaxRate:       "0.200000000000000000",
						// CommissionMaxChangeRate: "0.010000000000000000",
						// MinSelfDelegation:       "1",
						// TotalSignedBlock:        1,
						// TotalActiveBlock:        1,
						// ImpreciseUpTime:         big.NewFloat(0.1),
						// VotedGovProposal:        big.NewInt(1),
						// DailyEditQuota:          1,
					}, nil)

				// mockValidatorsView.On(
				// 				// 	"Upsert",
				// 				// 	[]validator_view.ValidatorRow{
				// 				// 		{
				// 				// 			// MaybeId:                 int64(3),
				// 				// 			OperatorAddress:         "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
				// 				// 			ConsensusNodeAddress:    "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
				// 				// 			InitialDelegatorAddress: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
				// 				// 			TendermintPubkey:        "Kpox5fS2po0sJUHmzllExuJ4uZ5nm0bbCp6UQKESsnE=",
				// 				// 			TendermintAddress:       "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
				// 				// 			Status:                  "Attention",
				// 				// 			Jailed:                  false,
				// 				// 			JoinedAtBlockHeight:     1,
				// 				// 			Power:                   "",
				// 				// 			Moniker:                 "mymonicker",
				// 				// 			Identity:                "myidentity",
				// 				// 			Website:                 "mywebsite",
				// 				// 			SecurityContact:         "mysecuritycontact",
				// 				// 			Details:                 "mydetails",
				// 				// 			CommissionRate:          "0.100000000000000000",
				// 				// 			CommissionMaxRate:       "0.200000000000000000",
				// 				// 			CommissionMaxChangeRate: "0.010000000000000000",
				// 				// 			MinSelfDelegation:       "1",
				// 				// 			TotalSignedBlock:        1,
				// 				// 			TotalActiveBlock:        1,
				// 				// 			// ImpreciseUpTime:         0.1,
				// 				// 			// VotedGovProposal:        1,
				// 				// 			DailyEditQuota: 1,
				// 				// 		},
				// 				// 	},
				// 				// ).Return(nil)

				return mocks
			},
		},
	}

	for _, tc := range testCases {
		mockRDbConn := NewMockRDbConn()
		mockTx := NewMockRDbTx()
		mockRDbConn.On("Begin").Return(mockTx, nil)
		mocks := tc.MockFunc(tc.Events)

		projection := NewValidatorProjection(mockRDbConn)
		err := projection.HandleEvents(1, tc.Events)

		assert.NoError(t, err)

		for _, m := range mocks {
			m.AssertExpectations(t)
		}

		fmt.Println(tc.Name, "Passed")
	}
}
