package account_test

import (
	"fmt"
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
	"github.com/stretchr/testify/assert"
	testify_mock "github.com/stretchr/testify/mock"
)

func NewAccountProjection(rdbConn rdb.Conn, client cosmosapp.Client) *account.Account {
	return account.NewAccount(
		nil,
		rdbConn,
		client,
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
			Name: "HandleAccountTransferred",
			Events: []entity_event.Event{
				&event_usecase.AccountTransferred{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.ACCOUNT_TRANSFERRED,
						Version:     1,
						BlockHeight: 1,
					}),
					Sender:    "Sender",
					Recipient: "Recipient",
					Amount:    coin.Coins{},
				},
			},
			MockFunc: func(mockClient *cosmosapp.MockClient) (mocks []*testify_mock.Mock) {
				mockClient.On("Account", "Recipient").Return(
					&cosmosapp.Account{
						Type:    "AccountType",
						Address: "Recipient",
						MaybePubkey: &cosmosapp.PubKey{
							Type: "PubKeyType",
							Key:  "Key",
						},
						AccountNumber: "AccountNumber",
						Sequence:      "Sequence",
						MaybeModuleAccount: &cosmosapp.ModuleAccount{
							Name:        "",
							Permissions: []string{},
						},
						MaybeDelayedVestingAccount: &cosmosapp.DelayedVestingAccount{
							OriginalVesting:  []cosmosapp.VestingBalance{},
							DelegatedFree:    []cosmosapp.VestingBalance{},
							DelegatedVesting: []cosmosapp.VestingBalance{},
							EndTime:          "",
						},
						MaybeContinuousVestingAccount: &cosmosapp.ContinuousVestingAccount{
							OriginalVesting:  []cosmosapp.VestingBalance{},
							DelegatedFree:    []cosmosapp.VestingBalance{},
							DelegatedVesting: []cosmosapp.VestingBalance{},
							StartTime:        "",
							EndTime:          "",
						},
						MaybePeriodicVestingAccount: &cosmosapp.PeriodicVestingAccount{
							OriginalVesting:  []cosmosapp.VestingBalance{},
							DelegatedFree:    []cosmosapp.VestingBalance{},
							DelegatedVesting: []cosmosapp.VestingBalance{},
							StartTime:        "",
							EndTime:          "",
							VestingPeriods:   []cosmosapp.VestingPeriod{},
						},
					},
					nil,
				)

				mockClient.On("Balances", "Recipient").Return(
					coin.Coins{
						coin.Coin{
							Denom:  "Denom",
							Amount: coin.NewInt(100),
						},
					},
					nil,
				)

				mockAccountsView := account_view.NewMockAccountsView(nil).(*account_view.MockAccountsView)
				mocks = append(mocks, &mockAccountsView.Mock)

				account.NewAccountsView = func(_ *rdb.Handle) account_view.Accounts {
					return mockAccountsView
				}

				pubkey := "Key"

				mockAccountsView.On(
					"Upsert",
					&account_view.AccountRow{
						Address:        "Recipient",
						Type:           "AccountType",
						MaybeName:      (*string)(nil),
						MaybePubkey:    &pubkey,
						AccountNumber:  "AccountNumber",
						SequenceNumber: "Sequence",
						Balance: coin.Coins{
							{
								Denom:  "Denom",
								Amount: coin.NewInt(100),
							},
						},
					},
				).Return(nil)

				mockClient.On("Account", "Sender").Return(
					&cosmosapp.Account{
						Type:    "AccountType",
						Address: "Sender",
						MaybePubkey: &cosmosapp.PubKey{
							Type: "PubKeyType",
							Key:  pubkey,
						},
						AccountNumber: "AccountNumber",
						Sequence:      "Sequence",
						MaybeModuleAccount: &cosmosapp.ModuleAccount{
							Name:        "",
							Permissions: []string{},
						},
						MaybeDelayedVestingAccount: &cosmosapp.DelayedVestingAccount{
							OriginalVesting:  []cosmosapp.VestingBalance{},
							DelegatedFree:    []cosmosapp.VestingBalance{},
							DelegatedVesting: []cosmosapp.VestingBalance{},
							EndTime:          "",
						},
						MaybeContinuousVestingAccount: &cosmosapp.ContinuousVestingAccount{
							OriginalVesting:  []cosmosapp.VestingBalance{},
							DelegatedFree:    []cosmosapp.VestingBalance{},
							DelegatedVesting: []cosmosapp.VestingBalance{},
							StartTime:        "",
							EndTime:          "",
						},
						MaybePeriodicVestingAccount: &cosmosapp.PeriodicVestingAccount{
							OriginalVesting:  []cosmosapp.VestingBalance{},
							DelegatedFree:    []cosmosapp.VestingBalance{},
							DelegatedVesting: []cosmosapp.VestingBalance{},
							StartTime:        "",
							EndTime:          "",
							VestingPeriods:   []cosmosapp.VestingPeriod{},
						},
					},
					nil,
				)

				mockClient.On("Balances", "Sender").Return(
					coin.Coins{
						coin.Coin{
							Denom:  "Denom",
							Amount: coin.NewInt(1000),
						},
					},
					nil,
				)

				mockAccountsView.On(
					"Upsert",
					&account_view.AccountRow{
						Address:        "Sender",
						Type:           "AccountType",
						MaybeName:      (*string)(nil),
						MaybePubkey:    &pubkey,
						AccountNumber:  "AccountNumber",
						SequenceNumber: "Sequence",
						Balance: coin.Coins{
							{
								Denom:  "Denom",
								Amount: coin.NewInt(1000),
							},
						},
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
