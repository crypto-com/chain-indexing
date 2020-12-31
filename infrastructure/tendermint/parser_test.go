package tendermint_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	infrastructure_tendermint_test "github.com/crypto-com/chain-indexing/infrastructure/tendermint/test"
)

var _ = Describe("Parser", func() {
	Describe("ParseBlockResp", func() {
		It("should return parse Evidence when there is DuplicateVoteEvidence", func() {
			blockReader := strings.NewReader(infrastructure_tendermint_test.BLOCK_WITH_DUPLICATED_VOTE_EVIDENCE)

			_, _, err := ParseBlockResp(blockReader)
			Expect(err).To(BeNil())
			//Expect(block.)
		})
	})
})
