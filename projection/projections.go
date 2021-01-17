package projection

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/projection/crossfire"

	"github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/crypto-com/chain-indexing/projection/account"
	"github.com/crypto-com/chain-indexing/projection/account_message"
	"github.com/crypto-com/chain-indexing/projection/account_transaction"
	"github.com/crypto-com/chain-indexing/projection/block"
	"github.com/crypto-com/chain-indexing/projection/blockevent"
	"github.com/crypto-com/chain-indexing/projection/transaction"
	"github.com/crypto-com/chain-indexing/projection/validator"
	"github.com/crypto-com/chain-indexing/projection/validatorstats"
)

func InitProjection(name string, params InitParams) projection_entity.Projection {
	switch name {
	case "Account":
		return account.NewAccount(params.Logger, params.RdbConn, params.CosmosAppClient)
	case "AccountTransaction":
		return account_transaction.NewAccountTransaction(params.Logger, params.RdbConn, params.AccountAddressPrefix)
	case "AccountMessage":
		return account_message.NewAccountMessage(params.Logger, params.RdbConn)
	case "Block":
		return block.NewBlock(params.Logger, params.RdbConn)
	case "BlockEvent":
		return blockevent.NewBlockEvent(params.Logger, params.RdbConn)
	case "Transaction":
		return transaction.NewTransaction(params.Logger, params.RdbConn)
	case "Validator":
		return validator.NewValidator(
			params.Logger, params.RdbConn, params.ConsNodeAddressPrefix,
		)
	case "ValidatorStats":
		return validatorstats.NewValidatorStats(params.Logger, params.RdbConn)
	case "Crossfire":
		return crossfire.NewCrossfire(
			params.Logger,
			params.RdbConn,
			params.ConsNodeAddressPrefix,
			params.ValidatorAddressPrefix,
			params.AccountAddressPrefix,
			params.PhaseOneStartTime,
			params.PhaseTwoStartTime,
			params.PhaseThreeStartTime,
			params.CompetitionEndTime,
			params.AdminAddress,
			params.NetworkUpgradeProposalID,
			params.ParticipantsListURL,
			params.JackpotOneStartTime,
			params.JackpotTwoStartTime,
			params.JackpotThreeStartTime,
			params.JackpotFourStartTime,
			params.JackpotFourEndTime,
		)
	// register more projections here
	default:
		panic(fmt.Sprintf("Unrecognized projection: %s", name))
	}
}

type InitParams struct {
	Logger  applogger.Logger
	RdbConn rdb.Conn

	CosmosAppClient        cosmosapp.Client
	ConsNodeAddressPrefix  string
	ValidatorAddressPrefix string
	AccountAddressPrefix   string

	// Crossfire related fields
	JackpotOneStartTime      int64
	JackpotTwoStartTime      int64
	JackpotThreeStartTime    int64
	JackpotFourStartTime     int64
	JackpotFourEndTime       int64
	PhaseOneStartTime        int64
	PhaseTwoStartTime        int64
	PhaseThreeStartTime      int64
	CompetitionEndTime       int64
	AdminAddress             string
	NetworkUpgradeProposalID string
	ParticipantsListURL      string
}
