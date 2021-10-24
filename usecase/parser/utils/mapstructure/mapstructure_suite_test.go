package mapstructure_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMapstructure(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mapstructure Suite")
}
