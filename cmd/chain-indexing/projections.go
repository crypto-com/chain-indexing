package main

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/account_message"
	"github.com/crypto-com/chain-indexing/appinterface/projection/block"
	"github.com/crypto-com/chain-indexing/appinterface/projection/blockevent"
	"github.com/crypto-com/chain-indexing/appinterface/projection/crossfire"
	"github.com/crypto-com/chain-indexing/appinterface/projection/transaction"
	"github.com/crypto-com/chain-indexing/appinterface/projection/validator"
	"github.com/crypto-com/chain-indexing/appinterface/projection/validatorstats"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
)

func initProjections(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	config *Config,
) []projection_entity.Projection {
	return []projection_entity.Projection{
		block.NewBlock(logger, rdbConn),
		transaction.NewTransaction(logger, rdbConn),
		blockevent.NewBlockEvent(logger, rdbConn),
		validator.NewValidator(
			logger, rdbConn, config.Blockchain.ConNodeAddressPrefix,
		),
		validatorstats.NewValidatorStats(logger, rdbConn),
		account_message.NewAccountMessage(logger, rdbConn),

		// NOTICE: crossfire dry-run projection is only for main-net competition
		// the logic and view tables could be removed after the competition is ended.
		crossfire.NewCrossfire(
			logger,
			rdbConn,
			config.Blockchain.ConNodeAddressPrefix,
			config.Blockchain.ValidatorAddressPrefix,
			config.Crossfire.PhaseOneStartTime,
			config.Crossfire.PhaseTwoStartTime,
			config.Crossfire.PhaseThreeStartTime,
			config.Crossfire.CompetitionEndTime,
			config.Crossfire.AdminAddress,
			config.Crossfire.NetworkUpgradeProposalID,
			config.Crossfire.ParticipantsListURL,
			config.Blockchain.AccountAddressPrefix,
		),

		// register more projections here
	}
}
