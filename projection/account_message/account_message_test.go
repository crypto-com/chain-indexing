package account_message_test

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/internal/utctime"
	"github.com/crypto-com/chain-indexing/projection/account_message"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	usecase_model "github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/stretchr/testify/assert"
	testify_mock "github.com/stretchr/testify/mock"
)

func NewAccountMessageProjection(rdbConn rdb.Conn) *account_message.AccountMessage {
	return account_message.NewAccountMessage(
		nil,
		rdbConn,
		"accountAddressPrefix",
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
		MockFunc func(mockConn *test.MockRDbConn) []*testify_mock.Mock
	}{
		{
			Name: "HandleAccountTransferred",
			Events: []entity_event.Event{
				&event_usecase.BlockCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.ACCOUNT_TRANSFERRED,
						Version:     1,
						BlockHeight: 1,
					}),
					Block: &usecase_model.Block{
						Height:          1,
						Hash:            "Hash",
						Time:            utctime.UTCTime{},
						AppHash:         "AppHash",
						ProposerAddress: "ProposerAddress",
						Txs:             nil,
						Signatures:      nil,
						Evidences:       nil,
					},
				},
				&event_usecase.MsgSend{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_SEND,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					FromAddress: "FromAddress",
					ToAddress:   "ToAddress",
					Amount: coin.Coins{
						coin.Coin{
							Denom:  "Denom",
							Amount: coin.NewInt(100),
						},
					},
				},
			},
			MockFunc: func(mockCoon *test.MockRDbConn) (mocks []*testify_mock.Mock) {
				mockTx := NewMockRDbTx()
				mocks = append(mocks, &mockTx.Mock)
				mockCoon.On("Begin").Return(mockTx, nil)

				mockExecResult := &test.MockRDbExecResult{}
				mocks = append(mocks, &mockExecResult.Mock)
				mockExecResult.On("RowsAffected").Return(int64(1))
				mockTx.On(
					"Exec",
					"INSERT INTO view_account_messages_total AS totals (identity,total) VALUES ($1,$2) ON CONFLICT (identity) DO UPDATE SET total = totals.total + EXCLUDED.total",
					"FromAddress:-",
					int64(1),
				).Return(mockExecResult, nil)

				mockTx.On(
					"Exec",
					"INSERT INTO view_account_messages_total AS totals (identity,total) VALUES ($1,$2) ON CONFLICT (identity) DO UPDATE SET total = totals.total + EXCLUDED.total",
					"FromAddress:MsgSend",
					int64(1),
				).Return(mockExecResult, nil)

				mockTx.On(
					"Exec",
					"INSERT INTO view_account_messages_total AS totals (identity,total) VALUES ($1,$2) ON CONFLICT (identity) DO UPDATE SET total = totals.total + EXCLUDED.total",
					"ToAddress:-",
					int64(1),
				).Return(mockExecResult, nil)

				mockTx.On(
					"Exec",
					"INSERT INTO view_account_messages_total AS totals (identity,total) VALUES ($1,$2) ON CONFLICT (identity) DO UPDATE SET total = totals.total + EXCLUDED.total",
					"ToAddress:MsgSend",
					int64(1),
				).Return(mockExecResult, nil)

				mockExecResult2 := &test.MockRDbExecResult{}
				mocks = append(mocks, &mockExecResult2.Mock)
				mockExecResult2.On("RowsAffected").Return(int64(2))

				mockTx.On(
					"Exec",
					"INSERT INTO view_account_messages (block_height,block_hash,block_time,account,transaction_hash,success,message_index,message_type,data) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9),($10,$11,$12,$13,$14,$15,$16,$17,$18)",
					int64(1),
					"Hash",
					int64(0),
					"FromAddress",
					"TxHash",
					true,
					int(0),
					"MsgSend",
					testify_mock.AnythingOfType("string"),
					int64(1),
					"Hash",
					int64(0),
					"ToAddress",
					"TxHash",
					true,
					int(0),
					"MsgSend",
					testify_mock.AnythingOfType("string"),
				).Return(mockExecResult2, nil)

				mockUpdateLastHandledEventHeight(mockTx)

				return mocks
			},
		},
	}

	for _, tc := range testCases {
		mockRDbConn := NewMockRDbConn()
		mocks := tc.MockFunc(mockRDbConn)
		mocks = append(mocks, &mockRDbConn.Mock)

		projection := NewAccountMessageProjection(mockRDbConn)
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
		"AccountMessage",
	).Return(mockRowResult, nil)

	mockExecResult := &test.MockRDbExecResult{}
	mockExecResult.On("RowsAffected").Return(int64(1))
	mockTx.On(
		"Exec",
		"UPDATE projections SET id = $1, last_handled_event_height = $2 WHERE id = $3",
		"AccountMessage",
		int64(1),
		"AccountMessage",
	).Return(mockExecResult, nil)
}
