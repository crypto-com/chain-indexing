package parser_test

import (
	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseEndBlockEventsCommands", func() {
	It("should return GravityEthereumSendToCosmosHandled commands when end_block_events has ethereum_send_to_cosmos_handled event", func() {
		blockResults := mustParseBlockResultsResp(usecase_parser_test.END_BLOCK_ETHEREUM_SEND_TO_COSMOS_HANDLED_BLOCK_RESULTS_RESP)

		// TODO
		cmds, err := parser.ParseEndBlockEventsCommands(
			blockResults.Height,
			blockResults.EndBlockEvents,
		)
		Expect(err).To(BeNil())
		Expect(cmds).To(HaveLen(3))
		expectedBlockHeight := int64(630)
		Expect(cmds[1]).To(Equal(
			command_usecase.NewGravityHandleEthereumSendToCosmos(
				expectedBlockHeight,
				model.GravityEthereumSendToCosmosHandledEventParams{
					Module:   "gravity",
					Sender:   "0x5E44D43F4Aa0B3ED004eaaD4eF21a83DFF2ef6E5",
					Receiver: "tcrc13yux6z8mh6w5t3v4uq7clewnh35znrgdgye0k2",
					Amount: coin.MustNewCoins(
						coin.MustNewCoin("gravity0x564A1c3AF089D02D0B6C311C650eA3768424cbfa", coin.NewInt(250)),
					),
					BridgeChainId:         42,
					EthereumTokenContract: "0x564A1c3AF089D02D0B6C311C650eA3768424cbfa",
					Nonce:                 6,
					EthereumEventVoteRecordId: []byte{
						5, 0, 0, 0, 0, 0, 0, 0, 6, 192, 153, 206, 71, 79, 238, 213, 132, 157, 104, 158, 56, 54, 255, 37, 161, 144, 49, 44, 59, 103, 183, 62, 196, 85, 136, 16, 73, 135, 49, 242, 167,
					},
				},
			),
		))
	})
})
