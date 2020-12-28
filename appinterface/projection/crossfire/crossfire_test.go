package crossfire_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

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

			err := projection.HandleEvents(anyHeight, []event_entity.Event{event})
			Expect(err).To(BeNil())
			Expect(crossfireValidatorsView.Count()).To(Equal(int64(1)))
		})
	})
})
