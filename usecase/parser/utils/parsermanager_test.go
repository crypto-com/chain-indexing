package utils_test

import (
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParseManager", func() {
	It("should proper version height", func() {

		pm := utils.NewCosmosParserManager(
			utils.CosmosParserManagerParams{
				Logger: nil,
				Config: utils.CosmosParserManagerConfig{
					CosmosVersionBlockHeight: utils.CosmosVersionBlockHeight{
						V0_43_0: 123,
					},
				},
			},
		)

		Expect(pm.GetCosmosV0_43_0BlockHeight()).To(Equal(utils.ParserBlockHeight(123)))
	})

	It("should get parser A", func() {

		pm := utils.NewCosmosParserManager(
			utils.CosmosParserManagerParams{
				Logger: nil,
				Config: utils.CosmosParserManagerConfig{
					CosmosVersionBlockHeight: utils.CosmosVersionBlockHeight{
						V0_43_0: 0,
					},
				},
			},
		)

		parserKey := utils.CosmosParserKey("parser")

		pm.RegisterParser(parserKey, 0, test.ParserA)

		p := pm.GetParser(parserKey, 0)
		cmds := p(utils.CosmosParserParams{})

		Expect(cmds[0].Name()).To(Equal("commandA"))

		p = pm.GetParser(parserKey, 1)
		cmds = p(utils.CosmosParserParams{})

		Expect(cmds[0].Name()).To(Equal("commandA"))
	})

	It("should able to get different parser with different height", func() {

		pm := utils.NewCosmosParserManager(
			utils.CosmosParserManagerParams{
				Logger: nil,
				Config: utils.CosmosParserManagerConfig{
					CosmosVersionBlockHeight: utils.CosmosVersionBlockHeight{
						V0_43_0: 0,
					},
				},
			},
		)

		parserKey := utils.CosmosParserKey("parser")

		pm.RegisterParser(parserKey, 0, test.ParserA)

		pm.RegisterParser(parserKey, 10, test.ParserB)

		p := pm.GetParser(parserKey, 10)
		cmds := p(utils.CosmosParserParams{})

		Expect(cmds[0].Name()).To(Equal("commandB"))

		p = pm.GetParser(parserKey, 1)
		cmds = p(utils.CosmosParserParams{})

		Expect(cmds[0].Name()).To(Equal("commandA"))
	})

	It("should able to get overridden parser", func() {

		pm := utils.NewCosmosParserManager(
			utils.CosmosParserManagerParams{
				Logger: nil,
				Config: utils.CosmosParserManagerConfig{
					CosmosVersionBlockHeight: utils.CosmosVersionBlockHeight{
						V0_43_0: 0,
					},
				},
			},
		)

		parserKey := utils.CosmosParserKey("parser")

		pm.RegisterParser(parserKey, 0, test.ParserA)

		p := pm.GetParser(parserKey, 1)
		cmds := p(utils.CosmosParserParams{})

		Expect(cmds[0].Name()).To(Equal("commandA"))

		pm.RegisterParser(parserKey, 0, test.ParserB)

		p = pm.GetParser(parserKey, 1)
		cmds = p(utils.CosmosParserParams{})

		Expect(cmds[0].Name()).To(Equal("commandB"))
	})
})
