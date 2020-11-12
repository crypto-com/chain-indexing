package parser_test

import (
	"strings"

	"github.com/crypto-com/chainindex/usecase/domain/createtransaction"

	"github.com/crypto-com/chainindex/entity/command"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/infrastructure/tendermint"
	"github.com/crypto-com/chainindex/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chainindex/usecase/parser/test"
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
		It("should parse Transaction commands correctly", func() {
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(usecase_parser_test.ONE_TX_TWO_MSG_BLOCK_RESP))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(usecase_parser_test.ONE_TX_TWO_MSG_BLOCK_RESULTS_RESP))

			cmds, err := parser.ParseTransactionCommands(block, blockResults)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			expectedBlockHeight := int64(343358)
			Expect(cmds).To(Equal([]command.Command{createtransaction.NewCommand(
				expectedBlockHeight,
				createtransaction.Params{
					TxHash:    "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416",
					Code:      0,
					Log:       "[{\"msgIndex\":0,\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"send\"},{\"key\":\"sender\",\"value\":\"tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3\"},{\"key\":\"module\",\"value\":\"bank\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3\"},{\"key\":\"sender\",\"value\":\"tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3\"},{\"key\":\"amount\",\"value\":\"1000basetcro\"}]}]},{\"msgIndex\":1,\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"send\"},{\"key\":\"sender\",\"value\":\"tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3\"},{\"key\":\"module\",\"value\":\"bank\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3\"},{\"key\":\"sender\",\"value\":\"tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3\"},{\"key\":\"amount\",\"value\":\"2000basetcro\"}]}]}]",
					MsgCount:  2,
					GasWanted: "200000",
					GasUsed:   "80148",
				},
			)}))
		})
	})
})
