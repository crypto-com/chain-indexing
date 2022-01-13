package chainstats

import (
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	entity_projection "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	"github.com/crypto-com/chain-indexing/projection/chainstats/view"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ entity_projection.Projection = &ChainStats{}

const GENESIS_BLOCK_TIME = "genesis_block_time"
const TOTAL_BLOCK_TIME = "total_block_time"
const TOTAL_BLOCK_COUNT = "total_block_count"

type ChainStats struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger

	maybeGenesisBlockTime *int64 // Unix Nano
	maybeTotalBlockCount  *big.Int

	migrationHelper migrationhelper.MigrationHelper
}

func NewChainStats(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	migrationHelper migrationhelper.MigrationHelper,
) *ChainStats {
	return &ChainStats{
		rdbprojectionbase.NewRDbBase(
			rdbConn.ToHandle(),
			"ChainStats",
		),

		rdbConn,
		logger,

		nil,
		nil,

		migrationHelper,
	}
}

func (_ *ChainStats) GetEventsToListen() []string {
	return []string{
		event_usecase.GENESIS_CREATED,
		event_usecase.BLOCK_CREATED,
	}
}

// TODO: should change it to projection folder name to `chain_stats`, then we can remove it
const (
	MIGRATION_DIRECOTRY = "projection/chainstats/migrations"
)

func (projection *ChainStats) OnInit() error {
	if projection.migrationHelper != nil {
		projection.migrationHelper.Migrate()
	}

	chainStatsView := view.NewChainStats(projection.rdbConn.ToHandle())

	var getErr error
	rawGenesisBlockTime, getErr := chainStatsView.FindBy(GENESIS_BLOCK_TIME)
	if getErr != nil {
		return fmt.Errorf("error getting genesis block time on init: %v", getErr)
	}

	if rawGenesisBlockTime == "" {
		return nil
	}

	unixNanoTime, parseErr := strconv.ParseInt(rawGenesisBlockTime, 10, 64)
	if parseErr != nil {
		return fmt.Errorf("error parsing genesis block time on init: %v", parseErr)
	}
	projection.maybeGenesisBlockTime = &unixNanoTime

	rawTotalBlockCount, getErr := chainStatsView.FindBy(TOTAL_BLOCK_COUNT)
	if getErr != nil {
		return fmt.Errorf("error getting total block count on init: %v", getErr)
	}
	var ok bool
	if projection.maybeTotalBlockCount, ok = new(big.Int).SetString(rawTotalBlockCount, 10); !ok {
		return errors.New("error parsing total block count as big.Int on init")
	}

	return nil
}

func (projection *ChainStats) HandleEvents(height int64, events []event_entity.Event) error {
	rdbTx, err := projection.rdbConn.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %v", err)
	}

	committed := false
	defer func() {
		if !committed {
			_ = rdbTx.Rollback()
		}
	}()

	rdbTxHandle := rdbTx.ToHandle()
	chainStatsView := view.NewChainStats(rdbTxHandle)

	event := events[0]
	var maybeGenesisBlockTime *int64
	if genesisCreatedEvent, ok := event.(*event_usecase.GenesisCreated); ok {
		genesisBlockTime, parseErr := utctime.Parse(
			time.RFC3339, genesisCreatedEvent.Genesis.GenesisTime,
		)
		if parseErr != nil {
			return fmt.Errorf("error pasring genesis block time: %v", parseErr)
		}

		genesisBlockTimeUnixNano := genesisBlockTime.UnixNano()
		maybeGenesisBlockTime = &genesisBlockTimeUnixNano

		var setErr error
		if setErr = chainStatsView.Set(
			GENESIS_BLOCK_TIME, strconv.FormatInt(genesisBlockTime.UnixNano(), 10),
		); setErr != nil {
			return fmt.Errorf("error setting genesis block time: %v", setErr)
		}
		if setErr = chainStatsView.Set(TOTAL_BLOCK_TIME, "0"); setErr != nil {
			return fmt.Errorf("error setting total block time: %v", setErr)
		}
		if setErr = chainStatsView.Set(TOTAL_BLOCK_COUNT, "0"); setErr != nil {
			return fmt.Errorf("error setting total block count: %v", setErr)
		}
	} else if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
		blockTime := blockCreatedEvent.Block.Time

		totalBlockTime := blockTime.UnixNano() - *projection.maybeGenesisBlockTime
		totalBlockCount := new(big.Int).Add(projection.maybeTotalBlockCount, big.NewInt(1))

		var setErr error
		if setErr = chainStatsView.Set(TOTAL_BLOCK_TIME, strconv.FormatInt(totalBlockTime, 10)); setErr != nil {
			return fmt.Errorf("error setting total block time: %v", setErr)
		}
		if setErr = chainStatsView.Set(TOTAL_BLOCK_COUNT, totalBlockCount.String()); setErr != nil {
			return fmt.Errorf("error setting total block count: %v", setErr)
		}
	}

	if err = projection.UpdateLastHandledEventHeight(rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err = rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true

	if maybeGenesisBlockTime != nil {
		projection.maybeGenesisBlockTime = maybeGenesisBlockTime
		projection.maybeTotalBlockCount = big.NewInt(0)
	} else {
		projection.maybeTotalBlockCount = new(big.Int).Add(projection.maybeTotalBlockCount, big.NewInt(1))
	}
	return nil
}
