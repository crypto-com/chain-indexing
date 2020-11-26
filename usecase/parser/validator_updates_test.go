package parser_test

import (
	"github.com/crypto-com/chainindex/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/entity/command"
	command_usecase "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chainindex/usecase/parser/test"
)

var _ = Describe("ParseValidatorUpdatesCommands", func() {
	Describe("MsgSend", func() {

		It("should return commands corresponding to events in validator_updates", func() {
			blockResults := mustParseBlockResultsResp(usecase_parser_test.VALIDATOR_UPDATES_CREATE_VALIDATOR_BLOCK_RESULTS_RESP)

			cmds, err := parser.ParseValidatorUpdatesCommands(
				blockResults.Height,
				blockResults.ValidatorUpdates,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(2))
			expectedBlockHeight := int64(17)
			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewChangePower(
					expectedBlockHeight,
					model.PowerChangeParams{
						TendermintPubkey: "jZmiyA+S/yVqVuN2Px/9OqB/xgMPaj4mPdHpUOg/Kj0=",
						Power:            "312",
					},
				),
				command_usecase.NewChangePower(
					expectedBlockHeight,
					model.PowerChangeParams{
						TendermintPubkey: "LNa+qkaUeJ97z/uLAKv1YTLMspaGxSkQyipkAmtwivo=",
						Power:            "212",
					},
				),
			}))
		})

		It("should return 0 power commands when poser is not defined", func() {
			blockResults := mustParseBlockResultsResp(usecase_parser_test.VALIDATOR_UPDATES_VALIDATOR_SLASHED_BLOCK_RESULTS_RESP)

			cmds, err := parser.ParseValidatorUpdatesCommands(
				blockResults.Height,
				blockResults.ValidatorUpdates,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(2))
			expectedBlockHeight := int64(28)
			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewChangePower(
					expectedBlockHeight,
					model.PowerChangeParams{
						TendermintPubkey: "jZmiyA+S/yVqVuN2Px/9OqB/xgMPaj4mPdHpUOg/Kj0=",
						Power:            "0",
					},
				),
				command_usecase.NewChangePower(
					expectedBlockHeight,
					model.PowerChangeParams{
						TendermintPubkey: "LNa+qkaUeJ97z/uLAKv1YTLMspaGxSkQyipkAmtwivo=",
						Power:            "0",
					},
				),
			}))
		})
	})
})
