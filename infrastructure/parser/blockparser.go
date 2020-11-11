package parser

import (
	entity_command "github.com/crypto-com/chainindex/entity/command"
	usecase_command "github.com/crypto-com/chainindex/usecase/command"
	usecase_model "github.com/crypto-com/chainindex/usecase/model"
)

func ParseBlockToCommands(block *usecase_model.Block, rawBlock *usecase_model.RawBlock) ([]entity_command.Command, error) {
	var err error
	var cmds []entity_command.Command

	createRawBlockCommand := ParseCreateRawBlockCommand(rawBlock)
	cmds = append(cmds, createRawBlockCommand)

	createBlockCommand := ParseCreateBlockCommand(block)
	cmds = append(cmds, createBlockCommand)

	signBlockCommands := ParseSignBlockCommands(block)
	cmds = append(cmds, signBlockCommands...)

	return cmds, err
}

func ParseCreateRawBlockCommand(rawBlock *usecase_model.RawBlock) *usecase_command.CreateRawBlock {
	return usecase_command.NewCreateRawBlock(rawBlock)
}

func ParseCreateBlockCommand(block *usecase_model.Block) *usecase_command.CreateBlock {
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

func ParseSignBlockCommands(block *usecase_model.Block) []entity_command.Command {
	commands := make([]entity_command.Command, len(block.Signatures))

	for _, signature := range block.Signatures {
		commands = append(commands, usecase_command.NewSignBlock(
			block.Height,
			signature.ValidatorAddress,
			block.ProposerAddress == signature.ValidatorAddress,
			signature.Timestamp,
			signature.Signature,
		))
	}

	return commands
}
