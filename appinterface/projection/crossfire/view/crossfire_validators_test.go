package view_test

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/crossfire/view"
	"github.com/crypto-com/chain-indexing/internal/utctime"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	. "github.com/crypto-com/chain-indexing/test"
)

var _ = Describe("Crossfire Validators", func() {
	WithTestPgxConn(func(conn *pg.PgxConn, migrate *pg.Migrate) {
		BeforeEach(func() {
			_ = migrate.Reset()
			migrate.MustUp()
		})

		AfterEach(func() {
			migrate.MustReset()
		})

		Describe("Insert & FindBy", func() {
			It("should insert validator into view and find the target validator", func() {
				var err error

				id := int64(1)
				var validator = view.CrossfireValidatorRow{
					MaybeId:                         &id,
					OperatorAddress:                 "tcrocncl14m5a4kxt2e82uqqs5gtqza29dm5wqzyalddug5",
					ConsensusNodeAddress:            "tcrocnclcons1khkxmphc7sv0fqrej3rltsslrstud78cam9ekl",
					InitialDelegatorAddress:         "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
					TendermintPubkey:                "na51D8RmKXyWrid9I6wtdxgP6f1Nl3EyNNEzqxVquoM=",
					TendermintAddress:               "B5EC6D86F8F418F480799447F5C21F1C17C6F8F8",
					Status:                          "Unbonded",
					Jailed:                          false,
					JoinedAtBlockHeight:             1,
					JoinedAtBlockTime:               utctime.FromUnixNano(int64(1000000)),
					Moniker:                         "Testing",
					Identity:                        "foo",
					Website:                         "www.example.com",
					SecurityContact:                 "foo@example.com",
					Details:                         "example",
					TaskPhase1NodeSetup:             "Completed",
					TaskPhase2KeepNodeActive:        "Incompleted",
					TaskPhase2ProposalVote:          "Incompleted",
					TaskPhase2NetworkUpgrade:        "Incompleted",
					RankTaskPhase1n2CommitmentCount: 0,
					RankTaskPhase3CommitmentCount:   0,
					RankTaskHighestTxSent:           0,
				}

				crossfireValidatorsView := view.NewCrossfireValidators(conn.ToHandle())

				err = crossfireValidatorsView.Upsert(&validator)
				Expect(err).To(BeNil())

				var actual *view.CrossfireValidatorRow
				actual, err = crossfireValidatorsView.FindBy(view.CrossfireValidatorIdentity{
					MaybeConsensusNodeAddress: &validator.ConsensusNodeAddress,
				})
				Expect(err).To(BeNil())
				Expect(*actual).To(Equal(validator))
			})
		})

		Describe("LastJoinedBlockHeight", func() {
			It("should insert validator into view and get the last joined info", func() {
				var err error

				id := int64(1)
				fakeBlockHeight := int64(123)
				fakeBlockTime := utctime.FromUnixNano(int64(1000000))
				var validator = view.CrossfireValidatorRow{
					MaybeId:                         &id,
					OperatorAddress:                 "tcrocncl14m5a4kxt2e82uqqs5gtqza29dm5wqzyalddug5",
					ConsensusNodeAddress:            "tcrocnclcons1khkxmphc7sv0fqrej3rltsslrstud78cam9ekl",
					InitialDelegatorAddress:         "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
					TendermintPubkey:                "na51D8RmKXyWrid9I6wtdxgP6f1Nl3EyNNEzqxVquoM=",
					TendermintAddress:               "B5EC6D86F8F418F480799447F5C21F1C17C6F8F8",
					Status:                          "Unbonded",
					Jailed:                          false,
					JoinedAtBlockHeight:             fakeBlockHeight,
					JoinedAtBlockTime:               fakeBlockTime,
					Moniker:                         "Testing",
					Identity:                        "foo",
					Website:                         "www.example.com",
					SecurityContact:                 "foo@example.com",
					Details:                         "example",
					TaskPhase1NodeSetup:             "Completed",
					TaskPhase2KeepNodeActive:        "Incompleted",
					TaskPhase2ProposalVote:          "Incompleted",
					TaskPhase2NetworkUpgrade:        "Incompleted",
					RankTaskPhase1n2CommitmentCount: 0,
					RankTaskPhase3CommitmentCount:   0,
					RankTaskHighestTxSent:           0,
				}

				crossfireValidatorsView := view.NewCrossfireValidators(conn.ToHandle())

				err = crossfireValidatorsView.Upsert(&validator)
				Expect(err).To(BeNil())

				_, height, time, err := crossfireValidatorsView.LastJoinedBlockHeight(validator.OperatorAddress, validator.ConsensusNodeAddress)
				Expect(err).To(BeNil())
				Expect(height).To(Equal(fakeBlockHeight))
				Expect(time).To(Equal(fakeBlockTime))
			})
		})

		Describe("Insert & List", func() {
			It("should insert validator into view and list the validators", func() {
				var err error

				id := int64(1)
				var validator = view.CrossfireValidatorRow{
					MaybeId:                         &id,
					OperatorAddress:                 "tcrocncl14m5a4kxt2e82uqqs5gtqza29dm5wqzyalddug5",
					ConsensusNodeAddress:            "tcrocnclcons1khkxmphc7sv0fqrej3rltsslrstud78cam9ekl",
					InitialDelegatorAddress:         "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
					TendermintPubkey:                "na51D8RmKXyWrid9I6wtdxgP6f1Nl3EyNNEzqxVquoM=",
					TendermintAddress:               "B5EC6D86F8F418F480799447F5C21F1C17C6F8F8",
					Status:                          "Unbonded",
					Jailed:                          false,
					JoinedAtBlockHeight:             1,
					JoinedAtBlockTime:               utctime.FromUnixNano(int64(1000000)),
					Moniker:                         "Testing",
					Identity:                        "foo",
					Website:                         "www.example.com",
					SecurityContact:                 "foo@example.com",
					Details:                         "example",
					TaskPhase1NodeSetup:             "Completed",
					TaskPhase2KeepNodeActive:        "Incompleted",
					TaskPhase2ProposalVote:          "Incompleted",
					TaskPhase2NetworkUpgrade:        "Incompleted",
					RankTaskPhase1n2CommitmentCount: 0,
					RankTaskPhase3CommitmentCount:   0,
					RankTaskHighestTxSent:           0,
				}

				crossfireValidatorsView := view.NewCrossfireValidators(conn.ToHandle())

				err = crossfireValidatorsView.Upsert(&validator)
				Expect(err).To(BeNil())

				list, err := crossfireValidatorsView.List()
				Expect(err).To(BeNil())
				var targetList []view.CrossfireValidatorRow
				targetList = append(targetList, view.CrossfireValidatorRow{
					MaybeId:                         &id,
					OperatorAddress:                 "tcrocncl14m5a4kxt2e82uqqs5gtqza29dm5wqzyalddug5",
					ConsensusNodeAddress:            "tcrocnclcons1khkxmphc7sv0fqrej3rltsslrstud78cam9ekl",
					InitialDelegatorAddress:         "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
					TendermintPubkey:                "na51D8RmKXyWrid9I6wtdxgP6f1Nl3EyNNEzqxVquoM=",
					TendermintAddress:               "B5EC6D86F8F418F480799447F5C21F1C17C6F8F8",
					Status:                          "Unbonded",
					Jailed:                          false,
					JoinedAtBlockHeight:             1,
					JoinedAtBlockTime:               utctime.FromUnixNano(int64(1000000)),
					Moniker:                         "Testing",
					Identity:                        "foo",
					Website:                         "www.example.com",
					SecurityContact:                 "foo@example.com",
					Details:                         "example",
					TaskPhase1NodeSetup:             "Completed",
					TaskPhase2KeepNodeActive:        "Incompleted",
					TaskPhase2ProposalVote:          "Incompleted",
					TaskPhase2NetworkUpgrade:        "Incompleted",
					RankTaskPhase1n2CommitmentCount: 0,
					RankTaskPhase3CommitmentCount:   0,
					RankTaskHighestTxSent:           0,
				})
				Expect(list).To(Equal(targetList))
			})
		})
	})
})
