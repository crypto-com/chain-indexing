package createblock_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCreateBlock(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CreateBlock Suite")
}
