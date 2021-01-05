package crossfire_test

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/crossfire"
	"github.com/crypto-com/chain-indexing/appinterface/projection/crossfire/constants"
	"github.com/crypto-com/chain-indexing/appinterface/projection/crossfire/view"
	. "github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	entity_projection "github.com/crypto-com/chain-indexing/entity/projection"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	. "github.com/crypto-com/chain-indexing/internal/logger/test"
	"github.com/crypto-com/chain-indexing/internal/primptr"
	"github.com/crypto-com/chain-indexing/internal/utctime"
	. "github.com/crypto-com/chain-indexing/test"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	usecase_model "github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Crossfire", func() {
	It("should implement projection", func() {
		fakeLogger := NewFakeLogger()
		fakeRdbConn := NewFakeRDbConn()
		var _ entity_projection.Projection = crossfire.NewCrossfire(
			fakeLogger,
			fakeRdbConn,
			"tcrocnclcons",
			"tcrocncl",
			1,
			1,
			1,
			1,
			"foo",
			"14",
			"https://raw.githubusercontent.com/foreseaz/random/master/participants.json",
		)
	})

	WithTestPgxConn(func(pgConn *pg.PgxConn, pgMigrate *pg.Migrate) {

		anyHeight := int64(1)

		validatorCreatedEvent := event_usecase.NewMsgCreateValidator(event_usecase.MsgCommonParams{
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
			DelegatorAddress:  "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
			ValidatorAddress:  "tcrocncl14m5a4kxt2e82uqqs5gtqza29dm5wqzyalddug5",
			TendermintPubkey:  "na51D8RmKXyWrid9I6wtdxgP6f1Nl3EyNNEzqxVquoM=",
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
			DelegatorAddress:  "tcro1f6qcvp33dc79xzpuwll7mln5lnepuqv8d7led9",
			ValidatorAddress:  "tcrocncl1f6qcvp33dc79xzpuwll7mln5lnepuqv8cpuq4x",
			TendermintPubkey:  "na51D8RmKXyWrid9I6wtdxgP6f1Nl3EyNNEzqxVquoM=",
			Amount:            coin.MustNewCoinFromString("222"),
		})

		phaseOneBlockCreatedEvent := event_usecase.NewBlockCreated(&usecase_model.Block{
			Height:          anyHeight,
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

		phaseTwoBlockCreatedEvent := event_usecase.NewBlockCreated(&usecase_model.Block{
			Height:          anyHeight,
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

		txCreatedParams := model.CreateTransactionParams{
			TxHash:   "A6D4C1F59A9D232747CA4F8A484F1F3B14A0075E801DF2A25F472B4280505B74",
			Code:     0,
			Log:      "{\"events\":[]}",
			MsgCount: 1,
			Signers: []model.TransactionSigner{
				{
					Type:            "/cosmos.crypto.secp256k1.PubKey",
					Pubkeys:         []string{"tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh"},
					AccountSequence: uint64(1),
				},
				{
					Type:            "/cosmos.crypto.secp256k1.PubKey",
					Pubkeys:         []string{"tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh"},
					AccountSequence: uint64(2),
				},
				{
					Type:            "/cosmos.crypto.secp256k1.PubKey",
					Pubkeys:         []string{"tcro1khkxmphc7sv0fqrej3rltsslrstud78c0jl6l6"},
					AccountSequence: uint64(1),
				},
				{
					Type:            "/cosmos.crypto.secp256k1.PubKey",
					Pubkeys:         []string{"tcro1f6qcvp33dc79xzpuwll7mln5lnepuqv8d7led9"},
					AccountSequence: uint64(2),
				},
				{
					Type:            "/cosmos.crypto.secp256k1.PubKey",
					Pubkeys:         []string{"tcro1f6qcvp33dc79xzpuwll7mln5lnepuqv8d7led9"},
					AccountSequence: uint64(3),
				},
				{
					Type: "/cosmos.crypto.multisig.LegacyAminoPubKey",
					Pubkeys: []string{
						"tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
						"tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
						"tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
					},
					MaybeThreshold:  primptr.Int(2),
					AccountSequence: uint64(1),
				},
			},
			Fee:           coin.MustNewCoinFromString("1000"),
			FeePayer:      "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
			FeeGranter:    "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
			GasWanted:     200000,
			GasUsed:       10000,
			Memo:          "Test memo",
			TimeoutHeight: int64(10),
		}

		transactionCreatedEvent := event_usecase.NewTransactionCreated(anyHeight, txCreatedParams)
		BeforeEach(func() {
			_ = pgMigrate.Reset()
			pgMigrate.MustUp()
		})

		AfterEach(func() {
			_ = pgMigrate.Reset()
		})

		It("should project crossfire_validators view when event is MsgCreateValidator", func() {
			crossfireValidatorsView := view.NewCrossfireValidators(pgConn.ToHandle())

			blockCreatedEvent := event_usecase.NewBlockCreated(&usecase_model.Block{
				Height:          anyHeight,
				Hash:            "B69554A020537DA8E7C7610A318180C09BFEB91229BB85D4A78DDA2FACF68A48",
				Time:            utctime.FromUnixNano(int64(1000000)),
				AppHash:         "24474D86CBFA7E6328D473C17A9E46CD5A80FFE82A348A74844BF3E2BA2B3AF1",
				ProposerAddress: "F9E6FFB9B536956201AA138224FD888D03775AB4",
				Txs: []string{
					"AAAMZqICtpjLrA3uEe3Rkg6cqDgQl0iBwG1Wm8ORZRzKL9EBAE1R0oP73H8AhHG1E1dke4ppXA/+43AyxWX7oVBTlQGnAABsiqVxCipneyzo9n1t7xauih3rXxssZTChLv/oPHUaSAAAAgAAAAAAAAAAAMw4kh8pZRDC442WIIUIo6+fIE1cTHkwEiv++LwFjNLwNtPTO1AVJx+dGivu9FZK3cN8r0mvLPVunExBxwKqKDkwVUbVSaIejuSc6Wm38rp+ekZjIExhf41FIwkocStPzBY4VYQ31YE7pmNnKocfCEg9CcKH+MIFYJ0FSxoEiMratLXjKAD6B/aMZ3pRCxxL/YxFNs+YKzGbgdKcLtydZ1B/OTGUeNU3SLNQJGmJeTL0AOwvQ4j7KkiQFZGa8uoQQEJo9rXjIg+xPk7a+zBdwC91EY/ZZz7pdVQMUgHBeKgsVvy0o78BOPWfuziBK5xVsoz5K8ZlMYQNPC5mY+bguMKYNJ5PQ9rzn+LseYT5jhalUsQrPABhEFOoQVUH1id3rszDkLLTn2/d89N/JJGY1+mL+upWWAsmJw1yTHqibSJ7RbiGffnh93MCOHeRe5OLrfmDFfhqOLNt9yuMYNdyJ+noVsfI7Ws9Kpxe2SnuCWBs9yOgM0l7UTMuIokqhGCkarXge/DUWLUcV694Jr8qcJT3OtoQohN+p+Xj66nowJLgFW7xBdFsT4vv7xI8giFxUpB+JkLgRZz2d2eam3iLDCHR+sNyvIDuXDUXhKM6aIykGgvHVHr3bBJRy/JPZFC0A3kqmheMnJpbV6kHXaIbmbgeQeXc+wxq1swFVxkwp14zbRmHwSGVGRgihjmoFYl9MyorQzNFETgp28gq2AKrCHNnIRs63m2cD5X4jZaFzfYAv9ifpOKqRtDgIFG0olge/ig00FFL6KdHySg1Qaab7g==",
				},
				Signatures: []usecase_model.BlockSignature{
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
				"https://raw.githubusercontent.com/foreseaz/random/master/participants.json",
			)

			Expect(crossfireValidatorsView.Count()).To(Equal(int64(0)))

			// project two events
			err := projection.HandleEvents(anyHeight, []event_entity.Event{blockCreatedEvent, validatorCreatedEvent})
			Expect(err).To(BeNil())

			count, err := crossfireValidatorsView.Count()
			Expect(err).To(BeNil())
			Expect(count).To(Equal(int64(1)))

			// list all
			list, err := crossfireValidatorsView.List()
			Expect(err).To(BeNil())
			var targetList []view.CrossfireValidatorRow
			id := int64(1)
			targetList = append(targetList, view.CrossfireValidatorRow{
				MaybeId:                         &id,
				OperatorAddress:                 "tcrocncl14m5a4kxt2e82uqqs5gtqza29dm5wqzyalddug5",
				ConsensusNodeAddress:            "tcrocnclcons1khkxmphc7sv0fqrej3rltsslrstud78cam9ekl",
				InitialDelegatorAddress:         "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
				TendermintPubkey:                "na51D8RmKXyWrid9I6wtdxgP6f1Nl3EyNNEzqxVquoM=",
				TendermintAddress:               "B5EC6D86F8F418F480799447F5C21F1C17C6F8F8",
				Status:                          "Unbonded",
				Jailed:                          false,
				JoinedAtBlockHeight:             1,
				JoinedAtBlockTime:               utctime.FromUnixNano(int64(1000000)),
				Moniker:                         "Testing",
				Identity:                        "foo",
				Website:                         "www.example.com",
				SecurityContact:                 "foo@example.com",
				Details:                         "example",
				TaskPhase1NodeSetup:             "Completed",
				TaskPhase2KeepNodeActive:        "Incompleted",
				TaskPhase2ProposalVote:          "Incompleted",
				TaskPhase2NetworkUpgrade:        "Incompleted",
				RankTaskPhase1n2CommitmentCount: 0,
				RankTaskPhase3CommitmentCount:   0,
				RankTaskHighestTxSent:           0,
			})
			Expect(list).To(Equal(targetList))
		})

		It("should check task setup Missed correctly", func() {
			crossfireValidatorsView := view.NewCrossfireValidators(pgConn.ToHandle())

			blockCreatedEvent := event_usecase.NewBlockCreated(&usecase_model.Block{
				Height:          anyHeight,
				Hash:            "B69554A020537DA8E7C7610A318180C09BFEB91229BB85D4A78DDA2FACF68A48",
				Time:            utctime.FromUnixNano(int64(1611547300000000000)), // passed phase 2
				AppHash:         "24474D86CBFA7E6328D473C17A9E46CD5A80FFE82A348A74844BF3E2BA2B3AF1",
				ProposerAddress: "F9E6FFB9B536956201AA138224FD888D03775AB4",
				Txs: []string{
					"AAAMZqICtpjLrA3uEe3Rkg6cqDgQl0iBwG1Wm8ORZRzKL9EBAE1R0oP73H8AhHG1E1dke4ppXA/+43AyxWX7oVBTlQGnAABsiqVxCipneyzo9n1t7xauih3rXxssZTChLv/oPHUaSAAAAgAAAAAAAAAAAMw4kh8pZRDC442WIIUIo6+fIE1cTHkwEiv++LwFjNLwNtPTO1AVJx+dGivu9FZK3cN8r0mvLPVunExBxwKqKDkwVUbVSaIejuSc6Wm38rp+ekZjIExhf41FIwkocStPzBY4VYQ31YE7pmNnKocfCEg9CcKH+MIFYJ0FSxoEiMratLXjKAD6B/aMZ3pRCxxL/YxFNs+YKzGbgdKcLtydZ1B/OTGUeNU3SLNQJGmJeTL0AOwvQ4j7KkiQFZGa8uoQQEJo9rXjIg+xPk7a+zBdwC91EY/ZZz7pdVQMUgHBeKgsVvy0o78BOPWfuziBK5xVsoz5K8ZlMYQNPC5mY+bguMKYNJ5PQ9rzn+LseYT5jhalUsQrPABhEFOoQVUH1id3rszDkLLTn2/d89N/JJGY1+mL+upWWAsmJw1yTHqibSJ7RbiGffnh93MCOHeRe5OLrfmDFfhqOLNt9yuMYNdyJ+noVsfI7Ws9Kpxe2SnuCWBs9yOgM0l7UTMuIokqhGCkarXge/DUWLUcV694Jr8qcJT3OtoQohN+p+Xj66nowJLgFW7xBdFsT4vv7xI8giFxUpB+JkLgRZz2d2eam3iLDCHR+sNyvIDuXDUXhKM6aIykGgvHVHr3bBJRy/JPZFC0A3kqmheMnJpbV6kHXaIbmbgeQeXc+wxq1swFVxkwp14zbRmHwSGVGRgihjmoFYl9MyorQzNFETgp28gq2AKrCHNnIRs63m2cD5X4jZaFzfYAv9ifpOKqRtDgIFG0olge/ig00FFL6KdHySg1Qaab7g==",
				},
				Signatures: []usecase_model.BlockSignature{},
			})

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
				"https://raw.githubusercontent.com/foreseaz/random/master/participants.json",
			)

			Expect(crossfireValidatorsView.Count()).To(Equal(int64(0)))

			// project two events
			err := projection.HandleEvents(anyHeight, []event_entity.Event{blockCreatedEvent, validatorCreatedEvent})
			Expect(err).To(BeNil())

			// list all
			list, err := crossfireValidatorsView.List()
			Expect(err).To(BeNil())
			var targetList []view.CrossfireValidatorRow
			id := int64(1)
			targetList = append(targetList, view.CrossfireValidatorRow{
				MaybeId:                         &id,
				OperatorAddress:                 "tcrocncl14m5a4kxt2e82uqqs5gtqza29dm5wqzyalddug5",
				ConsensusNodeAddress:            "tcrocnclcons1khkxmphc7sv0fqrej3rltsslrstud78cam9ekl",
				InitialDelegatorAddress:         "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
				TendermintPubkey:                "na51D8RmKXyWrid9I6wtdxgP6f1Nl3EyNNEzqxVquoM=",
				TendermintAddress:               "B5EC6D86F8F418F480799447F5C21F1C17C6F8F8",
				Status:                          "Unbonded",
				Jailed:                          false,
				JoinedAtBlockHeight:             1,
				JoinedAtBlockTime:               utctime.FromUnixNano(int64(1611547300000000000)),
				Moniker:                         "Testing",
				Identity:                        "foo",
				Website:                         "www.example.com",
				SecurityContact:                 "foo@example.com",
				Details:                         "example",
				TaskPhase1NodeSetup:             "Missed",
				TaskPhase2KeepNodeActive:        "Incompleted",
				TaskPhase2ProposalVote:          "Incompleted",
				TaskPhase2NetworkUpgrade:        "Incompleted",
				RankTaskPhase1n2CommitmentCount: 0,
				RankTaskPhase3CommitmentCount:   0,
				RankTaskHighestTxSent:           0,
			})
			Expect(list).To(Equal(targetList))
		})

		It("should project crossfire_validators view when event is MsgSubmitSoftwareUpgradeProposal", func() {
			crossfireChainStatsView := view.NewCrossfireChainStats(pgConn.ToHandle())

			blockCreatedEvent := event_usecase.NewBlockCreated(&usecase_model.Block{
				Height:          anyHeight,
				Hash:            "B69554A020537DA8E7C7610A318180C09BFEB91229BB85D4A78DDA2FACF68A48",
				Time:            utctime.FromUnixNano(int64(1612056800000000000)),
				AppHash:         "24474D86CBFA7E6328D473C17A9E46CD5A80FFE82A348A74844BF3E2BA2B3AF1",
				ProposerAddress: "F9E6FFB9B536956201AA138224FD888D03775AB4",
				Txs: []string{
					"AAAMZqICtpjLrA3uEe3Rkg6cqDgQl0iBwG1Wm8ORZRzKL9EBAE1R0oP73H8AhHG1E1dke4ppXA/+43AyxWX7oVBTlQGnAABsiqVxCipneyzo9n1t7xauih3rXxssZTChLv/oPHUaSAAAAgAAAAAAAAAAAMw4kh8pZRDC442WIIUIo6+fIE1cTHkwEiv++LwFjNLwNtPTO1AVJx+dGivu9FZK3cN8r0mvLPVunExBxwKqKDkwVUbVSaIejuSc6Wm38rp+ekZjIExhf41FIwkocStPzBY4VYQ31YE7pmNnKocfCEg9CcKH+MIFYJ0FSxoEiMratLXjKAD6B/aMZ3pRCxxL/YxFNs+YKzGbgdKcLtydZ1B/OTGUeNU3SLNQJGmJeTL0AOwvQ4j7KkiQFZGa8uoQQEJo9rXjIg+xPk7a+zBdwC91EY/ZZz7pdVQMUgHBeKgsVvy0o78BOPWfuziBK5xVsoz5K8ZlMYQNPC5mY+bguMKYNJ5PQ9rzn+LseYT5jhalUsQrPABhEFOoQVUH1id3rszDkLLTn2/d89N/JJGY1+mL+upWWAsmJw1yTHqibSJ7RbiGffnh93MCOHeRe5OLrfmDFfhqOLNt9yuMYNdyJ+noVsfI7Ws9Kpxe2SnuCWBs9yOgM0l7UTMuIokqhGCkarXge/DUWLUcV694Jr8qcJT3OtoQohN+p+Xj66nowJLgFW7xBdFsT4vv7xI8giFxUpB+JkLgRZz2d2eam3iLDCHR+sNyvIDuXDUXhKM6aIykGgvHVHr3bBJRy/JPZFC0A3kqmheMnJpbV6kHXaIbmbgeQeXc+wxq1swFVxkwp14zbRmHwSGVGRgihjmoFYl9MyorQzNFETgp28gq2AKrCHNnIRs63m2cD5X4jZaFzfYAv9ifpOKqRtDgIFG0olge/ig00FFL6KdHySg1Qaab7g==",
				},
				Signatures: []usecase_model.BlockSignature{
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

			softwareUpgradeEvent := event_usecase.NewMsgSubmitSoftwareUpgradeProposal(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      "A6D4C1F59A9D232747CA4F8A484F1F3B14A0075E801DF2A25F472B4280505B74",
				TxSuccess:   true,
				MsgIndex:    0,
			}, usecase_model.MsgSubmitSoftwareUpgradeProposalParams{
				MaybeProposalId: primptr.String("14"),
				Content: model.MsgSubmitSoftwareUpgradeProposalContent{
					Type:        "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal",
					Title:       "Upgrade Title",
					Description: "Upgrade Description",
					Plan: model.MsgSubmitSoftwareUpgradeProposalPlan{
						Name:   "Upgrade Name",
						Time:   utctime.FromUnixNano(int64(-6795364578871345152)),
						Height: 100000000,
						Info:   "Upgrade Info",
					},
				},
				ProposerAddress: "tcro15grftg88l0gdw4mg9t9pwnl0pde2asjzvfpkp4",
				InitialDeposit:  coin.Zero(),
			})

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
				"https://raw.githubusercontent.com/foreseaz/random/master/participants.json",
			)

			// Fire both events
			err := projection.HandleEvents(anyHeight, []event_entity.Event{blockCreatedEvent, softwareUpgradeEvent})
			Expect(err).To(BeNil())

			//Check updated blockheight
			targetBlockheight, err := crossfireChainStatsView.FindBy("network_upgrade:blockheight")
			Expect(err).To(BeNil())
			Expect(targetBlockheight).To(Equal(int64(100000000)))

			//Check updated timestamp
			targetTimestamp, err := crossfireChainStatsView.FindBy("network_upgrade:timestamp")
			Expect(err).To(BeNil())
			Expect(targetTimestamp).To(Equal(int64(-6795364578871345152)))
		})

		It("should project crossfire_validators view when event is MsgVoteCreated", func() {
			crossfireValidatorStatsView := view.NewCrossfireValidatorsStats(pgConn.ToHandle())
			crossfireValidatorView := view.NewCrossfireValidators(pgConn.ToHandle())

			blockCreatedEvent := event_usecase.NewBlockCreated(&usecase_model.Block{
				Height:          anyHeight,
				Hash:            "B69554A020537DA8E7C7610A318180C09BFEB91229BB85D4A78DDA2FACF68A48",
				Time:            utctime.FromUnixNano(int64(1612056800000000000)),
				AppHash:         "24474D86CBFA7E6328D473C17A9E46CD5A80FFE82A348A74844BF3E2BA2B3AF1",
				ProposerAddress: "F9E6FFB9B536956201AA138224FD888D03775AB4",
				Txs: []string{
					"AAAMZqICtpjLrA3uEe3Rkg6cqDgQl0iBwG1Wm8ORZRzKL9EBAE1R0oP73H8AhHG1E1dke4ppXA/+43AyxWX7oVBTlQGnAABsiqVxCipneyzo9n1t7xauih3rXxssZTChLv/oPHUaSAAAAgAAAAAAAAAAAMw4kh8pZRDC442WIIUIo6+fIE1cTHkwEiv++LwFjNLwNtPTO1AVJx+dGivu9FZK3cN8r0mvLPVunExBxwKqKDkwVUbVSaIejuSc6Wm38rp+ekZjIExhf41FIwkocStPzBY4VYQ31YE7pmNnKocfCEg9CcKH+MIFYJ0FSxoEiMratLXjKAD6B/aMZ3pRCxxL/YxFNs+YKzGbgdKcLtydZ1B/OTGUeNU3SLNQJGmJeTL0AOwvQ4j7KkiQFZGa8uoQQEJo9rXjIg+xPk7a+zBdwC91EY/ZZz7pdVQMUgHBeKgsVvy0o78BOPWfuziBK5xVsoz5K8ZlMYQNPC5mY+bguMKYNJ5PQ9rzn+LseYT5jhalUsQrPABhEFOoQVUH1id3rszDkLLTn2/d89N/JJGY1+mL+upWWAsmJw1yTHqibSJ7RbiGffnh93MCOHeRe5OLrfmDFfhqOLNt9yuMYNdyJ+noVsfI7Ws9Kpxe2SnuCWBs9yOgM0l7UTMuIokqhGCkarXge/DUWLUcV694Jr8qcJT3OtoQohN+p+Xj66nowJLgFW7xBdFsT4vv7xI8giFxUpB+JkLgRZz2d2eam3iLDCHR+sNyvIDuXDUXhKM6aIykGgvHVHr3bBJRy/JPZFC0A3kqmheMnJpbV6kHXaIbmbgeQeXc+wxq1swFVxkwp14zbRmHwSGVGRgihjmoFYl9MyorQzNFETgp28gq2AKrCHNnIRs63m2cD5X4jZaFzfYAv9ifpOKqRtDgIFG0olge/ig00FFL6KdHySg1Qaab7g==",
				},
				Signatures: []usecase_model.BlockSignature{
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

			msgVoteEvent := event_usecase.NewMsgVote(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      "A6D4C1F59A9D232747CA4F8A484F1F3B14A0075E801DF2A25F472B4280505B74",
				TxSuccess:   true,
				MsgIndex:    0,
			}, usecase_model.MsgVoteParams{
				ProposalId: "14",
				Voter:      "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
				Option:     "VOTE_OPTION_YES",
			})

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
				"https://raw.githubusercontent.com/foreseaz/random/master/participants.json",
			)

			// Fire both events
			err := projection.HandleEvents(anyHeight, []event_entity.Event{blockCreatedEvent, validatorCreatedEvent, msgVoteEvent})
			Expect(err).To(BeNil())

			//Check validators view count
			validatorViewCount, err := crossfireValidatorView.Count()
			Expect(err).To(BeNil())
			Expect(validatorViewCount).To(Equal(int64(1)))

			//check Validator status
			crossfireValidatorList, err := crossfireValidatorView.List()
			Expect(err).To(BeNil())
			Expect(crossfireValidatorList).To(HaveLen(1))
			Expect(crossfireValidatorList[0].TaskPhase2ProposalVote).To(Equal(constants.COMPLETED))

			//Check voted proposal id for voter
			voted_proposal_id, err := crossfireValidatorStatsView.FindBy("voted_proposal_id:tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh")
			Expect(err).To(BeNil())
			Expect(voted_proposal_id).To(Equal(int64(14)))

		})

		It("should project crossfire_validators view when event is TransactionCreated", func() {
			crossfireValidatorStatsView := view.NewCrossfireValidatorsStats(pgConn.ToHandle())
			crossfireValidatorView := view.NewCrossfireValidators(pgConn.ToHandle())

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
				"https://raw.githubusercontent.com/foreseaz/random/master/participants.json",
			)

			// Fire both events
			err := projection.HandleEvents(anyHeight, []event_entity.Event{phaseOneBlockCreatedEvent, validatorCreatedEvent, transactionCreatedEvent})
			errHandleNewBlock := projection.HandleEvents(anyHeight, []event_entity.Event{phaseTwoBlockCreatedEvent, transactionCreatedEvent})
			Expect(err).To(BeNil())
			Expect(errHandleNewBlock).To(BeNil())

			//Check validators view count
			validatorViewCount, err := crossfireValidatorView.Count()
			Expect(err).To(BeNil())
			Expect(validatorViewCount).To(Equal(int64(1)))

			//check Validator status
			crossfireValidatorList, err := crossfireValidatorView.List()
			Expect(err).To(BeNil())
			Expect(crossfireValidatorList).To(HaveLen(1))
			Expect(crossfireValidatorList[0].TaskPhase1NodeSetup).To(Equal(constants.COMPLETED))

			// total count for both address
			totalCountByAddress1, errTotalCountAddress1 := crossfireValidatorStatsView.FindBy("total_tx_sent:tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh")
			totalCountByAddress2, errTotalCountAddress2 := crossfireValidatorStatsView.FindBy("total_tx_sent:tcro1khkxmphc7sv0fqrej3rltsslrstud78c0jl6l6")

			// Phase One Count for both address
			phaseOneCountByAddress1, errPhaseOneCountAddress1 := crossfireValidatorStatsView.FindBy("phase_1_tx_sent:tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh")
			phaseOneCountByAddress2, errPhaseOneCountAddress2 := crossfireValidatorStatsView.FindBy("phase_1_tx_sent:tcro1khkxmphc7sv0fqrej3rltsslrstud78c0jl6l6")

			// Phase Two Count for both address
			phaseTwoCountByAddress1, errPhaseTwoCountAddress1 := crossfireValidatorStatsView.FindBy("phase_2_tx_sent:tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh")
			phaseTwoCountByAddress2, errPhaseTwoCountAddress2 := crossfireValidatorStatsView.FindBy("phase_2_tx_sent:tcro1khkxmphc7sv0fqrej3rltsslrstud78c0jl6l6")

			Expect(errTotalCountAddress1).To(BeNil())
			Expect(errTotalCountAddress2).To(BeNil())

			Expect(errPhaseOneCountAddress1).To(BeNil())
			Expect(errPhaseOneCountAddress2).To(BeNil())

			Expect(errPhaseTwoCountAddress1).To(BeNil())
			Expect(errPhaseTwoCountAddress2).To(BeNil())

			//Total Count checks
			Expect(totalCountByAddress1).To(Equal(int64(4)))
			Expect(totalCountByAddress2).To(Equal(int64(2)))

			//Phase One Count checks
			Expect(phaseOneCountByAddress1).To(Equal(int64(2)))
			Expect(phaseOneCountByAddress2).To(Equal(int64(1)))

			//Phase Two Count checks
			Expect(phaseTwoCountByAddress1).To(Equal(int64(2)))
			Expect(phaseTwoCountByAddress2).To(Equal(int64(1)))

		})

		It("should correctly compute Tx Sent Rank and update in DB", func() {
			crossfireValidatorView := view.NewCrossfireValidators(pgConn.ToHandle())

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
				"https://raw.githubusercontent.com/foreseaz/random/master/participants.json",
			)

			// Fire both events
			err := projection.HandleEvents(anyHeight, []event_entity.Event{phaseOneBlockCreatedEvent, validatorCreatedEvent, validator2CreatedEvent, transactionCreatedEvent})
			// errHandleNewBlock := projection.HandleEvents(anyHeight, []event_entity.Event{phaseTwoBlockCreatedEvent, transactionCreatedEvent})
			Expect(err).To(BeNil())
			// Expect(errHandleNewBlock).To(BeNil())

			//Check validators view count
			validatorViewCount, err := crossfireValidatorView.Count()
			Expect(err).To(BeNil())
			Expect(validatorViewCount).To(Equal(int64(2)))

			//check Validator status
			crossfireValidatorList, err := crossfireValidatorView.List()
			Expect(err).To(BeNil())
			Expect(crossfireValidatorList).To(HaveLen(2))
			Expect(crossfireValidatorList[0].TaskPhase1NodeSetup).To(Equal(constants.COMPLETED))

			Expect(crossfireValidatorList[0].OperatorAddress).To(Equal("tcrocncl14m5a4kxt2e82uqqs5gtqza29dm5wqzyalddug5"))
			Expect(crossfireValidatorList[0].InitialDelegatorAddress).To(Equal("tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh"))
			Expect(crossfireValidatorList[0].RankTaskHighestTxSent).To(Equal(int64(0)))

			Expect(crossfireValidatorList[1].OperatorAddress).To(Equal("tcrocncl1f6qcvp33dc79xzpuwll7mln5lnepuqv8cpuq4x"))
			Expect(crossfireValidatorList[1].InitialDelegatorAddress).To(Equal("tcro1f6qcvp33dc79xzpuwll7mln5lnepuqv8d7led9"))
			Expect(crossfireValidatorList[1].RankTaskHighestTxSent).To(Equal(int64(0)))

		})
	})
})
