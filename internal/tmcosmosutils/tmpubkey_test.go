package tmcosmosutils_test

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/crypto-com/chain-indexing/internal/tmcosmosutils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const tendermintPubKey = "na51D8RmKXyWrid9I6wtdxgP6f1Nl3EyNNEzqxVquoM="
const tendermintAddress = "B5EC6D86F8F418F480799447F5C21F1C17C6F8F8"
const validatorAddress = "tcrocncl1khkxmphc7sv0fqrej3rltsslrstud78c6dur8e"
const primaryAddress = "tcro1khkxmphc7sv0fqrej3rltsslrstud78c0jl6l6"
const accountAddress = "tcro1khkxmphc7sv0fqrej3rltsslrstud78c0jl6l6"
const consensusNodeAddress = "tcrocnclcons1khkxmphc7sv0fqrej3rltsslrstud78cam9ekl"
const consensusNodePubKey = "tcrocnclconspub1zcjduepqnkh82r7yvc5he94wya7j8tpdwuvql60afkthzv356ye6k9t2h2psr3u067"

var _ = Describe("tmcosmosutils", func() {
	Describe("TmAddressFromTmPubKey", func() {
		It("should work", func() {
			pubKey, _ := base64.StdEncoding.DecodeString(tendermintPubKey)
			Expect(tmcosmosutils.TmAddressFromTmPubKey(
				pubKey,
			)).To(Equal(tendermintAddress))
		})
	})

	Describe("ConsensusNodeAddressFromTmPubKey", func() {
		It("should work", func() {
			pubKey, _ := base64.StdEncoding.DecodeString(tendermintPubKey)
			Expect(tmcosmosutils.ConsensusNodeAddressFromTmPubKey(
				"tcrocnclcons", pubKey,
			)).To(Equal(consensusNodeAddress))
		})
	})

	Describe("ConsensusNodePubKeyFromTmPubKey", func() {
		It("should work", func() {
			pubKey, _ := base64.StdEncoding.DecodeString(tendermintPubKey)
			Expect(tmcosmosutils.ConsensusNodePubKeyFromTmPubKey(
				"tcrocnclconspub", pubKey,
			)).To(Equal(consensusNodePubKey))
		})
	})

	Describe("ValidatorAddressFromPubAddress", func() {
		It("should work from council node address", func() {
			Expect(tmcosmosutils.ValidatorAddressFromPubAddress(
				"tcrocncl", consensusNodeAddress,
			)).To(Equal(validatorAddress))
		})
		It("should work from tcro.. address to validator address", func() {
			Expect(tmcosmosutils.ValidatorAddressFromPubAddress(
				"tcrocncl", primaryAddress,
			)).To(Equal(validatorAddress))
		})
		It("should work from tcro.. address to council node address", func() {
			Expect(tmcosmosutils.ValidatorAddressFromPubAddress(
				"tcrocnclcons", primaryAddress,
			)).To(Equal(consensusNodeAddress))
		})
	})

	Describe("ConsensusNodeAddressFromConsensusNodePubKey", func() {
		It("should work", func() {
			Expect(tmcosmosutils.ConsensusNodeAddressFromConsensusNodePubKey(
				"tcrocnclcons", consensusNodePubKey,
			)).To(Equal(consensusNodeAddress))
		})
	})

	Describe("AddressFromPubKey", func() {
		It("should work", func() {
			pubkey, _ := base64.StdEncoding.DecodeString("A3ill3YNyWvcMstrbssC9SpzhMm+tCMWPB7bgOqWQZYk")
			Expect(tmcosmosutils.AccountAddressFromPubKey(
				"tcro", pubkey,
			)).To(Equal("tcro1p4fzn6ta24c6ek4v2qls6y5uug44ku9tnypcaf"))
		})
	})

	Describe("ValidatorAddressFromAccountAddress", func() {
		It("should work from account address to validator address", func() {
			Expect(tmcosmosutils.ValidatorAddressFromAccountAddress(
				"tcrocncl", accountAddress,
			)).To(Equal(validatorAddress))
		})
	})

	Describe("AddressFromPubKey", func() {
		It("should work", func() {
			pubkey, _ := base64.StdEncoding.DecodeString("A3ill3YNyWvcMstrbssC9SpzhMm+tCMWPB7bgOqWQZYk")
			Expect(tmcosmosutils.AccountAddressFromPubKey(
				"tcro", pubkey,
			)).To(Equal("tcro1p4fzn6ta24c6ek4v2qls6y5uug44ku9tnypcaf"))
		})
	})

	Describe("PubKeyFromCosmosPubKey", func() {
		It("should work", func() {
			pubKey, err := tmcosmosutils.PubKeyFromCosmosPubKey(
				"tcropub1addwnpepqvg3kxeltzkv5xdhcw36sts6wkwgfq5yjmw98nt878u06ap634pzqdgcjrz",
			)
			Expect(err).To(BeNil())
			Expect(hex.EncodeToString(pubKey)).To(Equal("03111b1b3f58acca19b7c3a3a82e1a759c84828496dc53cd67f1f8fd743a8d4220"))
		})
	})

	Describe("MultiSigAddressFromPubKeys", func() {
		It("should work", func() {
			pubKey1, _ := tmcosmosutils.PubKeyFromCosmosPubKey(
				"tcropub1addwnpepqvg3kxeltzkv5xdhcw36sts6wkwgfq5yjmw98nt878u06ap634pzqdgcjrz",
			)
			pubKey2, _ := tmcosmosutils.PubKeyFromCosmosPubKey(
				"tcropub1addwnpepqv7jajxxuack203a3ut2usxnxldmjpukhlkc8yv68fnjlk5j58qg68yxxs9",
			)
			pubKey3, _ := tmcosmosutils.PubKeyFromCosmosPubKey(
				"tcropub1addwnpepq0ck3tt7qwg2jt6mce7eagn4jmh660c6lt3y8gl85f75ex7qycrvgtw6272",
			)

			threshold := 2
			sortPubKeys := true
			Expect(tmcosmosutils.MultiSigAddressFromPubKeys(
				"tcro", [][]byte{
					pubKey1, pubKey2, pubKey3,
				}, threshold, sortPubKeys,
			)).To(Equal("tcro1xc5uw8j6h3cjd7m2l9pn7xzg97q5pv9mdvp8ly"))
		})

		It("should follow the provided public key orders when unsorted", func() {
			pubKey1, _ := base64.StdEncoding.DecodeString("Az0uyMbncWU+PY8WrkDTN9u5B5a/7YORmjpnL9qSocCN")
			pubKey2, _ := base64.StdEncoding.DecodeString("AxEbGz9YrMoZt8OjqC4adZyEgoSW3FPNZ/H4/XQ6jUIg")
			pubKey3, _ := base64.StdEncoding.DecodeString("A/ForX4DkKkvW8Z9nqJ1lu+tPxr64kOj56J9TJvAJgbE")

			threshold := 2
			sortPubKeys := false
			Expect(tmcosmosutils.MultiSigAddressFromPubKeys(
				"tcro", [][]byte{
					pubKey1, pubKey2, pubKey3,
				}, threshold, sortPubKeys,
			)).To(Equal("tcro1xc5uw8j6h3cjd7m2l9pn7xzg97q5pv9mdvp8ly"))
		})

		It("should follow the provided public key orders when unsorted", func() {
			pubKey1, _ := base64.StdEncoding.DecodeString("AxEbGz9YrMoZt8OjqC4adZyEgoSW3FPNZ/H4/XQ6jUIg")
			pubKey2, _ := base64.StdEncoding.DecodeString("Az0uyMbncWU+PY8WrkDTN9u5B5a/7YORmjpnL9qSocCN")
			pubKey3, _ := base64.StdEncoding.DecodeString("A/ForX4DkKkvW8Z9nqJ1lu+tPxr64kOj56J9TJvAJgbE")

			threshold := 2
			sortPubKeys := false
			Expect(tmcosmosutils.MultiSigAddressFromPubKeys(
				"tcro", [][]byte{
					pubKey1, pubKey2, pubKey3,
				}, threshold, sortPubKeys,
			)).To(Equal("tcro1se7jsq9ax3qqm3uyc0aullneu25fxm56u8ryqw"))
		})
	})
})
