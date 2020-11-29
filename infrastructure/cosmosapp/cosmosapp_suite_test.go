package cosmosapp_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCosmosapp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cosmosapp Suite")
}
