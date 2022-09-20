package parser_test

import (
	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/external/utctime"
	types "github.com/crypto-com/chain-indexing/projection/block_raw_event/constants"
	"github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseTxsResultsBlockEventsCommands", func() {
	It("should return CreateSend commands when txs_results has simple send transaction", func() {
		block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_SEND_BLOCK_RESP)
		blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_MSG_SEND_BLOCK_RESULTS_RESP)

		cmds, err := parser.ParseBlockResultsTxsResults(
			block,
			blockResults,
		)

		Expect(err).To(BeNil())
		Expect(cmds).To(HaveLen(6))
		expectedBlockHeight := int64(377673)
		Expect(cmds).To(Equal([]command.Command{
			command_usecase.NewCreateBlockRawEvent(
				expectedBlockHeight,
				model.CreateBlockRawEventParams{
					BlockHash:  "BBC28EC0167AC0D8CCD5D7D0ECE6F2A6485751F0E68B9BF80E0FC112C64C0AF8",
					BlockTime:  utctime.FromUnixNano(1605173821926253966),
					FromResult: types.TXS_RESULTS,
					Data: model.DataParams{
						Type: "transfer",
						Content: model.BlockResultsEvent{
							Type: "transfer",
							Attributes: []model.BlockResultsEventAttribute{
								{
									Key:   "recipient",
									Value: "tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha",
									Index: true,
								},
								{
									Key:   "sender",
									Value: "tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv",
									Index: true,
								},
								{
									Key:   "amount",
									Value: "8000000basetcro",
									Index: true,
								},
							},
						},
					},
				},
			),
			command_usecase.NewCreateBlockRawEvent(
				expectedBlockHeight,
				model.CreateBlockRawEventParams{
					BlockHash:  "BBC28EC0167AC0D8CCD5D7D0ECE6F2A6485751F0E68B9BF80E0FC112C64C0AF8",
					BlockTime:  utctime.FromUnixNano(1605173821926253966),
					FromResult: types.TXS_RESULTS,
					Data: model.DataParams{
						Type: "message",
						Content: model.BlockResultsEvent{
							Type: "message",
							Attributes: []model.BlockResultsEventAttribute{
								{
									Key:   "sender",
									Value: "tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv",
									Index: true,
								},
							},
						},
					},
				},
			),
			command_usecase.NewCreateBlockRawEvent(
				expectedBlockHeight,
				model.CreateBlockRawEventParams{
					BlockHash:  "BBC28EC0167AC0D8CCD5D7D0ECE6F2A6485751F0E68B9BF80E0FC112C64C0AF8",
					BlockTime:  utctime.FromUnixNano(1605173821926253966),
					FromResult: types.TXS_RESULTS,
					Data: model.DataParams{
						Type: "message",
						Content: model.BlockResultsEvent{
							Type: "message",
							Attributes: []model.BlockResultsEventAttribute{
								{
									Key:   "action",
									Value: "send",
									Index: true,
								},
							},
						},
					},
				},
			),
			command_usecase.NewCreateBlockRawEvent(
				expectedBlockHeight,
				model.CreateBlockRawEventParams{
					BlockHash:  "BBC28EC0167AC0D8CCD5D7D0ECE6F2A6485751F0E68B9BF80E0FC112C64C0AF8",
					BlockTime:  utctime.FromUnixNano(1605173821926253966),
					FromResult: types.TXS_RESULTS,
					Data: model.DataParams{
						Type: "transfer",
						Content: model.BlockResultsEvent{
							Type: "transfer",
							Attributes: []model.BlockResultsEventAttribute{
								{
									Key:   "recipient",
									Value: "tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv",
									Index: true,
								},
								{
									Key:   "sender",
									Value: "tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv",
									Index: true,
								},
								{
									Key:   "amount",
									Value: "1000000000basetcro",
									Index: true,
								},
							},
						},
					},
				},
			),
			command_usecase.NewCreateBlockRawEvent(
				expectedBlockHeight,
				model.CreateBlockRawEventParams{
					BlockHash:  "BBC28EC0167AC0D8CCD5D7D0ECE6F2A6485751F0E68B9BF80E0FC112C64C0AF8",
					BlockTime:  utctime.FromUnixNano(1605173821926253966),
					FromResult: types.TXS_RESULTS,
					Data: model.DataParams{
						Type: "message",
						Content: model.BlockResultsEvent{
							Type: "message",
							Attributes: []model.BlockResultsEventAttribute{
								{
									Key:   "sender",
									Value: "tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv",
									Index: true,
								},
							},
						},
					},
				},
			),
			command_usecase.NewCreateBlockRawEvent(
				expectedBlockHeight,
				model.CreateBlockRawEventParams{
					BlockHash:  "BBC28EC0167AC0D8CCD5D7D0ECE6F2A6485751F0E68B9BF80E0FC112C64C0AF8",
					BlockTime:  utctime.FromUnixNano(1605173821926253966),
					FromResult: types.TXS_RESULTS,
					Data: model.DataParams{
						Type: "message",
						Content: model.BlockResultsEvent{
							Type: "message",
							Attributes: []model.BlockResultsEventAttribute{
								{
									Key:   "module",
									Value: "bank",
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
