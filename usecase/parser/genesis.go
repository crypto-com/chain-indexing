package parser

import (
	"github.com/crypto-com/chainindex/entity/command"
	command_usecase "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/model/genesis"
)

func ParseGenesisCommands(genesis *genesis.Genesis) ([]command.Command, error) {
	return []command.Command{
		command_usecase.NewCreateGenesis(*genesis),
	}, nil
}
