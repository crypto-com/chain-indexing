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
		Runner:      mock,
		TypeConv:    &pg.PgxTypeConv{},
		StmtBuilder: sq.StatementBuilderType{},
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

func TestIBCChannel_HandleEvents(t *testing.T) {
	testCases := []struct {
		Name     string
		Events   []entity_event.Event
		MockFunc func(mockConn *test.MockRDbConn, mockClient *cosmosapp.MockClient) []*testify_mock.Mock
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
			MockFunc: func(mockCoon *test.MockRDbConn, mockClient *cosmosapp.MockClient) (mocks []*testify_mock.Mock) {
				mockTx := NewMockRDbTx()
				mocks = append(mocks, &mockTx.Mock)
				mockCoon.On("Begin").Return(mockTx, nil)

				mockClient.On("Account", "Recipient").Return(
					&cosmosapp.Account{
						Type:    "AccountType",
						Address: "Recipient",
						MaybePubkey: &cosmosapp.PubKey{
							Type: "PubKeyType",
							Key:  "Key",
						},
						AccountNumber: "",
						Sequence:      "",
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

				mockExecResult := &test.MockRDbExecResult{}
				mocks = append(mocks, &mockExecResult.Mock)
				mockExecResult.On("RowsAffected").Return(int64(1))
				mockTx.On(
					"Exec",
					"INSERT INTO view_accounts (address,account_type,name,pubkey,account_number,sequence_number,balance) VALUES ($1,$2,$3,$4,$5,$6,$7) ON CONFLICT(address) DO UPDATE SET balance = EXCLUDED.balance",
					"Recipient",
					"AccountType",
					(*string)(nil),
					testify_mock.Anything,
					"",
					"",
					"[{\"denom\":\"Denom\",\"amount\":\"100\"}]",
				).Return(mockExecResult, nil)

				mockClient.On("Account", "Sender").Return(
					&cosmosapp.Account{
						Type:    "AccountType",
						Address: "Sender",
						MaybePubkey: &cosmosapp.PubKey{
							Type: "PubKeyType",
							Key:  "Key",
						},
						AccountNumber: "",
						Sequence:      "",
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

				mockTx.On(
					"Exec",
					"INSERT INTO view_accounts (address,account_type,name,pubkey,account_number,sequence_number,balance) VALUES ($1,$2,$3,$4,$5,$6,$7) ON CONFLICT(address) DO UPDATE SET balance = EXCLUDED.balance",
					"Sender",
					"AccountType",
					(*string)(nil),
					testify_mock.Anything,
					"",
					"",
					"[{\"denom\":\"Denom\",\"amount\":\"1000\"}]",
				).Return(mockExecResult, nil)

				mockUpdateLastHandledEventHeight(mockTx)

				return mocks
			},
		},
	}

	for _, tc := range testCases {
		mockRDbConn := NewMockRDbConn()
		mockClient := &cosmosapp.MockClient{}
		mocks := tc.MockFunc(mockRDbConn, mockClient)
		mocks = append(mocks, &mockRDbConn.Mock, &mockClient.Mock)

		projection := NewAccountProjection(mockRDbConn, mockClient)
		err := projection.HandleEvents(1, tc.Events)
		assert.NoError(t, err)

		for _, m := range mocks {
			m.AssertExpectations(t)
		}

		fmt.Println(tc.Name, "Passed")
	}
}

func mockUpdateLastHandledEventHeight(mockTx *test.MockRDbTx) {
	mockRowResult := &test.MockRDbRowResult{}

	var lastHandledEventHeight int64
	mockRowResult.On("Scan", &lastHandledEventHeight).Return(nil)
	mockTx.On(
		"QueryRow",
		"SELECT last_handled_event_height FROM projections WHERE id = $1",
		"Account",
	).Return(mockRowResult, nil)

	mockExecResult := &test.MockRDbExecResult{}
	mockExecResult.On("RowsAffected").Return(int64(1))
	mockTx.On(
		"Exec",
		"UPDATE projections SET id = $1, last_handled_event_height = $2 WHERE id = $3",
		"Account",
		int64(1),
		"Account",
	).Return(mockExecResult, nil)
}
