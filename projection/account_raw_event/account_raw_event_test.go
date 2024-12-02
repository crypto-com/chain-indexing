package account_raw_event_test

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	entity_event "github.com/crypto-com/chain-indexing/entity/event"

	"github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/projection/account_raw_event"
	account_raw_event_view "github.com/crypto-com/chain-indexing/projection/account_raw_event/view"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	usecase_model "github.com/crypto-com/chain-indexing/usecase/model"
	testify_mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func NewAccountRawEventProjection(rdbConn rdb.Conn) *account_raw_event.AccountRawEvent {
	return account_raw_event.NewAccountRawEvent(
		nil,
		rdbConn,
		"cro",
		nil,
	)
}

func NewMockRDbConn() *test.MockRDbConn {
	mock := test.NewMockRDbConn()
	mock.On("ToHandle").Return(&rdb.Handle{
		Runner:   mock,
		TypeConv: &pg.PgxTypeConv{},
		StmtBuilder: &rdb.StatementBuilder{
			StatementBuilderType: sq.StatementBuilderType{},
			PlaceholderFormat:    nil,
		},
	})

	return mock
}

func NewMockRDbTx() *test.MockRDbTx {
	mockTx := &test.MockRDbTx{}
	mockTx.On("ToHandle").Return(&rdb.Handle{
		Runner:      mockTx,
		TypeConv:    &pg.PgxTypeConv{},
		StmtBuilder: pg.PostgresStmtBuilder,
	}).Maybe()
	mockTx.On("Rollback").Return(nil).Maybe()
	mockTx.On("Commit").Return(nil).Maybe()

	return mockTx
}
func TestAccountRawEvents_HandleEvents(t *testing.T) {
	testCases := []struct {
		Name     string
		Events   []entity_event.Event
		MockFunc func(events []entity_event.Event) []*testify_mock.Mock
	}{
		{
			Name: "HandleCoinSpent",
			Events: []entity_event.Event{
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "coin_spent",
						Content: usecase_model.BlockResultsEvent{
							Type: "coin_spent",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "spender",
									Value: "CRO13UP2UE4TTSNP83A84VAUYN7449WUT09RLZZRXD",
									Index: true,
								},
								{
									Key:   "amount",
									Value: "6786516685basetcro",
									Index: true,
								},
							},
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockAccountRawEventsTotalView := account_raw_event_view.NewMockAccountRawEventsTotalView(nil).(*account_raw_event_view.MockAccountRawEventsTotalView)
				mocks = append(mocks, &mockAccountRawEventsTotalView.Mock)

				account_raw_event.NewAccountRawEventsTotal = func(_ *rdb.Handle) account_raw_event_view.AccountRawEventsTotal {
					return mockAccountRawEventsTotalView
				}

				mockAccountRawEventsTotalView.On(
					"Increment",
					"cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
					int64(1),
				).Return(nil)

				mockAccountRawEventsView := account_raw_event_view.NewMockAccountRawEventsView(nil).(*account_raw_event_view.MockAccountRawEventsView)
				mocks = append(mocks, &mockAccountRawEventsView.Mock)

				account_raw_event.NewAccountRawEvents = func(_ *rdb.Handle) account_raw_event_view.AccountRawEvents {
					return mockAccountRawEventsView
				}

				mockAccountRawEventsView.On(
					"InsertAll",
					[]account_raw_event_view.AccountRawEventRow{
						{
							Account:     "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "coin_spent",
									Content: usecase_model.BlockResultsEvent{
										Type: "coin_spent",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "spender",
												Value: "CRO13UP2UE4TTSNP83A84VAUYN7449WUT09RLZZRXD",
												Index: true,
											},
											{
												Key:   "amount",
												Value: "6786516685basetcro",
												Index: true,
											},
										},
									},
								},
							},
						},
					},
				).Return(nil)

				account_raw_event.UpdateLastHandledEventHeight = func(_ *account_raw_event.AccountRawEvent, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleCoinReceived",
			Events: []entity_event.Event{
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "coin_received",
						Content: usecase_model.BlockResultsEvent{
							Type: "coin_received",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "receiver",
									Value: "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
									Index: true,
								},
							},
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockAccountRawEventsTotalView := account_raw_event_view.NewMockAccountRawEventsTotalView(nil).(*account_raw_event_view.MockAccountRawEventsTotalView)
				mocks = append(mocks, &mockAccountRawEventsTotalView.Mock)

				account_raw_event.NewAccountRawEventsTotal = func(_ *rdb.Handle) account_raw_event_view.AccountRawEventsTotal {
					return mockAccountRawEventsTotalView
				}

				mockAccountRawEventsTotalView.On(
					"Increment",
					"cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
					int64(1),
				).Return(nil)

				mockAccountRawEventsView := account_raw_event_view.NewMockAccountRawEventsView(nil).(*account_raw_event_view.MockAccountRawEventsView)
				mocks = append(mocks, &mockAccountRawEventsView.Mock)

				account_raw_event.NewAccountRawEvents = func(_ *rdb.Handle) account_raw_event_view.AccountRawEvents {
					return mockAccountRawEventsView
				}

				mockAccountRawEventsView.On(
					"InsertAll",
					[]account_raw_event_view.AccountRawEventRow{
						{
							Account:     "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "coin_received",
									Content: usecase_model.BlockResultsEvent{
										Type: "coin_received",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "receiver",
												Value: "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
												Index: true,
											},
										},
									},
								},
							},
						},
					},
				).Return(nil)

				account_raw_event.UpdateLastHandledEventHeight = func(_ *account_raw_event.AccountRawEvent, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleTransfer",
			Events: []entity_event.Event{
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "transfer",
						Content: usecase_model.BlockResultsEvent{
							Type: "transfer",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "recipient",
									Value: "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
									Index: true,
								},
								{
									Key:   "sender",
									Value: "cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
									Index: true,
								},
								{
									Key:   "amount",
									Value: "1basecro",
									Index: true,
								},
							},
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockAccountRawEventsTotalView := account_raw_event_view.NewMockAccountRawEventsTotalView(nil).(*account_raw_event_view.MockAccountRawEventsTotalView)
				mocks = append(mocks, &mockAccountRawEventsTotalView.Mock)

				account_raw_event.NewAccountRawEventsTotal = func(_ *rdb.Handle) account_raw_event_view.AccountRawEventsTotal {
					return mockAccountRawEventsTotalView
				}

				mockAccountRawEventsTotalView.On(
					"Increment",
					"cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
					int64(1),
				).Return(nil)

				mockAccountRawEventsTotalView.On(
					"Increment",
					"cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
					int64(1),
				).Return(nil)

				mockAccountRawEventsView := account_raw_event_view.NewMockAccountRawEventsView(nil).(*account_raw_event_view.MockAccountRawEventsView)
				mocks = append(mocks, &mockAccountRawEventsView.Mock)

				account_raw_event.NewAccountRawEvents = func(_ *rdb.Handle) account_raw_event_view.AccountRawEvents {
					return mockAccountRawEventsView
				}

				mockAccountRawEventsView.On(
					"InsertAll",
					[]account_raw_event_view.AccountRawEventRow{
						{
							Account:     "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "transfer",
									Content: usecase_model.BlockResultsEvent{
										Type: "transfer",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "recipient",
												Value: "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
												Index: true,
											},
											{
												Key:   "sender",
												Value: "cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
												Index: true,
											},
											{
												Key:   "amount",
												Value: "1basecro",
												Index: true,
											},
										},
									},
								},
							},
						},
						{
							Account:     "cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "transfer",
									Content: usecase_model.BlockResultsEvent{
										Type: "transfer",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "recipient",
												Value: "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
												Index: true,
											},
											{
												Key:   "sender",
												Value: "cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
												Index: true,
											},
											{
												Key:   "amount",
												Value: "1basecro",
												Index: true,
											},
										},
									},
								},
							},
						},
					},
				).Return(nil)

				account_raw_event.UpdateLastHandledEventHeight = func(_ *account_raw_event.AccountRawEvent, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleCoinbase",
			Events: []entity_event.Event{
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "coinbase",
						Content: usecase_model.BlockResultsEvent{
							Type: "coinbase",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "minter",
									Value: "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
									Index: true,
								},
							},
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockAccountRawEventsTotalView := account_raw_event_view.NewMockAccountRawEventsTotalView(nil).(*account_raw_event_view.MockAccountRawEventsTotalView)
				mocks = append(mocks, &mockAccountRawEventsTotalView.Mock)

				account_raw_event.NewAccountRawEventsTotal = func(_ *rdb.Handle) account_raw_event_view.AccountRawEventsTotal {
					return mockAccountRawEventsTotalView
				}

				mockAccountRawEventsTotalView.On(
					"Increment",
					"cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
					int64(1),
				).Return(nil)

				mockAccountRawEventsView := account_raw_event_view.NewMockAccountRawEventsView(nil).(*account_raw_event_view.MockAccountRawEventsView)
				mocks = append(mocks, &mockAccountRawEventsView.Mock)

				account_raw_event.NewAccountRawEvents = func(_ *rdb.Handle) account_raw_event_view.AccountRawEvents {
					return mockAccountRawEventsView
				}

				mockAccountRawEventsView.On(
					"InsertAll",
					[]account_raw_event_view.AccountRawEventRow{
						{
							Account:     "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "coinbase",
									Content: usecase_model.BlockResultsEvent{
										Type: "coinbase",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "minter",
												Value: "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
												Index: true,
											},
										},
									},
								},
							},
						},
					},
				).Return(nil)

				account_raw_event.UpdateLastHandledEventHeight = func(_ *account_raw_event.AccountRawEvent, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleBurn",
			Events: []entity_event.Event{
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "burn",
						Content: usecase_model.BlockResultsEvent{
							Type: "burn",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "burner",
									Value: "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
									Index: true,
								},
							},
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockAccountRawEventsTotalView := account_raw_event_view.NewMockAccountRawEventsTotalView(nil).(*account_raw_event_view.MockAccountRawEventsTotalView)
				mocks = append(mocks, &mockAccountRawEventsTotalView.Mock)

				account_raw_event.NewAccountRawEventsTotal = func(_ *rdb.Handle) account_raw_event_view.AccountRawEventsTotal {
					return mockAccountRawEventsTotalView
				}

				mockAccountRawEventsTotalView.On(
					"Increment",
					"cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
					int64(1),
				).Return(nil)

				mockAccountRawEventsView := account_raw_event_view.NewMockAccountRawEventsView(nil).(*account_raw_event_view.MockAccountRawEventsView)
				mocks = append(mocks, &mockAccountRawEventsView.Mock)

				account_raw_event.NewAccountRawEvents = func(_ *rdb.Handle) account_raw_event_view.AccountRawEvents {
					return mockAccountRawEventsView
				}

				mockAccountRawEventsView.On(
					"InsertAll",
					[]account_raw_event_view.AccountRawEventRow{
						{
							Account:     "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "burn",
									Content: usecase_model.BlockResultsEvent{
										Type: "burn",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "burner",
												Value: "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
												Index: true,
											},
										},
									},
								},
							},
						},
					},
				).Return(nil)

				account_raw_event.UpdateLastHandledEventHeight = func(_ *account_raw_event.AccountRawEvent, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleProposerReward",
			Events: []entity_event.Event{
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "proposer_reward",
						Content: usecase_model.BlockResultsEvent{
							Type: "proposer_reward",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "validator",
									Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
									Index: true,
								},
							},
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockAccountRawEventsTotalView := account_raw_event_view.NewMockAccountRawEventsTotalView(nil).(*account_raw_event_view.MockAccountRawEventsTotalView)
				mocks = append(mocks, &mockAccountRawEventsTotalView.Mock)

				account_raw_event.NewAccountRawEventsTotal = func(_ *rdb.Handle) account_raw_event_view.AccountRawEventsTotal {
					return mockAccountRawEventsTotalView
				}

				mockAccountRawEventsTotalView.On(
					"Increment",
					"crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
					int64(1),
				).Return(nil)

				mockAccountRawEventsView := account_raw_event_view.NewMockAccountRawEventsView(nil).(*account_raw_event_view.MockAccountRawEventsView)
				mocks = append(mocks, &mockAccountRawEventsView.Mock)

				account_raw_event.NewAccountRawEvents = func(_ *rdb.Handle) account_raw_event_view.AccountRawEvents {
					return mockAccountRawEventsView
				}

				mockAccountRawEventsView.On(
					"InsertAll",
					[]account_raw_event_view.AccountRawEventRow{
						{
							Account:     "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "proposer_reward",
									Content: usecase_model.BlockResultsEvent{
										Type: "proposer_reward",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "validator",
												Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
												Index: true,
											},
										},
									},
								},
							},
						},
					},
				).Return(nil)

				account_raw_event.UpdateLastHandledEventHeight = func(_ *account_raw_event.AccountRawEvent, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleRewards",
			Events: []entity_event.Event{
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "rewards",
						Content: usecase_model.BlockResultsEvent{
							Type: "rewards",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "validator",
									Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
									Index: true,
								},
							},
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockAccountRawEventsTotalView := account_raw_event_view.NewMockAccountRawEventsTotalView(nil).(*account_raw_event_view.MockAccountRawEventsTotalView)
				mocks = append(mocks, &mockAccountRawEventsTotalView.Mock)

				account_raw_event.NewAccountRawEventsTotal = func(_ *rdb.Handle) account_raw_event_view.AccountRawEventsTotal {
					return mockAccountRawEventsTotalView
				}

				mockAccountRawEventsTotalView.On(
					"Increment",
					"crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
					int64(1),
				).Return(nil)

				mockAccountRawEventsView := account_raw_event_view.NewMockAccountRawEventsView(nil).(*account_raw_event_view.MockAccountRawEventsView)
				mocks = append(mocks, &mockAccountRawEventsView.Mock)

				account_raw_event.NewAccountRawEvents = func(_ *rdb.Handle) account_raw_event_view.AccountRawEvents {
					return mockAccountRawEventsView
				}

				mockAccountRawEventsView.On(
					"InsertAll",
					[]account_raw_event_view.AccountRawEventRow{
						{
							Account:     "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "rewards",
									Content: usecase_model.BlockResultsEvent{
										Type: "rewards",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "validator",
												Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
												Index: true,
											},
										},
									},
								},
							},
						},
					},
				).Return(nil)

				account_raw_event.UpdateLastHandledEventHeight = func(_ *account_raw_event.AccountRawEvent, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleCommission",
			Events: []entity_event.Event{
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "commission",
						Content: usecase_model.BlockResultsEvent{
							Type: "commission",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "validator",
									Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
									Index: true,
								},
							},
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockAccountRawEventsTotalView := account_raw_event_view.NewMockAccountRawEventsTotalView(nil).(*account_raw_event_view.MockAccountRawEventsTotalView)
				mocks = append(mocks, &mockAccountRawEventsTotalView.Mock)

				account_raw_event.NewAccountRawEventsTotal = func(_ *rdb.Handle) account_raw_event_view.AccountRawEventsTotal {
					return mockAccountRawEventsTotalView
				}

				mockAccountRawEventsTotalView.On(
					"Increment",
					"crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
					int64(1),
				).Return(nil)

				mockAccountRawEventsView := account_raw_event_view.NewMockAccountRawEventsView(nil).(*account_raw_event_view.MockAccountRawEventsView)
				mocks = append(mocks, &mockAccountRawEventsView.Mock)

				account_raw_event.NewAccountRawEvents = func(_ *rdb.Handle) account_raw_event_view.AccountRawEvents {
					return mockAccountRawEventsView
				}

				mockAccountRawEventsView.On(
					"InsertAll",
					[]account_raw_event_view.AccountRawEventRow{
						{
							Account:     "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "commission",
									Content: usecase_model.BlockResultsEvent{
										Type: "commission",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "validator",
												Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
												Index: true,
											},
										},
									},
								},
							},
						},
					},
				).Return(nil)

				account_raw_event.UpdateLastHandledEventHeight = func(_ *account_raw_event.AccountRawEvent, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleSlash",
			Events: []entity_event.Event{
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "slash",
						Content: usecase_model.BlockResultsEvent{
							Type: "slash",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "address",
									Value: "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
									Index: true,
								},
							},
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockAccountRawEventsTotalView := account_raw_event_view.NewMockAccountRawEventsTotalView(nil).(*account_raw_event_view.MockAccountRawEventsTotalView)
				mocks = append(mocks, &mockAccountRawEventsTotalView.Mock)

				account_raw_event.NewAccountRawEventsTotal = func(_ *rdb.Handle) account_raw_event_view.AccountRawEventsTotal {
					return mockAccountRawEventsTotalView
				}

				mockAccountRawEventsTotalView.On(
					"Increment",
					"cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
					int64(1),
				).Return(nil)

				mockAccountRawEventsView := account_raw_event_view.NewMockAccountRawEventsView(nil).(*account_raw_event_view.MockAccountRawEventsView)
				mocks = append(mocks, &mockAccountRawEventsView.Mock)

				account_raw_event.NewAccountRawEvents = func(_ *rdb.Handle) account_raw_event_view.AccountRawEvents {
					return mockAccountRawEventsView
				}

				mockAccountRawEventsView.On(
					"InsertAll",
					[]account_raw_event_view.AccountRawEventRow{
						{
							Account:     "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "slash",
									Content: usecase_model.BlockResultsEvent{
										Type: "slash",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "address",
												Value: "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
												Index: true,
											},
										},
									},
								},
							},
						},
					},
				).Return(nil)

				account_raw_event.UpdateLastHandledEventHeight = func(_ *account_raw_event.AccountRawEvent, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleCompleteUnbonding",
			Events: []entity_event.Event{
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "complete_unbonding",
						Content: usecase_model.BlockResultsEvent{
							Type: "complete_unbonding",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "delegator",
									Value: "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
									Index: true,
								},
								{
									Key:   "validator",
									Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
									Index: true,
								},
							},
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockAccountRawEventsTotalView := account_raw_event_view.NewMockAccountRawEventsTotalView(nil).(*account_raw_event_view.MockAccountRawEventsTotalView)
				mocks = append(mocks, &mockAccountRawEventsTotalView.Mock)

				account_raw_event.NewAccountRawEventsTotal = func(_ *rdb.Handle) account_raw_event_view.AccountRawEventsTotal {
					return mockAccountRawEventsTotalView
				}

				mockAccountRawEventsTotalView.On(
					"Increment",
					"cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
					int64(1),
				).Return(nil)

				mockAccountRawEventsTotalView.On(
					"Increment",
					"crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
					int64(1),
				).Return(nil)

				mockAccountRawEventsView := account_raw_event_view.NewMockAccountRawEventsView(nil).(*account_raw_event_view.MockAccountRawEventsView)
				mocks = append(mocks, &mockAccountRawEventsView.Mock)

				account_raw_event.NewAccountRawEvents = func(_ *rdb.Handle) account_raw_event_view.AccountRawEvents {
					return mockAccountRawEventsView
				}

				mockAccountRawEventsView.On(
					"InsertAll",
					[]account_raw_event_view.AccountRawEventRow{
						{
							Account:     "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "complete_unbonding",
									Content: usecase_model.BlockResultsEvent{
										Type: "complete_unbonding",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "delegator",
												Value: "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
												Index: true,
											},
											{
												Key:   "validator",
												Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
												Index: true,
											},
										},
									},
								},
							},
						},
						{
							Account:     "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "complete_unbonding",
									Content: usecase_model.BlockResultsEvent{
										Type: "complete_unbonding",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "delegator",
												Value: "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
												Index: true,
											},
											{
												Key:   "validator",
												Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
												Index: true,
											},
										},
									},
								},
							},
						},
					},
				).Return(nil)

				account_raw_event.UpdateLastHandledEventHeight = func(_ *account_raw_event.AccountRawEvent, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleEthereumSendToCosmosHandled",
			Events: []entity_event.Event{
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "ethereum_send_to_cosmos_handled",
						Content: usecase_model.BlockResultsEvent{
							Type: "ethereum_send_to_cosmos_handled",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "sender",
									Value: "cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
									Index: true,
								},
								{
									Key:   "receiver",
									Value: "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
									Index: true,
								},
							},
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockAccountRawEventsTotalView := account_raw_event_view.NewMockAccountRawEventsTotalView(nil).(*account_raw_event_view.MockAccountRawEventsTotalView)
				mocks = append(mocks, &mockAccountRawEventsTotalView.Mock)

				account_raw_event.NewAccountRawEventsTotal = func(_ *rdb.Handle) account_raw_event_view.AccountRawEventsTotal {
					return mockAccountRawEventsTotalView
				}

				mockAccountRawEventsTotalView.On(
					"Increment",
					"cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
					int64(1),
				).Return(nil)

				mockAccountRawEventsTotalView.On(
					"Increment",
					"cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
					int64(1),
				).Return(nil)

				mockAccountRawEventsView := account_raw_event_view.NewMockAccountRawEventsView(nil).(*account_raw_event_view.MockAccountRawEventsView)
				mocks = append(mocks, &mockAccountRawEventsView.Mock)

				account_raw_event.NewAccountRawEvents = func(_ *rdb.Handle) account_raw_event_view.AccountRawEvents {
					return mockAccountRawEventsView
				}

				mockAccountRawEventsView.On(
					"InsertAll",
					[]account_raw_event_view.AccountRawEventRow{
						{
							Account:     "cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "ethereum_send_to_cosmos_handled",
									Content: usecase_model.BlockResultsEvent{
										Type: "ethereum_send_to_cosmos_handled",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "sender",
												Value: "cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
												Index: true,
											},
											{
												Key:   "receiver",
												Value: "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
												Index: true,
											},
										},
									},
								},
							},
						},
						{
							Account:     "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "ethereum_send_to_cosmos_handled",
									Content: usecase_model.BlockResultsEvent{
										Type: "ethereum_send_to_cosmos_handled",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "sender",
												Value: "cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
												Index: true,
											},
											{
												Key:   "receiver",
												Value: "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
												Index: true,
											},
										},
									},
								},
							},
						},
					},
				).Return(nil)

				account_raw_event.UpdateLastHandledEventHeight = func(_ *account_raw_event.AccountRawEvent, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMessage",
			Events: []entity_event.Event{
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "message",
						Content: usecase_model.BlockResultsEvent{
							Type: "message",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "sender",
									Value: "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
									Index: true,
								},
							},
						},
					},
				},
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "message",
						Content: usecase_model.BlockResultsEvent{
							Type: "message",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "granter",
									Value: "cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
									Index: true,
								},
							},
						},
					},
				},
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "message",
						Content: usecase_model.BlockResultsEvent{
							Type: "message",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "grantee",
									Value: "cro1v7h324rm06r5admszs3d5jvq5fqnnuzpy7xxxx",
									Index: true,
								},
							},
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockAccountRawEventsTotalView := account_raw_event_view.NewMockAccountRawEventsTotalView(nil).(*account_raw_event_view.MockAccountRawEventsTotalView)
				mocks = append(mocks, &mockAccountRawEventsTotalView.Mock)

				account_raw_event.NewAccountRawEventsTotal = func(_ *rdb.Handle) account_raw_event_view.AccountRawEventsTotal {
					return mockAccountRawEventsTotalView
				}

				mockAccountRawEventsTotalView.On(
					"Increment",
					"cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
					int64(1),
				).Return(nil)

				mockAccountRawEventsTotalView.On(
					"Increment",
					"cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
					int64(1),
				).Return(nil)

				mockAccountRawEventsTotalView.On(
					"Increment",
					"cro1v7h324rm06r5admszs3d5jvq5fqnnuzpy7xxxx",
					int64(1),
				).Return(nil)

				mockAccountRawEventsView := account_raw_event_view.NewMockAccountRawEventsView(nil).(*account_raw_event_view.MockAccountRawEventsView)
				mocks = append(mocks, &mockAccountRawEventsView.Mock)

				account_raw_event.NewAccountRawEvents = func(_ *rdb.Handle) account_raw_event_view.AccountRawEvents {
					return mockAccountRawEventsView
				}

				mockAccountRawEventsView.On(
					"InsertAll",
					[]account_raw_event_view.AccountRawEventRow{
						{
							Account:     "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "message",
									Content: usecase_model.BlockResultsEvent{
										Type: "message",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "sender",
												Value: "cro13up2ue4ttsnp83a84vauyn7449wut09rlzzrxd",
												Index: true,
											},
										},
									},
								},
							},
						},

						{
							Account:     "cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "message",
									Content: usecase_model.BlockResultsEvent{
										Type: "message",
										Attributes: []usecase_model.BlockResultsEventAttribute{

											{
												Key:   "granter",
												Value: "cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
												Index: true,
											},
										},
									},
								},
							},
						},
						{
							Account:     "cro1v7h324rm06r5admszs3d5jvq5fqnnuzpy7xxxx",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "message",
									Content: usecase_model.BlockResultsEvent{
										Type: "message",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "grantee",
												Value: "cro1v7h324rm06r5admszs3d5jvq5fqnnuzpy7xxxx",
												Index: true,
											},
										},
									},
								},
							},
						},
					},
				).Return(nil)

				account_raw_event.UpdateLastHandledEventHeight = func(_ *account_raw_event.AccountRawEvent, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleWithdrawRewards",
			Events: []entity_event.Event{
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "withdraw_rewards",
						Content: usecase_model.BlockResultsEvent{
							Type: "withdraw_rewards",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "validator",
									Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
									Index: true,
								},
							},
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockAccountRawEventsTotalView := account_raw_event_view.NewMockAccountRawEventsTotalView(nil).(*account_raw_event_view.MockAccountRawEventsTotalView)
				mocks = append(mocks, &mockAccountRawEventsTotalView.Mock)

				account_raw_event.NewAccountRawEventsTotal = func(_ *rdb.Handle) account_raw_event_view.AccountRawEventsTotal {
					return mockAccountRawEventsTotalView
				}

				mockAccountRawEventsTotalView.On(
					"Increment",
					"crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
					int64(1),
				).Return(nil)

				mockAccountRawEventsView := account_raw_event_view.NewMockAccountRawEventsView(nil).(*account_raw_event_view.MockAccountRawEventsView)
				mocks = append(mocks, &mockAccountRawEventsView.Mock)

				account_raw_event.NewAccountRawEvents = func(_ *rdb.Handle) account_raw_event_view.AccountRawEvents {
					return mockAccountRawEventsView
				}

				mockAccountRawEventsView.On(
					"InsertAll",
					[]account_raw_event_view.AccountRawEventRow{
						{
							Account:     "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "withdraw_rewards",
									Content: usecase_model.BlockResultsEvent{
										Type: "withdraw_rewards",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "validator",
												Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
												Index: true,
											},
										},
									},
								},
							},
						},
					},
				).Return(nil)

				account_raw_event.UpdateLastHandledEventHeight = func(_ *account_raw_event.AccountRawEvent, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleSetWithdrawAddress",
			Events: []entity_event.Event{
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "set_withdraw_address",
						Content: usecase_model.BlockResultsEvent{
							Type: "set_withdraw_address",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "withdraw_address",
									Value: "cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
									Index: true,
								},
							},
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockAccountRawEventsTotalView := account_raw_event_view.NewMockAccountRawEventsTotalView(nil).(*account_raw_event_view.MockAccountRawEventsTotalView)
				mocks = append(mocks, &mockAccountRawEventsTotalView.Mock)

				account_raw_event.NewAccountRawEventsTotal = func(_ *rdb.Handle) account_raw_event_view.AccountRawEventsTotal {
					return mockAccountRawEventsTotalView
				}

				mockAccountRawEventsTotalView.On(
					"Increment",
					"cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
					int64(1),
				).Return(nil)

				mockAccountRawEventsView := account_raw_event_view.NewMockAccountRawEventsView(nil).(*account_raw_event_view.MockAccountRawEventsView)
				mocks = append(mocks, &mockAccountRawEventsView.Mock)

				account_raw_event.NewAccountRawEvents = func(_ *rdb.Handle) account_raw_event_view.AccountRawEvents {
					return mockAccountRawEventsView
				}

				mockAccountRawEventsView.On(
					"InsertAll",
					[]account_raw_event_view.AccountRawEventRow{
						{
							Account:     "cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "set_withdraw_address",
									Content: usecase_model.BlockResultsEvent{
										Type: "set_withdraw_address",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "withdraw_address",
												Value: "cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
												Index: true,
											},
										},
									},
								},
							},
						},
					},
				).Return(nil)

				account_raw_event.UpdateLastHandledEventHeight = func(_ *account_raw_event.AccountRawEvent, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleSetWithdrawAddress",
			Events: []entity_event.Event{
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "set_withdraw_address",
						Content: usecase_model.BlockResultsEvent{
							Type: "set_withdraw_address",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "withdraw_address",
									Value: "cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
									Index: true,
								},
							},
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockAccountRawEventsTotalView := account_raw_event_view.NewMockAccountRawEventsTotalView(nil).(*account_raw_event_view.MockAccountRawEventsTotalView)
				mocks = append(mocks, &mockAccountRawEventsTotalView.Mock)

				account_raw_event.NewAccountRawEventsTotal = func(_ *rdb.Handle) account_raw_event_view.AccountRawEventsTotal {
					return mockAccountRawEventsTotalView
				}

				mockAccountRawEventsTotalView.On(
					"Increment",
					"cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
					int64(1),
				).Return(nil)

				mockAccountRawEventsView := account_raw_event_view.NewMockAccountRawEventsView(nil).(*account_raw_event_view.MockAccountRawEventsView)
				mocks = append(mocks, &mockAccountRawEventsView.Mock)

				account_raw_event.NewAccountRawEvents = func(_ *rdb.Handle) account_raw_event_view.AccountRawEvents {
					return mockAccountRawEventsView
				}

				mockAccountRawEventsView.On(
					"InsertAll",
					[]account_raw_event_view.AccountRawEventRow{
						{
							Account:     "cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "set_withdraw_address",
									Content: usecase_model.BlockResultsEvent{
										Type: "set_withdraw_address",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "withdraw_address",
												Value: "cro1l3g8w5567d4dqtvw7nljestgm5envh0n0rwqmh",
												Index: true,
											},
										},
									},
								},
							},
						},
					},
				).Return(nil)

				account_raw_event.UpdateLastHandledEventHeight = func(_ *account_raw_event.AccountRawEvent, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleCompleteRedelegation",
			Events: []entity_event.Event{
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "complete_redelegation",
						Content: usecase_model.BlockResultsEvent{
							Type: "complete_redelegation",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "source_validator",
									Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
									Index: true,
								},
								{
									Key:   "destination_validator",
									Value: "crocncl149dyku4d4c36dmvvw583xqcflau3d9x303ffcj",
									Index: true,
								},
								{
									Key:   "delegator",
									Value: "cro1qfqd93p68lr65xlgt3547vyxthzsr27pxp66ac",
									Index: true,
								},
							},
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockAccountRawEventsTotalView := account_raw_event_view.NewMockAccountRawEventsTotalView(nil).(*account_raw_event_view.MockAccountRawEventsTotalView)
				mocks = append(mocks, &mockAccountRawEventsTotalView.Mock)

				account_raw_event.NewAccountRawEventsTotal = func(_ *rdb.Handle) account_raw_event_view.AccountRawEventsTotal {
					return mockAccountRawEventsTotalView
				}

				mockAccountRawEventsTotalView.On(
					"Increment",
					"crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
					int64(1),
				).Return(nil)

				mockAccountRawEventsTotalView.On(
					"Increment",
					"crocncl149dyku4d4c36dmvvw583xqcflau3d9x303ffcj",
					int64(1),
				).Return(nil)

				mockAccountRawEventsTotalView.On(
					"Increment",
					"cro1qfqd93p68lr65xlgt3547vyxthzsr27pxp66ac",
					int64(1),
				).Return(nil)

				mockAccountRawEventsView := account_raw_event_view.NewMockAccountRawEventsView(nil).(*account_raw_event_view.MockAccountRawEventsView)
				mocks = append(mocks, &mockAccountRawEventsView.Mock)

				account_raw_event.NewAccountRawEvents = func(_ *rdb.Handle) account_raw_event_view.AccountRawEvents {
					return mockAccountRawEventsView
				}

				mockAccountRawEventsView.On(
					"InsertAll",
					[]account_raw_event_view.AccountRawEventRow{
						{
							Account:     "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "complete_redelegation",
									Content: usecase_model.BlockResultsEvent{
										Type: "complete_redelegation",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "source_validator",
												Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
												Index: true,
											},
											{
												Key:   "destination_validator",
												Value: "crocncl149dyku4d4c36dmvvw583xqcflau3d9x303ffcj",
												Index: true,
											},
											{
												Key:   "delegator",
												Value: "cro1qfqd93p68lr65xlgt3547vyxthzsr27pxp66ac",
												Index: true,
											},
										},
									},
								},
							},
						},
						{
							Account:     "crocncl149dyku4d4c36dmvvw583xqcflau3d9x303ffcj",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "complete_redelegation",
									Content: usecase_model.BlockResultsEvent{
										Type: "complete_redelegation",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "source_validator",
												Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
												Index: true,
											},
											{
												Key:   "destination_validator",
												Value: "crocncl149dyku4d4c36dmvvw583xqcflau3d9x303ffcj",
												Index: true,
											},
											{
												Key:   "delegator",
												Value: "cro1qfqd93p68lr65xlgt3547vyxthzsr27pxp66ac",
												Index: true,
											},
										},
									},
								},
							},
						},
						{
							Account:     "cro1qfqd93p68lr65xlgt3547vyxthzsr27pxp66ac",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "complete_redelegation",
									Content: usecase_model.BlockResultsEvent{
										Type: "complete_redelegation",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "source_validator",
												Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
												Index: true,
											},
											{
												Key:   "destination_validator",
												Value: "crocncl149dyku4d4c36dmvvw583xqcflau3d9x303ffcj",
												Index: true,
											},
											{
												Key:   "delegator",
												Value: "cro1qfqd93p68lr65xlgt3547vyxthzsr27pxp66ac",
												Index: true,
											},
										},
									},
								},
							},
						},
					},
				).Return(nil)

				account_raw_event.UpdateLastHandledEventHeight = func(_ *account_raw_event.AccountRawEvent, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleCreateValidator",
			Events: []entity_event.Event{
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "create_validator",
						Content: usecase_model.BlockResultsEvent{
							Type: "create_validator",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "validator",
									Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
									Index: true,
								},
							},
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockAccountRawEventsTotalView := account_raw_event_view.NewMockAccountRawEventsTotalView(nil).(*account_raw_event_view.MockAccountRawEventsTotalView)
				mocks = append(mocks, &mockAccountRawEventsTotalView.Mock)

				account_raw_event.NewAccountRawEventsTotal = func(_ *rdb.Handle) account_raw_event_view.AccountRawEventsTotal {
					return mockAccountRawEventsTotalView
				}

				mockAccountRawEventsTotalView.On(
					"Increment",
					"crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
					int64(1),
				).Return(nil)

				mockAccountRawEventsView := account_raw_event_view.NewMockAccountRawEventsView(nil).(*account_raw_event_view.MockAccountRawEventsView)
				mocks = append(mocks, &mockAccountRawEventsView.Mock)

				account_raw_event.NewAccountRawEvents = func(_ *rdb.Handle) account_raw_event_view.AccountRawEvents {
					return mockAccountRawEventsView
				}

				mockAccountRawEventsView.On(
					"InsertAll",
					[]account_raw_event_view.AccountRawEventRow{
						{
							Account:     "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "create_validator",
									Content: usecase_model.BlockResultsEvent{
										Type: "create_validator",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "validator",
												Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
												Index: true,
											},
										},
									},
								},
							},
						},
					},
				).Return(nil)

				account_raw_event.UpdateLastHandledEventHeight = func(_ *account_raw_event.AccountRawEvent, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleDelegate",
			Events: []entity_event.Event{
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "delegate",
						Content: usecase_model.BlockResultsEvent{
							Type: "delegate",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "validator",
									Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
									Index: true,
								},
							},
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockAccountRawEventsTotalView := account_raw_event_view.NewMockAccountRawEventsTotalView(nil).(*account_raw_event_view.MockAccountRawEventsTotalView)
				mocks = append(mocks, &mockAccountRawEventsTotalView.Mock)

				account_raw_event.NewAccountRawEventsTotal = func(_ *rdb.Handle) account_raw_event_view.AccountRawEventsTotal {
					return mockAccountRawEventsTotalView
				}

				mockAccountRawEventsTotalView.On(
					"Increment",
					"crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
					int64(1),
				).Return(nil)

				mockAccountRawEventsView := account_raw_event_view.NewMockAccountRawEventsView(nil).(*account_raw_event_view.MockAccountRawEventsView)
				mocks = append(mocks, &mockAccountRawEventsView.Mock)

				account_raw_event.NewAccountRawEvents = func(_ *rdb.Handle) account_raw_event_view.AccountRawEvents {
					return mockAccountRawEventsView
				}

				mockAccountRawEventsView.On(
					"InsertAll",
					[]account_raw_event_view.AccountRawEventRow{
						{
							Account:     "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "delegate",
									Content: usecase_model.BlockResultsEvent{
										Type: "delegate",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "validator",
												Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
												Index: true,
											},
										},
									},
								},
							},
						},
					},
				).Return(nil)

				account_raw_event.UpdateLastHandledEventHeight = func(_ *account_raw_event.AccountRawEvent, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleUnbond",
			Events: []entity_event.Event{
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "unbond",
						Content: usecase_model.BlockResultsEvent{
							Type: "unbond",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "destination_validator",
									Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
									Index: true,
								},
							},
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockAccountRawEventsTotalView := account_raw_event_view.NewMockAccountRawEventsTotalView(nil).(*account_raw_event_view.MockAccountRawEventsTotalView)
				mocks = append(mocks, &mockAccountRawEventsTotalView.Mock)

				account_raw_event.NewAccountRawEventsTotal = func(_ *rdb.Handle) account_raw_event_view.AccountRawEventsTotal {
					return mockAccountRawEventsTotalView
				}

				mockAccountRawEventsTotalView.On(
					"Increment",
					"crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
					int64(1),
				).Return(nil)

				mockAccountRawEventsView := account_raw_event_view.NewMockAccountRawEventsView(nil).(*account_raw_event_view.MockAccountRawEventsView)
				mocks = append(mocks, &mockAccountRawEventsView.Mock)

				account_raw_event.NewAccountRawEvents = func(_ *rdb.Handle) account_raw_event_view.AccountRawEvents {
					return mockAccountRawEventsView
				}

				mockAccountRawEventsView.On(
					"InsertAll",
					[]account_raw_event_view.AccountRawEventRow{
						{
							Account:     "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "unbond",
									Content: usecase_model.BlockResultsEvent{
										Type: "unbond",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "destination_validator",
												Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
												Index: true,
											},
										},
									},
								},
							},
						},
					},
				).Return(nil)

				account_raw_event.UpdateLastHandledEventHeight = func(_ *account_raw_event.AccountRawEvent, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleRedelegate",
			Events: []entity_event.Event{
				&event_usecase.BlockRawEventCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_RAW_EVENT_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					BlockHash: "",
					BlockTime: utctime.UTCTime{},
					Data: usecase_model.DataParams{
						Type: "redelegate",
						Content: usecase_model.BlockResultsEvent{
							Type: "redelegate",
							Attributes: []usecase_model.BlockResultsEventAttribute{
								{
									Key:   "source_validator",
									Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
									Index: true,
								},
								{
									Key:   "destination_validator",
									Value: "crocncl149dyku4d4c36dmvvw583xqcflau3d9x303ffcj",
									Index: true,
								},
							},
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				mockAccountRawEventsTotalView := account_raw_event_view.NewMockAccountRawEventsTotalView(nil).(*account_raw_event_view.MockAccountRawEventsTotalView)
				mocks = append(mocks, &mockAccountRawEventsTotalView.Mock)

				account_raw_event.NewAccountRawEventsTotal = func(_ *rdb.Handle) account_raw_event_view.AccountRawEventsTotal {
					return mockAccountRawEventsTotalView
				}

				mockAccountRawEventsTotalView.On(
					"Increment",
					"crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
					int64(1),
				).Return(nil)

				mockAccountRawEventsTotalView.On(
					"Increment",
					"crocncl149dyku4d4c36dmvvw583xqcflau3d9x303ffcj",
					int64(1),
				).Return(nil)

				mockAccountRawEventsView := account_raw_event_view.NewMockAccountRawEventsView(nil).(*account_raw_event_view.MockAccountRawEventsView)
				mocks = append(mocks, &mockAccountRawEventsView.Mock)

				account_raw_event.NewAccountRawEvents = func(_ *rdb.Handle) account_raw_event_view.AccountRawEvents {
					return mockAccountRawEventsView
				}

				mockAccountRawEventsView.On(
					"InsertAll",
					[]account_raw_event_view.AccountRawEventRow{
						{
							Account:     "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "redelegate",
									Content: usecase_model.BlockResultsEvent{
										Type: "redelegate",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "source_validator",
												Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
												Index: true,
											},
											{
												Key:   "destination_validator",
												Value: "crocncl149dyku4d4c36dmvvw583xqcflau3d9x303ffcj",
												Index: true,
											},
										},
									},
								},
							},
						},
						{
							Account:     "crocncl149dyku4d4c36dmvvw583xqcflau3d9x303ffcj",
							BlockHeight: 1,
							BlockHash:   "",
							BlockTime:   utctime.UTCTime{},
							Data: []account_raw_event_view.AccountRawEventRowData{
								{
									Type: "redelegate",
									Content: usecase_model.BlockResultsEvent{
										Type: "redelegate",
										Attributes: []usecase_model.BlockResultsEventAttribute{
											{
												Key:   "source_validator",
												Value: "crocncl12razrhugyd75tpmek3t39gtelak20z8w68lzd7",
												Index: true,
											},
											{
												Key:   "destination_validator",
												Value: "crocncl149dyku4d4c36dmvvw583xqcflau3d9x303ffcj",
												Index: true,
											},
										},
									},
								},
							},
						},
					},
				).Return(nil)

				account_raw_event.UpdateLastHandledEventHeight = func(_ *account_raw_event.AccountRawEvent, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
	}
	for _, tc := range testCases {
		mockRDbConn := NewMockRDbConn()
		mockTx := NewMockRDbTx()
		mockRDbConn.On("Begin").Return(mockTx, nil)
		mocks := tc.MockFunc(tc.Events)

		projection := NewAccountRawEventProjection(mockRDbConn)
		err := projection.HandleEvents(1, tc.Events)

		require.NoError(t, err)

		for _, m := range mocks {
			m.AssertExpectations(t)
		}

		fmt.Println(tc.Name, "Passed")
	}
}
