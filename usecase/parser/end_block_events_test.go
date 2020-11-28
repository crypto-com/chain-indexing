package parser_test

import (
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

		cmds, err := parser.ParseEndBlockEventsCommands(
			blockResults.Height,
			blockResults.EndBlockEvents,
		)
		Expect(err).To(BeNil())
		Expect(cmds).To(HaveLen(1))
		expectedBlockHeight := int64(21575)
		expectedProposalId := "1"
		expectedResult := "proposal_rejected"
		Expect(cmds).To(Equal([]command.Command{
			command_usecase.NewEndProposal(
				expectedBlockHeight,
				expectedProposalId,
				expectedResult,
			),
		}))
	})

	It("should return InactiveProposal commands when end_blocks_events has proposal_inactive event", func() {
		blockResults := mustParseBlockResultsResp(usecase_parser_test.END_BLOCK_PROPOSAL_INACTIVED_BLOCK_RESULTS_RESP)

		cmds, err := parser.ParseEndBlockEventsCommands(
			blockResults.Height,
			blockResults.EndBlockEvents,
		)
		Expect(err).To(BeNil())
		Expect(cmds).To(HaveLen(1))
		expectedBlockHeight := int64(21541)
		expectedProposalId := "2"
		expectedResult := "proposal_dropped"
		Expect(cmds).To(Equal([]command.Command{
			command_usecase.NewInactiveProposal(
				expectedBlockHeight,
				expectedProposalId,
				expectedResult,
			),
		}))
	})

	It("should return CompleteBonding commands when end_blocks_events has complete_unbonding event", func() {
		blockResults := mustParseBlockResultsResp(usecase_parser_test.END_BLOCK_COMPLETE_UNBONDING_BLOCK_RESULTS_RESP)

		cmds, err := parser.ParseEndBlockEventsCommands(
			blockResults.Height,
			blockResults.EndBlockEvents,
		)
		Expect(err).To(BeNil())
		Expect(cmds).To(HaveLen(1))
		expectedBlockHeight := int64(477415)
		Expect(cmds).To(Equal([]command.Command{
			command_usecase.NewCreateCompleteBonding(
				expectedBlockHeight,
				model.CompleteBondingParams{
					Delegator: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
					Validator: "tcrocncl1sruzd529lhjju6hfcwd2fxp3v0e7p0vqqtme76",
					Amount:    coin.MustNewCoinFromString("5"),
				},
			),
		}))
	})
})
