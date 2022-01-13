package event_test

import (
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/entity/event/test"
	"github.com/crypto-com/chain-indexing/external/primptr"
	. "github.com/crypto-com/chain-indexing/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	appinterface_event "github.com/crypto-com/chain-indexing/appinterface/event"
	"github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
)

var _ = Describe("RdbEventStore", func() {
	WithTestPgxConn(func(pgxConn *pg.PgxConn, migrate rdb.Migrate) {
		BeforeEach(func() {
			_ = migrate.Reset()
			migrate.MustUp()
		})

		AfterEach(func() {
			_ = migrate.Reset()
		})

		Describe("Insert", func() {
			It("should insert an event properly without any error", func() {
				registry := event.NewRegistry()
				store := appinterface_event.NewRDbStore(pgxConn.ToHandle(), registry)

				event := test.NewFakeEvent()
				err := store.Insert(event)
				Expect(err).To(BeNil())
			})
		})

		Describe("InsertAll", func() {
			It("should insert multiple events properly without any error", func() {
				registry := event.NewRegistry()
				store := appinterface_event.NewRDbStore(pgxConn.ToHandle(), registry)

				events := []event.Event{test.NewFakeEvent(), test.NewFakeEvent()}
				err := store.InsertAll(events)
				Expect(err).To(BeNil())
			})
		})

		Describe("GetLatestHeight", func() {
			It("should return nil when events table does not have any record", func() {
				registry := event.NewRegistry()
				store := appinterface_event.NewRDbStore(pgxConn.ToHandle(), registry)

				actual, err := store.GetLatestHeight()
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(primptr.Int64Nil()))
			})

			It("should get 1 after insert fake event with height 1", func() {
				registry := event.NewRegistry()
				store := appinterface_event.NewRDbStore(pgxConn.ToHandle(), registry)

				// Insert an event with latestHeight 1
				mockEvent := test.NewMockEvent()
				mockEvent.On("Height").Return(int64(1))
				mockEvent.On("Name").Return("MockEvent")
				mockEvent.On("Version").Return(0)
				mockEvent.On("UUID").Return("mock-event-id")
				mockEvent.On("ToJSON").Return("\"MockEvent\"", nil)
				err := store.Insert(mockEvent)
				Expect(err).To(BeNil())

				latestHeight, err := store.GetLatestHeight()
				Expect(err).To(BeNil())
				Expect(latestHeight).To(Equal(primptr.Int64(1)))
			})
		})
	})
})
