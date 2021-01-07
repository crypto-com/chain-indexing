package crossfire_test

import (
	"net/http"

	"github.com/crypto-com/chain-indexing/appinterface/projection/crossfire"
	. "github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	. "github.com/crypto-com/chain-indexing/internal/logger/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Crossfire Chain Stats", func() {
	It("should get participants correctly", func() {

		server := ghttp.NewServer()

		server.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/"),
				ghttp.RespondWith(http.StatusOK, `[{
						"operatorAddress": "tcrocncl14m5a4kxt2e82uqqs5gtqza29dm5wqzyalddug5",
						"primaryAddress": "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
						"moniker": "Bambarello"
					  },
					  {
						"operatorAddress": "tcrocncl1n4t5q77kn9vf73s7ljs96m85jgg49yqpg0chrj",
						"primaryAddress": "tcro1f6qcvp33dc79xzpuwll7mln5lnepuqv8d7led9",
						"moniker": "eric-node"
					  },
					  {
						"operatorAddress": "tcrocncl1f6qcvp33dc79xzpuwll7mln5lnepuqv8cpuq4x",
						"primaryAddress": "tcro1432x4lc5mrgm30c9xx35unmn9ultemm5nt40vq",
						"moniker": "erasde"
					  }]`,
				),
			))
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
			server.URL(),
		)

		participants, err := projection.Client.Participants()
		Expect(err).To(BeNil())
		// should match the 200 participants JSON
		Expect(len(*participants)).To(Equal(3))
	})
})
