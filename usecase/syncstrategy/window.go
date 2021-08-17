package syncstrategy

import (
	"github.com/crypto-com/chain-indexing/entity/command"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"golang.org/x/sync/errgroup"
)

var _ Strategy = &Window{}

// Simple sync strategy by using a window of n concurrent block syncs at a time
type Window struct {
	logger applogger.Logger

	size int
}

func NewWindow(logger applogger.Logger, size int) *Window {
	return &Window{
		logger.WithFields(applogger.LogFields{
			"module": "WindowStrategy",
			"size":   size,
		}),

		size,
	}
}

func (window *Window) Sync(
	currentHeight int64,
	latestHeight int64,
	worker SyncBlockWorker,
) ([][]command.Command, SyncedHeight, error) {
	beginHeight := currentHeight
	var endHeight int64
	if latestHeight-currentHeight+1 < int64(window.size) {
		endHeight = latestHeight
	} else {
		endHeight = beginHeight + int64(window.size) - 1
	}

	logger := window.logger.WithFields(applogger.LogFields{
		"beginHeight": beginHeight,
		"endHeight":   endHeight,
	})
	logger.Debug("spawning goroutines for sync block workers")

	workersErrGroup := errgroup.Group{}

	commandWindow := newUnsafeCommandWindow(beginHeight, endHeight)

	for height := beginHeight; height <= endHeight; height += 1 {
		height := height
		workersErrGroup.Go(func() error {
			commands, err := worker(height)
			if err != nil {
				logger.Errorf("received error from sync block worker #%d: %v", height, err)
				return err
			}

			commandWindow.Put(height, commands)

			return nil
		})
	}

	if err := workersErrGroup.Wait(); err != nil {
		return nil, beginHeight - 1, err
	}

	return commandWindow.Export(), endHeight, nil
}

// An concurrency-unsafe command window. Never use it in multiple goroutines.
type unsafeCommandWindow struct {
	beginHeight int64

	blocksCommands [][]command.Command
}

func newUnsafeCommandWindow(beginHeight int64, endHeight int64) *unsafeCommandWindow {
	size := endHeight - beginHeight + 1
	blocksCommands := make([][]command.Command, size)
	for i := int64(0); i < size; i += 1 {
		blocksCommands[i] = make([]command.Command, 0)
	}
	return &unsafeCommandWindow{
		beginHeight,

		blocksCommands,
	}
}

func (window *unsafeCommandWindow) Put(blockHeight int64, commands []command.Command) {
	window.blocksCommands[blockHeight-window.beginHeight] = commands
}

func (window *unsafeCommandWindow) Export() [][]command.Command {
	return window.blocksCommands
}
