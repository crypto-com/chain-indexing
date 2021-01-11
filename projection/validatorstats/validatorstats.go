package validatorstats

import (
	"fmt"
	"strings"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	entity_projection "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/crypto-com/chain-indexing/projection/validatorstats/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ entity_projection.Projection = &ValidatorStats{}

const TOTAL_REWARD = "total_reward"
const TOTAL_DELEGATE = "total_delegate"

type ValidatorStats struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger
}

func NewValidatorStats(logger applogger.Logger, rdbConn rdb.Conn) *ValidatorStats {
	return &ValidatorStats{
		rdbprojectionbase.NewRDbBase(rdbConn.ToHandle(), "ValidatorStats"),

		rdbConn,
		logger,
	}
}

func (_ *ValidatorStats) GetEventsToListen() []string {
	return []string{
		event_usecase.MSG_CREATE_VALIDATOR_CREATED,
		event_usecase.BLOCK_PROPOSER_REWARDED,
		event_usecase.BLOCK_REWARDED,
		event_usecase.MSG_DELEGATE_CREATED,
		event_usecase.MSG_UNDELEGATE_CREATED,
	}
}

func (projection *ValidatorStats) OnInit() error {
	return nil
}

func (projection *ValidatorStats) HandleEvents(height int64, events []event_entity.Event) error {
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
	validatorStatsView := view.NewValidatorStats(rdbTxHandle)

	rawTotalReward, err := validatorStatsView.FindBy(TOTAL_REWARD)
	if err != nil {
		return fmt.Errorf("error getting total reward metrics: %v", err)
	}
	if rawTotalReward == "" {
		rawTotalReward = "0"
	}
	totalReward := coin.MustNewCoinFromString(rawTotalReward)

	rawTotalDelegate, err := validatorStatsView.FindBy(TOTAL_DELEGATE)
	if err != nil {
		return fmt.Errorf("error getting total reward metrics: %v", err)
	}
	if rawTotalDelegate == "" {
		rawTotalDelegate = "0"
	}
	totalDelegate := coin.MustNewCoinFromString(rawTotalDelegate)

	for _, event := range events {
		if createValidatorEvent, ok := event.(*event_usecase.MsgCreateValidator); ok {
			totalDelegate, err = totalReward.Add(createValidatorEvent.Amount)
			if err != nil {
				return fmt.Errorf("error adding validator initial delegate: %v", err)
			}
		} else if blockProposerRewardedEvent, ok := event.(*event_usecase.BlockProposerRewarded); ok {
			totalReward, err = totalReward.Add(
				coin.MustNewCoinFromString(TrimDecimalPlaces(blockProposerRewardedEvent.Amount)),
			)
			if err != nil {
				return fmt.Errorf("error adding rewards: %v", err)
			}
		} else if blockRewardedEvent, ok := event.(*event_usecase.BlockRewarded); ok {
			totalReward, err = totalReward.Add(
				coin.MustNewCoinFromString(TrimDecimalPlaces(blockRewardedEvent.Amount)),
			)
			if err != nil {
				return fmt.Errorf("error adding rewards: %v", err)
			}
		} else if delegateEvent, ok := event.(*event_usecase.MsgDelegate); ok {
			totalDelegate, err = totalDelegate.Add(delegateEvent.Amount)
			if err != nil {
				return fmt.Errorf("error adding delegate: %v", err)
			}
		} else if undelegateEvent, ok := event.(*event_usecase.MsgUndelegate); ok {
			totalDelegate, err = totalDelegate.Add(undelegateEvent.Amount.Negate())
			if err != nil {
				return fmt.Errorf("error subtracting delegate: %v", err)
			}
		}
	}

	if err = validatorStatsView.Set(TOTAL_REWARD, totalReward.String()); err != nil {
		return fmt.Errorf("error updating total reward")
	}
	if err = validatorStatsView.Set(TOTAL_DELEGATE, totalDelegate.String()); err != nil {
		return fmt.Errorf("error updating total delegate")
	}

	if err = projection.UpdateLastHandledEventHeight(rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err = rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true
	return nil
}

func TrimDecimalPlaces(s string) string {
	parts := strings.Split(s, ".")
	return parts[0]
}
