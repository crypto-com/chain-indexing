package tendermint_test

import (
	"fmt"
	"net/http"

	infrastructure_tendermint_test "github.com/crypto-com/chain-indexing/infrastructure/tendermint/test"

	jsoniter "github.com/json-iterator/go"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	"github.com/crypto-com/chain-indexing/appinterface/tendermint"
	"github.com/crypto-com/chain-indexing/external/utctime"
	. "github.com/crypto-com/chain-indexing/infrastructure/tendermint"

	usecase_model "github.com/crypto-com/chain-indexing/usecase/model"
)

var _ = Describe("HTTPClient", func() {
	var server *ghttp.Server

	BeforeEach(func() {
		server = ghttp.NewServer()
	})

	It("should implement Client", func() {
		var _ tendermint.Client = NewHTTPClient("http://localhost:26657", true)
	})

	Describe("RawBlockResults", func() {
		It("should return nil Events when there are no transactions nor events", func() {
			anyBlockHeight := int64(1)
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/block_results", fmt.Sprintf("height=%d", anyBlockHeight)),
					ghttp.RespondWith(http.StatusOK, infrastructure_tendermint_test.BLOCK_RESULTS_EMPTY_EVENTS_JSON),
				),
			)

			client := NewHTTPClient(server.URL(), true)

			blockResults, err := client.BlockResults(anyBlockHeight)
			Expect(err).To(BeNil())
			Expect(*blockResults).To(Equal(usecase_model.BlockResults{
				Height:           anyBlockHeight,
				TxsResults:       []usecase_model.BlockResultsTxsResult{},
				BeginBlockEvents: []usecase_model.BlockResultsEvent{},
				EndBlockEvents:   []usecase_model.BlockResultsEvent{},
				ValidatorUpdates: []usecase_model.BlockResultsValidatorUpdate{},
				ConsensusParamUpdates: usecase_model.BlockResultsConsensusParamUpdates{
					Block: usecase_model.BlockResultsConsensusParamUpdatesBlock{
						MaxBytes: "22020096",
						MaxGas:   "-1",
					},
					Evidence: usecase_model.BlockResultsConsensusParamUpdatesEvidence{
						MaxAgeNumBlocks: "100000",
						MaxAgeDuration:  "172800000000000",
					},
					Validator: usecase_model.BlockResultsConsensusParamsUpdatesValidator{
						PubKeyTypes: []string{"ed25519"},
					},
				},
			}))
		})

		It("should return parsed block results", func() {
			anyBlockHeight := int64(3813)
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/block_results", fmt.Sprintf("height=%d", anyBlockHeight)),
					ghttp.RespondWith(http.StatusOK, infrastructure_tendermint_test.BLOCK_RESULTS_JSON),
				),
			)

			client := NewHTTPClient(server.URL(), true)

			blockResults, err := client.BlockResults(anyBlockHeight)
			Expect(err).To(BeNil())
			expected := "{\"height\":367216,\"txsResults\":[{\"code\":0,\"data\":\"\\n\\n\\n\\u0008delegate\",\"log\":[{\"msgIndex\":0,\"events\":[{\"type\":\"delegate\",\"attributes\":[{\"key\":\"validator\",\"value\":\"tcrocncl1pm27djcs5djxjsxw3unrkv3m3jtxdexktw5epu\",\"index\":false},{\"key\":\"amount\",\"value\":\"19302674761\",\"index\":false}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"delegate\",\"index\":false},{\"key\":\"sender\",\"value\":\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\",\"index\":false},{\"key\":\"module\",\"value\":\"staking\",\"index\":false},{\"key\":\"sender\",\"value\":\"tcro1pm27djcs5djxjsxw3unrkv3m3jtxdexk73hqel\",\"index\":false}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro1pm27djcs5djxjsxw3unrkv3m3jtxdexk73hqel\",\"index\":false},{\"key\":\"sender\",\"value\":\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\",\"index\":false},{\"key\":\"amount\",\"value\":\"1913979901basetcro\",\"index\":false}]}]}],\"rawLog\":\"[{\\\"events\\\":[{\\\"type\\\":\\\"delegate\\\",\\\"attributes\\\":[{\\\"key\\\":\\\"validator\\\",\\\"value\\\":\\\"tcrocncl1pm27djcs5djxjsxw3unrkv3m3jtxdexktw5epu\\\"},{\\\"key\\\":\\\"amount\\\",\\\"value\\\":\\\"19302674761\\\"}]},{\\\"type\\\":\\\"message\\\",\\\"attributes\\\":[{\\\"key\\\":\\\"action\\\",\\\"value\\\":\\\"delegate\\\"},{\\\"key\\\":\\\"sender\\\",\\\"value\\\":\\\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\\\"},{\\\"key\\\":\\\"module\\\",\\\"value\\\":\\\"staking\\\"},{\\\"key\\\":\\\"sender\\\",\\\"value\\\":\\\"tcro1pm27djcs5djxjsxw3unrkv3m3jtxdexk73hqel\\\"}]},{\\\"type\\\":\\\"transfer\\\",\\\"attributes\\\":[{\\\"key\\\":\\\"recipient\\\",\\\"value\\\":\\\"tcro1pm27djcs5djxjsxw3unrkv3m3jtxdexk73hqel\\\"},{\\\"key\\\":\\\"sender\\\",\\\"value\\\":\\\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\\\"},{\\\"key\\\":\\\"amount\\\",\\\"value\\\":\\\"1913979901basetcro\\\"}]}]}]\",\"info\":\"\",\"gasWanted\":\"200000\",\"gasUsed\":\"143179\",\"events\":[{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha\",\"index\":true},{\"key\":\"sender\",\"value\":\"tcro1pm27djcs5djxjsxw3unrkv3m3jtxdexk73hqel\",\"index\":true},{\"key\":\"amount\",\"value\":\"20000basetcro\",\"index\":true}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"tcro1pm27djcs5djxjsxw3unrkv3m3jtxdexk73hqel\",\"index\":true}]}],\"codespace\":\"\"},{\"code\":0,\"data\":\"\\n\\n\\n\\u0008delegate\",\"log\":[{\"msgIndex\":0,\"events\":[{\"type\":\"delegate\",\"attributes\":[{\"key\":\"validator\",\"value\":\"tcrocncl10gsqs8jzdlrem80shp0x6wx0jw7qu7m8cd29y5\",\"index\":false},{\"key\":\"amount\",\"value\":\"17338013566\",\"index\":false}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"delegate\",\"index\":false},{\"key\":\"sender\",\"value\":\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\",\"index\":false},{\"key\":\"module\",\"value\":\"staking\",\"index\":false},{\"key\":\"sender\",\"value\":\"tcro10gsqs8jzdlrem80shp0x6wx0jw7qu7m8djfuuh\",\"index\":false}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro10gsqs8jzdlrem80shp0x6wx0jw7qu7m8djfuuh\",\"index\":false},{\"key\":\"sender\",\"value\":\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\",\"index\":false},{\"key\":\"amount\",\"value\":\"2310941249basetcro\",\"index\":false}]}]}],\"rawLog\":\"[{\\\"events\\\":[{\\\"type\\\":\\\"delegate\\\",\\\"attributes\\\":[{\\\"key\\\":\\\"validator\\\",\\\"value\\\":\\\"tcrocncl10gsqs8jzdlrem80shp0x6wx0jw7qu7m8cd29y5\\\"},{\\\"key\\\":\\\"amount\\\",\\\"value\\\":\\\"17338013566\\\"}]},{\\\"type\\\":\\\"message\\\",\\\"attributes\\\":[{\\\"key\\\":\\\"action\\\",\\\"value\\\":\\\"delegate\\\"},{\\\"key\\\":\\\"sender\\\",\\\"value\\\":\\\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\\\"},{\\\"key\\\":\\\"module\\\",\\\"value\\\":\\\"staking\\\"},{\\\"key\\\":\\\"sender\\\",\\\"value\\\":\\\"tcro10gsqs8jzdlrem80shp0x6wx0jw7qu7m8djfuuh\\\"}]},{\\\"type\\\":\\\"transfer\\\",\\\"attributes\\\":[{\\\"key\\\":\\\"recipient\\\",\\\"value\\\":\\\"tcro10gsqs8jzdlrem80shp0x6wx0jw7qu7m8djfuuh\\\"},{\\\"key\\\":\\\"sender\\\",\\\"value\\\":\\\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\\\"},{\\\"key\\\":\\\"amount\\\",\\\"value\\\":\\\"2310941249basetcro\\\"}]}]}]\",\"info\":\"\",\"gasWanted\":\"200000\",\"gasUsed\":\"143737\",\"events\":[{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha\",\"index\":true},{\"key\":\"sender\",\"value\":\"tcro10gsqs8jzdlrem80shp0x6wx0jw7qu7m8djfuuh\",\"index\":true},{\"key\":\"amount\",\"value\":\"20000basetcro\",\"index\":true}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"tcro10gsqs8jzdlrem80shp0x6wx0jw7qu7m8djfuuh\",\"index\":true}]}],\"codespace\":\"\"}],\"beginBlockEvents\":[{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha\",\"index\":true},{\"key\":\"sender\",\"value\":\"tcro1m3h30wlvsf8llruxtpukdvsy0km2kum87lx9mq\",\"index\":true},{\"key\":\"amount\",\"value\":\"17449528321basetcro\",\"index\":true}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"tcro1m3h30wlvsf8llruxtpukdvsy0km2kum87lx9mq\",\"index\":true}]},{\"type\":\"mint\",\"attributes\":[{\"key\":\"bonded_ratio\",\"value\":\"0.000809196054376644\",\"index\":true},{\"key\":\"inflation\",\"value\":\"0.013755821936855184\",\"index\":true},{\"key\":\"annual_provisions\",\"value\":\"110133046994204576.138016526579386288\",\"index\":true},{\"key\":\"amount\",\"value\":\"17449528321\",\"index\":true}]}],\"endBlockEvents\":[{\"type\":\"commission\",\"attributes\":[{\"key\":\"amount\",\"value\":\"87247841.605000000000000000basetcro\",\"index\":true},{\"key\":\"validator\",\"value\":\"tcrocncl18p07yvmphymscz6tl4a7zmh93g0k6vy72ww4s4\",\"index\":true}]},{\"type\":\"rewards\",\"attributes\":[{\"key\":\"amount\",\"value\":\"872478416.050000000000000000basetcro\",\"index\":true},{\"key\":\"validator\",\"value\":\"tcrocncl18p07yvmphymscz6tl4a7zmh93g0k6vy72ww4s4\",\"index\":true}]}],\"validatorUpdates\":[{\"pubkey\":{\"type\":\"tendermint.crypto.PublicKey_Ed25519\",\"pubkey\":\"SE5zeTjcYPXVrfcOva61QWokSZFfQu2h316fR6bB2dY=\"},\"address\":\"CA721C3A05F500838DDD1B16F4E2D2D09E463218\",\"power\":138525202},{\"pubkey\":{\"type\":\"tendermint.crypto.PublicKey_Ed25519\",\"pubkey\":\"Epmo3U6yXlxSDQzWZ8yBPOMHw2R85lc26RK98Rlo0oM=\"},\"address\":\"E067FCE33F7FDBD0CE4872F8E240A7AD6E654726\",\"power\":112904113}],\"consensusParamUpdates\":{\"block\":{\"maxBytes\":\"22020096\",\"maxGas\":\"-1\"},\"evidence\":{\"maxAgeNumBlocks\":\"100000\",\"maxAgeDuration\":\"172800000000000\",\"maxBytes\":\"\"},\"validator\":{\"pubKeyTypes\":[\"ed25519\"]}}}"
			Expect(jsoniter.MarshalToString(blockResults)).To(Equal(expected))
		})

		It("should return parsed block results when validator updates has no power", func() {
			anyBlockHeight := int64(3813)
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/block_results", fmt.Sprintf("height=%d", anyBlockHeight)),
					ghttp.RespondWith(http.StatusOK, infrastructure_tendermint_test.BLOCK_RESULTS_VALIDATOR_UPDATES_WITHOUT_POWER_JSON),
				),
			)

			client := NewHTTPClient(server.URL(), true)

			blockResults, err := client.BlockResults(anyBlockHeight)
			Expect(err).To(BeNil())
			Expect(jsoniter.MarshalToString(blockResults)).To(Equal("{\"height\":4794,\"txsResults\":[],\"beginBlockEvents\":[],\"endBlockEvents\":[],\"validatorUpdates\":[{\"pubkey\":{\"type\":\"tendermint.crypto.PublicKey_Ed25519\",\"pubkey\":\"CpCz+c19SHaNWW31P+7blzyHo0sQMn4uk8gIej+pXW8=\"},\"address\":\"02D50170AFB2B3718477AA607469BEA96C80CE88\",\"power\":null}],\"consensusParamUpdates\":{\"block\":{\"maxBytes\":\"22020096\",\"maxGas\":\"-1\"},\"evidence\":{\"maxAgeNumBlocks\":\"100000\",\"maxAgeDuration\":\"172800000000000\",\"maxBytes\":\"\"},\"validator\":{\"pubKeyTypes\":[\"ed25519\"]}}}"))
		})
	})

	Describe("Block", func() {
		It("should return block response", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/block", "height=100"),
					ghttp.RespondWith(http.StatusOK, infrastructure_tendermint_test.BLOCK_JSON),
				),
			)

			blockHeight := int64(100)
			client := NewHTTPClient(server.URL(), true)
			block, _, err := client.Block(blockHeight)
			Expect(err).To(BeNil())
			blockTime, _ := utctime.Parse("2006-01-02T15:04:05.000000000Z", "2020-10-15T09:33:42.195143319Z")
			signature0Time, _ := utctime.Parse("2006-01-02T15:04:05.00000000Z", "2020-10-15T09:33:42.18646236Z")
			signature1Time, _ := utctime.Parse("2006-01-02T15:04:05.000000000Z", "2020-10-15T09:33:42.195143319Z")
			signature2Time, _ := utctime.Parse("2006-01-02T15:04:05.000000000Z", "2020-10-15T09:33:42.206633743Z")

			Expect(*block).To(Equal(usecase_model.Block{
				Height:          blockHeight,
				Hash:            "82C25937191D1CF73BE9222CB04CE35B7A1366CC5BB08D9BB9AB457712E4F2D1",
				Time:            blockTime,
				AppHash:         "6AE0920938F76727054BC2531247632C5C0521E2B91EA3A9864EA4FF55023D77",
				ProposerAddress: "384E5F30F02538C0A34CBFF32F8D5554671C9029",
				Txs:             []string{},
				Signatures: []usecase_model.BlockSignature{
					{
						BlockIdFlag:      2,
						ValidatorAddress: "384E5F30F02538C0A34CBFF32F8D5554671C9029",
						Timestamp:        signature0Time,
						Signature:        "Lnjz+jTGTk6DzqvbvdjFIG5zyzzpioiN37/B2pKi6ac/Kf2a+kmfeQA3CXnPO5GY/AoImfbVcCQs4asTSDCqCA==",
					},
					{
						BlockIdFlag:      2,
						ValidatorAddress: "4D9F47C5A19D5550685D2D55DF2FF8C5D7504CEB",
						Timestamp:        signature1Time,
						Signature:        "OpYxiF7QTAaG4NG2/iHWvts8ISQED6F+pNU5Cv2ts8c8mFW9J+g0oig3IXhvVG011uQjcr0CVw5P0eOXx1vYDg==",
					},
					{
						BlockIdFlag:      2,
						ValidatorAddress: "6D9E4B4995037D608E365CE90436C24580ABCC33",
						Timestamp:        signature2Time,
						Signature:        "jQd4JNrvX6DKmqDZ9VqoKtxRIxQHrvPWd4XW+ayrtVakiIMCWoVf1GMvxLbXYg68CyjmbuAX2VhCD0gSnj3pAw==",
					},
				},
				Evidences: []usecase_model.BlockEvidence{},
			}))
		})
	})

	Describe("Genesis", func() {
		It("should return genesis response", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/genesis_chunked"),
					ghttp.RespondWith(http.StatusOK, infrastructure_tendermint_test.GENESIS_MIXED_NUMBER_AND_STRING_JSON),
				),
			)
			client := NewHTTPClient(server.URL(), true)
			_, err := client.GenesisChunked()
			Expect(err).To(BeNil())
		})
	})
})
