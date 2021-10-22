package parser_test

import (
	"github.com/crypto-com/chain-indexing/internal/json"
	"github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseEndBlockEventsCommands", func() {
	It("should return CreateCronosSendToIbc commands when end_block_events has cronos send to IBC transaction", func() {
		block, _ := mustParseBlockResp(usecase_parser_test.BLOCK_RESULTS_TXS_RESULTS_CREATE_SEND_TO_IBC_BLOCK_RESP)
		blockResults := mustParseBlockResultsResp(usecase_parser_test.BLOCK_RESULTS_TXS_RESULTS_CREATE_SEND_TO_IBC_BLOCK_RESULTS_RESP)

		// TODO
		cmds, err := parser.ParseBlockResultsTxsResults(
			block,
			blockResults,
		)
		Expect(err).To(BeNil())
		Expect(cmds).To(HaveLen(1))
		expectedBlockHeight := int64(50813)
		Expect(cmds[0]).To(Equal(
			command_usecase.NewCreateCronosSendToIBC(
				expectedBlockHeight,
				model.CronosSendToIBCParams{
					TxHash:         "E0DFC9314BAFC8A95991D49764CAFCDDE233FEB3892E2A053C7AA6BB1708241E",
					EthereumTxHash: "0x33c03179001d0c780416fbdbfcbab2e7778dc8692ab5202699225c9e2c592242",
					SourcePort:     "transfer",
					SourceChannel:  "channel-0",
					Token: model.CronosSendToIBCToken{
						Amount: json.NewUint64(1000000000),
						Denom:  "transfer/channel-0/basetcro",
					},
					Sender:   "tcrc13yux6z8mh6w5t3v4uq7clewnh35znrgdgye0k2",
					Receiver: "tcro1tzhdkuc328cgh2hycyfddtdpqfwwu42ywyfvkj",
					TimeoutHeight: model.CronosSendToIBCHeight{
						RevisionNumber: uint64(0),
						RevisionHeight: uint64(0),
					},
					TimeoutTimestamp:   "1633722759099392805",
					PacketDataHex:      "7b22616d6f756e74223a2231303030303030303030222c2264656e6f6d223a227472616e736665722f6368616e6e656c2d302f626173657463726f222c227265636569766572223a227463726f31747a68646b756333323863676832687963796664647464707166777775343279777966766b6a222c2273656e646572223a22746372633133797578367a386d6836773574337634757137636c65776e6833357a6e726764677965306b32227d",
					PacketSequence:     28,
					DestinationPort:    "transfer",
					DestinationChannel: "channel-129",
					ChannelOrdering:    "ORDER_UNORDERED",
					ConnectionID:       "connection-0",
				},
			),
		))
	})
})
