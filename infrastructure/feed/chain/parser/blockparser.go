package parser

import (
	entity_command "github.com/crypto-com/chainindex/entity/command"
	usecase_command "github.com/crypto-com/chainindex/usecase/command"
	usecase_model "github.com/crypto-com/chainindex/usecase/model"
)

func ParseBlockToCommands(
	block *usecase_model.Block,
	rawBlock *usecase_model.RawBlock,
	_ *usecase_model.BlockResults,
) ([]entity_command.Command, error) {
	var err error
	var commands []entity_command.Command

	createRawBlockCommand := ParseCreateRawBlockCommand(rawBlock)
	commands = append(commands, &createRawBlockCommand)

	createBlockCommand := ParseCreateBlockCommand(block)
	commands = append(commands, &createBlockCommand)

	return commands, err
}

func ParseCreateRawBlockCommand(rawBlock *usecase_model.RawBlock) usecase_command.CreateRawBlock {
	return usecase_command.NewCreateRawBlock(rawBlock)
}

func ParseCreateBlockCommand(block *usecase_model.Block) usecase_command.CreateBlock {
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

	return usecase_command.NewCreateBlock(modelBlock)
}
