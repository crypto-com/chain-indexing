package bridge_activity_matcher

import (
	"errors"
	"fmt"
	"time"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/types"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/view"
	projection_usecase "github.com/crypto-com/chain-indexing/usecase/projection"
	"github.com/mitchellh/mapstructure"
)

var (
	NewBridgePendingActivitiesView = view.NewBridgePendingActivitiesView
	NewBridgeActivitiesView        = view.NewBridgeActivitiesView
	NewPgxConnPool                 = func(config *pg.PgxConnPoolConfig, logger applogger.Logger) (rdb.Conn, error) {
		return pg.NewPgxConnPool(config, logger)
	}
)
var _ projection_entity.CronJob = &BridgeActivityMatcher{}

type Config struct {
	Interval           time.Duration             `mapstructure:"interval"`
	CounterpartyChains []CounterpartyChainConfig `mapstructure:"counterparty_chains"`
}

type CounterpartyChainConfig struct {
	Name     string                     `mapstructure:"name"`
	Database CounterpartyDatabaseConfig `mapstructure:"database"`
}

type CounterpartyDatabaseConfig struct {
	SSL                 bool          `mapstructure:"ssl"`
	Host                string        `mapstructure:"host"`
	Port                int32         `mapstructure:"port"`
	Username            string        `mapstructure:"username"`
	Password            string        `mapstructure:"password"`
	Name                string        `mapstructure:"name"`
	Schema              string        `mapstructure:"schema"`
	MaxConns            int32         `mapstructure:"pool_max_conns"`
	MinConns            int32         `mapstructure:"pool_min_conns"`
	MaxConnLifeTime     time.Duration `mapstructure:"pool_max_conn_lifetime"`
	MaxConnIdleTime     time.Duration `mapstructure:"pool_max_conn_idle_time"`
	HealthCheckInterval time.Duration `mapstructure:"pool_health_check_interval"`
}

type counterpartyChainRDbConfig struct {
	Name string
	Conn rdb.Conn
}

func ConfigFromInterface(data interface{}) (Config, error) {
	config := Config{}

	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
		),
		Result: &config,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		return config, fmt.Errorf("error creating projection config decoder: %v", decoderErr)
	}

	if err := decoder.Decode(data); err != nil {
		return config, fmt.Errorf("error decoding projection BridgePendingActivity config: %v", err)
	}

	return config, nil
}

type BridgeActivityMatcher struct {
	projection_usecase.Base

	config Config

	thisRDbConn            rdb.Conn
	counterpartyRDbConfigs []counterpartyChainRDbConfig
	logger                 applogger.Logger

	migrationHelper migrationhelper.MigrationHelper
}

const (
	MIGRATION_DIRECOTRY = "projection/bridge_activity/bridge_activity_matcher/migrations"
)

func New(
	config Config,
	logger applogger.Logger,
	rdbConn rdb.Conn,
	migrationHelper migrationhelper.MigrationHelper,
) *BridgeActivityMatcher {
	return &BridgeActivityMatcher{
		Base: projection_usecase.NewBase("BridgeActivityMatcher"),

		config: config,

		thisRDbConn:            rdbConn,
		counterpartyRDbConfigs: make([]counterpartyChainRDbConfig, 0),
		logger: logger.WithFields(applogger.LogFields{
			"module": "BridgeActivityMatcher",
		}),

		migrationHelper: migrationHelper,
	}
}

func (cronJob *BridgeActivityMatcher) Id() string {
	return "BridgeActivityMatcher"
}

func (cronJob *BridgeActivityMatcher) Config() *Config {
	return &cronJob.config
}

func (cronJob *BridgeActivityMatcher) OnInit() error {
	config := cronJob.Config()

	// TODO: Refactor to generic rdbConn creator
	for _, counterpartyChain := range config.CounterpartyChains {
		rdbConn, rdbConnErr := NewPgxConnPool(&pg.PgxConnPoolConfig{
			ConnConfig: pg.ConnConfig{
				Host:          counterpartyChain.Database.Host,
				Port:          counterpartyChain.Database.Port,
				MaybeUsername: &counterpartyChain.Database.Username,
				MaybePassword: &counterpartyChain.Database.Password,
				Database:      counterpartyChain.Database.Name,
				SSL:           counterpartyChain.Database.SSL,
			},
			MaybeMaxConns:          &counterpartyChain.Database.MaxConns,
			MaybeMinConns:          &counterpartyChain.Database.MinConns,
			MaybeMaxConnLifeTime:   &counterpartyChain.Database.MaxConnLifeTime,
			MaybeMaxConnIdleTime:   &counterpartyChain.Database.MaxConnIdleTime,
			MaybeHealthCheckPeriod: &counterpartyChain.Database.HealthCheckInterval,
		}, cronJob.logger)
		if rdbConnErr != nil {
			return fmt.Errorf(
				"error when initializing %s indexing DB connection: %w",
				counterpartyChain.Name, rdbConnErr,
			)
		}

		cronJob.counterpartyRDbConfigs = append(cronJob.counterpartyRDbConfigs, counterpartyChainRDbConfig{
			Name: counterpartyChain.Name,
			Conn: rdbConn,
		})
	}

	if cronJob.migrationHelper != nil {
		cronJob.migrationHelper.Migrate()
	}

	return nil
}

func (cronJob *BridgeActivityMatcher) Interval() time.Duration {
	return cronJob.Config().Interval
}

