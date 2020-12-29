package crossfire_test

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/crossfire"
	"github.com/crypto-com/chain-indexing/appinterface/projection/crossfire/view"
	. "github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	entity_projection "github.com/crypto-com/chain-indexing/entity/projection"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	. "github.com/crypto-com/chain-indexing/internal/logger/test"
	. "github.com/crypto-com/chain-indexing/test"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	usecase_model "github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Crossfire", func() {
	It("should implement projection", func() {
		fakeLogger := NewFakeLogger()
		fakeRdbConn := NewFakeRDbConn()
		var _ entity_projection.Projection = crossfire.NewCrossfire(fakeLogger, fakeRdbConn, "tcrocnclcons")
	})

	WithTestPgxConn(func(pgConn *pg.PgxConn, pgMigrate *pg.Migrate) {
		BeforeEach(func() {
			_ = pgMigrate.Reset()
			pgMigrate.MustUp()
		})

		AfterEach(func() {
			_ = pgMigrate.Reset()
		})

		It("should project crossfire_validators view when event is MsgCreateValidator", func() {
			crossfireValidatorsView := view.NewCrossfireValidators(pgConn.ToHandle())

			anyHeight := int64(1)
			anyCoin, _ := coin.NewCoinFromInt(1)
			event := event_usecase.NewMsgCreateValidator(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      "A6D4C1F59A9D232747CA4F8A484F1F3B14A0075E801DF2A25F472B4280505B74",
				TxSuccess:   true,
				MsgIndex:    0,
			}, usecase_model.MsgCreateValidatorParams{
				Description: usecase_model.MsgValidatorDescription{
					Moniker:         "Jail_Testing",
					Identity:        "foo",
					Website:         "www.example.com",
					SecurityContact: "foo@example.com",
					Details:         "example",
				},
				Commission: usecase_model.MsgValidatorCommission{
					Rate:          "0.100000000000000000",
					MaxRate:       "0.200000000000000000",
					MaxChangeRate: "0.010000000000000000",
				},
				MinSelfDelegation: "1",
				DelegatorAddress:  "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
				ValidatorAddress:  "tcrocncl14m5a4kxt2e82uqqs5gtqza29dm5wqzyalddug5",
				Pubkey:            "tcrocnclconspub1zcjduepq673fv82rsss7um343qems34uxjquwealhrntcq4naj83kk9d8syqdswgxv",
				Amount:            anyCoin,
			})

			fakeLogger := NewFakeLogger()
			projection := crossfire.NewCrossfire(fakeLogger, pgConn, "tcrocnclcons")

			Expect(crossfireValidatorsView.Count()).To(Equal(int64(0)))

			// insert one
			err := projection.HandleEvents(anyHeight, []event_entity.Event{event})
			Expect(err).To(BeNil())

			count, err := crossfireValidatorsView.Count()
			Expect(err).To(BeNil())
			Expect(count).To(Equal(int64(1)))

			// list all
			list, err := crossfireValidatorsView.List()
			Expect(err).To(BeNil())
			var targetList []view.CrossfireValidatorRow
			id := int64(1)
			targetList = append(targetList, view.CrossfireValidatorRow{
				MaybeId:                             &id,
				OperatorAddress:                     "tcrocncl14m5a4kxt2e82uqqs5gtqza29dm5wqzyalddug5",
				ConsensusNodeAddress:                "tcrocnclcons1u4jfqxk5femyyt0s5s55xuywv8ehnu34gcuaad",
				InitialDelegatorAddress:             "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
				Status:                              "Unbonded",
				Jailed:                              false,
				JoinedAtBlockHeight:                 1,
				Moniker:                             "Jail_Testing",
				Identity:                            "foo",
				Website:                             "www.example.com",
				SecurityContact:                     "foo@example.com",
				Details:                             "example",
				Phase1TaskNodeSetup:                 "Incompleted",
				Phase1TaskBlockValidCommit:          "Incompleted",
				Phase2TaskKeepNodeActive:            "Incompleted",
				Phase2TaskProposalVote:              "Incompleted",
				Phase2TaskNetworkUpgrade:            "Incompleted",
				Phase2TaskNetworkUpgradeBlockCommit: "Incompleted",
				Phase1n2TaskCommitmentCountRank:     0,
				Phase3TaskCommitmentCountRank:       0,
				TaskHighestTxSentRank:               0,
			})
			Expect(list).To(Equal(targetList))
		})
	})
})
