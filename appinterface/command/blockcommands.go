package command

import (
	"fmt"
	"github.com/crypto-com/chainindex/appinterface/command/types"
	"github.com/crypto-com/chainindex/appinterface/event"
	tenderminttypes "github.com/crypto-com/chainindex/appinterface/tendermint/types"
)

// CreateBlockCommand
type CreateBlockCommand struct {
	Command
	Payload *types.Block
}

// NewCreateBlockCommand create a CreateBlockCommand
func NewCreateBlockCommand(Payload *types.Block) CreateBlockCommand {
	cmd := NewCommand("CreateBlock", 1)
	return CreateBlockCommand{
		cmd,
		Payload,
	}
}

// Exec creates an event
func (c *CreateBlockCommand) Exec() (*event.Event, error) {
	var err error
	if c.Payload == nil {
		return nil, fmt.Errorf("cannot create event with empty command payload: %v", err)
	}

	// create new event
	evt := event.NewEvent("BlockCreated", c.Version, c.Payload)
	return &evt, nil
}

// CreateRawBlockCommand
type CreateRawBlockCommand struct {
	Command
	Payload *tenderminttypes.BlockResp
}

// NewCreateRawBlockCommand create a CreateRawBlockCommand
func NewCreateRawBlockCommand(Payload *tenderminttypes.BlockResp) CreateRawBlockCommand {
	cmd := NewCommand("CreateRawBlock", 1)
	return CreateRawBlockCommand{
		cmd,
		Payload,
	}
}

// Exec create an event
func (c *CreateRawBlockCommand) Exec() (*event.Event, error) {
	var err error
	if c.Payload == nil {
		return nil, fmt.Errorf("cannot create event with empty command payload: %v", err)
	}

	// create new event
	evt := event.NewEvent("RawBlockCreated", c.Version, c.Payload)
	return &evt, nil
}
