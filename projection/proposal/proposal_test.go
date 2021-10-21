package proposal_test

import (
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
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/crypto-com/chain-indexing/internal/primptr"
	"github.com/crypto-com/chain-indexing/internal/utctime"
	"github.com/crypto-com/chain-indexing/projection/proposal"
	"github.com/crypto-com/chain-indexing/projection/proposal/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	usecase_event "github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

func NewProposalProjection(rdbConn rdb.Conn) *proposal.Proposal {

	return proposal.NewProposal(
		nil,
		rdbConn,
		"",
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
						OperatorAddress: "OperatorAddress",
					}, nil).
					Twice()

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
						MaybeProposerOperatorAddress: primptr.String("OperatorAddress"),
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
					On("Insert", &view.DepositorRow{
						ProposalId:                    "MaybeProposalId",
						DepositorAddress:              "ProposerAddress",
						MaybeDepositorOperatorAddress: primptr.String("OperatorAddress"),
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
