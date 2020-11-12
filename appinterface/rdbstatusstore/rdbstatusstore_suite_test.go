package rdbstatusstore_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRdbstatusstore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rdbstatusstore Suite")
}
