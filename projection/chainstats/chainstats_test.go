package chainstats_test

import (
	"path"
	"runtime"
	"time"

	"github.com/crypto-com/chain-indexing/usecase/model/genesis"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	. "github.com/crypto-com/chain-indexing/external/logger/test"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/projection/chainstats"
	chainstats_view "github.com/crypto-com/chain-indexing/projection/chainstats/view"
	. "github.com/crypto-com/chain-indexing/test"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	model_usecase "github.com/crypto-com/chain-indexing/usecase/model"
)

var CHAINSTATS_MIGRATIONS_PATH = func() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("error retrieving file directory")
	}

	return path.Join(filename, "../migrations")
}()

var _ = Describe("Chain Stats", func() {
	Describe("genesis time parsing learning test", func() {
		It("should parse correctly different variations of time format", func() {
			Expect(
				utctime.MustParse(time.RFC3339, "2021-03-25T01:00:00Z"),
			).To(Equal(int64(1616634000000000000)))

			Expect(
				utctime.MustParse("2006-01-02T15:04:05.000000Z", "2020-12-23T07:30:28.674523Z"),
			).To(Equal(int64(1608708628674523000)))
			Expect(
				utctime.MustParse(time.RFC3339, "2020-12-23T07:30:28.674523Z"),
			).To(Equal(int64(1608708628674523000)))
		})
	})

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
				"chainstats_schema_migrations",
				CHAINSTATS_MIGRATIONS_PATH,
			)
			projectionMigrate.MustUp()
		})

		AfterEach(func() {
			reset()
		})

		It("should update the last projection handled block height with the fired event block height", func() {
			genesisEvent := event_usecase.NewGenesisCreated(genesis.Genesis{
				GenesisTime:   "2021-03-25T01:00:00Z",
				ChainID:       "test-chain-1",
				InitialHeight: "1",
				ConsensusParams: genesis.ConsensusParams{
					Block: genesis.Block{
						MaxBytes:   "500000",
						MaxGas:     "81500000",
						TimeIotaMS: "1000",
					},
					Evidence: genesis.ConsensusParamsEvidence{
						MaxAgeNumBlocks: "403200",
						MaxAgeDuration:  "2419200000000000",
						MaxBytes:        "150000",
					},
					Validator: genesis.ConsensusParamsValidator{PubKeyTypes: []string{"ed25519"}},
					Version:   genesis.Chainmain{},
				},
				AppHash:    "",
				AppState:   genesis.AppState{},
				Validators: nil,
			})

			blockHeight1 := int64(1)
			blockHeight1Event := event_usecase.NewBlockCreated(&model_usecase.Block{
				Height:          blockHeight1,
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
			projection := chainstats.NewChainStats(fakeLogger, pgxConn, nil)

			handleGenesisEventsErr := projection.HandleEvents(0, []event_entity.Event{
				genesisEvent,
			})
			Expect(handleGenesisEventsErr).To(BeNil())
			genesisHeight := primptr.Int64(0)
			Expect(projection.GetLastHandledEventHeight()).To(Equal(genesisHeight))

			handleBlockEventsErr := projection.HandleEvents(blockHeight1, []event_entity.Event{
				blockHeight1Event,
			})
			Expect(handleBlockEventsErr).To(BeNil())
			Expect(projection.GetLastHandledEventHeight()).To(Equal(primptr.Int64(blockHeight1)))
		})

		It("should update update the total block height after handling BlockCreated event", func() {
			chainStatsView := chainstats_view.NewChainStats(pgxConn.ToHandle())

			genesisEvent := event_usecase.NewGenesisCreated(genesis.Genesis{
				GenesisTime:   "2021-03-25T01:00:00Z",
				ChainID:       "test-chain-1",
				InitialHeight: "1",
				ConsensusParams: genesis.ConsensusParams{
					Block: genesis.Block{
						MaxBytes:   "500000",
						MaxGas:     "81500000",
						TimeIotaMS: "1000",
					},
					Evidence: genesis.ConsensusParamsEvidence{
						MaxAgeNumBlocks: "403200",
						MaxAgeDuration:  "2419200000000000",
						MaxBytes:        "150000",
					},
					Validator: genesis.ConsensusParamsValidator{PubKeyTypes: []string{"ed25519"}},
					Version:   genesis.Chainmain{},
				},
				AppHash:    "",
				AppState:   genesis.AppState{},
				Validators: nil,
			})

			blockHeight1 := int64(1)
			blockHeight1Event := event_usecase.NewBlockCreated(&model_usecase.Block{
				Height:          blockHeight1,
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

			chainStatsProjection := chainstats.NewChainStats(fakeLogger, pgxConn, nil)

			totalBlockCountBeforeHandling, findByErrBeforeHandling := chainStatsView.
				FindBy(chainstats.TOTAL_BLOCK_COUNT)

			Expect(totalBlockCountBeforeHandling).To(BeEmpty())
			Expect(findByErrBeforeHandling).To(BeNil())

			handleGenesisEventsErr := chainStatsProjection.HandleEvents(blockHeight1, []event_entity.Event{
				genesisEvent,
			})
			Expect(handleGenesisEventsErr).To(BeNil())

			totalBlockCountAfterHandlingGenesis, findByErrAfterHandlingGenesis := chainStatsView.
				FindBy(chainstats.TOTAL_BLOCK_COUNT)

			Expect(chainStatsProjection.GetLastHandledEventHeight()).To(Equal(primptr.Int64(blockHeight1)))
			Expect(totalBlockCountAfterHandlingGenesis).To(Equal("0"))
			Expect(findByErrAfterHandlingGenesis).To(BeNil())

			handleBlockHeight1EventsErr := chainStatsProjection.HandleEvents(blockHeight1, []event_entity.Event{
				blockHeight1Event,
			})
			Expect(handleBlockHeight1EventsErr).To(BeNil())

			totalBlockCountAfterHandlingBlock, findByErrAfterHandlingBlock := chainStatsView.
				FindBy(chainstats.TOTAL_BLOCK_COUNT)

			Expect(chainStatsProjection.GetLastHandledEventHeight()).To(Equal(primptr.Int64(blockHeight1)))
			Expect(totalBlockCountAfterHandlingBlock).To(Equal("1"))
			Expect(findByErrAfterHandlingBlock).To(BeNil())
		})
	})
})
