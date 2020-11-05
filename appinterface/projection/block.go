package projection

// TODO
// import (
// 	"github.com/crypto-com/chainindex/appinterface/projection/view"
// 	dddevent "github.com/crypto-com/chainindex/ddd/event"
// 	"github.com/crypto-com/chainindex/entity/event"
// )

// // TODO: Listen to council node related events and project council node
// type Block struct {
// 	blocksView *view.Blocks
// }

// func NewBlock(blocksView *view.Blocks) *Block {
// 	return &Block{
// 		blocksView,
// 	}
// }

// func (_ *Block) Id() string {
// 	return "Block"
// }

// func (_ *Block) GetEventsToListen() []string {
// 	return []string{event.BLOCK_CREATED}
// }

// func (projection *Block) OnInit() error {
// 	// TODO
// 	return nil
// }

// func (projection *Block) HandleEvent(evt dddevent.Event) error {
// 	if blockCreatedEvt, ok := evt.(*event.BlockCreated); ok {
// 		committedCouncilNodes := make([]view.BlockCommittedCouncilNode, 0)
// 		for _, signature := range blockCreatedEvt.Block.Signatures {
// 			committedCouncilNodes = append(committedCouncilNodes, view.BlockCommittedCouncilNode{
// 				Address:    signature.ValidatorAddress,
// 				Time:       signature.Timestamp,
// 				Signature:  signature.Signature,
// 				IsProposer: blockCreatedEvt.Block.ProposerAddress == signature.ValidatorAddress,
// 			})
// 		}

// 		if err := projection.blocksView.Insert(&view.Block{
// 			Height:                blockCreatedEvt.Block.Height,
// 			Hash:                  blockCreatedEvt.Block.Hash,
// 			Time:                  blockCreatedEvt.Block.Time,
// 			AppHash:               blockCreatedEvt.Block.AppHash,
// 			TransactionCount:      len(blockCreatedEvt.Block.Txs),
// 			CommittedCouncilNodes: committedCouncilNodes,
// 		}); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
