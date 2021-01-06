package crossfire_test

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/crossfire"
	. "github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	. "github.com/crypto-com/chain-indexing/internal/logger/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Crossfire Chain Stats", func() {
	It("should get participants correctly", func() {
		fakeLogger := NewFakeLogger()
		fakeRdbConn := NewFakeRDbConn()
		projection := crossfire.NewCrossfire(
			fakeLogger,
			fakeRdbConn,
			"tcrocnclcons",
			"tcrocncl",
			1,
			1,
			1,
			1,
			"foo",
			"14",
			"https://raw.githubusercontent.com/foreseaz/random/master/",
		)

		participants, err := projection.Client.Participants()
		Expect(err).To(BeNil())
		// should match the 200 participants JSON
		Expect(len(*participants)).To(Equal(200))
	})
})
