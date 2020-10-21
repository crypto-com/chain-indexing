package tendermint_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTendermint(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tendermint Suite")
}
