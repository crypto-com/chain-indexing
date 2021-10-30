package validator

import (
	"fmt"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/internal/utctime"
	"github.com/crypto-com/chain-indexing/projection/validator/view"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

// a simple tool to keep record of incremental record
type privTotalIncrementalMap struct {
	data map[string]int64
}

func privNewTotalIncrementalMap() *privTotalIncrementalMap {
	return &privTotalIncrementalMap{
		data: make(map[string]int64),
	}
}
func (totalMap *privTotalIncrementalMap) Increment(key string, value int64) {
	if _, ok := totalMap.data[key]; !ok {
		totalMap.data[key] = 0
	}
	totalMap.data[key] += value
}
func (totalMap *privTotalIncrementalMap) IncrementByOne(key string) {
	totalMap.Increment(key, 1)
}
func (totalMap *privTotalIncrementalMap) Set(key string, value int64) {
	totalMap.data[key] = value
}
func (totalMap *privTotalIncrementalMap) Persist(validatorActivitiesTotalView *view.ValidatorActivitiesTotal) error {
	for key, value := range totalMap.data {
		if err := validatorActivitiesTotalView.Increment(key, value); err != nil {
			return fmt.Errorf("error incrementing total of `%s`: %w", key, err)
		}
	}
	return nil
}

func (projection *Validator) projectValidatorActivitiesView(
	validatorsView *view.Validators,
	validatorActivitiesView *view.ValidatorActivities,
	validatorActivitiesTotalView *view.ValidatorActivitiesTotal,
	blockHash string,
	blockTime utctime.UTCTime,
	events []event_entity.Event,
) error {
	activityRows := make([]view.ValidatorActivityRow, 0)
	totalIncrementalMap := privNewTotalIncrementalMap()
	for _, event := range events {
		if createGenesisValidatorEvent, ok := event.(*event_usecase.CreateGenesisValidator); ok {
			activityRows = append(activityRows, view.ValidatorActivityRow{
				BlockHeight:          createGenesisValidatorEvent.BlockHeight,
				BlockHash:            blockHash,
				BlockTime:            blockTime,
				MaybeTransactionHash: nil,
				OperatorAddress:      createGenesisValidatorEvent.ValidatorAddress,
				Success:              true,
				Data: view.ValidatorActivityRowData{
					Type:    createGenesisValidatorEvent.Name(),
					Content: createGenesisValidatorEvent,
				},
			})

			totalIncrementalMap.IncrementByOne("-")
			totalIncrementalMap.IncrementByOne(createGenesisValidatorEvent.ValidatorAddress)
			totalIncrementalMap.IncrementByOne(
				fmt.Sprintf("%s:%s", createGenesisValidatorEvent.ValidatorAddress, createGenesisValidatorEvent.Name()),
			)
			totalIncrementalMap.IncrementByOne(fmt.Sprintf("-:%s", createGenesisValidatorEvent.Name()))
		} else if createValidatorEvent, ok := event.(*event_usecase.MsgCreateValidator); ok {
			activityRows = append(activityRows, view.ValidatorActivityRow{
				BlockHeight:          createValidatorEvent.BlockHeight,
				BlockHash:            blockHash,
				BlockTime:            blockTime,
				MaybeTransactionHash: primptr.String(createValidatorEvent.TxHash()),
				OperatorAddress:      createValidatorEvent.ValidatorAddress,
				Success:              createValidatorEvent.TxSuccess(),
				Data: view.ValidatorActivityRowData{
					Type:    createValidatorEvent.MsgType(),
					Content: createValidatorEvent,
				},
			})

			totalIncrementalMap.IncrementByOne("-")
			totalIncrementalMap.IncrementByOne(createValidatorEvent.ValidatorAddress)
			totalIncrementalMap.IncrementByOne(
				fmt.Sprintf("%s:%s", createValidatorEvent.ValidatorAddress, createValidatorEvent.Name()),
			)
			totalIncrementalMap.IncrementByOne(fmt.Sprintf("-:%s", createValidatorEvent.Name()))
		} else if editValidatorEvent, ok := event.(*event_usecase.MsgEditValidator); ok {
			activityRows = append(activityRows, view.ValidatorActivityRow{
				BlockHeight:          editValidatorEvent.BlockHeight,
				BlockHash:            blockHash,
				BlockTime:            blockTime,
				MaybeTransactionHash: primptr.String(editValidatorEvent.TxHash()),
				OperatorAddress:      editValidatorEvent.ValidatorAddress,
				Success:              editValidatorEvent.TxSuccess(),
				Data: view.ValidatorActivityRowData{
					Type:    editValidatorEvent.MsgType(),
					Content: editValidatorEvent,
				},
			})

			totalIncrementalMap.IncrementByOne("-")
			totalIncrementalMap.IncrementByOne(editValidatorEvent.ValidatorAddress)
			totalIncrementalMap.IncrementByOne(
				fmt.Sprintf("%s:%s", editValidatorEvent.ValidatorAddress, editValidatorEvent.Name()),
			)
			totalIncrementalMap.IncrementByOne(fmt.Sprintf("-:%s", editValidatorEvent.Name()))
		} else if delegateEvent, ok := event.(*event_usecase.MsgDelegate); ok {
			activityRows = append(activityRows, view.ValidatorActivityRow{
				BlockHeight:          delegateEvent.BlockHeight,
				BlockHash:            blockHash,
				BlockTime:            blockTime,
				MaybeTransactionHash: primptr.String(delegateEvent.TxHash()),
				OperatorAddress:      delegateEvent.ValidatorAddress,
				Success:              delegateEvent.TxSuccess(),
				Data: view.ValidatorActivityRowData{
					Type:    delegateEvent.MsgType(),
					Content: delegateEvent,
				},
			})

			totalIncrementalMap.IncrementByOne("-")
			totalIncrementalMap.IncrementByOne(delegateEvent.ValidatorAddress)
			totalIncrementalMap.IncrementByOne(
				fmt.Sprintf("%s:%s", delegateEvent.ValidatorAddress, delegateEvent.Name()),
			)
			totalIncrementalMap.IncrementByOne(fmt.Sprintf("-:%s", delegateEvent.Name()))
		} else if redelegateEvent, ok := event.(*event_usecase.MsgBeginRedelegate); ok {
			activityRows = append(activityRows, view.ValidatorActivityRow{
				BlockHeight:          redelegateEvent.BlockHeight,
				BlockHash:            blockHash,
				BlockTime:            blockTime,
				MaybeTransactionHash: primptr.String(redelegateEvent.TxHash()),
				OperatorAddress:      redelegateEvent.ValidatorSrcAddress,
				Success:              redelegateEvent.TxSuccess(),
				Data: view.ValidatorActivityRowData{
					Type:    redelegateEvent.MsgType(),
					Content: redelegateEvent,
				},
			})
			activityRows = append(activityRows, view.ValidatorActivityRow{
				BlockHeight:          redelegateEvent.BlockHeight,
				BlockHash:            blockHash,
				BlockTime:            blockTime,
				MaybeTransactionHash: primptr.String(redelegateEvent.TxHash()),
				OperatorAddress:      redelegateEvent.ValidatorDstAddress,
				Success:              redelegateEvent.TxSuccess(),
				Data: view.ValidatorActivityRowData{
					Type:    redelegateEvent.MsgType(),
					Content: redelegateEvent,
				},
			})

			totalIncrementalMap.Increment("-", int64(2))
			totalIncrementalMap.IncrementByOne(redelegateEvent.ValidatorSrcAddress)
			totalIncrementalMap.IncrementByOne(
				fmt.Sprintf("%s:%s", redelegateEvent.ValidatorSrcAddress, redelegateEvent.Name()),
			)
			totalIncrementalMap.IncrementByOne(fmt.Sprintf("-:%s", redelegateEvent.Name()))
			totalIncrementalMap.IncrementByOne(redelegateEvent.ValidatorDstAddress)
			totalIncrementalMap.IncrementByOne(fmt.Sprintf("-:%s", redelegateEvent.Name()))
		} else if undelegateEvent, ok := event.(*event_usecase.MsgUndelegate); ok {
			activityRows = append(activityRows, view.ValidatorActivityRow{
				BlockHeight:          undelegateEvent.BlockHeight,
				BlockHash:            blockHash,
				BlockTime:            blockTime,
				MaybeTransactionHash: primptr.String(undelegateEvent.TxHash()),
				OperatorAddress:      undelegateEvent.ValidatorAddress,
				Success:              undelegateEvent.TxSuccess(),
				Data: view.ValidatorActivityRowData{
					Type:    undelegateEvent.MsgType(),
					Content: undelegateEvent,
				},
			})

			totalIncrementalMap.IncrementByOne("-")
			totalIncrementalMap.IncrementByOne(undelegateEvent.ValidatorAddress)
			totalIncrementalMap.IncrementByOne(
				fmt.Sprintf("%s:%s", undelegateEvent.ValidatorAddress, undelegateEvent.Name()),
			)
			totalIncrementalMap.IncrementByOne(fmt.Sprintf("-:%s", undelegateEvent.Name()))
		} else if withdrawDelegatorRewardEvent, ok := event.(*event_usecase.MsgWithdrawDelegatorReward); ok {
			activityRows = append(activityRows, view.ValidatorActivityRow{
				BlockHeight:          withdrawDelegatorRewardEvent.BlockHeight,
				BlockHash:            blockHash,
				BlockTime:            blockTime,
				MaybeTransactionHash: primptr.String(withdrawDelegatorRewardEvent.TxHash()),
				OperatorAddress:      withdrawDelegatorRewardEvent.ValidatorAddress,
				Success:              withdrawDelegatorRewardEvent.TxSuccess(),
				Data: view.ValidatorActivityRowData{
					Type:    withdrawDelegatorRewardEvent.MsgType(),
					Content: withdrawDelegatorRewardEvent,
				},
			})

			totalIncrementalMap.IncrementByOne("-")
			totalIncrementalMap.IncrementByOne(withdrawDelegatorRewardEvent.ValidatorAddress)
			totalIncrementalMap.IncrementByOne(
				fmt.Sprintf("%s:%s",
					withdrawDelegatorRewardEvent.ValidatorAddress,
					withdrawDelegatorRewardEvent.Name(),
				),
			)
			totalIncrementalMap.IncrementByOne(fmt.Sprintf("-:%s", withdrawDelegatorRewardEvent.Name()))
		} else if withdrawValidatorCommissionEvent, ok := event.(*event_usecase.MsgWithdrawValidatorCommission); ok {
			activityRows = append(activityRows, view.ValidatorActivityRow{
				BlockHeight:          withdrawValidatorCommissionEvent.BlockHeight,
				BlockHash:            blockHash,
				BlockTime:            blockTime,
				MaybeTransactionHash: primptr.String(withdrawValidatorCommissionEvent.TxHash()),
				OperatorAddress:      withdrawValidatorCommissionEvent.ValidatorAddress,
				Success:              withdrawValidatorCommissionEvent.TxSuccess(),
				Data: view.ValidatorActivityRowData{
					Type:    withdrawValidatorCommissionEvent.MsgType(),
					Content: withdrawValidatorCommissionEvent,
				},
			})

			totalIncrementalMap.IncrementByOne("-")
			totalIncrementalMap.IncrementByOne(withdrawValidatorCommissionEvent.ValidatorAddress)
			totalIncrementalMap.IncrementByOne(
				fmt.Sprintf("%s:%s",
					withdrawValidatorCommissionEvent.ValidatorAddress,
					withdrawValidatorCommissionEvent.Name(),
				),
			)
			totalIncrementalMap.IncrementByOne(
				fmt.Sprintf("-:%s", withdrawValidatorCommissionEvent.Name()),
			)
		} else if validatorJailedEvent, ok := event.(*event_usecase.ValidatorJailed); ok {
			validatorRow, err := validatorsView.FindBy(view.ValidatorIdentity{
				MaybeConsensusNodeAddress: &validatorJailedEvent.ConsensusNodeAddress,
			})
			if err != nil {
				return fmt.Errorf(
					"error getting existing validator `%s`: %v", validatorJailedEvent.ConsensusNodeAddress, err,
				)
			}
			activityRows = append(activityRows, view.ValidatorActivityRow{
				BlockHeight:          validatorJailedEvent.BlockHeight,
				BlockHash:            blockHash,
				BlockTime:            blockTime,
				MaybeTransactionHash: nil,
				OperatorAddress:      validatorRow.OperatorAddress,
				Success:              true,
				Data: view.ValidatorActivityRowData{
					Type:    validatorJailedEvent.Name(),
					Content: validatorJailedEvent,
				},
			})

			totalIncrementalMap.IncrementByOne("-")
			totalIncrementalMap.IncrementByOne(validatorRow.OperatorAddress)
			totalIncrementalMap.IncrementByOne(
				fmt.Sprintf("%s:%s", validatorRow.OperatorAddress, validatorJailedEvent.Name()),
			)
			totalIncrementalMap.IncrementByOne(fmt.Sprintf("-:%s", validatorJailedEvent.Name()))
		} else if validatorSlashedEvent, ok := event.(*event_usecase.ValidatorSlashed); ok {
			validatorRow, err := validatorsView.FindBy(view.ValidatorIdentity{
				MaybeConsensusNodeAddress: &validatorSlashedEvent.ConsensusNodeAddress,
			})
			if err != nil {
				return fmt.Errorf(
					"error getting existing validator `%s`: %v", validatorSlashedEvent.ConsensusNodeAddress, err)
			}
			activityRows = append(activityRows, view.ValidatorActivityRow{
				BlockHeight:          validatorSlashedEvent.BlockHeight,
				BlockHash:            blockHash,
				BlockTime:            blockTime,
				MaybeTransactionHash: nil,
				OperatorAddress:      validatorRow.OperatorAddress,
				Success:              true,
				Data: view.ValidatorActivityRowData{
					Type:    validatorSlashedEvent.Name(),
					Content: validatorSlashedEvent,
				},
			})

			totalIncrementalMap.IncrementByOne("-")
			totalIncrementalMap.IncrementByOne(validatorRow.OperatorAddress)
			totalIncrementalMap.IncrementByOne(
				fmt.Sprintf("%s:%s", validatorRow.OperatorAddress, validatorSlashedEvent.Name()),
			)
			totalIncrementalMap.IncrementByOne(fmt.Sprintf("-:%s", validatorSlashedEvent.Name()))
		} else if unjailEvent, ok := event.(*event_usecase.MsgUnjail); ok {
			activityRows = append(activityRows, view.ValidatorActivityRow{
				BlockHeight:          unjailEvent.BlockHeight,
				BlockHash:            blockHash,
				BlockTime:            blockTime,
				MaybeTransactionHash: primptr.String(unjailEvent.TxHash()),
				OperatorAddress:      unjailEvent.ValidatorAddr,
				Success:              unjailEvent.TxSuccess(),
				Data: view.ValidatorActivityRowData{
					Type:    unjailEvent.MsgType(),
					Content: unjailEvent,
				},
			})

			totalIncrementalMap.IncrementByOne("-")
			totalIncrementalMap.IncrementByOne(unjailEvent.ValidatorAddr)
			totalIncrementalMap.IncrementByOne(
				fmt.Sprintf("%s:%s", unjailEvent.ValidatorAddr, unjailEvent.Name()),
			)
			totalIncrementalMap.IncrementByOne(fmt.Sprintf("-:%s", unjailEvent.Name()))
		}
	}

	if err := validatorActivitiesView.InsertAll(activityRows); err != nil {
		return fmt.Errorf("error inserting validator activities into view: %w", err)
	}
	if err := totalIncrementalMap.Persist(validatorActivitiesTotalView); err != nil {
		return err
	}
	return nil
}
