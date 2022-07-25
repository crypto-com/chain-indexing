package test

import (
	"github.com/crypto-com/chain-indexing/entity/command"
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
)

type commandA struct{}

func (c commandA) Name() string {
	return "commandA"
}
func (c commandA) Version() int {
	return 0
}
func (c commandA) Exec() (entity_event.Event, error) {
	return nil, nil
}

type commandB struct{}

func (c commandB) Name() string {
	return "commandB"
}
func (c commandB) Version() int {
	return 0
}
func (c commandB) Exec() (entity_event.Event, error) {
	return nil, nil
}

func ParserA(_ utils.CosmosParserParams) ([]command.Command, []string) {
	return []command.Command{
		commandA{},
	}, []string{}
}

func ParserB(_ utils.CosmosParserParams) ([]command.Command, []string) {
	return []command.Command{
		commandB{},
	}, []string{}
}
