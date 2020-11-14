package rdbstatusstore_test

import (
	"github.com/crypto-com/chainindex/appinterface/rdbstatusstore"
	"github.com/crypto-com/chainindex/infrastructure/pg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/crypto-com/chainindex/test"
)

var _ = Describe("RdbStatusStore", func() {
	WithTestPgxConn(func(pgxConn *pg.PgxConn, pgMigrate *pg.Migrate) {
		BeforeEach(func() {
			_ = pgMigrate.Reset()
			pgMigrate.MustUp()
		})

		AfterEach(func() {
			_ = pgMigrate.Reset()
		})

		Describe("GetLastIndexedBlockHeight", func() {
			It("should return the initial last_indexed_block_height 1 with no error", func() {
				store := rdbstatusstore.NewRDbStatusStoreImpl(pgxConn.ToHandle())
				height, err := store.GetLastIndexedBlockHeight()

				Expect(err).To(BeNil())
				Expect(height).To(Equal(int64(1)))
			})
		})

		Describe("UpdateLastIndexedBlockHeight", func() {
			It("should return the updated last indexed block height properly with no error", func() {
				store := rdbstatusstore.NewRDbStatusStoreImpl(pgxConn.ToHandle())

				newHeight := int64(100)
				err := store.UpdateLastIndexedBlockHeight(newHeight)
				Expect(err).To(BeNil())

				currentHeight, err := store.GetLastIndexedBlockHeight()
				Expect(err).To(BeNil())
				Expect(currentHeight).To(Equal(int64(100)))
			})
		})
	})
})
