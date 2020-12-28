package crossfire_test

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/crossfire"
	. "github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	entity_projection "github.com/crypto-com/chain-indexing/entity/projection"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	. "github.com/crypto-com/chain-indexing/internal/logger/test"
	. "github.com/crypto-com/chain-indexing/test"
	. "github.com/onsi/ginkgo"
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
		})
	})
})
