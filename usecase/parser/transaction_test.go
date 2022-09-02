package parser_test

import (
	"strings"

	"github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/entity/command"
	FakeLogger "github.com/crypto-com/chain-indexing/external/logger/test"
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
			mockClient := cosmosapp.NewMockClient()
			fakeLogger := FakeLogger.NewFakeLogger()
			blockResults := mustParseBlockResultsResp(usecase_parser_test.ONE_TX_TWO_MSG_BLOCK_RESULTS_RESP)

			tx := mustParseTxsResp(usecase_parser_test.ONE_TX_TWO_MSG_TXS_RESP)
			txs := []model.Tx{*tx}

			anyAccountAddressPrefix := "tcro"

			cmds, err := parser.ParseTransactionCommands(
				fakeLogger,
				txs,
				mockClient,
				blockResults,
				anyAccountAddressPrefix,
				[]string{},
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
							MaybeKeyInfo: &model.TransactionSignerKeyInfo{
								Type:       "/cosmos.crypto.secp256k1.PubKey",
								IsMultiSig: false,
								Pubkeys: []string{
									"AgiLen9uwpvsreYibwgnQtzupil7kyNJl4oTG3Wl6oIE",
								},
								MaybeThreshold: nil,
							},
							Address:         "tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3",
							AccountSequence: 10167,
						},
						{
							MaybeKeyInfo: &model.TransactionSignerKeyInfo{
								Type:       "/cosmos.crypto.secp256k1.PubKey",
								IsMultiSig: false,
								Pubkeys: []string{
									"A8PSgaKFkq3Ogb7jCU8A6uJpMsvGgvuiObkPR9rJ/nA2",
								},
								MaybeThreshold: nil,
							},
							Address:         "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3",
							AccountSequence: 10186,
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
			mockClient := cosmosapp.NewMockClient()
			fakeLogger := FakeLogger.NewFakeLogger()
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_WITH_FEE_BLOCK_RESULTS_RESP)

			tx := mustParseTxsResp(usecase_parser_test.TX_WITH_FEE_TXS_RESP)
			txs := []model.Tx{*tx}

			anyAccountAddressPrefix := "tcro"

			cmds, err := parser.ParseTransactionCommands(
				fakeLogger,
				txs,
				mockClient,
				blockResults,
				anyAccountAddressPrefix,
				[]string{},
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
							MaybeKeyInfo: &model.TransactionSignerKeyInfo{
								Type:       "/cosmos.crypto.secp256k1.PubKey",
								IsMultiSig: false,
								Pubkeys: []string{
									"Ax+Rgmd2ta8FxUOoFJ9Dvo3782nMWJzdYP0Jcyrk5XwO",
								},
								MaybeThreshold: nil,
							},
							Address:         "tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv",
							AccountSequence: 59,
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
			mockClient := cosmosapp.NewMockClient()
			fakeLogger := FakeLogger.NewFakeLogger()
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_FAILED_WITH_FEE_BLOCK_RESULTS_RESP)

			tx := mustParseTxsResp(usecase_parser_test.TX_FAILED_WITH_FEE_TXS_RESP)
			txs := []model.Tx{*tx}

			anyAccountAddressPrefix := "tcro"

			cmds, err := parser.ParseTransactionCommands(
				fakeLogger,
				txs,
				mockClient,
				blockResults,
				anyAccountAddressPrefix,
				[]string{},
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
							MaybeKeyInfo: &model.TransactionSignerKeyInfo{
								Type:       "/cosmos.crypto.secp256k1.PubKey",
								IsMultiSig: false,
								Pubkeys: []string{
									"Ax+Rgmd2ta8FxUOoFJ9Dvo3782nMWJzdYP0Jcyrk5XwO",
								},
								MaybeThreshold: nil,
							},
							Address:         "tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv",
							AccountSequence: 59,
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
			mockClient := cosmosapp.NewMockClient()
			fakeLogger := FakeLogger.NewFakeLogger()
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(usecase_parser_test.TX_FAILED_WITHOUT_FEE_BLOCK_RESULTS_RESP))

			tx := mustParseTxsResp(usecase_parser_test.TX_FAILED_WITHOUT_FEE_TXS_RESP)
			txs := []model.Tx{*tx}

			anyAccountAddressPrefix := "tcro"

			cmds, err := parser.ParseTransactionCommands(
				fakeLogger,
				txs,
				mockClient,
				blockResults,
				anyAccountAddressPrefix,
				[]string{},
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
							MaybeKeyInfo: &model.TransactionSignerKeyInfo{
								Type:       "/cosmos.crypto.secp256k1.PubKey",
								IsMultiSig: false,
								Pubkeys: []string{
									"AhLYCDVbpM12Jafqp0poKEdIEpeTn03mJ5+mIgRz4PWa",
								},
								MaybeThreshold: nil,
							},
							Address:         "tcro1l38ze5fmgrgzw6rn3afx3gtpare3jgs8ke4n69",
							AccountSequence: 5,
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
			mockClient := cosmosapp.NewMockClient()
			fakeLogger := FakeLogger.NewFakeLogger()
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_WITH_MEMO_TIMEOUT_HEIGHT_BLOCK_RESULTS_RESP)

			tx := mustParseTxsResp(usecase_parser_test.TX_WITH_MEMO_TIMEOUT_HEIGHT_TXS_RESP)
			txs := []model.Tx{*tx}

			anyAccountAddressPrefix := "tcro"

			cmds, err := parser.ParseTransactionCommands(
				fakeLogger,
				txs,
				mockClient,
				blockResults,
				anyAccountAddressPrefix,
				[]string{},
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
							MaybeKeyInfo: &model.TransactionSignerKeyInfo{
								Type:       "/cosmos.crypto.secp256k1.PubKey",
								IsMultiSig: false,
								Pubkeys: []string{
									"A1mhVLohDEidpGYmpNYxxvikcaL72jQhZN1fxKFYkB8m",
								},
								MaybeThreshold: nil,
							},
							Address:         "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
							AccountSequence: 25,
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
			mockClient := cosmosapp.NewMockClient()
			fakeLogger := FakeLogger.NewFakeLogger()
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_FAILED_WITH_MEMO_TIMEOUT_HEIGHT_BLOCK_RESULTS_RESP)

			tx := mustParseTxsResp(usecase_parser_test.TX_FAILED_WITH_MEMO_TIMEOUT_HEIGHT_TXS_RESP)
			txs := []model.Tx{*tx}

			anyAccountAddressPrefix := "tcro"

			cmds, err := parser.ParseTransactionCommands(
				fakeLogger,
				txs,
				mockClient,
				blockResults,
				anyAccountAddressPrefix,
				[]string{},
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
							MaybeKeyInfo: &model.TransactionSignerKeyInfo{
								Type:       "/cosmos.crypto.secp256k1.PubKey",
								IsMultiSig: false,
								Pubkeys: []string{
									"A1mhVLohDEidpGYmpNYxxvikcaL72jQhZN1fxKFYkB8m",
								},
								MaybeThreshold: nil,
							},
							Address:         "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
							AccountSequence: 26,
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

		It("should parse Transaction commands when the signer public key is empty", func() {
			mockClient := cosmosapp.NewMockClient()
			fakeLogger := FakeLogger.NewFakeLogger()
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_SIGNER_EMPTY_PUBKEY_BLOCK_RESULTS_RESP)

			tx1 := mustParseTxsResp(usecase_parser_test.TX_SIGNER_EMPTY_PUBKEY_TXS_RESP_1)
			tx2 := mustParseTxsResp(usecase_parser_test.TX_SIGNER_EMPTY_PUBKEY_TXS_RESP_2)
			tx3 := mustParseTxsResp(usecase_parser_test.TX_SIGNER_EMPTY_PUBKEY_TXS_RESP_3)
			txs := []model.Tx{*tx1, *tx2, *tx3}

			anyAccountAddressPrefix := "cosmos"

			cmds, err := parser.ParseTransactionCommands(
				fakeLogger,
				txs,
				mockClient,
				blockResults,
				anyAccountAddressPrefix,
				[]string{},
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(3))
			expectedBlockHeight := int64(9399574)
			Expect(cmds[2]).To(Equal(command_usecase.NewCreateTransaction(
				expectedBlockHeight,
				model.CreateTransactionParams{
					TxHash:   "AC2BD4EF48A13B81641AC1CD3A60894718A66FC597A1353B8ACC4E5B6311DB1D",
					Index:    2,
					Code:     0,
					Log:      "[{\"msgIndex\":0,\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"cosmos1k4r9xelttshp5jq7c9a2khwpyj0r69j7lzhu77\"},{\"key\":\"amount\",\"value\":\"900uatom\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"cosmos10v888l3d44rh65mjku0m96dgkegnwcxwpm3har\"},{\"key\":\"amount\",\"value\":\"900uatom\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.bank.v1beta1.MsgSend\"},{\"key\":\"sender\",\"value\":\"cosmos10v888l3d44rh65mjku0m96dgkegnwcxwpm3har\"},{\"key\":\"module\",\"value\":\"bank\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"cosmos1k4r9xelttshp5jq7c9a2khwpyj0r69j7lzhu77\"},{\"key\":\"sender\",\"value\":\"cosmos10v888l3d44rh65mjku0m96dgkegnwcxwpm3har\"},{\"key\":\"amount\",\"value\":\"900uatom\"}]}]}]",
					MsgCount: 1,
					Signers: []model.TransactionSigner{
						{
							MaybeKeyInfo:    nil,
							Address:         "",
							AccountSequence: 35,
						},
					},
					Fee:           coin.MustParseCoinsNormalized("500uatom"),
					FeePayer:      "",
					FeeGranter:    "",
					GasWanted:     200000,
					GasUsed:       66595,
					Memo:          "",
					TimeoutHeight: 0,
				},
			)))
		})
		It("should parse Transaction commands when the signer public key is in state", func() {
			mockClient := cosmosapp.MockClient{}
			fakeLogger := FakeLogger.NewFakeLogger()
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_SIGNER_PUBKEY_IN_STATE_BLOCK_RESULTS_RESP)

			tx := mustParseTxsResp(usecase_parser_test.TX_SIGNER_PUBKEY_IN_STATE_TXS_RESP)
			txs := []model.Tx{*tx}

			anyAccountAddressPrefix := "basecro"

			mockClient.On("Account", "cro1rrnm2rkhrkgelj7j9pxeu7sggv8fde3tkymym7").Return(
				&cosmosapp.Account{
					Type:    "AccountType",
					Address: "cro1rrnm2rkhrkgelj7j9pxeu7sggv8fde3tkymym7",
					MaybePubkey: &cosmosapp.PubKey{
						Type: "/cosmos.crypto.secp256k1.PubKey",
						Key:  "AomWNLM+dBB76InhfghTzlUDOPevNG2AClk286KuSODS",
					},
					AccountNumber: "AccountNumber",
					Sequence:      "1",
					MaybeModuleAccount: &cosmosapp.ModuleAccount{
						Name:        "",
						Permissions: []string{},
					},
					MaybeDelayedVestingAccount: &cosmosapp.DelayedVestingAccount{
						OriginalVesting:  []cosmosapp.VestingBalance{},
						DelegatedFree:    []cosmosapp.VestingBalance{},
						DelegatedVesting: []cosmosapp.VestingBalance{},
						EndTime:          "",
					},
					MaybeContinuousVestingAccount: &cosmosapp.ContinuousVestingAccount{
						OriginalVesting:  []cosmosapp.VestingBalance{},
						DelegatedFree:    []cosmosapp.VestingBalance{},
						DelegatedVesting: []cosmosapp.VestingBalance{},
						StartTime:        "",
						EndTime:          "",
					},
					MaybePeriodicVestingAccount: &cosmosapp.PeriodicVestingAccount{
						OriginalVesting:  []cosmosapp.VestingBalance{},
						DelegatedFree:    []cosmosapp.VestingBalance{},
						DelegatedVesting: []cosmosapp.VestingBalance{},
						StartTime:        "",
						EndTime:          "",
						VestingPeriods:   []cosmosapp.VestingPeriod{},
					},
				},
				nil,
			)
			mockClient.On("Account", "cro1442fdq2t62vqchraj6ujxnhq3gkzq3ra9nt4lc").Return(
				&cosmosapp.Account{
					Type:    "AccountType",
					Address: "cro1442fdq2t62vqchraj6ujxnhq3gkzq3ra9nt4lc",
					MaybePubkey: &cosmosapp.PubKey{
						Type: "/cosmos.crypto.secp256k1.PubKey",
						Key:  "ApTiQlAs/mTr9a1RQwmm5G+bXe2MvGRypncY9pAHcWKO",
					},
					AccountNumber: "AccountNumber",
					Sequence:      "1",
					MaybeModuleAccount: &cosmosapp.ModuleAccount{
						Name:        "",
						Permissions: []string{},
					},
					MaybeDelayedVestingAccount: &cosmosapp.DelayedVestingAccount{
						OriginalVesting:  []cosmosapp.VestingBalance{},
						DelegatedFree:    []cosmosapp.VestingBalance{},
						DelegatedVesting: []cosmosapp.VestingBalance{},
						EndTime:          "",
					},
					MaybeContinuousVestingAccount: &cosmosapp.ContinuousVestingAccount{
						OriginalVesting:  []cosmosapp.VestingBalance{},
						DelegatedFree:    []cosmosapp.VestingBalance{},
						DelegatedVesting: []cosmosapp.VestingBalance{},
						StartTime:        "",
						EndTime:          "",
					},
					MaybePeriodicVestingAccount: &cosmosapp.PeriodicVestingAccount{
						OriginalVesting:  []cosmosapp.VestingBalance{},
						DelegatedFree:    []cosmosapp.VestingBalance{},
						DelegatedVesting: []cosmosapp.VestingBalance{},
						StartTime:        "",
						EndTime:          "",
						VestingPeriods:   []cosmosapp.VestingPeriod{},
					},
				},
				nil,
			)

			cmds, err := parser.ParseTransactionCommands(
				fakeLogger,
				txs,
				&mockClient,
				blockResults,
				anyAccountAddressPrefix,
				[]string{"cro1rrnm2rkhrkgelj7j9pxeu7sggv8fde3tkymym7", "cro1442fdq2t62vqchraj6ujxnhq3gkzq3ra9nt4lc"},
			)

			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			expectedBlockHeight := int64(324)
			Expect(cmds[0]).To(Equal(command_usecase.NewCreateTransaction(
				expectedBlockHeight,
				model.CreateTransactionParams{
					TxHash:   "C98FF5B5B95DC21F6D614D914FDA0A546DAE65E826ABFFE5C89BCF56A6F4112C",
					Index:    0,
					Code:     0,
					Log:      "[{\"msgIndex\":0,\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"cro1mzh0ps49m7ur7y8fwhy5xtt06mq7fnk3xuqset\"},{\"key\":\"amount\",\"value\":\"2basecro\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"cro1rrnm2rkhrkgelj7j9pxeu7sggv8fde3tkymym7\"},{\"key\":\"amount\",\"value\":\"1basecro\"},{\"key\":\"spender\",\"value\":\"cro1442fdq2t62vqchraj6ujxnhq3gkzq3ra9nt4lc\"},{\"key\":\"amount\",\"value\":\"1basecro\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.bank.v1beta1.MsgMultiSend\"},{\"key\":\"sender\",\"value\":\"cro1rrnm2rkhrkgelj7j9pxeu7sggv8fde3tkymym7\"},{\"key\":\"sender\",\"value\":\"cro1442fdq2t62vqchraj6ujxnhq3gkzq3ra9nt4lc\"},{\"key\":\"module\",\"value\":\"bank\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"cro1mzh0ps49m7ur7y8fwhy5xtt06mq7fnk3xuqset\"},{\"key\":\"amount\",\"value\":\"2basecro\"}]}]}]",
					MsgCount: 1,
					Signers: []model.TransactionSigner{
						{
							MaybeKeyInfo: &model.TransactionSignerKeyInfo{
								Type:       "/cosmos.crypto.secp256k1.PubKey",
								IsMultiSig: false,
								Pubkeys: []string{
									"AomWNLM+dBB76InhfghTzlUDOPevNG2AClk286KuSODS",
								},
								MaybeThreshold: nil,
							},
							Address:         "cro1rrnm2rkhrkgelj7j9pxeu7sggv8fde3tkymym7",
							AccountSequence: 1,
						},
						{
							MaybeKeyInfo: &model.TransactionSignerKeyInfo{
								Type:       "/cosmos.crypto.secp256k1.PubKey",
								IsMultiSig: false,
								Pubkeys: []string{
									"ApTiQlAs/mTr9a1RQwmm5G+bXe2MvGRypncY9pAHcWKO",
								},
								MaybeThreshold: nil,
							},
							Address:         "cro1442fdq2t62vqchraj6ujxnhq3gkzq3ra9nt4lc",
							AccountSequence: 1,
						},
					},
					Fee:           coin.MustParseCoinsNormalized("0basecro"),
					FeePayer:      "",
					FeeGranter:    "",
					GasWanted:     200000,
					GasUsed:       77224,
					Memo:          "",
					TimeoutHeight: 0,
				},
			)))
		})
	})
})
