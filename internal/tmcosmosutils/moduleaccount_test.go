package tmcosmosutils_test

import (
	"github.com/crypto-com/chain-indexing/internal/tmcosmosutils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("tmcosmosutils", func() {
	Describe("NewModuleAccounts", func() {
		It("should work", func() {
			moduleAccounts := tmcosmosutils.NewModuleAccounts("tcro")

			Expect(moduleAccounts.FeeCollector).To(Equal("tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha"))
			Expect(moduleAccounts.Mint).To(Equal("tcro1m3h30wlvsf8llruxtpukdvsy0km2kum87lx9mq"))
			Expect(moduleAccounts.Distribution).To(Equal("tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l"))
			Expect(moduleAccounts.Gov).To(Equal("tcro10d07y265gmmuvt4z0w9aw880jnsr700jvvjc2n"))
			Expect(moduleAccounts.BondedTokensPool).To(Equal("tcro1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3r4gj9h"))
			Expect(moduleAccounts.NotBondedTokensPool).To(Equal("tcro1tygms3xhhs3yv487phx3dw4a95jn7t7lh45rnr"))
			Expect(moduleAccounts.IBCTransfer).To(Equal("tcro1yl6hdjhmkf37639730gffanpzndzdpmhc3h5tr"))
		})

		It("should work for mainnet", func() {
			moduleAccounts := tmcosmosutils.NewModuleAccounts("cro")

			Expect(moduleAccounts.FeeCollector).To(Equal("cro17xpfvakm2amg962yls6f84z3kell8c5lgztehv"))
			Expect(moduleAccounts.Mint).To(Equal("cro1m3h30wlvsf8llruxtpukdvsy0km2kum8s20pm3"))
			Expect(moduleAccounts.Distribution).To(Equal("cro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8lyv94w"))
			Expect(moduleAccounts.Gov).To(Equal("cro10d07y265gmmuvt4z0w9aw880jnsr700jzemu2z"))
			Expect(moduleAccounts.BondedTokensPool).To(Equal("cro1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3dqpk9x"))
			Expect(moduleAccounts.NotBondedTokensPool).To(Equal("cro1tygms3xhhs3yv487phx3dw4a95jn7t7leqa8nj"))
			Expect(moduleAccounts.IBCTransfer).To(Equal("cro1yl6hdjhmkf37639730gffanpzndzdpmhky7stj"))
		})
	})
})
