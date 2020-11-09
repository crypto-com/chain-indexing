package command

// Executor interface defines command executor signature
type Executor interface {
	Exec() error
	ExecAll(cmds []Command) error
}
