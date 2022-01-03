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
	"github.com/crypto-com/chain-indexing/projection/blockevent"
	"github.com/crypto-com/chain-indexing/projection/blockevent/view"
	. "github.com/crypto-com/chain-indexing/test"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	model_usecase "github.com/crypto-com/chain-indexing/usecase/model"
)

var BLOCK_MIGRATIONS_PATH = func() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("error retrieving file directory")
	}

	return path.Join(filename, "../migrations")
}()

var _ = Describe("Block Events", func() {
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
				"blockevent_schema_migrations",
				BLOCK_MIGRATIONS_PATH,
			)
			projectionMigrate.MustUp()
		})

		AfterEach(func() {
			reset()
		})

		It("should update the last projection handled block height with the fired event block height", func() {
			anyHeight := int64(1)
			event := event_usecase.NewBlockCreated(&model_usecase.Block{
				Height:          anyHeight,
				Hash:            "B69554A020537DA8E7C7610A318180C09BFEB91229BB85D4A78DDA2FACF68A48",
				Time:            utctime.FromUnixNano(int64(1000000)),
				AppHash:         "24474D86CBFA7E6328D473C17A9E46CD5A80FFE82A348A74844BF3E2BA2B3AF1",
				ProposerAddress: "F9E6FFB9B536956201AA138224FD888D03775AB4",
				Txs: []string{
					"AAAMZqICtpjLrA3uEe3Rkg6cqDgQl0iBwG1Wm8ORZRzKL9EBAE1R0oP73H8AhHG1E1dke4ppXA/+43AyxWX7oVBTlQGnAABsiqVxCipneyzo9n1t7xauih3rXxssZTChLv/oPHUaSAAAAgAAAAAAAAAAAMw4kh8pZRDC442WIIUIo6+fIE1cTHkwEiv++LwFjNLwNtPTO1AVJx+dGivu9FZK3cN8r0mvLPVunExBxwKqKDkwVUbVSaIejuSc6Wm38rp+ekZjIExhf41FIwkocStPzBY4VYQ31YE7pmNnKocfCEg9CcKH+MIFYJ0FSxoEiMratLXjKAD6B/aMZ3pRCxxL/YxFNs+YKzGbgdKcLtydZ1B/OTGUeNU3SLNQJGmJeTL0AOwvQ4j7KkiQFZGa8uoQQEJo9rXjIg+xPk7a+zBdwC91EY/ZZz7pdVQMUgHBeKgsVvy0o78BOPWfuziBK5xVsoz5K8ZlMYQNPC5mY+bguMKYNJ5PQ9rzn+LseYT5jhalUsQrPABhEFOoQVUH1id3rszDkLLTn2/d89N/JJGY1+mL+upWWAsmJw1yTHqibSJ7RbiGffnh93MCOHeRe5OLrfmDFfhqOLNt9yuMYNdyJ+noVsfI7Ws9Kpxe2SnuCWBs9yOgM0l7UTMuIokqhGCkarXge/DUWLUcV694Jr8qcJT3OtoQohN+p+Xj66nowJLgFW7xBdFsT4vv7xI8giFxUpB+JkLgRZz2d2eam3iLDCHR+sNyvIDuXDUXhKM6aIykGgvHVHr3bBJRy/JPZFC0A3kqmheMnJpbV6kHXaIbmbgeQeXc+wxq1swFVxkwp14zbRmHwSGVGRgihjmoFYl9MyorQzNFETgp28gq2AKrCHNnIRs63m2cD5X4jZaFzfYAv9ifpOKqRtDgIFG0olge/ig00FFL6KdHySg1Qaab7g==",
				},
				Signatures: []model_usecase.BlockSignature{
					{
						BlockIdFlag:      2,
						ValidatorAddress: "F9E6FFB9B536956201AA138224FD888D03775AB4",
						Timestamp:        utctime.FromUnixNano(int64(1000000)),
						Signature:        "ZW2pUcKFN/oPQCmdCouchXmgpPyd/Ddo45dhHEMwsBeHTBuSJh15zUMmfl5FZsPHeKC8citFvOm/52bgl5XHCw==",
					},
					{
						BlockIdFlag:      2,
						ValidatorAddress: "031E3891DDB94FC7C7C132B7CD9736738110C889",
						Timestamp:        utctime.FromUnixNano(int64(2000000)),
						Signature:        "uhWDC9NDT86FbRVGbOM2lGY8sVkWU51JJ9F8gPwTfK0ebcui1R34oM+jhPKdStn/4sq4qDgzbsN66cQ5kl8NAw==",
					},
				},
			})

			fakeLogger := NewFakeLogger()
			projection := blockevent.NewBlockEvent(fakeLogger, pgxConn, nil)
			err := projection.HandleEvents(anyHeight, []event_entity.Event{event})
			Expect(err).To(BeNil())

			Expect(projection.GetLastHandledEventHeight()).To(Equal(primptr.Int64(anyHeight)))
		})

		It("should insert the block event record after handling event", func() {
			blockEventsView := view.NewBlockEvents(pgxConn.ToHandle())

			anyHeight := int64(1)
			event := event_usecase.NewBlockCreated(&model_usecase.Block{
				Height:          anyHeight,
				Hash:            "B69554A020537DA8E7C7610A318180C09BFEB91229BB85D4A78DDA2FACF68A48",
				Time:            utctime.FromUnixNano(int64(1000000)),
				AppHash:         "24474D86CBFA7E6328D473C17A9E46CD5A80FFE82A348A74844BF3E2BA2B3AF1",
				ProposerAddress: "F9E6FFB9B536956201AA138224FD888D03775AB4",
				Txs: []string{
					"AAAMZqICtpjLrA3uEe3Rkg6cqDgQl0iBwG1Wm8ORZRzKL9EBAE1R0oP73H8AhHG1E1dke4ppXA/+43AyxWX7oVBTlQGnAABsiqVxCipneyzo9n1t7xauih3rXxssZTChLv/oPHUaSAAAAgAAAAAAAAAAAMw4kh8pZRDC442WIIUIo6+fIE1cTHkwEiv++LwFjNLwNtPTO1AVJx+dGivu9FZK3cN8r0mvLPVunExBxwKqKDkwVUbVSaIejuSc6Wm38rp+ekZjIExhf41FIwkocStPzBY4VYQ31YE7pmNnKocfCEg9CcKH+MIFYJ0FSxoEiMratLXjKAD6B/aMZ3pRCxxL/YxFNs+YKzGbgdKcLtydZ1B/OTGUeNU3SLNQJGmJeTL0AOwvQ4j7KkiQFZGa8uoQQEJo9rXjIg+xPk7a+zBdwC91EY/ZZz7pdVQMUgHBeKgsVvy0o78BOPWfuziBK5xVsoz5K8ZlMYQNPC5mY+bguMKYNJ5PQ9rzn+LseYT5jhalUsQrPABhEFOoQVUH1id3rszDkLLTn2/d89N/JJGY1+mL+upWWAsmJw1yTHqibSJ7RbiGffnh93MCOHeRe5OLrfmDFfhqOLNt9yuMYNdyJ+noVsfI7Ws9Kpxe2SnuCWBs9yOgM0l7UTMuIokqhGCkarXge/DUWLUcV694Jr8qcJT3OtoQohN+p+Xj66nowJLgFW7xBdFsT4vv7xI8giFxUpB+JkLgRZz2d2eam3iLDCHR+sNyvIDuXDUXhKM6aIykGgvHVHr3bBJRy/JPZFC0A3kqmheMnJpbV6kHXaIbmbgeQeXc+wxq1swFVxkwp14zbRmHwSGVGRgihjmoFYl9MyorQzNFETgp28gq2AKrCHNnIRs63m2cD5X4jZaFzfYAv9ifpOKqRtDgIFG0olge/ig00FFL6KdHySg1Qaab7g==",
				},
				Signatures: []model_usecase.BlockSignature{
					{
						BlockIdFlag:      2,
						ValidatorAddress: "F9E6FFB9B536956201AA138224FD888D03775AB4",
						Timestamp:        utctime.FromUnixNano(int64(1000000)),
						Signature:        "ZW2pUcKFN/oPQCmdCouchXmgpPyd/Ddo45dhHEMwsBeHTBuSJh15zUMmfl5FZsPHeKC8citFvOm/52bgl5XHCw==",
					},
					{
						BlockIdFlag:      2,
						ValidatorAddress: "031E3891DDB94FC7C7C132B7CD9736738110C889",
						Timestamp:        utctime.FromUnixNano(int64(2000000)),
						Signature:        "uhWDC9NDT86FbRVGbOM2lGY8sVkWU51JJ9F8gPwTfK0ebcui1R34oM+jhPKdStn/4sq4qDgzbsN66cQ5kl8NAw==",
					},
				},
			})

			anyAmount := coin.MustParseDecCoins("1000basetcro")
			eventBlockEvent := event_usecase.NewBlockRewarded(anyHeight, "validator", anyAmount)
			blockEventListFilter := view.BlockEventsListFilter{
				MaybeBlockHeight: primptr.Int64(anyHeight),
			}
			blockEventListOrder := view.BlockEventsListOrder{
				Height: "ASC",
			}
			paginationPtr := &pagination_interface.Pagination{}

			fakeLogger := NewFakeLogger()

			blockEventProjection := blockevent.NewBlockEvent(fakeLogger, pgxConn, nil)

			blockEventRow, mayBePaginationResult, listErr := blockEventsView.
				List(blockEventListFilter, blockEventListOrder, paginationPtr)

			Expect(blockEventRow).To(HaveLen(0))
			Expect(mayBePaginationResult).To(BeNil())
			Expect(listErr).To(BeNil())

			handleEventsErr := blockEventProjection.HandleEvents(anyHeight, []event_entity.Event{event, eventBlockEvent})
			Expect(handleEventsErr).To(BeNil())

			blockEventRowAfterHandling, maybePaginationResultAfterHandling, errAfterHandling := blockEventsView.
				List(blockEventListFilter, blockEventListOrder, paginationPtr)

			Expect(blockEventProjection.GetLastHandledEventHeight()).To(Equal(primptr.Int64(anyHeight)))
			Expect(blockEventRowAfterHandling).To(HaveLen(1))
			Expect(blockEventRowAfterHandling[0].BlockHeight).To(Equal(anyHeight))
			Expect(blockEventRowAfterHandling[0].BlockTime).To(Equal(utctime.FromUnixNano(int64(1000000))))
			Expect(blockEventRowAfterHandling[0].MaybeId).To(Equal(primptr.Int64(1)))
			Expect(blockEventRowAfterHandling[0].BlockHash).To(Equal("B69554A020537DA8E7C7610A318180C09BFEB91229BB85D4A78DDA2FACF68A48"))
			Expect(maybePaginationResultAfterHandling).To(BeNil())
			Expect(errAfterHandling).To(BeNil())
		})

		It("should update projection last handled event height when there is no matching event at the height", func() {
			anyHeight := int64(1)

			fakeLogger := NewFakeLogger()
			projection := blockevent.NewBlockEvent(fakeLogger, pgxConn, nil)

			Expect(projection.GetLastHandledEventHeight()).To(BeNil())

			err := projection.HandleEvents(anyHeight, []event_entity.Event{})
			Expect(err).To(BeNil())

			Expect(projection.GetLastHandledEventHeight()).To(Equal(primptr.Int64(anyHeight)))
		})
	})
})
