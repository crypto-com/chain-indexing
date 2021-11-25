package account

import (
	cosmosapp_interface "github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	appprojection "github.com/crypto-com/chain-indexing/projection"
)

func NewTestAccount(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	cosmosClient cosmosapp_interface.Client,
	config *appprojection.Config,
) *Account {
	return &Account{
		nil,
		rdbConn,
		logger,
		cosmosClient,
		config,
	}
}
