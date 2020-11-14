package coin_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCoin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Coin Suite")
}
