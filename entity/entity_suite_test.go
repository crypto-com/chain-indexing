package entity_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEntity(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Entity Suite")
}
