package usecase_parser_test

import (
	"github.com/crypto-com/chain-indexing/usecase/parser"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
)

func InitParserManager() *utils.CosmosParserManager {
	pm := utils.NewCosmosParserManager(
		utils.CosmosParserManagerParams{
			Logger: nil,
			Config: utils.CosmosParserManagerConfig{
				CosmosVersionBlockHeight: utils.CosmosVersionBlockHeight{
					V0_42_7: 0,
				},
			},
		})

	parser.InitParsers(pm)

	return pm
}
