package command

import "github.com/crypto-com/chainindex/ddd"

type Command interface {
	Name() string
	Version() int

	Exec() []ddd.Event
}
