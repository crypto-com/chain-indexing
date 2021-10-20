package projection

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/crypto-com/chain-indexing/projection/account"
	"github.com/crypto-com/chain-indexing/projection/account_message"
	"github.com/crypto-com/chain-indexing/projection/account_transaction"
	"github.com/crypto-com/chain-indexing/projection/block"
	"github.com/crypto-com/chain-indexing/projection/blockevent"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/bridge_pending_activity"
	"github.com/crypto-com/chain-indexing/projection/chainstats"
	"github.com/crypto-com/chain-indexing/projection/ibc_channel"
	"github.com/crypto-com/chain-indexing/projection/ibc_channel_message"
	"github.com/crypto-com/chain-indexing/projection/nft"
	"github.com/crypto-com/chain-indexing/projection/proposal"
	"github.com/crypto-com/chain-indexing/projection/transaction"
	"github.com/crypto-com/chain-indexing/projection/validator"
	"github.com/crypto-com/chain-indexing/projection/validatorstats"
)

func InitProjection(name string, params InitProjectionParams) projection_entity.Projection {
	switch name {
	case "Account":
		return account.NewAccount(params.Logger, params.RdbConn, params.CosmosAppClient)
	case "AccountTransaction":
		return account_transaction.NewAccountTransaction(params.Logger, params.RdbConn, params.AccountAddressPrefix)
	case "AccountMessage":
		return account_message.NewAccountMessage(params.Logger, params.RdbConn, params.AccountAddressPrefix)
	case "Block":
		return block.NewBlock(params.Logger, params.RdbConn)
	case "BlockEvent":
		return blockevent.NewBlockEvent(params.Logger, params.RdbConn)
	case "ChainStats":
		return chainstats.NewChainStats(params.Logger, params.RdbConn)
	case "Proposal":
		return proposal.NewProposal(params.Logger, params.RdbConn, params.ConsNodeAddressPrefix)
	case "Transaction":
		return transaction.NewTransaction(params.Logger, params.RdbConn)
	case "Validator":
		return validator.NewValidator(
			params.Logger, params.RdbConn, params.ConsNodeAddressPrefix,
		)
	case "ValidatorStats":
		return validatorstats.NewValidatorStats(params.Logger, params.RdbConn)
	case "NFT":
		return nft.NewNFT(params.Logger, params.RdbConn, nft.Config{
			EnableDrop:       false,
			DropDataAccessor: "",
		})
	case "CryptoComNFT":
		return nft.NewNFT(params.Logger, params.RdbConn, nft.Config{
			EnableDrop:       true,
			DropDataAccessor: "dropId",
		})
	case "IBCChannel":
		return ibc_channel.NewIBCChannel(params.Logger, params.RdbConn, ibc_channel.Config{
			EnableTxMsgTrace: false,
		})
	case "IBCChannelTxMsgTrace":
		return ibc_channel.NewIBCChannel(params.Logger, params.RdbConn, ibc_channel.Config{
			EnableTxMsgTrace: true,
		})
	case "IBCChannelMessage":
		return ibc_channel_message.NewIBCChannelMessage(params.Logger, params.RdbConn)
	case "BridgePendingActivity":
		return bridge_pending_activity.NewBridgePendingActivity(params.Logger, params.RdbConn)
	// register more projections here
	default:
		panic(fmt.Sprintf("Unrecognized projection: %s", name))
	}
}

type InitProjectionParams struct {
	Logger  applogger.Logger
	RdbConn rdb.Conn

	CosmosAppClient       cosmosapp.Client
	AccountAddressPrefix  string
	ConsNodeAddressPrefix string
}
