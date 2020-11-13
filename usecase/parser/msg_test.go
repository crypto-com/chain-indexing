package parser_test

import (
	"strings"

	"github.com/crypto-com/chainindex/entity/command"
	"github.com/crypto-com/chainindex/infrastructure/tendermint"
	"github.com/crypto-com/chainindex/usecase/coin"
	"github.com/crypto-com/chainindex/usecase/domain/createmsgsend"
	usecase_parser_test "github.com/crypto-com/chainindex/usecase/parser/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/usecase/parser"
)

var _ = Describe("Msg", func() {
	Describe("ParseMsgCommands", func() {
		It("should parse Msg commands when there is bank.MsgSend in the transaction", func() {
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(usecase_parser_test.TX_MSG_SEND_BLOCK_RESP))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(usecase_parser_test.TX_MSG_SEND_BLOCK_RESULTS_RESP))

			cmds := parser.ParseMsgToCommands(
				block,
				blockResults,
			)
			Expect(cmds).To(HaveLen(1))
			expectedBlockHeight := int64(377673)
			expectedTxHash := "2A2A64A310B3D0E84C9831F4353E188A6E63BF451975C859DF40C54047AC6324"
			expectedMsgIndex := 0
			Expect(cmds).To(Equal([]command.Command{createmsgsend.NewCommand(
				expectedBlockHeight,
				expectedTxHash,
				expectedMsgIndex,
				createmsgsend.Params{
					FromAddress: "tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv",
					ToAddress:   "tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv",
					Amount:      coin.MustNewCoinFromString("1000000000"),
				},
			)}))
		})

		It("should parse Msg commands when there are multiple bank.MsgSend in one transaction", func() {
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(usecase_parser_test.ONE_TX_TWO_MSG_SEND_BLOCK_RESP))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(usecase_parser_test.ONE_TX_TWO_MSG_SEND_BLOCK_RESULTS_RESP))

			cmds := parser.ParseMsgToCommands(
				block,
				blockResults,
			)
			Expect(cmds).To(HaveLen(2))
			expectedBlockHeight := int64(343358)
			expectedTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			Expect(cmds).To(Equal([]command.Command{createmsgsend.NewCommand(
				expectedBlockHeight,
				expectedTxHash,
				0, // msgIndex
				createmsgsend.Params{
					FromAddress: "tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3",
					ToAddress:   "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3",
					Amount:      coin.MustNewCoinFromString("1000"),
				},
			), createmsgsend.NewCommand(
				expectedBlockHeight,
				expectedTxHash,
				1, // msgIndex
				createmsgsend.Params{
					FromAddress: "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3",
					ToAddress:   "tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3",
					Amount:      coin.MustNewCoinFromString("2000"),
				},
			)}))
		})

		// TODO: MsgMultiSend
	})
})
