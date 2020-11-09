package rdbeventstore_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRdbeventstore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rdbeventstore Suite")
}
