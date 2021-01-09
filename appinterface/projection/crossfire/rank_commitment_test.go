package crossfire_test

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/crossfire/constants"
	crossfire_test "github.com/crypto-com/chain-indexing/appinterface/projection/crossfire/test"
	"github.com/crypto-com/chain-indexing/internal/utctime"
	"net/http"

	"github.com/crypto-com/chain-indexing/appinterface/projection/crossfire"
	"github.com/crypto-com/chain-indexing/appinterface/projection/crossfire/view"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	. "github.com/crypto-com/chain-indexing/internal/logger/test"
	. "github.com/crypto-com/chain-indexing/test"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	usecase_model "github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var rankServer *ghttp.Server
var _ = Describe("Crossfire", func() {
	WithTestPgxConn(func(pgConn *pg.PgxConn, pgMigrate *pg.Migrate) {
		anyHeight := int64(1)
		validator1CreatedEvent := event_usecase.NewMsgCreateValidator(event_usecase.MsgCommonParams{
			BlockHeight: anyHeight,
			TxHash:      "A6D4C1F59A9D232747CA4F8A484F1F3B14A0075E801DF2A25F472B4280505B74",
			TxSuccess:   true,
			MsgIndex:    0,
		}, usecase_model.MsgCreateValidatorParams{
			Description: usecase_model.MsgValidatorDescription{
				Moniker:         "Testing",
				Identity:        "foo",
				Website:         "www.example.com",
				SecurityContact: "foo@example.com",
				Details:         "example",
			},
			Commission: usecase_model.MsgValidatorCommission{
				Rate:          "0.100000000000000000",
				MaxRate:       "0.200000000000000000",
				MaxChangeRate: "0.010000000000000000",
			},
			MinSelfDelegation: "1",
			DelegatorAddress:  "tcro1n4t5q77kn9vf73s7ljs96m85jgg49yqpasmwm3",
			ValidatorAddress:  "tcrocncl1n4t5q77kn9vf73s7ljs96m85jgg49yqpg0chrj",
			TendermintPubkey:  "Og8ZfQTHFgTBGD5qoyo5NpyJCJRddC+WuSPtyZtlE7E=",
			Amount:            coin.MustNewCoinFromString("100"),
		})

		validator2CreatedEvent := event_usecase.NewMsgCreateValidator(event_usecase.MsgCommonParams{
			BlockHeight: anyHeight,
			TxHash:      "A6D4C1F59A9D232747CA4F8A484F1F3B14A0075E801DF2A25F472B4280505B74",
			TxSuccess:   true,
			MsgIndex:    0,
		}, usecase_model.MsgCreateValidatorParams{
			Description: usecase_model.MsgValidatorDescription{
				Moniker:         "nebkas.ro",
				Identity:        "foo",
				Website:         "www.example.com",
				SecurityContact: "foo@example.com",
				Details:         "example",
			},
			Commission: usecase_model.MsgValidatorCommission{
				Rate:          "0.100000000000000000",
				MaxRate:       "0.200000000000000000",
				MaxChangeRate: "0.010000000000000000",
			},
			MinSelfDelegation: "1",
			DelegatorAddress:  "tcro15xr8daqzpu0wf8t6hx95zlxmqwzmf4eaph3yzv",
			ValidatorAddress:  "tcrocncl15xr8daqzpu0wf8t6hx95zlxmqwzmf4ea5gja60",
			TendermintPubkey:  "BuuPYme7R4eH/nWs2p+sS1UpCQwy+QJgBZuhGICH8Es=",
			Amount:            coin.MustNewCoinFromString("222"),
		})

		validator3CreatedEvent := event_usecase.NewMsgCreateValidator(event_usecase.MsgCommonParams{
			BlockHeight: anyHeight,
			TxHash:      "A6D4C1F59A9D232747CA4F8A484F1F3B14A0075E801DF2A25F472B4280505B74",
			TxSuccess:   true,
			MsgIndex:    0,
		}, usecase_model.MsgCreateValidatorParams{
			Description: usecase_model.MsgValidatorDescription{
				Moniker:         "nebksdfas.ro",
				Identity:        "foo",
				Website:         "www.example.com",
				SecurityContact: "foo@example.com",
				Details:         "example",
			},
			Commission: usecase_model.MsgValidatorCommission{
				Rate:          "0.100000000000000000",
				MaxRate:       "0.200000000000000000",
				MaxChangeRate: "0.010000000000000000",
			},
			MinSelfDelegation: "1",
			DelegatorAddress:  "tcro197ujxhaeyyv309f39c0s2gn0af0pps5pden6h7",
			ValidatorAddress:  "tcrocncl197ujxhaeyyv309f39c0s2gn0af0pps5pcxsr0a",
			TendermintPubkey:  "wWw0e9tZcVmev/NyJlZv5Apd7U5IONoyx3U/9rD5fHI=",
			Amount:            coin.MustNewCoinFromString("333"),
		})

		phaseOneBlockCreatedEvent := event_usecase.NewBlockCreated(&usecase_model.Block{
			Height:          int64(1),
			Hash:            "B69554A020537DA8E7C7610A318180C09BFEB91229BB85D4A78DDA2FACF68A48",
			Time:            utctime.FromUnixNano(int64(1610942500000000000)),
			AppHash:         "24474D86CBFA7E6328D473C17A9E46CD5A80FFE82A348A74844BF3E2BA2B3AF1",
			ProposerAddress: "F9E6FFB9B536956201AA138224FD888D03775AB4",
			Txs: []string{
				"AAAMZqICtpjLrA3uEe3Rkg6cqDgQl0iBwG1Wm8ORZRzKL9EBAE1R0oP73H8AhHG1E1dke4ppXA/+43AyxWX7oVBTlQGnAABsiqVxCipneyzo9n1t7xauih3rXxssZTChLv/oPHUaSAAAAgAAAAAAAAAAAMw4kh8pZRDC442WIIUIo6+fIE1cTHkwEiv++LwFjNLwNtPTO1AVJx+dGivu9FZK3cN8r0mvLPVunExBxwKqKDkwVUbVSaIejuSc6Wm38rp+ekZjIExhf41FIwkocStPzBY4VYQ31YE7pmNnKocfCEg9CcKH+MIFYJ0FSxoEiMratLXjKAD6B/aMZ3pRCxxL/YxFNs+YKzGbgdKcLtydZ1B/OTGUeNU3SLNQJGmJeTL0AOwvQ4j7KkiQFZGa8uoQQEJo9rXjIg+xPk7a+zBdwC91EY/ZZz7pdVQMUgHBeKgsVvy0o78BOPWfuziBK5xVsoz5K8ZlMYQNPC5mY+bguMKYNJ5PQ9rzn+LseYT5jhalUsQrPABhEFOoQVUH1id3rszDkLLTn2/d89N/JJGY1+mL+upWWAsmJw1yTHqibSJ7RbiGffnh93MCOHeRe5OLrfmDFfhqOLNt9yuMYNdyJ+noVsfI7Ws9Kpxe2SnuCWBs9yOgM0l7UTMuIokqhGCkarXge/DUWLUcV694Jr8qcJT3OtoQohN+p+Xj66nowJLgFW7xBdFsT4vv7xI8giFxUpB+JkLgRZz2d2eam3iLDCHR+sNyvIDuXDUXhKM6aIykGgvHVHr3bBJRy/JPZFC0A3kqmheMnJpbV6kHXaIbmbgeQeXc+wxq1swFVxkwp14zbRmHwSGVGRgihjmoFYl9MyorQzNFETgp28gq2AKrCHNnIRs63m2cD5X4jZaFzfYAv9ifpOKqRtDgIFG0olge/ig00FFL6KdHySg1Qaab7g==",
			},
			Signatures: []usecase_model.BlockSignature{
				{
					BlockIdFlag:      2,
					ValidatorAddress: "703B26AEA0867B03572719D22F4B8E6D93CA838C", // node 1
					Timestamp:        utctime.FromUnixNano(int64(1000000)),
					Signature:        "ZW2pUcKFN/oPQCmdCouchXmgpPyd/Ddo45dhHEMwsBeHTBuSJh15zUMmfl5FZsPHeKC8citFvOm/52bgl5XHCw==",
				},
				{
					BlockIdFlag:      2,
					ValidatorAddress: "AEA0F558C9616A7089791D1AE4C08DC5F69A0A0B", // node 2
					Timestamp:        utctime.FromUnixNano(int64(2000000)),
					Signature:        "uhWDC9NDT86FbRVGbOM2lGY8sVkWU51JJ9F8gPwTfK0ebcui1R34oM+jhPKdStn/4sq4qDgzbsN66cQ5kl8NAw==",
				},
			},
		})

		phaseTwoBlockCreatedEvent := event_usecase.NewBlockCreated(&usecase_model.Block{
			Height:          int64(100000000),
			Hash:            "B69554A020537DA8E7C7610A318180C09BFEB91229BB85D4A78DDA2FACF68A48",
			Time:            utctime.FromUnixNano(int64(1611547500000000000)),
			AppHash:         "24474D86CBFA7E6328D473C17A9E46CD5A80FFE82A348A74844BF3E2BA2B3AF1",
			ProposerAddress: "F9E6FFB9B536956201AA138224FD888D03775AB4",
			Txs: []string{
				"AAAMZqICtpjLrA3uEe3Rkg6cqDgQl0iBwG1Wm8ORZRzKL9EBAE1R0oP73H8AhHG1E1dke4ppXA/+43AyxWX7oVBTlQGnAABsiqVxCipneyzo9n1t7xauih3rXxssZTChLv/oPHUaSAAAAgAAAAAAAAAAAMw4kh8pZRDC442WIIUIo6+fIE1cTHkwEiv++LwFjNLwNtPTO1AVJx+dGivu9FZK3cN8r0mvLPVunExBxwKqKDkwVUbVSaIejuSc6Wm38rp+ekZjIExhf41FIwkocStPzBY4VYQ31YE7pmNnKocfCEg9CcKH+MIFYJ0FSxoEiMratLXjKAD6B/aMZ3pRCxxL/YxFNs+YKzGbgdKcLtydZ1B/OTGUeNU3SLNQJGmJeTL0AOwvQ4j7KkiQFZGa8uoQQEJo9rXjIg+xPk7a+zBdwC91EY/ZZz7pdVQMUgHBeKgsVvy0o78BOPWfuziBK5xVsoz5K8ZlMYQNPC5mY+bguMKYNJ5PQ9rzn+LseYT5jhalUsQrPABhEFOoQVUH1id3rszDkLLTn2/d89N/JJGY1+mL+upWWAsmJw1yTHqibSJ7RbiGffnh93MCOHeRe5OLrfmDFfhqOLNt9yuMYNdyJ+noVsfI7Ws9Kpxe2SnuCWBs9yOgM0l7UTMuIokqhGCkarXge/DUWLUcV694Jr8qcJT3OtoQohN+p+Xj66nowJLgFW7xBdFsT4vv7xI8giFxUpB+JkLgRZz2d2eam3iLDCHR+sNyvIDuXDUXhKM6aIykGgvHVHr3bBJRy/JPZFC0A3kqmheMnJpbV6kHXaIbmbgeQeXc+wxq1swFVxkwp14zbRmHwSGVGRgihjmoFYl9MyorQzNFETgp28gq2AKrCHNnIRs63m2cD5X4jZaFzfYAv9ifpOKqRtDgIFG0olge/ig00FFL6KdHySg1Qaab7g==",
			},
			Signatures: []usecase_model.BlockSignature{
				{
					BlockIdFlag:      2,
					ValidatorAddress: "703B26AEA0867B03572719D22F4B8E6D93CA838C", // node 1
					Timestamp:        utctime.FromUnixNano(int64(1000000)),
					Signature:        "ZW2pUcKFN/oPQCmdCouchXmgpPyd/Ddo45dhHEMwsBeHTBuSJh15zUMmfl5FZsPHeKC8citFvOm/52bgl5XHCw==",
				},
				{
					BlockIdFlag:      2,
					ValidatorAddress: "AEA0F558C9616A7089791D1AE4C08DC5F69A0A0B", // node 2
					Timestamp:        utctime.FromUnixNano(int64(2000000)),
					Signature:        "uhWDC9NDT86FbRVGbOM2lGY8sVkWU51JJ9F8gPwTfK0ebcui1R34oM+jhPKdStn/4sq4qDgzbsN66cQ5kl8NAw==",
				},
				{
					BlockIdFlag:      2,
					ValidatorAddress: "0FB7AE9AC2E3F148CA130341B6CD4DB3682E2D54", // node 3
					Timestamp:        utctime.FromUnixNano(int64(2000000)),
					Signature:        "uhWDC9NDT86FbRVGbOM2lGY8sVkWU51JJ9F8gPwTfK0ebcui1R34oM+jhPKdStn/4sq4qDgzbsN66cQ5kl8NAw==",
				},
			},
		})

		phaseThreeBlockCreatedEvent1 := event_usecase.NewBlockCreated(&usecase_model.Block{
			Height:          int64(100000000),
			Hash:            "B69554A020537DA8E7C7610A318180C09BFEB91229BB85D4A78DDA2FACF68A48",
			Time:            utctime.FromUnixNano(int64(1612756900000000000)),
			AppHash:         "24474D86CBFA7E6328D473C17A9E46CD5A80FFE82A348A74844BF3E2BA2B3AF1",
			ProposerAddress: "F9E6FFB9B536956201AA138224FD888D03775AB4",
			Txs: []string{
				"AAAMZqICtpjLrA3uEe3Rkg6cqDgQl0iBwG1Wm8ORZRzKL9EBAE1R0oP73H8AhHG1E1dke4ppXA/+43AyxWX7oVBTlQGnAABsiqVxCipneyzo9n1t7xauih3rXxssZTChLv/oPHUaSAAAAgAAAAAAAAAAAMw4kh8pZRDC442WIIUIo6+fIE1cTHkwEiv++LwFjNLwNtPTO1AVJx+dGivu9FZK3cN8r0mvLPVunExBxwKqKDkwVUbVSaIejuSc6Wm38rp+ekZjIExhf41FIwkocStPzBY4VYQ31YE7pmNnKocfCEg9CcKH+MIFYJ0FSxoEiMratLXjKAD6B/aMZ3pRCxxL/YxFNs+YKzGbgdKcLtydZ1B/OTGUeNU3SLNQJGmJeTL0AOwvQ4j7KkiQFZGa8uoQQEJo9rXjIg+xPk7a+zBdwC91EY/ZZz7pdVQMUgHBeKgsVvy0o78BOPWfuziBK5xVsoz5K8ZlMYQNPC5mY+bguMKYNJ5PQ9rzn+LseYT5jhalUsQrPABhEFOoQVUH1id3rszDkLLTn2/d89N/JJGY1+mL+upWWAsmJw1yTHqibSJ7RbiGffnh93MCOHeRe5OLrfmDFfhqOLNt9yuMYNdyJ+noVsfI7Ws9Kpxe2SnuCWBs9yOgM0l7UTMuIokqhGCkarXge/DUWLUcV694Jr8qcJT3OtoQohN+p+Xj66nowJLgFW7xBdFsT4vv7xI8giFxUpB+JkLgRZz2d2eam3iLDCHR+sNyvIDuXDUXhKM6aIykGgvHVHr3bBJRy/JPZFC0A3kqmheMnJpbV6kHXaIbmbgeQeXc+wxq1swFVxkwp14zbRmHwSGVGRgihjmoFYl9MyorQzNFETgp28gq2AKrCHNnIRs63m2cD5X4jZaFzfYAv9ifpOKqRtDgIFG0olge/ig00FFL6KdHySg1Qaab7g==",
			},
			Signatures: []usecase_model.BlockSignature{
				{
					BlockIdFlag:      2,
					ValidatorAddress: "703B26AEA0867B03572719D22F4B8E6D93CA838C", // node 1
					Timestamp:        utctime.FromUnixNano(int64(1000000)),
					Signature:        "ZW2pUcKFN/oPQCmdCouchXmgpPyd/Ddo45dhHEMwsBeHTBuSJh15zUMmfl5FZsPHeKC8citFvOm/52bgl5XHCw==",
				},
				{
					BlockIdFlag:      2,
					ValidatorAddress: "0FB7AE9AC2E3F148CA130341B6CD4DB3682E2D54", // node 3
					Timestamp:        utctime.FromUnixNano(int64(2000000)),
					Signature:        "uhWDC9NDT86FbRVGbOM2lGY8sVkWU51JJ9F8gPwTfK0ebcui1R34oM+jhPKdStn/4sq4qDgzbsN66cQ5kl8NAw==",
				},
			},
		})

		phaseThreeBlockCreatedEvent2 := event_usecase.NewBlockCreated(&usecase_model.Block{
			Height:          int64(100000000),
			Hash:            "B69554A020537DA8E7C7610A318180C09BFEB91229BB85D4A78DDA2FACF68A48",
			Time:            utctime.FromUnixNano(int64(1612757000000000000)),
			AppHash:         "24474D86CBFA7E6328D473C17A9E46CD5A80FFE82A348A74844BF3E2BA2B3AF1",
			ProposerAddress: "F9E6FFB9B536956201AA138224FD888D03775AB4",
			Txs: []string{
				"AAAMZqICtpjLrA3uEe3Rkg6cqDgQl0iBwG1Wm8ORZRzKL9EBAE1R0oP73H8AhHG1E1dke4ppXA/+43AyxWX7oVBTlQGnAABsiqVxCipneyzo9n1t7xauih3rXxssZTChLv/oPHUaSAAAAgAAAAAAAAAAAMw4kh8pZRDC442WIIUIo6+fIE1cTHkwEiv++LwFjNLwNtPTO1AVJx+dGivu9FZK3cN8r0mvLPVunExBxwKqKDkwVUbVSaIejuSc6Wm38rp+ekZjIExhf41FIwkocStPzBY4VYQ31YE7pmNnKocfCEg9CcKH+MIFYJ0FSxoEiMratLXjKAD6B/aMZ3pRCxxL/YxFNs+YKzGbgdKcLtydZ1B/OTGUeNU3SLNQJGmJeTL0AOwvQ4j7KkiQFZGa8uoQQEJo9rXjIg+xPk7a+zBdwC91EY/ZZz7pdVQMUgHBeKgsVvy0o78BOPWfuziBK5xVsoz5K8ZlMYQNPC5mY+bguMKYNJ5PQ9rzn+LseYT5jhalUsQrPABhEFOoQVUH1id3rszDkLLTn2/d89N/JJGY1+mL+upWWAsmJw1yTHqibSJ7RbiGffnh93MCOHeRe5OLrfmDFfhqOLNt9yuMYNdyJ+noVsfI7Ws9Kpxe2SnuCWBs9yOgM0l7UTMuIokqhGCkarXge/DUWLUcV694Jr8qcJT3OtoQohN+p+Xj66nowJLgFW7xBdFsT4vv7xI8giFxUpB+JkLgRZz2d2eam3iLDCHR+sNyvIDuXDUXhKM6aIykGgvHVHr3bBJRy/JPZFC0A3kqmheMnJpbV6kHXaIbmbgeQeXc+wxq1swFVxkwp14zbRmHwSGVGRgihjmoFYl9MyorQzNFETgp28gq2AKrCHNnIRs63m2cD5X4jZaFzfYAv9ifpOKqRtDgIFG0olge/ig00FFL6KdHySg1Qaab7g==",
			},
			Signatures: []usecase_model.BlockSignature{
				{
					BlockIdFlag:      2,
					ValidatorAddress: "703B26AEA0867B03572719D22F4B8E6D93CA838C", // node 1
					Timestamp:        utctime.FromUnixNano(int64(1000000)),
					Signature:        "ZW2pUcKFN/oPQCmdCouchXmgpPyd/Ddo45dhHEMwsBeHTBuSJh15zUMmfl5FZsPHeKC8citFvOm/52bgl5XHCw==",
				},
				{
					BlockIdFlag:      2,
					ValidatorAddress: "AEA0F558C9616A7089791D1AE4C08DC5F69A0A0B", // node 2
					Timestamp:        utctime.FromUnixNano(int64(2000000)),
					Signature:        "uhWDC9NDT86FbRVGbOM2lGY8sVkWU51JJ9F8gPwTfK0ebcui1R34oM+jhPKdStn/4sq4qDgzbsN66cQ5kl8NAw==",
				},
			},
		})

		BeforeEach(func() {
			_ = pgMigrate.Reset()
			pgMigrate.MustUp()
			rankServer = ghttp.NewServer()
			rankServer.RouteToHandler(
				"GET",
				"/participants.json",
				ghttp.RespondWith(http.StatusOK, crossfire_test.COMMIT_RANK_PARTICIPANTS_SAMPLE_JSON),
			)
		})

		AfterEach(func() {
			_ = pgMigrate.Reset()
			rankServer.Close()
		})

		It("should correctly compute phase 1 and phase 2 commitment rank", func() {
			crossfireValidatorsView := view.NewCrossfireValidators(pgConn.ToHandle())
			crossfireValidatorsStatsView := view.NewCrossfireValidatorsStats(pgConn.ToHandle())
			crossfireChainStatsView := view.NewCrossfireChainStats(pgConn.ToHandle())
			fakeLogger := NewFakeLogger()

			projection := crossfire.NewCrossfire(
				fakeLogger,
				pgConn,
				"tcrocnclcons",
				"tcrocncl",
				1610942400000000000,
				1611547200000000000,
				1612756800000000000,
				1613361599000000000,
				"tcro15grftg88l0gdw4mg9t9pwnl0pde2asjzvfpkp4",
				"14",
				rankServer.URL()+"/participants.json",
			)

			err := projection.HandleEvents(anyHeight, []event_entity.Event{validator1CreatedEvent, validator2CreatedEvent, validator3CreatedEvent})
			Expect(err).To(BeNil())

			// 1. store 3 validators
			opAddr1 := "tcrocncl1n4t5q77kn9vf73s7ljs96m85jgg49yqpg0chrj" // 703B26AEA0867B03572719D22F4B8E6D93CA838C
			actual1, err := crossfireValidatorsView.FindBy(view.CrossfireValidatorIdentity{
				MaybeOperatorAddress: &opAddr1,
			})
			Expect(err).To(BeNil())
			Expect(actual1.TendermintAddress).To(Equal("703B26AEA0867B03572719D22F4B8E6D93CA838C"))

			opAddr2 := "tcrocncl15xr8daqzpu0wf8t6hx95zlxmqwzmf4ea5gja60" // AEA0F558C9616A7089791D1AE4C08DC5F69A0A0B
			actual2, err := crossfireValidatorsView.FindBy(view.CrossfireValidatorIdentity{
				MaybeOperatorAddress: &opAddr2,
			})
			Expect(err).To(BeNil())
			Expect(actual2.TendermintAddress).To(Equal("AEA0F558C9616A7089791D1AE4C08DC5F69A0A0B"))

			opAddr3 := "tcrocncl197ujxhaeyyv309f39c0s2gn0af0pps5pcxsr0a" // 0FB7AE9AC2E3F148CA130341B6CD4DB3682E2D54
			actual3, err := crossfireValidatorsView.FindBy(view.CrossfireValidatorIdentity{
				MaybeOperatorAddress: &opAddr3,
			})
			Expect(err).To(BeNil())
			Expect(actual3.TendermintAddress).To(Equal("0FB7AE9AC2E3F148CA130341B6CD4DB3682E2D54"))

			// 2. store a phase 1 block that contains validator 1 & 2's commitment
			err = projection.HandleEvents(anyHeight, []event_entity.Event{phaseOneBlockCreatedEvent})
			Expect(err).To(BeNil())

			phase1BlockCount, err := crossfireChainStatsView.FindBy(constants.PHASE_1_BLOCK_COUNT)
			Expect(err).To(BeNil())
			Expect(phase1BlockCount).To(Equal(int64(1)))

			node1CommitCount, err := crossfireValidatorsStatsView.FindBy(
				constants.ValidatorCommitmentKey(
					actual1.OperatorAddress,
					constants.PHASE_1N2_COMMIT_PREFIX,
				),
			)
			Expect(err).To(BeNil())
			Expect(node1CommitCount).To(Equal(int64(1)))

			node2CommitCount, err := crossfireValidatorsStatsView.FindBy(
				constants.ValidatorCommitmentKey(
					actual2.OperatorAddress,
					constants.PHASE_1N2_COMMIT_PREFIX,
				),
			)
			Expect(err).To(BeNil())
			Expect(node2CommitCount).To(Equal(int64(1)))

			actual1, err = crossfireValidatorsView.FindBy(view.CrossfireValidatorIdentity{
				MaybeOperatorAddress: &opAddr1,
			})
			Expect(err).To(BeNil())
			Expect(actual1.RankTaskPhase1n2CommitmentCount).To(Equal(int64(1))) // node1 rank #1

			actual2, err = crossfireValidatorsView.FindBy(view.CrossfireValidatorIdentity{
				MaybeOperatorAddress: &opAddr2,
			})
			Expect(err).To(BeNil())
			Expect(actual2.RankTaskPhase1n2CommitmentCount).To(Equal(int64(1))) // node2 rank #1

			// 3. store a phase 2 block that contains validator 1, 2 & 3's commitment
			err = projection.HandleEvents(anyHeight, []event_entity.Event{phaseTwoBlockCreatedEvent})
			Expect(err).To(BeNil())

			phase2BlockCount, err := crossfireChainStatsView.FindBy(constants.PHASE_2_BLOCK_COUNT)
			Expect(err).To(BeNil())
			Expect(phase2BlockCount).To(Equal(int64(1)))

			node1CommitCount, err = crossfireValidatorsStatsView.FindBy(
				constants.ValidatorCommitmentKey(
					actual1.OperatorAddress,
					constants.PHASE_1N2_COMMIT_PREFIX,
				),
			)
			Expect(err).To(BeNil())
			Expect(node1CommitCount).To(Equal(int64(2))) // node1 2 commitments

			node2CommitCount, err = crossfireValidatorsStatsView.FindBy(
				constants.ValidatorCommitmentKey(
					actual2.OperatorAddress,
					constants.PHASE_1N2_COMMIT_PREFIX,
				),
			)
			Expect(err).To(BeNil())
			Expect(node2CommitCount).To(Equal(int64(2))) // node2 2 commitments

			node3CommitCount, err := crossfireValidatorsStatsView.FindBy(
				constants.ValidatorCommitmentKey(
					actual3.OperatorAddress,
					constants.PHASE_1N2_COMMIT_PREFIX,
				),
			)
			Expect(err).To(BeNil())
			Expect(node3CommitCount).To(Equal(int64(1))) // node3 1 commitment

			actual1, err = crossfireValidatorsView.FindBy(view.CrossfireValidatorIdentity{
				MaybeOperatorAddress: &opAddr1,
			})
			Expect(err).To(BeNil())
			Expect(actual1.RankTaskPhase1n2CommitmentCount).To(Equal(int64(1))) // node1 rank #1

			actual2, err = crossfireValidatorsView.FindBy(view.CrossfireValidatorIdentity{
				MaybeOperatorAddress: &opAddr2,
			})
			Expect(err).To(BeNil())
			Expect(actual2.RankTaskPhase1n2CommitmentCount).To(Equal(int64(1))) // node2 rank #1

			actual3, err = crossfireValidatorsView.FindBy(view.CrossfireValidatorIdentity{
				MaybeOperatorAddress: &opAddr3,
			})
			Expect(err).To(BeNil())
			Expect(actual3.RankTaskPhase1n2CommitmentCount).To(Equal(int64(2))) // node3 rank #2
		})

		It("should correctly compute phase 3 commitment rank", func() {
			crossfireValidatorsView := view.NewCrossfireValidators(pgConn.ToHandle())
			crossfireValidatorsStatsView := view.NewCrossfireValidatorsStats(pgConn.ToHandle())
			crossfireChainStatsView := view.NewCrossfireChainStats(pgConn.ToHandle())
			fakeLogger := NewFakeLogger()

			projection := crossfire.NewCrossfire(
				fakeLogger,
				pgConn,
				"tcrocnclcons",
				"tcrocncl",
				1610942400000000000,
				1611547200000000000,
				1612756800000000000,
				1613361599000000000,
				"tcro15grftg88l0gdw4mg9t9pwnl0pde2asjzvfpkp4",
				"14",
				rankServer.URL()+"/participants.json",
			)

			err := projection.HandleEvents(anyHeight, []event_entity.Event{validator1CreatedEvent, validator2CreatedEvent, validator3CreatedEvent})
			Expect(err).To(BeNil())

			// 1. store 3 validators
			opAddr1 := "tcrocncl1n4t5q77kn9vf73s7ljs96m85jgg49yqpg0chrj" // 703B26AEA0867B03572719D22F4B8E6D93CA838C
			actual1, err := crossfireValidatorsView.FindBy(view.CrossfireValidatorIdentity{
				MaybeOperatorAddress: &opAddr1,
			})
			Expect(err).To(BeNil())
			Expect(actual1.TendermintAddress).To(Equal("703B26AEA0867B03572719D22F4B8E6D93CA838C"))

			opAddr2 := "tcrocncl15xr8daqzpu0wf8t6hx95zlxmqwzmf4ea5gja60" // AEA0F558C9616A7089791D1AE4C08DC5F69A0A0B
			actual2, err := crossfireValidatorsView.FindBy(view.CrossfireValidatorIdentity{
				MaybeOperatorAddress: &opAddr2,
			})
			Expect(err).To(BeNil())
			Expect(actual2.TendermintAddress).To(Equal("AEA0F558C9616A7089791D1AE4C08DC5F69A0A0B"))

			opAddr3 := "tcrocncl197ujxhaeyyv309f39c0s2gn0af0pps5pcxsr0a" // 0FB7AE9AC2E3F148CA130341B6CD4DB3682E2D54
			actual3, err := crossfireValidatorsView.FindBy(view.CrossfireValidatorIdentity{
				MaybeOperatorAddress: &opAddr3,
			})
			Expect(err).To(BeNil())
			Expect(actual3.TendermintAddress).To(Equal("0FB7AE9AC2E3F148CA130341B6CD4DB3682E2D54"))

			// 2. store a phase 3 block that contains validator 1 & 3's commitment
			err = projection.HandleEvents(anyHeight, []event_entity.Event{phaseThreeBlockCreatedEvent1})
			Expect(err).To(BeNil())

			phase3BlockCount, err := crossfireChainStatsView.FindBy(constants.PHASE_3_BLOCK_COUNT)
			Expect(err).To(BeNil())
			Expect(phase3BlockCount).To(Equal(int64(1)))

			node1CommitCount, err := crossfireValidatorsStatsView.FindBy(
				constants.ValidatorCommitmentKey(
					actual1.OperatorAddress,
					constants.PHASE_3_COMMIT_PREFIX,
				),
			)
			Expect(err).To(BeNil())
			Expect(node1CommitCount).To(Equal(int64(1)))

			node3CommitCount, err := crossfireValidatorsStatsView.FindBy(
				constants.ValidatorCommitmentKey(
					actual3.OperatorAddress,
					constants.PHASE_3_COMMIT_PREFIX,
				),
			)
			Expect(err).To(BeNil())
			Expect(node3CommitCount).To(Equal(int64(1)))

			actual1, err = crossfireValidatorsView.FindBy(view.CrossfireValidatorIdentity{
				MaybeOperatorAddress: &opAddr1,
			})
			Expect(err).To(BeNil())
			Expect(actual1.RankTaskPhase3CommitmentCount).To(Equal(int64(1))) // node1 rank #1

			actual3, err = crossfireValidatorsView.FindBy(view.CrossfireValidatorIdentity{
				MaybeOperatorAddress: &opAddr3,
			})
			Expect(err).To(BeNil())
			Expect(actual3.RankTaskPhase3CommitmentCount).To(Equal(int64(1))) // node3 rank #1

			// 3. store a phase 3 block that contains validator 1 & 2's commitment
			err = projection.HandleEvents(anyHeight, []event_entity.Event{phaseThreeBlockCreatedEvent2})
			Expect(err).To(BeNil())

			phase2BlockCount, err := crossfireChainStatsView.FindBy(constants.PHASE_3_BLOCK_COUNT)
			Expect(err).To(BeNil())
			Expect(phase2BlockCount).To(Equal(int64(2)))

			node1CommitCount, err = crossfireValidatorsStatsView.FindBy(
				constants.ValidatorCommitmentKey(
					actual1.OperatorAddress,
					constants.PHASE_3_COMMIT_PREFIX,
				),
			)
			Expect(err).To(BeNil())
			Expect(node1CommitCount).To(Equal(int64(2))) // node1 2 commitments

			node2CommitCount, err := crossfireValidatorsStatsView.FindBy(
				constants.ValidatorCommitmentKey(
					actual2.OperatorAddress,
					constants.PHASE_3_COMMIT_PREFIX,
				),
			)
			Expect(err).To(BeNil())
			Expect(node2CommitCount).To(Equal(int64(1))) // node2 1 commitments

			node3CommitCount, err = crossfireValidatorsStatsView.FindBy(
				constants.ValidatorCommitmentKey(
					actual3.OperatorAddress,
					constants.PHASE_3_COMMIT_PREFIX,
				),
			)
			Expect(err).To(BeNil())
			Expect(node3CommitCount).To(Equal(int64(1))) // node3 1 commitment

			actual1, err = crossfireValidatorsView.FindBy(view.CrossfireValidatorIdentity{
				MaybeOperatorAddress: &opAddr1,
			})
			Expect(err).To(BeNil())
			Expect(actual1.RankTaskPhase3CommitmentCount).To(Equal(int64(1))) // node1 rank #1

			actual2, err = crossfireValidatorsView.FindBy(view.CrossfireValidatorIdentity{
				MaybeOperatorAddress: &opAddr2,
			})
			Expect(err).To(BeNil())
			Expect(actual2.RankTaskPhase3CommitmentCount).To(Equal(int64(2))) // node2 rank #2

			actual3, err = crossfireValidatorsView.FindBy(view.CrossfireValidatorIdentity{
				MaybeOperatorAddress: &opAddr3,
			})
			Expect(err).To(BeNil())
			Expect(actual3.RankTaskPhase3CommitmentCount).To(Equal(int64(2))) // node3 rank #2
		})
	})
})
