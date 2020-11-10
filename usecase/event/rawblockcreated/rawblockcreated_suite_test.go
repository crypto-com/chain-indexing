package rawblockcreated_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRawBlockCreated(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RawBlockCreated Suite")
}
