package chainindex_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestChainindex(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Chainindex Suite")
}
