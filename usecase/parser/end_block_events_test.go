package parser_test

import (
	"github.com/crypto-com/chain-indexing/external/utctime"
	types "github.com/crypto-com/chain-indexing/projection/block_raw_event/constants"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/entity/command"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseEndBlockEventsCommands", func() {
	It("should return EndProposal commands when end_block_events has proposal_active event", func() {
		blockResults := mustParseBlockResultsResp(usecase_parser_test.END_BLOCK_PROPOSAL_REJECTED_BLOCK_RESULTS_RESP)
		block, _ := mustParseBlockResp(usecase_parser_test.END_BLOCK_COMPLETE_UNBONDING_BLOCK_RESP)

		cmds, err := parser.ParseEndBlockEventsCommands(
			blockResults.Height,
			block.Hash,
			block.Time,
			blockResults.EndBlockEvents,
		)
		Expect(err).To(BeNil())
		Expect(cmds).To(HaveLen(2))
		expectedBlockHeight := int64(21575)
		expectedProposalId := "1"
		expectedResult := "proposal_rejected"
		Expect(cmds).To(Equal([]command.Command{
			command_usecase.NewEndProposal(
				expectedBlockHeight,
				expectedProposalId,
				expectedResult,
			),
			command_usecase.NewCreateBlockRawEvent(
				expectedBlockHeight,
				model.CreateBlockRawEventParams{
					BlockHash:  "8703C54C9FE1C2D6D05DAC79D795E120F385F5F43E5CDC17B73090E9DA40CEA9",
					BlockTime:  utctime.FromUnixNano(1631893335936780880),
					FromResult: types.END_BLOCK_EVENT,
					Data: model.DataParams{
						Type: "active_proposal",
						Content: model.BlockResultsEvent{
							Type: "active_proposal",
							Attributes: []model.BlockResultsEventAttribute{
								{
									Key:   "proposal_id",
									Value: "1",
									Index: true,
								},
								{
									Key:   "proposal_result",
									Value: "proposal_rejected",
									Index: true,
								},
							},
						},
					},
				},
			),
		}))
	})

	It("should return EndProposal commands when end_blocks_events has proposal_active passed event", func() {
		blockResults := mustParseBlockResultsResp(usecase_parser_test.END_BLOCK_PROPOSAL_PASSED_BLOCK_RESULTS_RESP)
		block, _ := mustParseBlockResp(usecase_parser_test.END_BLOCK_COMPLETE_UNBONDING_BLOCK_RESP)

		cmds, err := parser.ParseEndBlockEventsCommands(
			blockResults.Height,
			block.Hash,
			block.Time,
			blockResults.EndBlockEvents,
		)
		Expect(err).To(BeNil())
		Expect(cmds).To(HaveLen(2))
		expectedBlockHeight := int64(520186)
		expectedProposalId := "7"
		expectedResult := "proposal_passed"
		Expect(cmds).To(Equal([]command.Command{
			command_usecase.NewEndProposal(
				expectedBlockHeight,
				expectedProposalId,
				expectedResult,
			),
			command_usecase.NewCreateBlockRawEvent(
				expectedBlockHeight,
				model.CreateBlockRawEventParams{
					BlockHash:  "8703C54C9FE1C2D6D05DAC79D795E120F385F5F43E5CDC17B73090E9DA40CEA9",
					BlockTime:  utctime.FromUnixNano(1631893335936780880),
					FromResult: types.END_BLOCK_EVENT,
					Data: model.DataParams{
						Type: "active_proposal",
						Content: model.BlockResultsEvent{
							Type: "active_proposal",
							Attributes: []model.BlockResultsEventAttribute{
								{
									Key:   "proposal_id",
									Value: "7",
									Index: true,
								},
								{
									Key:   "proposal_result",
									Value: "proposal_passed",
									Index: true,
								},
							},
						},
					},
				},
			),
		}))
	})

	It("should return InactiveProposal commands when end_blocks_events has proposal_inactive event", func() {
		blockResults := mustParseBlockResultsResp(usecase_parser_test.END_BLOCK_PROPOSAL_INACTIVED_BLOCK_RESULTS_RESP)
		block, _ := mustParseBlockResp(usecase_parser_test.END_BLOCK_COMPLETE_UNBONDING_BLOCK_RESP)

		cmds, err := parser.ParseEndBlockEventsCommands(
			blockResults.Height,
			block.Hash,
			block.Time,
			blockResults.EndBlockEvents,
		)
		Expect(err).To(BeNil())
		Expect(cmds).To(HaveLen(2))
		expectedBlockHeight := int64(21541)
		expectedProposalId := "2"
		expectedResult := "proposal_dropped"
		Expect(cmds).To(Equal([]command.Command{
			command_usecase.NewInactiveProposal(
				expectedBlockHeight,
				expectedProposalId,
				expectedResult,
			),
			command_usecase.NewCreateBlockRawEvent(
				expectedBlockHeight,
				model.CreateBlockRawEventParams{
					BlockHash:  "8703C54C9FE1C2D6D05DAC79D795E120F385F5F43E5CDC17B73090E9DA40CEA9",
					BlockTime:  utctime.FromUnixNano(1631893335936780880),
					FromResult: types.END_BLOCK_EVENT,
					Data: model.DataParams{
						Type: "inactive_proposal",
						Content: model.BlockResultsEvent{
							Type: "inactive_proposal",
							Attributes: []model.BlockResultsEventAttribute{
								{
									Key:   "proposal_id",
									Value: "2",
									Index: true,
								},
								{
									Key:   "proposal_result",
									Value: "proposal_dropped",
									Index: true,
								},
							},
						},
					},
				},
			),
		}))
	})

	It("should return CompleteBonding commands when end_blocks_events has complete_unbonding event", func() {
		blockResults := mustParseBlockResultsResp(usecase_parser_test.END_BLOCK_COMPLETE_UNBONDING_BLOCK_RESULTS_RESP)
		block, _ := mustParseBlockResp(usecase_parser_test.END_BLOCK_COMPLETE_UNBONDING_BLOCK_RESP)

		cmds, err := parser.ParseEndBlockEventsCommands(
			blockResults.Height,
			block.Hash,
			block.Time,
			blockResults.EndBlockEvents,
		)
		Expect(err).To(BeNil())
		Expect(cmds).To(HaveLen(2))
		expectedBlockHeight := int64(477415)
		Expect(cmds).To(Equal([]command.Command{
			command_usecase.NewCreateCompleteBonding(
				expectedBlockHeight,
				model.CompleteBondingParams{
					Delegator: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
					Validator: "tcrocncl1sruzd529lhjju6hfcwd2fxp3v0e7p0vqqtme76",
					Amount:    coin.MustParseCoinsNormalized("5basetcro"),
				},
			),
			command_usecase.NewCreateBlockRawEvent(
				expectedBlockHeight,
				model.CreateBlockRawEventParams{
					BlockHash:  "8703C54C9FE1C2D6D05DAC79D795E120F385F5F43E5CDC17B73090E9DA40CEA9",
					BlockTime:  utctime.FromUnixNano(1631893335936780880),
					FromResult: types.END_BLOCK_EVENT,
					Data: model.DataParams{
						Type: "complete_unbonding",
						Content: model.BlockResultsEvent{
							Type: "complete_unbonding",
							Attributes: []model.BlockResultsEventAttribute{
								{
									Key:   "amount",
									Value: "5basetcro",
									Index: true,
								},
								{
									Key:   "validator",
									Value: "tcrocncl1sruzd529lhjju6hfcwd2fxp3v0e7p0vqqtme76",
									Index: true,
								},
								{
									Key:   "delegator",
									Value: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
									Index: true,
								},
							},
						},
					},
				},
			),
		}))
	})
})