func (cronJob *BridgeActivityMatcher) Exec() error {
	thisBridgePendingActivities := NewBridgePendingActivitiesView(cronJob.thisRDbConn.ToHandle())

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

	for _, counterpartyRDbConfig := range cronJob.counterpartyRDbConfigs {
		counterpartyBridgePendingActivities := NewBridgePendingActivitiesView(counterpartyRDbConfig.Conn.ToHandle())

		counterpartyAllUnprocessedOutgoing, counterpartyAllUnprocessedOutgoingErr := counterpartyBridgePendingActivities.
			ListAllUnprocessedOutgoing()
		if counterpartyAllUnprocessedOutgoingErr != nil {
			return fmt.Errorf(
				"error querying %s unprocessed outgoing pending activities of this projection: %v",
				counterpartyRDbConfig.Name, thisAllUnprocessedOutgoingErr,
			)
		}

		if handleThisUnprocessedOutgoingErr := cronJob.HandleOutgoing(
			counterpartyAllUnprocessedOutgoing,
			counterpartyBridgePendingActivities,
		); handleThisUnprocessedOutgoingErr != nil {
			return fmt.Errorf(
				"error handling %s unprocessed outgoing bridge pending activities: %v",
				counterpartyRDbConfig.Name, handleThisUnprocessedOutgoingErr,
			)
		}

		counterpartyAllUnprocessedIncoming, counterpartyAllUnprocessedIncomingErr := counterpartyBridgePendingActivities.
			ListAllUnprocessedIncoming()
		if counterpartyAllUnprocessedIncomingErr != nil {
			return fmt.Errorf(
				"error querying unprocessed %s incoming pending activities of projection: %v",
				counterpartyRDbConfig.Name, thisAllUnprocessedOutgoingErr,
			)
		}

		if handleCounterpartyUnprocessedIncomingErr := cronJob.HandleIncoming(
			counterpartyAllUnprocessedIncoming,
			counterpartyBridgePendingActivities,
		); handleCounterpartyUnprocessedIncomingErr != nil {
			return fmt.Errorf(
				"error handling %s unprocessed incoming bridge pending activities: %v",
				counterpartyRDbConfig.Name, handleCounterpartyUnprocessedIncomingErr,
			)
		}
	}

	return nil
}

func (cronJob *BridgeActivityMatcher) HandleOutgoing(
	rows []view.BridgePendingActivityReadRow,
	bridgePendingActivities view.BridgePendingActivities,
) error {
	if len(rows) == 0 {
		return nil
	}

	for _, row := range rows {
		if err := cronJob.handleOutgoingRow(row, bridgePendingActivities); err != nil {
			return err
		}
	}

	return nil
}

func (cronJob *BridgeActivityMatcher) handleOutgoingRow(
	row view.BridgePendingActivityReadRow,
	bridgePendingActivities view.BridgePendingActivities,
) error {
	logger := cronJob.logger.WithFields(applogger.LogFields{
		"rowType": types.DIRECTION_OUTGOING,
		"rowId":   row.Id,
	})

	logger.Info("going to handle outgoing row")

	thisRDbTx, thisRDbConnBeginErr := cronJob.thisRDbConn.Begin()
	if thisRDbConnBeginErr != nil {
		return fmt.Errorf("error beginning transaction: %v", thisRDbConnBeginErr)
	}

	committed := false
	defer func() {
		if !committed {
			_ = thisRDbTx.Rollback()
		}
	}()

	thisRDbTxHandle := thisRDbTx.ToHandle()

	bridgeActivities := NewBridgeActivitiesView(thisRDbTxHandle)

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

	if err := thisRDbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true

	return nil
}

func (cronJob *BridgeActivityMatcher) HandleIncoming(
	rows []view.BridgePendingActivityReadRow,
	bridgePendingActivities view.BridgePendingActivities,
) error {
	if len(rows) == 0 {
		return nil
	}

	for _, row := range rows {
		if err := cronJob.handleIncomingRow(row, bridgePendingActivities); err != nil {
			return err
		}
	}

	return nil
}

func (cronJob *BridgeActivityMatcher) handleIncomingRow(
	row view.BridgePendingActivityReadRow,
	bridgePendingActivities view.BridgePendingActivities,
) error {
	logger := cronJob.logger.WithFields(applogger.LogFields{
		"rowType": types.DIRECTION_INCOMING,
		"rowId":   row.Id,
	})

	logger.Info("going to handle incoming row")

	thisRDbTx, thisRDbConnBeginErr := cronJob.thisRDbConn.Begin()
	if thisRDbConnBeginErr != nil {
		return fmt.Errorf("error beginning transaction: %v", thisRDbConnBeginErr)
	}
	committed := false
	defer func() {
		if !committed {
			_ = thisRDbTx.Rollback()
		}
	}()

	thisRDbTxHandle := thisRDbTx.ToHandle()

	bridgeActivities := NewBridgeActivitiesView(thisRDbTxHandle)

	commit := func() error {
		if err := thisRDbTx.Commit(); err != nil {
			return fmt.Errorf("error committing changes: %v", err)
		}
		committed = true

		return nil
	}

	mutExistingActivity, existingActivityErr := bridgeActivities.FindByLinkId(row.LinkId)
	if existingActivityErr != nil {
		if errors.Is(existingActivityErr, rdb.ErrNoRows) {
			logger.Infof(
				"error querying existing activity by link id %s: not found, will try again later.",
				row.LinkId,
			)
			return nil
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

		return nil
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
	committed = true

	logger.Info("successfully handled incoming row")
	return nil
}
