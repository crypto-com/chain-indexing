package parser_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseCreateEvidenceCommands", func() {

	It("should return commands corresponding to evidences in block", func() {
		block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
			usecase_parser_test.EVIDENCE_BLOCK_RESP,
		))

		cmds, err := parser.ParseEvidencesCommands(
			block.Height,
			block.Evidences,
		)
		Expect(err).To(BeNil())
		Expect(cmds).To(HaveLen(1))
		cmd := cmds[0]
		Expect(cmd.Name()).To(Equal("CreateEvidence"))

		untypedEvent, _ := cmd.Exec()
		typedEvent := untypedEvent.(*event.EvidenceSubmitted)

		expectedBlockHeight := int64(58346)
		expectedInfractionHeight := int64(58344)
		Expect(typedEvent.BlockHeight).To(Equal(expectedBlockHeight))
		Expect(typedEvent.InfractionHeight).To(Equal(expectedInfractionHeight))
		Expect(typedEvent.TendermintAddress).To(Equal("ABB5C6BEC847C7343FE1D80197E95AED7A1F91F5"))
	})

	It("should return 0 evidence commands when evidences is empty block", func() {
		block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
			usecase_parser_test.EVIDENCE_EMPTY_BLOCK_RESP,
		))

		cmds, err := parser.ParseEvidencesCommands(
			block.Height,
			block.Evidences,
		)
		Expect(err).To(BeNil())
		Expect(cmds).To(HaveLen(0))
	})
})
