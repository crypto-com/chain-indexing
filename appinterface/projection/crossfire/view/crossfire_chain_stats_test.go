package view_test

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/crossfire/constants"
	"github.com/crypto-com/chain-indexing/appinterface/projection/crossfire/view"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	. "github.com/crypto-com/chain-indexing/test"
)

var _ = Describe("Crossfire Chain Stats", func() {
	WithTestPgxConn(func(conn *pg.PgxConn, migrate *pg.Migrate) {
		BeforeEach(func() {
			_ = migrate.Reset()
			migrate.MustUp()
		})

		AfterEach(func() {
			migrate.MustReset()
		})

		Describe("Increment", func() {
			It("should run increment correctly", func() {
				var err error

				key := constants.PHASE1_BLOCK_COUNT
				crossfireChainStatsView := view.NewCrossfireChainStats(conn.ToHandle())

				// increment
				err = crossfireChainStatsView.Increment(key)
				Expect(err).To(BeNil())
				// should be 1
				actual, err := crossfireChainStatsView.FindBy(key)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(int64(1)))

				// increment
				err = crossfireChainStatsView.Increment(key)
				Expect(err).To(BeNil())
				// should be 2
				actual, err = crossfireChainStatsView.FindBy(key)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(int64(2)))
			})
		})
	})
})
