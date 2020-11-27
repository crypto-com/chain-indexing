package parser

import (
	"fmt"

	"github.com/crypto-com/chainindex/entity/command"
	command_usecase "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model/genesis"
)

func ParseGenesisCommands(genesis *genesis.Genesis) ([]command.Command, error) {
	commands := []command.Command{
		command_usecase.NewCreateGenesis(*genesis),
	}
	for txIndex, genTx := range genesis.AppState.Genutil.GenTxs {
		for msgIndex, message := range genTx.Body.Messages {
			if message["@type"] == "/cosmos.staking.v1beta1.MsgCreateValidator" {
				msgCommonParams := event.MsgCommonParams{
					BlockHeight: 0,
					TxHash:      fmt.Sprintf("genesis-gentxs-%d", txIndex),
					TxSuccess:   true,
					MsgIndex:    msgIndex,
				}
				commands = append(commands, parseMsgCreateValidator(msgCommonParams, message)...)
			}
		}
	}
	return commands, nil
}
