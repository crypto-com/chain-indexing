package validator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	entity_projection "github.com/crypto-com/chain-indexing/entity/projection"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	. "github.com/crypto-com/chain-indexing/internal/logger/test"
	"github.com/crypto-com/chain-indexing/internal/primptr"
	"github.com/crypto-com/chain-indexing/internal/utctime"
	"github.com/crypto-com/chain-indexing/projection/validator"
	"github.com/crypto-com/chain-indexing/projection/validator/constants"
	validator_view "github.com/crypto-com/chain-indexing/projection/validator/view"
	. "github.com/crypto-com/chain-indexing/test"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	usecase_model "github.com/crypto-com/chain-indexing/usecase/model"
)

const prefixConsensusAddress string = "crocnclcons"

var _ = Describe("Validator Events", func() {
	It("should implement projection", func() {
		fakeLogger := NewFakeLogger()
		fakeRdbConn := NewFakeRDbConn()
		var _ entity_projection.Projection = validator.NewValidator(fakeLogger, fakeRdbConn, prefixConsensusAddress)
	})

	WithTestPgxConn(func(pgConn *pg.PgxConn, pgMigrate *pg.Migrate) {
		BeforeEach(func() {
			_ = pgMigrate.Reset()
			pgMigrate.MustUp()
		})

		AfterEach(func() {
			_ = pgMigrate.Reset()
		})

		It("should match the last projection handled block height with the fired event block height", func() {
			anyHeight := int64(1)

			description := model.MsgValidatorDescription{
				Moniker:         "mymonicker",
				Identity:        "myidentity",
				Website:         "mywebsite",
				SecurityContact: "mysecuritycontact",
				Details:         "mydetails",
			}
			commission := model.MsgValidatorCommission{
				Rate:          "0.100000000000000000",
				MaxRate:       "0.200000000000000000",
				MaxChangeRate: "0.010000000000000000",
			}

			createValidatorParams := usecase_model.MsgCreateValidatorParams{
				Description:       description,
				Commission:        commission,
				MinSelfDelegation: "1",
				DelegatorAddress:  "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
				ValidatorAddress:  "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
				TendermintPubkey:  "Kpox5fS2po0sJUHmzllExuJ4uZ5nm0bbCp6UQKESsnE=",
				Amount:            coin.MustParseCoinNormalized("10basetcro"),
			}
			event := event_usecase.NewBlockCreated(&usecase_model.Block{
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
						ValidatorAddress: "6B26ECB33BC875DFD4457867C183F6370D192B69",
						Timestamp:        utctime.FromUnixNano(int64(1000000)),
						Signature:        "ZW2pUcKFN/oPQCmdCouchXmgpPyd/Ddo45dhHEMwsBeHTBuSJh15zUMmfl5FZsPHeKC8citFvOm/52bgl5XHCw==",
					},
				},
			})

			eventNewValidatorCreated := event_usecase.NewMsgCreateValidator(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      "E69985AC8168383A81B7952DBE03EB9B3400FF80AEC0F362369DD7F38B1C2FE9",
				TxSuccess:   true,
				MsgIndex:    0,
			}, createValidatorParams)
			fakeLogger := NewFakeLogger()

			projection := validator.NewValidator(fakeLogger, pgConn, prefixConsensusAddress)

			err := projection.HandleEvents(anyHeight, []event_entity.Event{
				event,
				eventNewValidatorCreated,
			})
			Expect(err).To(BeNil())

			Expect(projection.GetLastHandledEventHeight()).To(Equal(primptr.Int64(anyHeight)))
		})

		It("should increment the validator view count after handling MsgCreateValidator event", func() {
			validatorView := validator_view.NewValidators(pgConn.ToHandle())

			anyHeight := int64(1)

			description := model.MsgValidatorDescription{
				Moniker:         "mymonicker",
				Identity:        "myidentity",
				Website:         "mywebsite",
				SecurityContact: "mysecuritycontact",
				Details:         "mydetails",
			}
			commission := model.MsgValidatorCommission{
				Rate:          "0.100000000000000000",
				MaxRate:       "0.200000000000000000",
				MaxChangeRate: "0.010000000000000000",
			}

			createValidatorParams := usecase_model.MsgCreateValidatorParams{
				Description:       description,
				Commission:        commission,
				MinSelfDelegation: "1",
				DelegatorAddress:  "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
				ValidatorAddress:  "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
				TendermintPubkey:  "Kpox5fS2po0sJUHmzllExuJ4uZ5nm0bbCp6UQKESsnE=",
				Amount:            coin.MustParseCoinNormalized("10basetcro"),
			}
			event := event_usecase.NewBlockCreated(&usecase_model.Block{
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
						ValidatorAddress: "6B26ECB33BC875DFD4457867C183F6370D192B69",
						Timestamp:        utctime.FromUnixNano(int64(1000000)),
						Signature:        "ZW2pUcKFN/oPQCmdCouchXmgpPyd/Ddo45dhHEMwsBeHTBuSJh15zUMmfl5FZsPHeKC8citFvOm/52bgl5XHCw==",
					},
				},
			})

			eventNewValidatorCreated := event_usecase.NewMsgCreateValidator(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      "E69985AC8168383A81B7952DBE03EB9B3400FF80AEC0F362369DD7F38B1C2FE9",
				TxSuccess:   true,
				MsgIndex:    0,
			}, createValidatorParams)
			countFilter := validator_view.CountFilter{
				MaybeStatus: []string{
					constants.UNBONDED,
					constants.BONDED,
					constants.JAILED,
					constants.UNBONDING,
				},
			}
			fakeLogger := NewFakeLogger()

			projection := validator.NewValidator(fakeLogger, pgConn, prefixConsensusAddress)

			validatorViewCountBeforeHandling, err := validatorView.Count(countFilter)

			//check before handling event
			Expect(validatorViewCountBeforeHandling).To(Equal(int64(0)))
			Expect(err).To(BeNil())

			//handle event below
			err = projection.HandleEvents(anyHeight, []event_entity.Event{event, eventNewValidatorCreated})
			Expect(err).To(BeNil())

			//check the list after event handling
			validatorViewCountAfterHandling, errAfterHandling := validatorView.Count(countFilter)

			//check before handling event
			Expect(errAfterHandling).To(BeNil())
			Expect(projection.GetLastHandledEventHeight()).To(Equal(primptr.Int64(anyHeight)))
			Expect(validatorViewCountAfterHandling).To(Equal(int64(1)))
		})

		It("should update projection last handled event height when there is no event at the height", func() {
			anyHeight := int64(1)

			fakeLogger := NewFakeLogger()
			projection := validator.NewValidator(fakeLogger, pgConn, prefixConsensusAddress)

			Expect(projection.GetLastHandledEventHeight()).To(BeNil())

			err := projection.HandleEvents(anyHeight, []event_entity.Event{})
			Expect(err).To(BeNil())

			Expect(projection.GetLastHandledEventHeight()).To(Equal(primptr.Int64(anyHeight)))
		})
	})
})
