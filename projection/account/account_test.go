package account_test

import (
	"fmt"
	"strconv"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/projection/account"
	account_view "github.com/crypto-com/chain-indexing/projection/account/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/model/genesis"
	"github.com/stretchr/testify/assert"
	testify_mock "github.com/stretchr/testify/mock"
)

func NewAccountProjection(rdbConn rdb.Conn, client cosmosapp.Client) *account.Account {
	return account.NewAccount(
		nil,
		rdbConn,
		"prefix",
		client,
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

func TestAccount_HandleEvents(t *testing.T) {
	testCases := []struct {
		Name     string
		Events   []entity_event.Event
		MockFunc func(mockClient *cosmosapp.MockClient) []*testify_mock.Mock
	}{
		{
			Name: "HandleGenesisCreatedEvent",
			Events: []entity_event.Event{
				&event_usecase.GenesisCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.GENESIS_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					Genesis: genesis.Genesis{
						GenesisTime:     "",
						ChainID:         "",
						InitialHeight:   "",
						ConsensusParams: genesis.ConsensusParams{},
						AppHash:         "",
						AppState: genesis.AppState{
							Auth: genesis.Auth{},
							Bank: genesis.Bank{
								Params: genesis.BankParams{},
								Balances: []genesis.Balance{
									{
										Address: "Address",
										Coins: []genesis.MinDeposit{
											{
												Denom:  "Denom",
												Amount: strconv.Itoa(10),
											},
										},
									},
								},
								Supply:        nil,
								DenomMetadata: nil,
							},
							Capability:   genesis.Capability{},
							Chainmain:    genesis.Chainmain{},
							Distribution: genesis.Distribution{},
							Evidence:     genesis.AppStateEvidence{},
							Genutil:      genesis.Genutil{},
							Gov:          genesis.Gov{},
							Ibc:          genesis.Ibc{},
							Mint:         genesis.Mint{},
							Params:       nil,
							Slashing:     genesis.Slashing{},
							Staking:      genesis.Staking{},
							Supply:       genesis.Supply{},
							Transfer:     genesis.Transfer{},
							Upgrade:      genesis.Upgrade{},
							Vesting:      genesis.Vesting{},
						},
						Validators: nil,
					},
				},
			},
			MockFunc: func(mockClient *cosmosapp.MockClient) (mocks []*testify_mock.Mock) {
				mockAccountsView := account_view.NewMockAccountsView(nil).(*account_view.MockAccountsView)
				mocks = append(mocks, &mockAccountsView.Mock)

				account.NewAccountsView = func(_ *rdb.Handle) account_view.Accounts {
					return mockAccountsView
				}

				mockAccountsView.On(
					"IncrementUsableBalance",
					"Address",
					"Denom",
					int64(10),
				).Return(nil)

				account.UpdateLastHandledEventHeight = func(_ *account.Account, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleCoinSpentEvent",
			Events: []entity_event.Event{
				&event_usecase.CoinSpent{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.COIN_SPENT,
						Version:     1,
						BlockHeight: 1,
					}),
					Address: "Address",
					Amount: []coin.Coin{
						{
							Denom:  "Denom",
							Amount: coin.NewInt(10),
						},
					},
				},
			},
			MockFunc: func(mockClient *cosmosapp.MockClient) (mocks []*testify_mock.Mock) {
				mockAccountsView := account_view.NewMockAccountsView(nil).(*account_view.MockAccountsView)
				mocks = append(mocks, &mockAccountsView.Mock)

				account.NewAccountsView = func(_ *rdb.Handle) account_view.Accounts {
					return mockAccountsView
				}

				mockAccountsView.On(
					"DecrementUsableBalance",
					"Address",
					"Denom",
					int64(10),
				).Return(nil)

				mockAccountsView.On(
					"InsertAccountEvent",
					account_view.AccountEvent{
						Address: "Address",
						Type:    "CoinSpent",
						Data: coin.Coins{
							coin.Coin{
								Denom:  "Denom",
								Amount: coin.NewInt(10),
							},
						},
						BlockHeight: 1,
					},
				).Return(nil)

				account.UpdateLastHandledEventHeight = func(_ *account.Account, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleCoinReceivedEvent",
			Events: []entity_event.Event{
				&event_usecase.CoinReceived{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.COIN_RECEIVED,
						Version:     1,
						BlockHeight: 1,
					}),
					Address: "Address",
					Amount: []coin.Coin{
						{
							Denom:  "Denom",
							Amount: coin.NewInt(10),
						},
					},
				},
			},
			MockFunc: func(mockClient *cosmosapp.MockClient) (mocks []*testify_mock.Mock) {
				mockAccountsView := account_view.NewMockAccountsView(nil).(*account_view.MockAccountsView)
				mocks = append(mocks, &mockAccountsView.Mock)

				account.NewAccountsView = func(_ *rdb.Handle) account_view.Accounts {
					return mockAccountsView
				}

				mockAccountsView.On(
					"IncrementUsableBalance",
					"Address",
					"Denom",
					int64(10),
				).Return(nil)

				mockAccountsView.On(
					"InsertAccountEvent",
					account_view.AccountEvent{
						Address: "Address",
						Type:    "CoinReceived",
						Data: coin.Coins{
							coin.Coin{
								Denom:  "Denom",
								Amount: coin.NewInt(10),
							},
						},
						BlockHeight: 1,
					},
				).Return(nil)

				account.UpdateLastHandledEventHeight = func(_ *account.Account, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleCoinMintEvent",
			Events: []entity_event.Event{
				&event_usecase.CoinMint{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.COIN_MINT,
						Version:     1,
						BlockHeight: 1,
					}),
					Address: "Address",
					Amount: []coin.Coin{
						{
							Denom:  "Denom",
							Amount: coin.NewInt(10),
						},
					},
				},
			},
			MockFunc: func(mockClient *cosmosapp.MockClient) (mocks []*testify_mock.Mock) {
				mockAccountsView := account_view.NewMockAccountsView(nil).(*account_view.MockAccountsView)
				mocks = append(mocks, &mockAccountsView.Mock)

				account.NewAccountsView = func(_ *rdb.Handle) account_view.Accounts {
					return mockAccountsView
				}

				mockAccountsView.On(
					"IncrementUsableBalance",
					"Address",
					"Denom",
					int64(10),
				).Return(nil)

				mockAccountsView.On(
					"InsertAccountEvent",
					account_view.AccountEvent{
						Address: "Address",
						Type:    "CoinMint",
						Data: coin.Coins{
							coin.Coin{
								Denom:  "Denom",
								Amount: coin.NewInt(10),
							},
						},
						BlockHeight: 1,
					},
				).Return(nil)

				account.UpdateLastHandledEventHeight = func(_ *account.Account, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleCoinBurnEvent",
			Events: []entity_event.Event{
				&event_usecase.CoinBurn{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.COIN_BURN,
						Version:     1,
						BlockHeight: 1,
					}),
					Address: "Address",
					Amount: []coin.Coin{
						{
							Denom:  "Denom",
							Amount: coin.NewInt(10),
						},
					},
				},
			},
			MockFunc: func(mockClient *cosmosapp.MockClient) (mocks []*testify_mock.Mock) {
				mockAccountsView := account_view.NewMockAccountsView(nil).(*account_view.MockAccountsView)
				mocks = append(mocks, &mockAccountsView.Mock)

				account.NewAccountsView = func(_ *rdb.Handle) account_view.Accounts {
					return mockAccountsView
				}

				mockAccountsView.On(
					"DecrementUsableBalance",
					"Address",
					"Denom",
					int64(10),
				).Return(nil)

				mockAccountsView.On(
					"InsertAccountEvent",
					account_view.AccountEvent{
						Address: "Address",
						Type:    "CoinBurn",
						Data: coin.Coins{
							coin.Coin{
								Denom:  "Denom",
								Amount: coin.NewInt(10),
							},
						},
						BlockHeight: 1,
					},
				).Return(nil)

				account.UpdateLastHandledEventHeight = func(_ *account.Account, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleTransactionFailedEvent",
			Events: []entity_event.Event{
				&event_usecase.TransactionFailed{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.TRANSACTION_FAILED,
						Version:     1,
						BlockHeight: 1,
					}),
					TxHash:   "",
					Index:    0,
					Code:     0,
					Log:      "",
					MsgCount: 0,
					Signers: []model.TransactionSigner{
						{
							TransactionSignerInfo: model.TransactionSignerInfo{
								Type:            "",
								IsMultiSig:      false,
								Pubkeys:         []string{
									"A5N3/S/r5nQjjvJpdXR5eY1QiebM9ULBsWVbCw/a6dA4",
								},
								MaybeThreshold:  nil,
								AccountSequence: 0,
							},
							Address: "prefix18mcwp6vtlvpgxy62eledk3chhjguw636muuqul",
						},
					},
					Senders: nil,
					Fee: []coin.Coin{
						{
							Denom:  "Denom",
							Amount: coin.NewInt(10),
						},
					},
					FeePayer:      "",
					FeeGranter:    "",
					GasWanted:     0,
					GasUsed:       0,
					Memo:          "",
					TimeoutHeight: 0,
				},
			},
			MockFunc: func(mockClient *cosmosapp.MockClient) (mocks []*testify_mock.Mock) {
				mockAccountsView := account_view.NewMockAccountsView(nil).(*account_view.MockAccountsView)
				mocks = append(mocks, &mockAccountsView.Mock)

				account.NewAccountsView = func(_ *rdb.Handle) account_view.Accounts {
					return mockAccountsView
				}

				mockAccountsView.On(
					"DecrementUsableBalance",
					"prefix18mcwp6vtlvpgxy62eledk3chhjguw636muuqul",
					"Denom",
					int64(10),
				).Return(nil)

				mockAccountsView.On(
					"InsertAccountEvent",
					account_view.AccountEvent{
						Address: "prefix18mcwp6vtlvpgxy62eledk3chhjguw636muuqul",
						Type:    "TransactionFailed",
						Data: coin.Coins{
							coin.Coin{
								Denom:  "Denom",
								Amount: coin.NewInt(10),
							},
						},
						BlockHeight: 1,
					},
				).Return(nil)

				account.UpdateLastHandledEventHeight = func(_ *account.Account, _ *rdb.Handle, _ int64) error {
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
		mockClient := &cosmosapp.MockClient{}
		mocks := tc.MockFunc(mockClient)
		mocks = append(mocks, &mockClient.Mock)

		projection := NewAccountProjection(mockRDbConn, mockClient)
		err := projection.HandleEvents(1, tc.Events)
		assert.NoError(t, err)

		for _, m := range mocks {
			m.AssertExpectations(t)
		}

		fmt.Println(tc.Name, "Passed")
	}
}
