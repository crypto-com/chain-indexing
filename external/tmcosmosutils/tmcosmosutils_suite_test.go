package tmcosmosutils_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTmcosmosutils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tmcosmosutils Suite")
}
