package blockevent_test

import (
	"path"
	"runtime"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	. "github.com/crypto-com/chain-indexing/external/logger/test"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/external/utctime"
	block_raw_event "github.com/crypto-com/chain-indexing/projection/block_raw_event"
	"github.com/crypto-com/chain-indexing/projection/block_raw_event/view"
	. "github.com/crypto-com/chain-indexing/test"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	model_usecase "github.com/crypto-com/chain-indexing/usecase/model"
)

var RAW_BLOCK_EVENT_MIGRATIONS_PATH = func() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("error retrieving file directory")
	}

	return path.Join(filename, "../migrations")
}()

var _ = Describe("Raw Block Events", func() {
	WithProjectionTestEnv(func(testEnv ProjectionTestEnv) {
		pgxConn := testEnv.Conn

		reset := func() {
			// Only needs to run Reset() on one of the migrate instance because Reset() drops everything in the Database
			_ = testEnv.RootMigrate.Reset()
		}

		BeforeEach(func() {
			reset()

			testEnv.RootMigrate.MustUp()
			projectionMigrate := testEnv.MigrateCreator(
				"blockrawevent_schema_migrations",
				RAW_BLOCK_EVENT_MIGRATIONS_PATH,
			)
			projectionMigrate.MustUp()
		})

		AfterEach(func() {
			reset()
		})

		It("should update the last projection handled block height with the fired event block height", func() {
			anyHeight := int64(1)
			event := event_usecase.NewBlockRawEventCreated(
				anyHeight,
				&model_usecase.CreateBlockRawEventParams{
					BlockHash:  "B69554A020537DA8E7C7610A318180C09BFEB91229BB85D4A78DDA2FACF68A48",
					BlockTime:  utctime.FromUnixNano(int64(1000000)),
					FromResult: "TxsResult",
					Data: model_usecase.DataParams{
						Type: "message",
						Content: model_usecase.BlockResultsEvent{
							Type: "message",
							Attributes: []model_usecase.BlockResultsEventAttribute{
								{
									Key:   "module",
									Value: "staking",
								},
								{
									Key:   "sender",
									Value: "cro183pry84s0hqxsxrdj3zsv6mp533t50stg59jsf",
								},
							},
						},
					},
				},
			)

			fakeLogger := NewFakeLogger()
			projection := block_raw_event.NewBlockRawEvent(fakeLogger, pgxConn, nil)
			err := projection.HandleEvents(anyHeight, []event_entity.Event{event})
			Expect(err).To(BeNil())

			Expect(projection.GetLastHandledEventHeight()).To(Equal(primptr.Int64(anyHeight)))
		})

		It("should insert the block event record after handling event", func() {
			blockRawEventsView := view.NewBlockRawEvents(pgxConn.ToHandle())

			anyHeight := int64(1)
			event := event_usecase.NewBlockRawEventCreated(
				anyHeight,
				&model_usecase.CreateBlockRawEventParams{
					BlockHash:  "B69554A020537DA8E7C7610A318180C09BFEB91229BB85D4A78DDA2FACF68A48",
					BlockTime:  utctime.FromUnixNano(int64(1000000)),
					FromResult: "TxsResult",
					Data: model_usecase.DataParams{
						Type: "message",
						Content: model_usecase.BlockResultsEvent{
							Type: "message",
							Attributes: []model_usecase.BlockResultsEventAttribute{
								{
									Key:   "module",
									Value: "staking",
								},
								{
									Key:   "sender",
									Value: "cro183pry84s0hqxsxrdj3zsv6mp533t50stg59jsf",
								},
							},
						},
					},
				},
			)

			blockRawEventListFilter := view.BlockRawEventsListFilter{
				MaybeBlockHeight: primptr.Int64(anyHeight),
			}
			blockRawEventListOrder := view.BlockRawEventsListOrder{
				Height: "ASC",
			}
			paginationPtr := &pagination_interface.Pagination{}

			fakeLogger := NewFakeLogger()

			blockRawEventProjection := block_raw_event.NewBlockRawEvent(fakeLogger, pgxConn, nil)

			blockRawEventRow, mayBePaginationResult, listErr := blockRawEventsView.
				List(blockRawEventListFilter, blockRawEventListOrder, paginationPtr)

			Expect(blockRawEventRow).To(HaveLen(0))
			Expect(mayBePaginationResult).To(BeNil())
			Expect(listErr).To(BeNil())

			handleEventsErr := blockRawEventProjection.HandleEvents(anyHeight, []event_entity.Event{event})
			Expect(handleEventsErr).To(BeNil())

			blockRawEventRowAfterHandling, maybePaginationResultAfterHandling, errAfterHandling := blockRawEventsView.
				List(blockRawEventListFilter, blockRawEventListOrder, paginationPtr)

			Expect(blockRawEventProjection.GetLastHandledEventHeight()).To(Equal(primptr.Int64(anyHeight)))
			Expect(blockRawEventRowAfterHandling).To(HaveLen(1))
			Expect(blockRawEventRowAfterHandling[0].BlockHeight).To(Equal(anyHeight))
			Expect(blockRawEventRowAfterHandling[0].BlockHash).To(Equal("B69554A020537DA8E7C7610A318180C09BFEB91229BB85D4A78DDA2FACF68A48"))
			Expect(blockRawEventRowAfterHandling[0].BlockTime).To(Equal(utctime.FromUnixNano(int64(1000000))))
			Expect(blockRawEventRowAfterHandling[0].MaybeId).To(Equal(primptr.Int64(1)))
			Expect(blockRawEventRowAfterHandling[0].FromResult).To(Equal("TxsResult"))
			Expect(maybePaginationResultAfterHandling).To(BeNil())
			Expect(errAfterHandling).To(BeNil())
		})

		It("should update projection last handled event height when there is no matching event at the height", func() {
			anyHeight := int64(1)

			fakeLogger := NewFakeLogger()
			projection := block_raw_event.NewBlockRawEvent(fakeLogger, pgxConn, nil)

			Expect(projection.GetLastHandledEventHeight()).To(BeNil())

			err := projection.HandleEvents(anyHeight, []event_entity.Event{})
			Expect(err).To(BeNil())

			Expect(projection.GetLastHandledEventHeight()).To(Equal(primptr.Int64(anyHeight)))
		})
	})
})
