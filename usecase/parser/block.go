package parser

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"

	entity_command "github.com/crypto-com/chain-indexing/entity/command"
	usecase_model "github.com/crypto-com/chain-indexing/usecase/model"
)

func ParseBlockToCommands(
	parserManager *utils.CosmosParserManager,
	txDecoder *utils.TxDecoder,
	block *usecase_model.Block,
	rawBlock *usecase_model.RawBlock,
	blockResults *usecase_model.BlockResults,
	accountAddressPrefix string,
	bondingDenom string,
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
		msgCommands, possibleSignerAddress, parseErr := ParseBlockTxsMsgToCommands(
			parserManager,
			txDecoder,
			block,
			blockResults,
			accountAddressPrefix,
			bondingDenom,
		)
		if parseErr != nil {
			return nil, fmt.Errorf("error parsing message commands: %v", parseErr)
		}

		transactionCommands, parseErr := ParseTransactionCommands(txDecoder, block, blockResults, accountAddressPrefix, possibleSignerAddress)
		if parseErr != nil {
			return nil, fmt.Errorf("error parsing transaction commands: %v", parseErr)
		}
		commands = append(commands, transactionCommands...)
		commands = append(commands, msgCommands...)

		txsAccountTransferCommands, parseErr := ParseTxAccountTransferCommands(
			block.Height,
			blockResults.TxsResults,
		)
		if parseErr != nil {
			return nil, fmt.Errorf("error parsing block_results account transfer commands: %v", parseErr)
		}
		commands = append(commands, txsAccountTransferCommands...)

		txsResultsCommand, parseTxsResultCommandErr := ParseBlockResultsTxsResults(block, blockResults)
		if parseTxsResultCommandErr != nil {
			return nil, fmt.Errorf("error parsing block_results txs_results commands: %v", parseErr)
		}
		commands = append(commands, txsResultsCommand...)
	}

	beginBlockEventsCommands, parseErr := ParseBeginBlockEventsCommands(
		block.Height,
		blockResults.BeginBlockEvents,
		bondingDenom,
	)
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing begin_block_events commands: %v", parseErr)
	}
	commands = append(commands, beginBlockEventsCommands...)

	endBlockEventsCommands, parseErr := ParseEndBlockEventsCommands(block.Height, blockResults.EndBlockEvents)
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing end_block_events commands: %v", parseErr)
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
