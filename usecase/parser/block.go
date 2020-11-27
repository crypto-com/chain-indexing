package parser

import (
	"fmt"

	"github.com/crypto-com/chainindex/usecase/command"

	entity_command "github.com/crypto-com/chainindex/entity/command"
	usecase_model "github.com/crypto-com/chainindex/usecase/model"
)

func ParseBlockToCommands(
	txDecoder *TxDecoder,
	block *usecase_model.Block,
	rawBlock *usecase_model.RawBlock,
	blockResults *usecase_model.BlockResults,
) ([]entity_command.Command, error) {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Sprintf("panic when parsing block at height %d: %v", block.Height, r))
		}
	}()

	var err error
	var commands []entity_command.Command

	createRawBlockCommand := ParseCreateRawBlockCommand(rawBlock)
	commands = append(commands, createRawBlockCommand)

	createBlockCommand := ParseCreateBlockCommand(block)
	commands = append(commands, createBlockCommand)

	if len(blockResults.TxsResults) > 0 {
		transactionCommands, parseErr := ParseTransactionCommands(txDecoder, block, blockResults)
		if parseErr != nil {
			return nil, fmt.Errorf("error parsing transaction commands: %v", parseErr)
		}
		commands = append(commands, transactionCommands...)

		msgCommands, parseErr := ParseBlockResultsTxsMsgToCommands(txDecoder, block, blockResults)
		if parseErr != nil {
			return nil, fmt.Errorf("error parsing message commands: %v", parseErr)
		}
		commands = append(commands, msgCommands...)

		txsAccountTransferCommands, parseErr := ParseTxAccountTransferCommands(
			block.Height,
			blockResults.TxsResults,
		)
		if parseErr != nil {
			return nil, fmt.Errorf("error parsing block_results account transfer commands: %v", parseErr)
		}
		commands = append(commands, txsAccountTransferCommands...)
	}

	beginBlockEventsCommands, parseErr := ParseBeginBlockEventsCommands(block.Height, blockResults.BeginBlockEvents)
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing block_results_events commands: %v", parseErr)
	}
	commands = append(commands, beginBlockEventsCommands...)

	endBlockEventsCommands, parseErr := ParseEndBlockEventsCommands(block.Height, blockResults.BeginBlockEvents)
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing block_results_events commands: %v", parseErr)
	}
	commands = append(commands, endBlockEventsCommands...)

	validatorUpdatesCommands, parseErr := ParseValidatorUpdatesCommands(block.Height, blockResults.ValidatorUpdates)
	commands = append(commands, validatorUpdatesCommands...)
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing validator_updates commands: %v", parseErr)
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
