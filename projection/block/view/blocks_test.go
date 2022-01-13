package view_test

import (
	"path"
	"runtime"

	random "github.com/brianvoe/gofakeit/v5"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/projection/block/view"
	. "github.com/crypto-com/chain-indexing/test"
)

var BLOCK_MIGRATIONS_PATH = func() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("error retrieving file directory")
	}

	return path.Join(filename, "../../migrations")
}()

var _ = Describe("Blocks", func() {
	WithProjectionTestEnv(func(testEnv ProjectionTestEnv) {
		pgxConn := testEnv.Conn

		reset := func() {
			// Only needs to run Reset() on one of the migrate instance because Reset() drops everything in the Database
			_ = testEnv.RootMigrate.Reset()
		}

		BeforeEach(func() {
			reset()

			testEnv.RootMigrate.MustUp()
			blockMigrate := testEnv.MigrateCreator(
				"block_schema_migrations",
				BLOCK_MIGRATIONS_PATH,
			)
			blockMigrate.MustUp()
		})

		AfterEach(func() {
			reset()
		})

		Describe("Insert", func() {
			It("should insert block into view", func() {
				var err error

				var block view.Block
				random.Struct(&block)

				blocksView := view.NewBlocks(pgxConn.ToHandle())

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
