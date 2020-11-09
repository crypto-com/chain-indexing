package parser

import (
	"github.com/crypto-com/chainindex/entity/command"
	usecase_command "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/model"
)

func ParseBlockToCommands(block *model.Block, rawBlock *model.RawBlock) ([]command.Command, error) {
	var err error
	var cmds []command.Command

	createRawBlockCommand := ParseCreateRawBlockCommand(rawBlock)
	cmds = append(cmds, &createRawBlockCommand)

	createBlockCommand := ParseCreateBlockCommand(block)
	cmds = append(cmds, &createBlockCommand)

	return cmds, err
}

func ParseCreateRawBlockCommand(rawBlock *model.RawBlock) usecase_command.CreateRawBlock {
	return usecase_command.NewCreateRawBlock(rawBlock)
}

func ParseCreateBlockCommand(block *model.Block) usecase_command.CreateBlock {
	var modelBlockSigs []model.BlockSignature
	for _, sig := range block.Signatures {
		modelBlockSigs = append(modelBlockSigs, model.BlockSignature{
			BlockIdFlag:      sig.BlockIdFlag,
			ValidatorAddress: sig.ValidatorAddress,
			Timestamp:        sig.Timestamp,
			Signature:        sig.Signature,
		})
	}

	modelBlock := &model.Block{
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
