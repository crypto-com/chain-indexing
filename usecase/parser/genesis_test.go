package parser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/entity/command"
	command_usecase "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chainindex/usecase/parser/test"
)

var _ = Describe("Parse Genesis", func() {
	It("should return genesis command corresponding to genesis response", func() {
		genesis := mustParseGenesisResp(usecase_parser_test.GENESIS_RESP)

		cmds, err := parser.ParseGenesisCommands(genesis)
		Expect(err).To(BeNil())
		Expect(cmds).To(Equal([]command.Command{
			command_usecase.NewCreateGenesis(*genesis),
		}))
	})
})
