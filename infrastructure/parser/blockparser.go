package parser

import (
	"github.com/crypto-com/chainindex/appinterface/tendermint/types"
	"github.com/crypto-com/chainindex/ddd/command"
	entity_command "github.com/crypto-com/chainindex/entity/command"
	"github.com/crypto-com/chainindex/entity/model"
)

func ParseBlockToCommands(block *types.Block, rawBlock *model.RawBlock) ([]command.Command, error) {
	var err error
	var cmds []command.Command

	createRawBlockCommand := ParseCreateRawBlockCommand(rawBlock)
	cmds = append(cmds, &createRawBlockCommand)

	createBlockCommand := ParseCreateBlockCommand(block)
	cmds = append(cmds, &createBlockCommand)

	return cmds, err
}

func ParseCreateRawBlockCommand(rawBlock *model.RawBlock) entity_command.CreateRawBlockCommand {
	return entity_command.NewCreateRawBlockCommand(rawBlock)
}

func ParseCreateBlockCommand(block *types.Block) entity_command.CreateBlockCommand {
	var modelBlockSigs []model.BlockSignature
	for _, sig := range block.Signatures {
		modelBlockSigs = append(modelBlockSigs, model.BlockSignature{
			BlockIdFlag:      sig.BlockIDFlag,
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

	return entity_command.NewCreateBlockCommand(modelBlock)
}
