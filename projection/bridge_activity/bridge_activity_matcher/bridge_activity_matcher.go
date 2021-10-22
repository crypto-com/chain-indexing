package bridge_activity_matcher

import (
	"errors"
	"fmt"
	"os"
	"time"

	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/types"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/view"
	projection_usecase "github.com/crypto-com/chain-indexing/usecase/projection"
)

var _ projection_entity.CronJob = &BridgeActivityMatcher{}

type Config struct {
	Interval               time.Duration `mapstructure:"interval"`
	CryptoOrgChainDatabase struct {
		SSL      bool   `mapstructure:"ssl"`
		Host     string `mapstructure:"host"`
		Port     int32  `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
		Schema   string `mapstructure:"schema"`
	} `mapstructure:"crypto_org_chain_database"`
}

type BridgeActivityMatcher struct {
	projection_usecase.Base

	thisRDbConn           rdb.Conn
	cryptoOrgChainRDbConn rdb.Conn
	logger                applogger.Logger
}

func New(logger applogger.Logger, rdbConn rdb.Conn) *BridgeActivityMatcher {
	var config Config
	return &BridgeActivityMatcher{
		Base: projection_usecase.NewBaseWithOptions("BridgeActivityMatcher", projection_usecase.Options{
			MaybeConfigPtr: &config,
		}),

		thisRDbConn:           rdbConn,
		cryptoOrgChainRDbConn: nil,
		logger: logger.WithFields(applogger.LogFields{
			"module": "BridgeActivityMatcher",
		}),
	}
}

func (cronJob *BridgeActivityMatcher) Id() string {
	return "BridgeActivityMatcher"
}

func (cronJob *BridgeActivityMatcher) Config() *Config {
	return cronJob.Base.Config().(*Config)
}

func (cronJob *BridgeActivityMatcher) OnInit() error {
	config := cronJob.Config()

	cryptoOrgChainDBPassword := os.Getenv("PROJECTION_BRIDGE_ACTIVITY_MATCHER_CRYPTO_ORG_CHAIN_DB_PASSWORD")
	if cryptoOrgChainDBPassword == "" {
		cryptoOrgChainDBPassword = config.CryptoOrgChainDatabase.Password
	}

	// TODO: Refactor to generic rdbConn creator
	var err error
	if cronJob.cryptoOrgChainRDbConn, err = pg.NewPgxConn(&pg.ConnConfig{
		Host:          config.CryptoOrgChainDatabase.Host,
		Port:          config.CryptoOrgChainDatabase.Port,
		MaybeUsername: &config.CryptoOrgChainDatabase.Username,
		MaybePassword: &cryptoOrgChainDBPassword,
		Database:      config.CryptoOrgChainDatabase.Name,
		SSL:           config.CryptoOrgChainDatabase.SSL,
	}, cronJob.logger); err != nil {
		return fmt.Errorf("error when initializing Crypto.org Chain indexing DB connection: %v", err)
	}

	return nil
}

func (cronJob *BridgeActivityMatcher) Interval() time.Duration {
	return cronJob.Config().Interval
}

func (cronJob *BridgeActivityMatcher) Exec() error {
	thisBridgePendingActivities := view.NewBridgePendingActivities(cronJob.thisRDbConn.ToHandle())
	cryptoOrgChainBridgePendingActivities := view.NewBridgePendingActivities(cronJob.cryptoOrgChainRDbConn.ToHandle())

	thisAllUnprocessedOutgoing, thisAllUnprocessedOutgoingErr := thisBridgePendingActivities.ListAllUnprocessedOutgoing()
	if thisAllUnprocessedOutgoingErr != nil {
		return fmt.Errorf(
			"error querying unprocessed outoing pending activities of this projection: %v",
			thisAllUnprocessedOutgoingErr,
		)
	}

	if handleThisUnprocessedOutgoingErr := cronJob.HandleOutgoing(
		thisAllUnprocessedOutgoing,
		thisBridgePendingActivities,
	); handleThisUnprocessedOutgoingErr != nil {
		return fmt.Errorf(
			"error handling this unprocessed outgoing bridge pending activities: %v",
			handleThisUnprocessedOutgoingErr,
		)
	}

	cryptoOrgChainAllUnprocessedOutgoing, cryptoOrgChainAllUnprocessedOutgoingErr := cryptoOrgChainBridgePendingActivities.
		ListAllUnprocessedOutgoing()
	if cryptoOrgChainAllUnprocessedOutgoingErr != nil {
		return fmt.Errorf(
			"error querying Crypto.org Chain unprocessed outgoing pending activities of this projection: %v",
			thisAllUnprocessedOutgoingErr,
		)
	}

	if handleThisUnprocessedOutgoingErr := cronJob.HandleOutgoing(
		cryptoOrgChainAllUnprocessedOutgoing,
		cryptoOrgChainBridgePendingActivities,
	); handleThisUnprocessedOutgoingErr != nil {
		return fmt.Errorf(
			"error handling Crypto.org Chain unprocessed outgoing bridge pending activities: %v",
			handleThisUnprocessedOutgoingErr,
		)
	}

	thisAllUnprocessedIncoming, thisAllUnprocessedIncomingErr := thisBridgePendingActivities.ListAllUnprocessedIncoming()
	if thisAllUnprocessedIncomingErr != nil {
		return fmt.Errorf(
			"error querying unprocessed incoming pending activities of this projection: %v",
			thisAllUnprocessedOutgoingErr,
		)
	}

	if handleThisUnprocessedIncomingErr := cronJob.HandleIncoming(
		thisAllUnprocessedIncoming,
		thisBridgePendingActivities,
	); handleThisUnprocessedIncomingErr != nil {
		return fmt.Errorf(
			"error handling this unprocessed incoming bridge pending activities: %v",
			handleThisUnprocessedIncomingErr,
		)
	}

	cryptoOrgChainAllUnprocessedIncoming, cryptoOrgChainAllUnprocessedIncomingErr := cryptoOrgChainBridgePendingActivities.
		ListAllUnprocessedIncoming()
	if cryptoOrgChainAllUnprocessedIncomingErr != nil {
		return fmt.Errorf(
			"error querying unprocessed Crypto.org Chain incoming pending activities of projection: %v",
			thisAllUnprocessedOutgoingErr,
		)
	}

	if handleCryptoOrgChainUnprocessedIncomingErr := cronJob.HandleIncoming(
		cryptoOrgChainAllUnprocessedIncoming,
		cryptoOrgChainBridgePendingActivities,
	); handleCryptoOrgChainUnprocessedIncomingErr != nil {
		return fmt.Errorf(
			"error handling Crypto.org Chain unprocessed incoming bridge pending activities: %v",
			handleCryptoOrgChainUnprocessedIncomingErr,
		)
	}

	return nil
}

func (cronJob *BridgeActivityMatcher) HandleOutgoing(
	rows []view.BridgePendingActivityReadRow,
	bridgePendingActivities *view.BridgePendingActivities,
) error {
	var currentThisRDbTx rdb.Tx
	var currentCommitted bool
	defer func() {
		if !currentCommitted && currentThisRDbTx != nil {
			_ = currentThisRDbTx.Rollback()
		}
	}()
	for _, row := range rows {
		var rdbConnBeginErr error

		currentCommitted = false
		currentThisRDbTx = nil
		if currentThisRDbTx, rdbConnBeginErr = cronJob.thisRDbConn.Begin(); rdbConnBeginErr != nil {
			return fmt.Errorf("error beginning transaction: %v", rdbConnBeginErr)
		}

		thisRDbTxHandle := currentThisRDbTx.ToHandle()

		bridgeActivities := view.NewBridgeActivities(thisRDbTxHandle)

		if insertErr := bridgeActivities.Insert(&view.BridgeActivityInsertRow{
			BridgeType:                           row.BridgeType,
			SourceBlockHeight:                    row.BlockHeight,
			SourceBlockTime:                      row.BlockTime,
			SourceTransactionId:                  *row.MaybeTransactionId,
			SourceChain:                          row.FromChainId,
			SourceAddress:                        *row.MaybeFromAddress,
			MaybeSourceSmartContractAddress:      row.MaybeFromSmartContractAddress,
			MaybeDestinationBlockHeight:          nil,
			MaybeDestinationBlockTime:            nil,
			MaybeDestinationTransactionId:        nil,
			DestinationChain:                     row.ToChainId,
			DestinationAddress:                   row.ToAddress,
			MaybeDestinationSmartContractAddress: row.MaybeToSmartContractAddress,
			MaybeChannelId:                       row.MaybeChannelId,
			LinkId:                               row.LinkId,
			Amount:                               row.Amount,
			MaybeDenom:                           row.MaybeDenom,
			MaybeBridgeFeeAmount:                 row.MaybeBridgeFeeAmount,
			MaybeBridgeFeeDenom:                  row.MaybeBridgeFeeDenom,
			Status:                               row.Status,
		}); insertErr != nil {
			return fmt.Errorf("error inserting outgoing activity into bridge activity table: %v", insertErr)
		}

		if updateToProcessedErr := bridgePendingActivities.UpdateToProcessed(row.Id); updateToProcessedErr != nil {
			return fmt.Errorf(
				"error updating pending activity row to processed into briding pending activity table: %v",
				updateToProcessedErr,
			)
		}

		if err := currentThisRDbTx.Commit(); err != nil {
			return fmt.Errorf("error committing changes: %v", err)
		}
		currentCommitted = true
	}

	return nil
}

func (cronJob *BridgeActivityMatcher) HandleIncoming(
	rows []view.BridgePendingActivityReadRow,
	bridgePendingActivities *view.BridgePendingActivities,
) error {
	logger := cronJob.logger.WithFields(applogger.LogFields{
		"source": bridgePendingActivities,
	})

	var currentThisRDbTx rdb.Tx
	var currentCommitted bool
	defer func() {
		if !currentCommitted && currentThisRDbTx != nil {
			_ = currentThisRDbTx.Rollback()
		}
	}()
	for _, row := range rows {
		logger.WithFields(applogger.LogFields{
			"rowId": row.Id,
		})
		logger.Info("going to handle incoming row")

		var rdbConnBeginErr error

		currentCommitted = false
		currentThisRDbTx = nil
		if currentThisRDbTx, rdbConnBeginErr = cronJob.thisRDbConn.Begin(); rdbConnBeginErr != nil {
			return fmt.Errorf("error beginning transaction: %v", rdbConnBeginErr)
		}

		thisRDbTxHandle := currentThisRDbTx.ToHandle()

		bridgeActivities := view.NewBridgeActivities(thisRDbTxHandle)

		commit := func() error {
			if err := currentThisRDbTx.Commit(); err != nil {
				return fmt.Errorf("error committing changes: %v", err)
			}
			currentCommitted = true

			return nil
		}

		mutExistingActivity, existingActivityErr := bridgeActivities.FindByLinkId(row.LinkId)
		if existingActivityErr != nil {
			if errors.Is(existingActivityErr, rdb.ErrNoRows) {
				logger.Infof(
					"error querying existing activity by link id %s: not found, will try again later.",
					row.LinkId,
				)
				continue
			}
			return fmt.Errorf("error querying existing activity by link id %s: %v", row.LinkId, existingActivityErr)
		}

		isCounterpartyPendingIncomingActivity := row.Direction != types.DIRECTION_RESPONSE
		isPendingActivityAResponse := !isCounterpartyPendingIncomingActivity
		// Destination block time is only updated after processed a pending activity at counterparty
		isActivityAlreadyProcessedByCounterparty := mutExistingActivity.MaybeDestinationBlockTime != nil

		if isActivityAlreadyProcessedByCounterparty && isPendingActivityAResponse {
			logger.Info("row already processed with counterparty activity")
			if updateToProcessedErr := bridgePendingActivities.UpdateToProcessed(row.Id); updateToProcessedErr != nil {
				return fmt.Errorf(
					"error updating pending activity row to processed into briding pending activity table: %v",
					updateToProcessedErr,
				)
			}
			if commitErr := commit(); commitErr != nil {
				return commitErr
			}

			continue
		}

		if isCounterpartyPendingIncomingActivity {
			mutExistingActivity.MaybeDestinationBlockTime = row.BlockTime
			mutExistingActivity.MaybeDestinationBlockHeight = primptr.Int64(row.BlockHeight)
			mutExistingActivity.MaybeDestinationTransactionId = row.MaybeTransactionId
			mutExistingActivity.MaybeDestinationSmartContractAddress = row.MaybeToSmartContractAddress
			mutExistingActivity.Status = row.Status
		} else {
			mutExistingActivity.Status = row.Status
		}
		if updateErr := bridgeActivities.Update(&mutExistingActivity); updateErr != nil {
			return fmt.Errorf("error updating activity into bridge activity table: %v", updateErr)
		}

		if updateToProcessedErr := bridgePendingActivities.UpdateToProcessed(row.Id); updateToProcessedErr != nil {
			return fmt.Errorf(
				"error updating pending activity row to processed into briding pending activity table: %v",
				updateToProcessedErr,
			)
		}

		if commitErr := commit(); commitErr != nil {
			return commitErr
		}

		logger.Info("successfully handled incoming row")
	}

	return nil
}
