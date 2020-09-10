package usecase_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUsecase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Usecase Suite")
}
