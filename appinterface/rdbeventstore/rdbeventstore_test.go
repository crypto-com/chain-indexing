package rdbeventstore_test

import (
	"github.com/crypto-com/chainindex/appinterface/rdbeventstore"
	"github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/entity/event/test"
	"github.com/crypto-com/chainindex/infrastructure/pg"
	"github.com/crypto-com/chainindex/internal/primptr"
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

		Describe("Insert", func() {
			It("should insert an event properly without any error", func() {
				store := rdbeventstore.NewRDbEventStore(pgxConn.ToHandle())

				evt := test.NewFakeEvent()
				err := store.Insert(evt)
				Expect(err).To(BeNil())
			})
		})

		Describe("InsertAll", func() {
			It("should insert multiple events properly without any error", func() {
				store := rdbeventstore.NewRDbEventStore(pgxConn.ToHandle())

				events := []event.Event{test.NewFakeEvent(), test.NewFakeEvent()}
				err := store.InsertAll(events)
				Expect(err).To(BeNil())
			})
		})

		Describe("GetLatestHeight", func() {
			It("should return nil when events table does not have any record", func() {
				store := rdbeventstore.NewRDbEventStore(pgxConn.ToHandle())

				emptyEventsTableHeight, err := store.GetLatestHeight()
				Expect(err).To(BeNil())
				Expect(emptyEventsTableHeight).To(Equal(primptr.Int64Nil()))
			})

			It("should get 1 after insert fake event with height 1", func() {
				store := rdbeventstore.NewRDbEventStore(pgxConn.ToHandle())

				// Insert one fake event with height 1
				evt := test.NewFakeEvent()
				err := store.Insert(evt)
				Expect(err).To(BeNil())

				// Get the height
				height, err := store.GetLatestHeight()
				Expect(err).To(BeNil())
				Expect(height).To(Equal(primptr.Int64(1)))
			})
		})
	})
})
