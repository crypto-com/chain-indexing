package blockevent

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/tmcosmosutils"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	"github.com/crypto-com/chain-indexing/projection/account_raw_event/view"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ projection.Projection = &AccountRawEvent{}

type AccountRawEvent struct {
	*rdbprojectionbase.Base

	accountAddressPrefix string

	rdbConn rdb.Conn
	logger  applogger.Logger

	migrationHelper migrationhelper.MigrationHelper
}

func NewAccountRawEvent(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	accountAddressPrefix string,
	migrationHelper migrationhelper.MigrationHelper,
) *AccountRawEvent {
	return &AccountRawEvent{
		rdbprojectionbase.NewRDbBase(
			rdbConn.ToHandle(),
			"AccountRawEvent",
		),

		accountAddressPrefix,

		rdbConn,
		logger,

		migrationHelper,
	}
}

func (_ *AccountRawEvent) GetEventsToListen() []string {
	return append([]string{
		event_usecase.BLOCK_CREATED,
		event_usecase.ACCOUNT_RAW_EVENT_CREATED,
	}, event_usecase.MSG_EVENTS...)
}

func (projection *AccountRawEvent) OnInit() error {
	if projection.migrationHelper != nil {
		projection.migrationHelper.Migrate()
	}
	return nil
}

func (projection *AccountRawEvent) HandleEvents(height int64, events []event_entity.Event) error {
	var err error

	var rdbTx rdb.Tx
	rdbTx, err = projection.rdbConn.Begin()
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
	eventsView := view.NewAccountRawEvents(rdbTxHandle)
	totalView := view.NewAccountRawEventsTotal(rdbTxHandle)

	var blockTime utctime.UTCTime
	var blockHash string
	eventRecords := make([]view.AccountRawEventRecord, 0)

	for _, event := range events {
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
			blockHash = blockCreatedEvent.Block.Hash
		}
	}

	for _, event := range events {
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
			blockHash = blockCreatedEvent.Block.Hash
		} else if typedEvent, ok := event.(*event_usecase.MsgSend); ok {

			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{typedEvent.FromAddress, typedEvent.ToAddress},
				Events:   typedEvent.MsgBase.Events,
			})

		} else if typedEvent, ok := event.(*event_usecase.MsgMultiSend); ok {
			involvedAccounts := make([]string, 0)
			for _, input := range typedEvent.Inputs {
				involvedAccounts = append(involvedAccounts, input.Address)
			}
			for _, output := range typedEvent.Outputs {
				involvedAccounts = append(involvedAccounts, output.Address)
			}

			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: involvedAccounts,
				Events:   typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgSetWithdrawAddress); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.DelegatorAddress,
					typedEvent.WithdrawAddress,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgWithdrawDelegatorReward); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.DelegatorAddress,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgWithdrawValidatorCommission); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.RecipientAddress,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgFundCommunityPool); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Depositor,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgSubmitParamChangeProposal); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.ProposerAddress,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgSubmitCommunityPoolSpendProposal); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.ProposerAddress,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgSubmitSoftwareUpgradeProposal); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.ProposerAddress,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgSubmitCancelSoftwareUpgradeProposal); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.ProposerAddress,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgDeposit); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Depositor,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgVote); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Voter,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgCreateValidator); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.DelegatorAddress,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgEditValidator); ok {
			accountAddress := tmcosmosutils.MustAccountAddressFromValidatorAddress(
				projection.accountAddressPrefix,
				typedEvent.ValidatorAddress,
			)

			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					accountAddress,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgDelegate); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.DelegatorAddress,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgUndelegate); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.DelegatorAddress,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgBeginRedelegate); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.DelegatorAddress,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgUnjail); ok {
			accountAddress := tmcosmosutils.MustAccountAddressFromValidatorAddress(
				projection.accountAddressPrefix,
				typedEvent.ValidatorAddr,
			)

			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					accountAddress,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgNFTIssueDenom); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Sender,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgNFTMintNFT); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Sender,
					typedEvent.Recipient,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgNFTTransferNFT); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Sender,
					typedEvent.Recipient,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgNFTEditNFT); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Sender,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgNFTBurnNFT); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Sender,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCCreateClient); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCUpdateClient); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCConnectionOpenInit); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCConnectionOpenAck); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCConnectionOpenTry); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCConnectionOpenConfirm); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCChannelOpenInit); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCChannelOpenAck); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCChannelOpenTry); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCChannelOpenConfirm); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCAcknowledgement); ok {
			accounts := []string{typedEvent.Params.Signer}

			if typedEvent.Params.MaybeFungibleTokenPacketData != nil {
				accounts = append(accounts, typedEvent.Params.MaybeFungibleTokenPacketData.Sender)
			}

			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: accounts,
				Events:   typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCRecvPacket); ok {
			accounts := []string{typedEvent.Params.Signer}

			if typedEvent.Params.MaybeFungibleTokenPacketData != nil {
				accounts = append(accounts, typedEvent.Params.MaybeFungibleTokenPacketData.Receiver)
			}

			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: accounts,
				Events:   typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCTransferTransfer); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{typedEvent.Params.Sender},
				Events:   typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCTimeout); ok {
			accounts := []string{typedEvent.Params.Signer}

			if typedEvent.Params.MaybeMsgTransfer != nil {
				accounts = append(accounts, typedEvent.Params.MaybeMsgTransfer.RefundReceiver)
			}

			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: accounts,
				Events:   typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCTimeoutOnClose); ok {
			accounts := []string{typedEvent.Params.Signer}

			if typedEvent.Params.MaybeMsgTransfer != nil {
				accounts = append(accounts, typedEvent.Params.MaybeMsgTransfer.RefundReceiver)
			}

			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: accounts,
				Events:   typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCChannelCloseInit); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{typedEvent.Params.Signer},
				Events:   typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCChannelCloseConfirm); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{typedEvent.Params.Signer},
				Events:   typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgGrant); ok {

			var accounts []string

			if typedEvent.Params.MaybeSendGrant != nil {
				accounts = []string{
					typedEvent.Params.MaybeSendGrant.Granter,
					typedEvent.Params.MaybeSendGrant.Grantee,
				}
			} else if typedEvent.Params.MaybeStakeGrant != nil {
				accounts = []string{
					typedEvent.Params.MaybeStakeGrant.Granter,
					typedEvent.Params.MaybeStakeGrant.Grantee,
				}
			} else if typedEvent.Params.MaybeGenericGrant != nil {
				accounts = []string{
					typedEvent.Params.MaybeGenericGrant.Granter,
					typedEvent.Params.MaybeGenericGrant.Grantee,
				}
			}

			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: accounts,
				Events:   typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgRevoke); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Params.Granter,
					typedEvent.Params.Grantee,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgExec); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Params.Grantee,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgGrantAllowance); ok {

			var accounts []string

			if typedEvent.Params.MaybeBasicAllowance != nil {
				accounts = []string{
					typedEvent.Params.MaybeBasicAllowance.Granter,
					typedEvent.Params.MaybeBasicAllowance.Grantee,
				}
			} else if typedEvent.Params.MaybePeriodicAllowance != nil {
				accounts = []string{
					typedEvent.Params.MaybePeriodicAllowance.Granter,
					typedEvent.Params.MaybePeriodicAllowance.Grantee,
				}
			} else if typedEvent.Params.MaybeAllowedMsgAllowance != nil {
				accounts = []string{
					typedEvent.Params.MaybeAllowedMsgAllowance.Granter,
					typedEvent.Params.MaybeAllowedMsgAllowance.Grantee,
				}
			}

			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: accounts,
				Events:   typedEvent.MsgBase.Events,
			})

		} else if typedEvent, ok := event.(*event_usecase.MsgRevokeAllowance); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Params.Granter,
					typedEvent.Params.Grantee,
				},
				Events: typedEvent.MsgBase.Events,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgCreateVestingAccount); ok {
			eventRecords = append(eventRecords, view.AccountRawEventRecord{
				Row: view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   "",
					BlockTime:   utctime.UTCTime{},
					Account:     "",
					Data:        view.AccountRawEventRowData{},
				},
				Accounts: []string{
					typedEvent.Params.FromAddress,
					typedEvent.Params.ToAddress,
				},
				Events: typedEvent.MsgBase.Events,
			})
		}
	}

	eventRows := make([]view.AccountRawEventRow, 0)

	for _, record := range eventRecords {
		for _, event := range record.Events {
			for _, account := range record.Accounts {
				eventRows = append(eventRows,
					view.AccountRawEventRow{

						BlockHeight: height,
						BlockHash:   blockHash,
						BlockTime:   blockTime,
						Account:     account,
						Data: view.AccountRawEventRowData{
							Type:    event.Type,
							Content: event,
						},
					},
				)
				// Calculate account raw event total
				if err := totalView.Increment(fmt.Sprintf("%s:-", account), 1); err != nil {
					return fmt.Errorf("error incrementing total account transaction of account: %w", err)
				}
			}
		}

	}

	if err := eventsView.InsertAll(eventRows); err != nil {
		return fmt.Errorf("error inserting account message: %w", err)
	}

	if err := projection.UpdateLastHandledEventHeight(rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err := rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true
	return nil
}
