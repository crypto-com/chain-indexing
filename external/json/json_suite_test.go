package json_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUtctime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Json Suite")
}
