package projection

import (
	"github.com/crypto-com/chainindex/appinterface/projection/view"
	"github.com/crypto-com/chainindex/ddd"
	"github.com/crypto-com/chainindex/entity/event"
)

// TODO: Listen to council node related events and project council node
type BlockProjection struct {
	blocksView view.Blocks
}

func (_ *BlockProjection) Id() string {
	return "Block"
}

func (_ *BlockProjection) GetEventsToListen() []string {
	return []string{event.BLOCK_CREATED}
}

func (projection *BlockProjection) OnInit() error {
	// TODO
	return nil
}

func (projection *BlockProjection) HandleEvent(evt ddd.Event) error {
	if blockCreatedEvt, ok := evt.(*event.BlockCreated); ok {
		committedCouncilNodes := make([]view.BlockCommittedCouncilNode, 0)
		for _, signature := range blockCreatedEvt.Block.Signatures {
			committedCouncilNodes = append(committedCouncilNodes, view.BlockCommittedCouncilNode{
				Address:    signature.ValidatorAddress,
				Time:       signature.Timestamp,
				Signature:  signature.Signature,
				IsProposer: blockCreatedEvt.Block.ProposerAddress == signature.ValidatorAddress,
			})
		}

		if err := projection.blocksView.Insert(&view.Block{
			Height:                blockCreatedEvt.Block.Height,
			Hash:                  blockCreatedEvt.Block.Hash,
			Time:                  blockCreatedEvt.Block.Time,
			AppHash:               blockCreatedEvt.Block.AppHash,
			TransactionCount:      len(blockCreatedEvt.Block.Txs),
			CommittedCouncilNodes: committedCouncilNodes,
		}); err != nil {
			return err
		}
	}
	return nil
}
