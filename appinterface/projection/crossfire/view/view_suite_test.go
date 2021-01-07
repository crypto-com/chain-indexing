package view_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCrossfireView(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Crossfire View Suite")
}
