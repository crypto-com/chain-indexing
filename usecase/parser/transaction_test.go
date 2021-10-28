package parser_test

import (
	"strings"

	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("TransactionParser", func() {
	Describe("TxHash", func() {
		It("should return transaction hash from hex encouded tx data", func() {
			txHex := "Cp4CCowBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEmwKK3Rjcm8xNjV0emNyaDJ5bDgzZzhxZXF4dWVnMmc1Z3pndTU3eTNmZTNrYzMSK3Rjcm8xODRsdGEybHN5dTQ3dnd5cDJlOHptdGNhM2s1eXE4NXA2YzR2cDMaEAoIYmFzZXRjcm8SBDEwMDAKjAEKHC9jb3Ntb3MuYmFuay52MWJldGExLk1zZ1NlbmQSbAordGNybzE4NGx0YTJsc3l1NDd2d3lwMmU4em10Y2EzazV5cTg1cDZjNHZwMxIrdGNybzE2NXR6Y3JoMnlsODNnOHFlcXh1ZWcyZzVnemd1NTd5M2ZlM2tjMxoQCghiYXNldGNybxIEMjAwMBKsAQpRCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohAgiLen9uwpvsreYibwgnQtzupil7kyNJl4oTG3Wl6oIEEgQKAggBGLdPClEKRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiEDw9KBooWSrc6BvuMJTwDq4mkyy8aC+6I5uQ9H2sn+cDYSBAoCCAEYyk8SBBDAmgwaQMtPcJacL5aryCBZz7bL4vKrOLFi07rejX0nMvBRA7BSd09ywefL+VMSkC/UwqhHC28pRTHhEDiNApbxrIYBVvIaQE0+gltCOfawUGDJU9nXJJkFLPmjMKJMYvt3UtTMjPR2bws7l78EzaUfrjtbmrkIokoxAW8GBgTuhEkC2Frr6Q0="
			Expect(parser.TxHash(txHex)).To(Equal(
				"4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416",
			))
		})
	})

	Describe("ParseTransactionCommands", func() {
		It("should parse Transaction commands when there is two Msg in one transaction", func() {
			txFeeParser := utils.NewTxDecoder()
			block, _ := mustParseBlockResp(usecase_parser_test.ONE_TX_TWO_MSG_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(usecase_parser_test.ONE_TX_TWO_MSG_BLOCK_RESULTS_RESP)
			anyAccountAddressPrefix := "tcro"

			cmds, err := parser.ParseTransactionCommands(
				txFeeParser,
				block,
				blockResults,
				anyAccountAddressPrefix,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			expectedBlockHeight := int64(343358)
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateTransaction(
				expectedBlockHeight,
				model.CreateTransactionParams{
					TxHash:   "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416",
					Code:     0,
					Log:      "[{\"msgIndex\":0,\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"send\"},{\"key\":\"sender\",\"value\":\"tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3\"},{\"key\":\"module\",\"value\":\"bank\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3\"},{\"key\":\"sender\",\"value\":\"tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3\"},{\"key\":\"amount\",\"value\":\"1000basetcro\"}]}]},{\"msgIndex\":1,\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"send\"},{\"key\":\"sender\",\"value\":\"tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3\"},{\"key\":\"module\",\"value\":\"bank\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3\"},{\"key\":\"sender\",\"value\":\"tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3\"},{\"key\":\"amount\",\"value\":\"2000basetcro\"}]}]}]",
					MsgCount: 2,
					Signers: []model.TransactionSigner{
						{
							TransactionSignerInfo: model.TransactionSignerInfo{
								Type:       "/cosmos.crypto.secp256k1.PubKey",
								IsMultiSig: false,
								Pubkeys: []string{
									"AgiLen9uwpvsreYibwgnQtzupil7kyNJl4oTG3Wl6oIE",
								},
								MaybeThreshold:  nil,
								AccountSequence: 10167,
							},
							Address: "tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3",
						},
						{
							TransactionSignerInfo: model.TransactionSignerInfo{
								Type:       "/cosmos.crypto.secp256k1.PubKey",
								IsMultiSig: false,
								Pubkeys: []string{
									"A8PSgaKFkq3Ogb7jCU8A6uJpMsvGgvuiObkPR9rJ/nA2",
								},
								MaybeThreshold:  nil,
								AccountSequence: 10186,
							},
							Address: "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3",
						},
					},
					Fee:           coin.NewEmptyCoins(),
					FeePayer:      "",
					FeeGranter:    "",
					GasWanted:     200000,
					GasUsed:       80148,
					Memo:          "",
					TimeoutHeight: 0,
				},
			)}))
		})

		It("should parse Transaction commands when there is transaction fee", func() {
			txFeeParser := utils.NewTxDecoder()
			block, _ := mustParseBlockResp(usecase_parser_test.TX_WITH_FEE_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_WITH_FEE_BLOCK_RESULTS_RESP)
			anyAccountAddressPrefix := "tcro"

			cmds, err := parser.ParseTransactionCommands(
				txFeeParser,
				block,
				blockResults,
				anyAccountAddressPrefix,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			expectedBlockHeight := int64(377673)
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateTransaction(
				expectedBlockHeight,
				model.CreateTransactionParams{
					TxHash:   "2A2A64A310B3D0E84C9831F4353E188A6E63BF451975C859DF40C54047AC6324",
					Code:     0,
					Log:      "[{\"msgIndex\":0,\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"send\"},{\"key\":\"sender\",\"value\":\"tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv\"},{\"key\":\"module\",\"value\":\"bank\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv\"},{\"key\":\"sender\",\"value\":\"tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv\"},{\"key\":\"amount\",\"value\":\"1000000000basetcro\"}]}]}]",
					MsgCount: 1,
					Signers: []model.TransactionSigner{
						{
							TransactionSignerInfo: model.TransactionSignerInfo{
								Type:       "/cosmos.crypto.secp256k1.PubKey",
								IsMultiSig: false,
								Pubkeys: []string{
									"Ax+Rgmd2ta8FxUOoFJ9Dvo3782nMWJzdYP0Jcyrk5XwO",
								},
								MaybeThreshold:  nil,
								AccountSequence: 59,
							},
							Address: "tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv",
						},
					},

					Fee:           coin.MustParseCoinsNormalized("8000000basetcro"),
					FeePayer:      "",
					FeeGranter:    "",
					GasWanted:     80000000,
					GasUsed:       62582,
					Memo:          "",
					TimeoutHeight: 0,
				},
			)}))
		})

		It("should parse Transaction commands when transaction failed with fee", func() {
			txFeeParser := utils.NewTxDecoder()
			block, _ := mustParseBlockResp(usecase_parser_test.TX_FAILED_WITH_FEE_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_FAILED_WITH_FEE_BLOCK_RESULTS_RESP)
			anyAccountAddressPrefix := "tcro"

			cmds, err := parser.ParseTransactionCommands(
				txFeeParser,
				block,
				blockResults,
				anyAccountAddressPrefix,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			expectedBlockHeight := int64(420301)
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateTransaction(
				expectedBlockHeight,
				model.CreateTransactionParams{
					TxHash:   "2A2A64A310B3D0E84C9831F4353E188A6E63BF451975C859DF40C54047AC6324",
					Code:     11,
					Log:      "out of gas in location: WriteFlat; gasWanted: 80000000, gasUsed: 80150021: out of gas",
					MsgCount: 1,
					Signers: []model.TransactionSigner{
						{
							TransactionSignerInfo: model.TransactionSignerInfo{
								Type:       "/cosmos.crypto.secp256k1.PubKey",
								IsMultiSig: false,
								Pubkeys: []string{
									"Ax+Rgmd2ta8FxUOoFJ9Dvo3782nMWJzdYP0Jcyrk5XwO",
								},
								MaybeThreshold:  nil,
								AccountSequence: 59,
							},
							Address: "tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv",
						},
					},
					Fee:           coin.MustParseCoinsNormalized("8000000basetcro"),
					FeePayer:      "",
					FeeGranter:    "",
					GasWanted:     80000000,
					GasUsed:       80150021,
					Memo:          "",
					TimeoutHeight: 0,
				},
			)}))
		})

		It("should parse Transaction commands when transaction failed without fee", func() {
			txFeeParser := utils.NewTxDecoder()
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(usecase_parser_test.TX_FAILED_WITHOUT_FEE_BLOCK_RESP))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(usecase_parser_test.TX_FAILED_WITHOUT_FEE_BLOCK_RESULTS_RESP))
			anyAccountAddressPrefix := "tcro"

			cmds, err := parser.ParseTransactionCommands(
				txFeeParser,
				block,
				blockResults,
				anyAccountAddressPrefix,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			expectedBlockHeight := int64(3245)
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateTransaction(
				expectedBlockHeight,
				model.CreateTransactionParams{
					TxHash:   "CDBA166168176BF7ECA2EAC9E9B49054F1BF4C8799B8C26CC0B9EE85CB93AF27",
					Code:     11,
					Log:      "out of gas in location: WriteFlat; gasWanted: 200000, gasUsed: 201420: out of gas",
					MsgCount: 5,
					Signers: []model.TransactionSigner{
						{
							TransactionSignerInfo: model.TransactionSignerInfo{
								Type:       "/cosmos.crypto.secp256k1.PubKey",
								IsMultiSig: false,
								Pubkeys: []string{
									"AhLYCDVbpM12Jafqp0poKEdIEpeTn03mJ5+mIgRz4PWa",
								},
								MaybeThreshold:  nil,
								AccountSequence: 5,
							},
							Address: "tcro1l38ze5fmgrgzw6rn3afx3gtpare3jgs8ke4n69",
						},
					},
					Fee:           coin.NewEmptyCoins(),
					FeePayer:      "",
					FeeGranter:    "",
					GasWanted:     200000,
					GasUsed:       201420,
					Memo:          "",
					TimeoutHeight: 0,
				},
			)}))
		})

		It("should parse Transaction commands when there is transaction memo and timeout_height", func() {
			txFeeParser := utils.NewTxDecoder()
			block, _ := mustParseBlockResp(usecase_parser_test.TX_WITH_MEMO_TIMEOUT_HEIGHT_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_WITH_MEMO_TIMEOUT_HEIGHT_BLOCK_RESULTS_RESP)
			anyAccountAddressPrefix := "tcro"

			cmds, err := parser.ParseTransactionCommands(
				txFeeParser,
				block,
				blockResults,
				anyAccountAddressPrefix,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			expectedBlockHeight := int64(492481)
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateTransaction(
				expectedBlockHeight,
				model.CreateTransactionParams{
					TxHash:   "314FB925A570DB56F69A9E58C05EB7CCBCBA444949FF14E5874D8B581322A952",
					Code:     0,
					Log:      "[{\"msgIndex\":0,\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"send\"},{\"key\":\"sender\",\"value\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\"},{\"key\":\"module\",\"value\":\"bank\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro1782gn9hzqavecukdaqqclvsnpck4mtz3vwzpxl\"},{\"key\":\"sender\",\"value\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\"},{\"key\":\"amount\",\"value\":\"100000000basetcro\"}]}]}]",
					MsgCount: 1,
					Signers: []model.TransactionSigner{
						{
							TransactionSignerInfo: model.TransactionSignerInfo{
								Type:       "/cosmos.crypto.secp256k1.PubKey",
								IsMultiSig: false,
								Pubkeys: []string{
									"A1mhVLohDEidpGYmpNYxxvikcaL72jQhZN1fxKFYkB8m",
								},
								MaybeThreshold:  nil,
								AccountSequence: 25,
							},
							Address: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						},
					},
					Fee:           coin.NewEmptyCoins(),
					FeePayer:      "",
					FeeGranter:    "",
					GasWanted:     200000,
					GasUsed:       50685,
					Memo:          "Test memo",
					TimeoutHeight: int64(500000),
				},
			)}))
		})

		It("should parse failed Transaction commands when there is transaction memo and timeout_height", func() {
			txFeeParser := utils.NewTxDecoder()
			block, _ := mustParseBlockResp(usecase_parser_test.TX_FAILED_WITH_MEMO_TIMEOUT_HEIGHT_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_FAILED_WITH_MEMO_TIMEOUT_HEIGHT_BLOCK_RESULTS_RESP)
			anyAccountAddressPrefix := "tcro"

			cmds, err := parser.ParseTransactionCommands(
				txFeeParser,
				block,
				blockResults,
				anyAccountAddressPrefix,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			expectedBlockHeight := int64(492759)
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateTransaction(
				expectedBlockHeight,
				model.CreateTransactionParams{
					TxHash:   "7CCAB9771B76F25E81C26E50265243014798172F9E8C06F8AD17442C61E592EC",
					Code:     11,
					Log:      "out of gas in location: ReadFlat; gasWanted: 50000, gasUsed: 50436: out of gas",
					MsgCount: 1,
					Signers: []model.TransactionSigner{
						{
							TransactionSignerInfo: model.TransactionSignerInfo{
								Type:       "/cosmos.crypto.secp256k1.PubKey",
								IsMultiSig: false,
								Pubkeys: []string{
									"A1mhVLohDEidpGYmpNYxxvikcaL72jQhZN1fxKFYkB8m",
								},
								MaybeThreshold:  nil,
								AccountSequence: 26,
							},
							Address: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						},
					},
					Fee:           coin.NewEmptyCoins(),
					FeePayer:      "",
					FeeGranter:    "",
					GasWanted:     50000,
					GasUsed:       50436,
					Memo:          "Test memo",
					TimeoutHeight: int64(500000),
				},
			)}))
		})

		It("should parse Transaction commands when the signer is multisig address", func() {
			txFeeParser := utils.NewTxDecoder()
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MULTISIG_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_MULTISIG_BLOCK_RESULTS_RESP)
			anyAccountAddressPrefix := "tcro"

			cmds, err := parser.ParseTransactionCommands(
				txFeeParser,
				block,
				blockResults,
				anyAccountAddressPrefix,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			expectedBlockHeight := int64(1014129)
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateTransaction(
				expectedBlockHeight,
				model.CreateTransactionParams{
					TxHash:   "82E32812C744066B2865FD5E5EADB791815127DE3746A4194B1FCEBA195BDCA6",
					Code:     0,
					Log:      "[{\"msgIndex\":0,\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"send\"},{\"key\":\"sender\",\"value\":\"tcro12ygwdvfvgt4c72e0mu7h6gmfv9ywh34r9kacjr\"},{\"key\":\"module\",\"value\":\"bank\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro12ygwdvfvgt4c72e0mu7h6gmfv9ywh34r9kacjr\"},{\"key\":\"sender\",\"value\":\"tcro12ygwdvfvgt4c72e0mu7h6gmfv9ywh34r9kacjr\"},{\"key\":\"amount\",\"value\":\"100000000basetcro\"}]}]}]",
					MsgCount: 1,
					Signers: []model.TransactionSigner{
						{
							TransactionSignerInfo: model.TransactionSignerInfo{
								Type:           "/cosmos.crypto.multisig.LegacyAminoPubKey",
								IsMultiSig:     true,
								MaybeThreshold: primptr.Int(3),
								Pubkeys: []string{
									"AyYeIUDy4m8rW6DgbRbX+k8uJn46trwyyuBE871lRsDE",
									"Ahe94UU90Bzry7/CnxzKJJ5XFJJqJ4u8cOv9rq632B/Z",
									"AgvNhfDEbHrUDP4gBpiEOmxMog+BHCEg4SB49KPUB7m+",
								},
								AccountSequence: 0,
							},
							Address: "tcro12ygwdvfvgt4c72e0mu7h6gmfv9ywh34r9kacjr",
						},
					},
					Fee:           coin.NewEmptyCoins(),
					FeePayer:      "",
					FeeGranter:    "",
					GasWanted:     200000,
					GasUsed:       78093,
					Memo:          "",
					TimeoutHeight: 0,
				},
			)}))
		})
	})
})
