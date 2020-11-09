package rdbeventstore_test

import (
	"github.com/crypto-com/chainindex/infrastructure/pg"
	. "github.com/crypto-com/chainindex/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RdbEventStore", func() {
	WithTestPgxConn(func(pgxConn *pg.PgxConn, pgMigrate *pg.Migrate) {
		BeforeEach(func() {
			_ = pgMigrate.Reset()
			pgMigrate.MustUp()
		})

		AfterEach(func() {
			_ = pgMigrate.Reset()
		})

		Describe("GetLatestHeight", func() {
			It("should get the last height record from the events table", func() {
				// TODO: unit test
				//store := rdbbase.NewRDbStoreImpl(rdbbase.DEFAULT_TABLE)
			})

			It("should pass!", func() {
				var err error = nil
				Expect(err).To(BeNil())
			})
		})

		It("should insert event and get latest event height", func() {
			// TODO: happy flow test
		})
	})
})
