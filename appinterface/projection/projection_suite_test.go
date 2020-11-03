package projection_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestProjection(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Projection Suite")
}
