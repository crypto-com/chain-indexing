package rdb_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRdb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rdb Suite")
}
