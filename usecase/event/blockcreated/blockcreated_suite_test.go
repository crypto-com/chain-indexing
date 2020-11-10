package blockcreated_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBlockCreated(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Blockcreated Suite")
}
