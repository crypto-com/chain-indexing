package parser

import (
	"fmt"

	"github.com/crypto-com/chainindex/usecase/command"

	entity_command "github.com/crypto-com/chainindex/entity/command"
	usecase_model "github.com/crypto-com/chainindex/usecase/model"
)

func ParseBlockToCommands(
	moduleAccounts *ModuleAccounts,
	block *usecase_model.Block,
	rawBlock *usecase_model.RawBlock,
	blockResults *usecase_model.BlockResults,
) ([]entity_command.Command, error) {
	var err error
	var commands []entity_command.Command

	createRawBlockCommand := ParseCreateRawBlockCommand(rawBlock)
	commands = append(commands, createRawBlockCommand)

	createBlockCommand := ParseCreateBlockCommand(block)
	commands = append(commands, createBlockCommand)

	if len(blockResults.TxsResults) > 0 {
		transactionCommands, err := ParseTransactionCommands(moduleAccounts, block, blockResults)
		if err != nil {
			return nil, fmt.Errorf("error parsing transaction commands: %v", err)
		}
		commands = append(commands, transactionCommands...)
	}

	return commands, err
}

func ParseCreateRawBlockCommand(rawBlock *usecase_model.RawBlock) *command.CreateRawBlock {
	return command.NewCreateRawBlock(rawBlock)
}

func ParseCreateBlockCommand(block *usecase_model.Block) *command.CreateBlock {
	var modelBlockSigs []usecase_model.BlockSignature
	for _, sig := range block.Signatures {
		modelBlockSigs = append(modelBlockSigs, usecase_model.BlockSignature{
			BlockIdFlag:      sig.BlockIdFlag,
			ValidatorAddress: sig.ValidatorAddress,
			Timestamp:        sig.Timestamp,
			Signature:        sig.Signature,
		})
	}

	modelBlock := &usecase_model.Block{
		Height:          block.Height,
		Hash:            block.Hash,
		Time:            block.Time,
		AppHash:         block.AppHash,
		ProposerAddress: block.ProposerAddress,
		Txs:             block.Txs,
		Signatures:      modelBlockSigs,
	}

	return command.NewCreateBLock(modelBlock)
}
