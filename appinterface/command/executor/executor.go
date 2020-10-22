package executor

// Executor interface defines command executor signature
type Executor interface {
	Exec() error
	ExecAll() error
}

// Exec process the command to event and save
//func Exec(c command.Command) error {}

