package view_test

import (
	random "github.com/brianvoe/gofakeit/v5"
	"github.com/crypto-com/chain-indexing/projection/block/view"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	. "github.com/crypto-com/chain-indexing/test"
)

var _ = Describe("Blocks", func() {
	WithTestPgxConn(func(conn *pg.PgxConn, migrate *pg.Migrate) {
		BeforeEach(func() {
			_ = migrate.Reset()
			migrate.MustUp()
		})

		AfterEach(func() {
			_ = migrate.Reset()
		})

		Describe("Insert", func() {
			It("should insert block into view", func() {
				var err error

				var block view.Block
				random.Struct(&block)

				blocksView := view.NewBlocks(conn.ToHandle())

				Expect(blocksView.Count()).To(Equal(int64(0)))

				err = blocksView.Insert(&block)
				Expect(err).To(BeNil())

				var actual *view.Block
				actual, err = blocksView.FindBy(&view.BlockIdentity{
					MaybeHeight: &block.Height,
				})
				Expect(err).To(BeNil())
				Expect(*actual).To(Equal(block))
			})
		})
	})
})
