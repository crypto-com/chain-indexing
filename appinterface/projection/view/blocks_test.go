package view_test

import (
	random "github.com/brianvoe/gofakeit/v5"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/crypto-com/chainindex/appinterface/projection/view"
	"github.com/crypto-com/chainindex/infrastructure/pg"
	. "github.com/crypto-com/chainindex/test"
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

				var block Block
				random.Struct(&block)

				blocksView := NewBlocks(conn.ToHandle())

				Expect(blocksView.Count()).To(Equal(0))

				err = blocksView.Insert(&block)
				Expect(err).To(BeNil())

				var actual *Block
				actual, err = blocksView.FindBy(&BlockIdentity{
					MaybeHeight: &block.Height,
				})
				Expect(err).To(BeNil())
				Expect(*actual).To(Equal(block))
			})
		})
	})
})
