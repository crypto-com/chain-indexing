package parser_test

import (
	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseEndBlockEventsCommands", func() {
	It("should return CreateRawBlockEvent commands when txs_results has events emitted", func() {
		block, _ := mustParseBlockResp(usecase_parser_test.BLOCK_RESULTS_TXS_RESULTS_RAW_EVENTS_BLOCK_RESP)
		blockResults := mustParseBlockResultsResp(usecase_parser_test.BLOCK_RESULTS_TXS_RESULTS_RAW_EVENTS_BLOCK_RESULTS_RESP)

		cmds, err := parser.ParseBlockResultsTxsResultsRawEvents(
			block,
			blockResults,
		)

		Expect(err).To(BeNil())
		Expect(cmds).To(HaveLen(8))
		expectedBlockHeight := int64(11015)
		Expect(cmds).To(Equal(
			[]command.Command{
				command_usecase.NewCreateRawBlockEvent(
					expectedBlockHeight,
					model.CreateRawBlockEventParams{
						BlockHash:  "C60E4496E326D24AABA166ACE7E86DC17AEDDC77106084B99D632030C65206EF",
						BlockTime:  utctime.FromUnixNano(1616702718680872778),
						FromResult: "TxsResults",
						Data: model.DataParams{
							Type: "tx",
							Content: model.BlockResultsEvent{
								Type: "tx",
								Attributes: []model.BlockResultsEventAttribute{
									{
										Key:   "acc_seq",
										Value: "cro1yjjlx5qsrj5rxn5xtd5rkm6dcqzlchxkrvsmg6/4118",
										Index: true,
									},
								},
							},
						},
					},
				),
				command_usecase.NewCreateRawBlockEvent(
					expectedBlockHeight,
					model.CreateRawBlockEventParams{
						BlockHash:  "C60E4496E326D24AABA166ACE7E86DC17AEDDC77106084B99D632030C65206EF",
						BlockTime:  utctime.FromUnixNano(1616702718680872778),
						FromResult: "TxsResults",
						Data: model.DataParams{
							Type: "tx",
							Content: model.BlockResultsEvent{
								Type: "tx",
								Attributes: []model.BlockResultsEventAttribute{
									{
										Key:   "signature",
										Value: "KfdBpc80DTs2LrWctzr1CX9f7vApKVfUbGKA3Y7hULkbXuKlw+BpArTgG6CCe4hZeEyr+3SX/3GUdfTbXk66yw==",
										Index: true,
									},
								},
							},
						},
					},
				),
				command_usecase.NewCreateRawBlockEvent(
					expectedBlockHeight,
					model.CreateRawBlockEventParams{
						BlockHash:  "C60E4496E326D24AABA166ACE7E86DC17AEDDC77106084B99D632030C65206EF",
						BlockTime:  utctime.FromUnixNano(1616702718680872778),
						FromResult: "TxsResults",
						Data: model.DataParams{
							Type: "transfer",
							Content: model.BlockResultsEvent{
								Type: "transfer",
								Attributes: []model.BlockResultsEventAttribute{
									{
										Key:   "recipient",
										Value: "cro17xpfvakm2amg962yls6f84z3kell8c5lgztehv",
										Index: true,
									},
									{
										Key:   "sender",
										Value: "cro1yjjlx5qsrj5rxn5xtd5rkm6dcqzlchxkrvsmg6",
										Index: true,
									},
									{
										Key:   "amount",
										Value: "10000basecro",
										Index: true,
									},
								},
							},
						},
					},
				),
				command_usecase.NewCreateRawBlockEvent(
					expectedBlockHeight,
					model.CreateRawBlockEventParams{
						BlockHash:  "C60E4496E326D24AABA166ACE7E86DC17AEDDC77106084B99D632030C65206EF",
						BlockTime:  utctime.FromUnixNano(1616702718680872778),
						FromResult: "TxsResults",
						Data: model.DataParams{
							Type: "message",
							Content: model.BlockResultsEvent{
								Type: "message",
								Attributes: []model.BlockResultsEventAttribute{
									{
										Key:   "sender",
										Value: "cro1yjjlx5qsrj5rxn5xtd5rkm6dcqzlchxkrvsmg6",
										Index: true,
									},
								},
							},
						},
					},
				),
				command_usecase.NewCreateRawBlockEvent(
					expectedBlockHeight,
					model.CreateRawBlockEventParams{
						BlockHash:  "C60E4496E326D24AABA166ACE7E86DC17AEDDC77106084B99D632030C65206EF",
						BlockTime:  utctime.FromUnixNano(1616702718680872778),
						FromResult: "TxsResults",
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
				command_usecase.NewCreateRawBlockEvent(
					expectedBlockHeight,
					model.CreateRawBlockEventParams{
						BlockHash:  "C60E4496E326D24AABA166ACE7E86DC17AEDDC77106084B99D632030C65206EF",
						BlockTime:  utctime.FromUnixNano(1616702718680872778),
						FromResult: "TxsResults",
						Data: model.DataParams{
							Type: "transfer",
							Content: model.BlockResultsEvent{
								Type: "transfer",
								Attributes: []model.BlockResultsEventAttribute{
									{
										Key:   "recipient",
										Value: "cro1valk6lj0k5lhw3qpk0qcw3gusqynlserl9rlme",
										Index: true,
									},
									{
										Key:   "sender",
										Value: "cro1yjjlx5qsrj5rxn5xtd5rkm6dcqzlchxkrvsmg6",
										Index: true,
									},
									{
										Key:   "amount",
										Value: "67786900000basecro",
										Index: true,
									},
								},
							},
						},
					},
				),
				command_usecase.NewCreateRawBlockEvent(
					expectedBlockHeight,
					model.CreateRawBlockEventParams{
						BlockHash:  "C60E4496E326D24AABA166ACE7E86DC17AEDDC77106084B99D632030C65206EF",
						BlockTime:  utctime.FromUnixNano(1616702718680872778),
						FromResult: "TxsResults",
						Data: model.DataParams{
							Type: "message",
							Content: model.BlockResultsEvent{
								Type: "message",
								Attributes: []model.BlockResultsEventAttribute{
									{
										Key:   "sender",
										Value: "cro1yjjlx5qsrj5rxn5xtd5rkm6dcqzlchxkrvsmg6",
										Index: true,
									},
								},
							},
						},
					},
				),
				command_usecase.NewCreateRawBlockEvent(
					expectedBlockHeight,
					model.CreateRawBlockEventParams{
						BlockHash:  "C60E4496E326D24AABA166ACE7E86DC17AEDDC77106084B99D632030C65206EF",
						BlockTime:  utctime.FromUnixNano(1616702718680872778),
						FromResult: "TxsResults",
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
			},
		))
	})
})
