package projection_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/ddd/projection"
	. "github.com/crypto-com/chainindex/ddd/projection/test"
	. "github.com/crypto-com/chainindex/test/fake"
)

var _ = Describe("ProjectionManager", func() {
	Describe("RegisterProjection", func() {
		It("should return Error when the projection with the same Id already exists", func() {
			manager := projection.NewManager(NewFakeRDbConn())

			anyProjection := NewFakeProjection()
			Expect(manager.RegisterProjection(anyProjection)).To(BeNil())
			Expect(manager.RegisterProjection(anyProjection)).To(MatchError("Projection `FakeProjection` already registered"))
		})

		It("should register projection to the manager", func() {
			manager := projection.NewManager(NewFakeRDbConn())

			anyProjection := NewFakeProjection()

			Expect(manager.IsProjectionRegistered(anyProjection)).To(BeFalse())

			Expect(manager.RegisterProjection(anyProjection)).To(BeNil())

			Expect(manager.IsProjectionRegistered(anyProjection)).To(BeTrue())
		})
	})
})
