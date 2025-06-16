package msg

import (
	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
)

func ParseMsgEventsToLog(
	parser utils.CosmosParser,
) utils.CosmosParser {
	return func(
		parserParams utils.CosmosParserParams,
	) ([]command.Command, []string) {
		events := utils.NewParsedTxsResultsEvents(parserParams.TxsResult.Events)
		log := events.ParsedEventToTxsResultLog()
		parserParams.TxsResult.Log = log

		return parser(utils.CosmosParserParams{
			AddressPrefix:   parserParams.AddressPrefix,
			StakingDenom:    parserParams.StakingDenom,
			TxsResult:       parserParams.TxsResult,
			MsgCommonParams: parserParams.MsgCommonParams,
			Msg:             parserParams.Msg,
			MsgIndex:        parserParams.MsgIndex,
			ParserManager:   parserParams.ParserManager,
		})
	}
}
