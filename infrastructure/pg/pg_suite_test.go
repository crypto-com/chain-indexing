package pg_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPg(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pg Suite")
}
