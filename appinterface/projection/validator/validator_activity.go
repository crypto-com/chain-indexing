package validator

import (
	"fmt"

	"github.com/crypto-com/chainindex/appinterface/projection/validator/view"
	event_entity "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/internal/primptr"
	"github.com/crypto-com/chainindex/internal/utctime"
	event_usecase "github.com/crypto-com/chainindex/usecase/event"
)

func (projection *Validator) projectValidatorActivitiesView(
	validatorsView *view.Validators,
	validatorActivitiesView *view.ValidatorActivities,
	blockHash string,
	blockTime utctime.UTCTime,
	event event_entity.Event,
) error {
	if createValidatorEvent, ok := event.(*event_usecase.MsgCreateValidator); ok {
		if err := validatorActivitiesView.Insert(&view.ValidatorActivityRow{
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
		}); err != nil {
			return fmt.Errorf("error inserting MsgCreateValiator into view: %v", err)
		}
	} else if editValidatorEvent, ok := event.(*event_usecase.MsgEditValidator); ok {
		if err := validatorActivitiesView.Insert(&view.ValidatorActivityRow{
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
		}); err != nil {
			return fmt.Errorf("error inserting MsgEditValidator into view: %v", err)
		}
	} else if delegateEvent, ok := event.(*event_usecase.MsgDelegate); ok {
		if err := validatorActivitiesView.Insert(&view.ValidatorActivityRow{
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
		}); err != nil {
			return fmt.Errorf("error inserting MsgDelegate into view: %v", err)
		}
	} else if redelegateEvent, ok := event.(*event_usecase.MsgBeginRedelegate); ok {
		if err := validatorActivitiesView.Insert(&view.ValidatorActivityRow{
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
		}); err != nil {
			return fmt.Errorf("error inserting MsgBeginRedelegate into view: %v", err)
		}
		if err := validatorActivitiesView.Insert(&view.ValidatorActivityRow{
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
		}); err != nil {
			return fmt.Errorf("error inserting MsgBeginRedelegate into view: %v", err)
		}
	} else if undelegateEvent, ok := event.(*event_usecase.MsgUndelegate); ok {
		if err := validatorActivitiesView.Insert(&view.ValidatorActivityRow{
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
		}); err != nil {
			return fmt.Errorf("error inserting MsgUndelegate into view: %v", err)
		}
	} else if withdrawDelegatorRewardEvent, ok := event.(*event_usecase.MsgWithdrawDelegatorReward); ok {
		if err := validatorActivitiesView.Insert(&view.ValidatorActivityRow{
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
		}); err != nil {
			return fmt.Errorf("error inserting MsgWithdrawDelegatorReward into view: %v", err)
		}
	} else if withdrawValidatorCommissionEvent, ok := event.(*event_usecase.MsgWithdrawValidatorCommission); ok {
		if err := validatorActivitiesView.Insert(&view.ValidatorActivityRow{
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
		}); err != nil {
			return fmt.Errorf("error inserting MsgEditValidator into view: %v", err)
		}
	} else if blockProposerRewardedEvent, ok := event.(*event_usecase.BlockProposerRewarded); ok {
		if err := validatorActivitiesView.Insert(&view.ValidatorActivityRow{
			BlockHeight:          blockProposerRewardedEvent.BlockHeight,
			BlockHash:            blockHash,
			BlockTime:            blockTime,
			MaybeTransactionHash: nil,
			OperatorAddress:      blockProposerRewardedEvent.Validator,
			Success:              true,
			Data: view.ValidatorActivityRowData{
				Type:    blockProposerRewardedEvent.Name(),
				Content: blockProposerRewardedEvent,
			},
		}); err != nil {
			return fmt.Errorf("error inserting BlockProposerRewarded into view: %v", err)
		}
	} else if blockRewardedEvent, ok := event.(*event_usecase.BlockRewarded); ok {
		if err := validatorActivitiesView.Insert(&view.ValidatorActivityRow{
			BlockHeight:          blockRewardedEvent.BlockHeight,
			BlockHash:            blockHash,
			BlockTime:            blockTime,
			MaybeTransactionHash: nil,
			OperatorAddress:      blockRewardedEvent.Validator,
			Success:              true,
			Data: view.ValidatorActivityRowData{
				Type:    blockRewardedEvent.Name(),
				Content: blockRewardedEvent,
			},
		}); err != nil {
			return fmt.Errorf("error inserting BlockRewarded into view: %v", err)
		}
	} else if blockCommissionedEvent, ok := event.(*event_usecase.BlockCommissioned); ok {
		if err := validatorActivitiesView.Insert(&view.ValidatorActivityRow{
			BlockHeight:          blockCommissionedEvent.BlockHeight,
			BlockHash:            blockHash,
			BlockTime:            blockTime,
			MaybeTransactionHash: nil,
			OperatorAddress:      blockCommissionedEvent.Validator,
			Success:              true,
			Data: view.ValidatorActivityRowData{
				Type:    blockCommissionedEvent.Name(),
				Content: blockCommissionedEvent,
			},
		}); err != nil {
			return fmt.Errorf("error inserting BlockCommissioned into view: %v", err)
		}
	} else if validatorJailedEvent, ok := event.(*event_usecase.ValidatorJailed); ok {
		validatorRow, err := validatorsView.FindBy(view.ValidatorIdentity{
			MaybeConsensusNodeAddress: &validatorJailedEvent.ConsensusNodeAddress,
		})
		if err != nil {
			return fmt.Errorf(
				"error getting existing validator `%s`: %v", validatorJailedEvent.ConsensusNodeAddress, err,
			)
		}
		if err := validatorActivitiesView.Insert(&view.ValidatorActivityRow{
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
		}); err != nil {
			return fmt.Errorf("error inserting BlockCommissioned into view: %v", err)
		}
	} else if validatorSlashedEvent, ok := event.(*event_usecase.ValidatorSlashed); ok {
		validatorRow, err := validatorsView.FindBy(view.ValidatorIdentity{
			MaybeConsensusNodeAddress: &validatorSlashedEvent.ConsensusNodeAddress,
		})
		if err != nil {
			return fmt.Errorf(
				"error getting existing validator `%s`: %v", validatorSlashedEvent.ConsensusNodeAddress, err)
		}
		if err := validatorActivitiesView.Insert(&view.ValidatorActivityRow{
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
		}); err != nil {
			return fmt.Errorf("error inserting BlockCommissioned into view: %v", err)
		}
	} else if unjailEvent, ok := event.(*event_usecase.MsgUnjail); ok {
		if err := validatorActivitiesView.Insert(&view.ValidatorActivityRow{
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
		}); err != nil {
			return fmt.Errorf("error inserting MsgEditValidator into view: %v", err)
		}
	}

	return nil
}
