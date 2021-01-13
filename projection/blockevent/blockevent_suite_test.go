package blockevent_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBlock(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Block Events Suite")
}
