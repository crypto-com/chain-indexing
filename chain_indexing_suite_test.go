package chain_indexing_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestChainIndexing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ChainIndexing Suite")
}
