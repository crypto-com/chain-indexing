package projection

import (
	"fmt"

	"github.com/crypto-com/chainindex/appinterface/projection/rdbbase"
	"github.com/crypto-com/chainindex/appinterface/projection/view"
	"github.com/crypto-com/chainindex/appinterface/rdb"
	entity_event "github.com/crypto-com/chainindex/entity/event"
	applogger "github.com/crypto-com/chainindex/internal/logger"
	usecase_event "github.com/crypto-com/chainindex/usecase/event"
)

// TODO: Listen to council node related events and project council node
type Block struct {
	*rdbbase.RDbBase

	logger applogger.Logger

	blocksView *view.Blocks
}

func NewBlock(logger applogger.Logger, rdbHandle *rdb.Handle, blocksView *view.Blocks) *Block {
	return &Block{
		rdbbase.NewRDbBase(rdbHandle, "Block"),
		logger,

		blocksView,
	}
}

func (_ *Block) GetEventsToListen() []string {
	return []string{usecase_event.BLOCK_CREATED}
}

func (projection *Block) OnInit() error {
	// TODO
	return nil
}

func (projection *Block) HandleEvents(evts []entity_event.Event) error {
	for _, evt := range evts {
		if blockCreatedEvt, ok := evt.(*usecase_event.BlockCreated); ok {
			return projection.handleBlockCreatedEvent(blockCreatedEvt)
		} else {
			return fmt.Errorf("received unexpecxted event %sV%d(%s)", evt.Name(), evt.Version(), evt.Id())
		}
	}
	return nil
}

func (projection *Block) handleBlockCreatedEvent(evt *usecase_event.BlockCreated) error {
	committedCouncilNodes := make([]view.BlockCommittedCouncilNode, 0)
	for _, signature := range evt.Block.Signatures {
		committedCouncilNodes = append(committedCouncilNodes, view.BlockCommittedCouncilNode{
			Address:    signature.ValidatorAddress,
			Time:       signature.Timestamp,
			Signature:  signature.Signature,
			IsProposer: evt.Block.ProposerAddress == signature.ValidatorAddress,
		})
	}

	if err := projection.blocksView.Insert(&view.Block{
		Height:                evt.Block.Height,
		Hash:                  evt.Block.Hash,
		Time:                  evt.Block.Time,
		AppHash:               evt.Block.AppHash,
		TransactionCount:      len(evt.Block.Txs),
		CommittedCouncilNodes: committedCouncilNodes,
	}); err != nil {
		return err
	}

	return nil
}
