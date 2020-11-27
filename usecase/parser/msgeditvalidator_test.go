package parser_test

import (
	"strings"

	"github.com/crypto-com/chainindex/internal/primptr"

	"github.com/crypto-com/chainindex/infrastructure/tendermint"
	command_usecase "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
	"github.com/crypto-com/chainindex/usecase/parser"
	. "github.com/onsi/ginkgo"

	usecase_parser_test "github.com/crypto-com/chainindex/usecase/parser/test"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgEditValidator", func() {

		It("should parse Msg commands when there is staking.MsgEditValidator in the transaction", func() {
			txDecoder := parser.NewTxDecoder("basetrcro")
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(usecase_parser_test.TX_MSG_EDIT_VALIDATOR_BLOCK_RESP))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(usecase_parser_test.TX_MSG_EDIT_VALIDATOR_BLOCK_RESULTS_RESP))

			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(4))

			thiscmd := cmds[0]
			Expect(thiscmd.Name()).To(Equal("CreateMsgEditValidator"))
			Expect(thiscmd).To(Equal(
				command_usecase.NewCreateMsgEditValidator(
					event.MsgCommonParams{
						BlockHeight: int64(504096),
						TxHash:      "3A570A84C89578D1659E096BE8E8EB946CEB630ED123037E0F333AA352475659",
						TxSuccess:   true,
						MsgIndex:    0,
					},

					model.MsgEditValidatorParams{
						Description: model.MsgValidatorDescription{
							Moniker:         "Edited Calvin Test Node",
							Identity:        "[do-not-modify]",
							Website:         "[do-not-modify]",
							SecurityContact: "[do-not-modify]",
							Details:         "[do-not-modify]",
						},
						ValidatorAddress:       "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
						MaybeCommissionRate:    nil,
						MaybeMinSelfDelegation: primptr.String("2"),
					},
				)))
		})
	})
})
