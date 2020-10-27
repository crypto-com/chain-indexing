package command

// Command interface provides a generic signature which any other
// type can implement
type CommandAdapter interface {
	Exec() error
	GetType() string
	GetVersion() uint64
}

// command struct defines the fields that a command needs
type Command struct {
	Type    string
	Version uint64
}

// NewCommand creates a new Command
func NewCommand(Type string, Version uint64) Command {
	return Command{
		Type,
		Version,
	}
}

// GetType return the type of a command
func (c Command) GetType() string {
	return c.Type
}

// GetVersion return the version of a command
func (c Command) GetVersion() uint64 {
	return c.Version
}
