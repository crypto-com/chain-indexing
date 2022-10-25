package nft_test

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	testify_mock "github.com/stretchr/testify/mock"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/projection/nft"
	"github.com/crypto-com/chain-indexing/projection/nft/constants"
	"github.com/crypto-com/chain-indexing/projection/nft/view"
	usecase_event "github.com/crypto-com/chain-indexing/usecase/event"
)

func NewNFTProjection(rdbConn rdb.Conn) *nft.NFT {

	return nft.NewNFT(
		nil,
		rdbConn,
		&nft.Config{
			EnableDrop: false,
		},
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
	mockTx.On("ToHandle").Return(nil).Maybe()
	mockTx.On("Rollback").Return(nil).Maybe()
	mockTx.On("Commit").Return(nil).Maybe()

	return mockTx
}

func TestNFT_HandleEvents(t *testing.T) {
	testCases := []struct {
		Name     string
		Events   []entity_event.Event
		MockFunc func(events []entity_event.Event) []*testify_mock.Mock
	}{
		{
			Name: "HandleMsgNFTIssueDenom",
			Events: []entity_event.Event{
				&usecase_event.MsgNFTIssueDenom{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_NFT_ISSUE_DENOM,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					DenomId:   "DenomId",
					DenomName: "DenomName",
					Schema:    "Schema",
					Sender:    "Sender",
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[0].(*usecase_event.MsgNFTIssueDenom)

				mockDenomsView := &view.MockDenomsView{}
				mocks = append(mocks, &mockDenomsView.Mock)
				mockDenomsView.
					On("Insert", &view.DenomRow{
						DenomId:              "DenomId",
						Name:                 "DenomName",
						Schema:               "Schema",
						Creator:              "Sender",
						CreatedAt:            utctime.UTCTime{},
						CreatedAtBlockHeight: 1,
					}).
					Return(nil)

				nft.NewDenoms = func(handle *rdb.Handle) view.Denoms {
					return mockDenomsView
				}

				mockDenomsTotalView := &view.MockDenomsTotalView{}
				mocks = append(mocks, &mockDenomsTotalView.Mock)
				mockDenomsTotalView.
					On("Increment", "-", int64(1)).
					Return(nil)
				mockDenomsTotalView.
					On("Increment", "DenomId", int64(1)).
					Return(nil)

				nft.NewDenomsTotal = func(handle *rdb.Handle) view.DenomsTotal {
					return mockDenomsTotalView
				}

				mockMessagesTotalView := &view.MockMessagesTotalView{}
				mocks = append(mocks, &mockMessagesTotalView.Mock)
				mockMessagesTotalView.
					On("IncrementAll", []string{
						"-:-:-:-",
						fmt.Sprintf("%s:-:-:-", ""),
						fmt.Sprintf("-:%s:-:-", "DenomId"),
						fmt.Sprintf("-:-:%s:-", ""),
						fmt.Sprintf("-:-:-:%s", "/chainmain.nft.v1.MsgIssueDenom"),
						fmt.Sprintf("%s:%s:-:-", "", "DenomId"),
						fmt.Sprintf("%s:-:%s:-", "", ""),
						fmt.Sprintf("%s:-:-:%s", "", "/chainmain.nft.v1.MsgIssueDenom"),
						fmt.Sprintf("-:%s:%s:-", "DenomId", ""),
						fmt.Sprintf("-:%s:-:%s", "DenomId", "/chainmain.nft.v1.MsgIssueDenom"),
						fmt.Sprintf("-:-:%s:%s", "", "/chainmain.nft.v1.MsgIssueDenom"),
						fmt.Sprintf("%s:%s:%s:-", "", "DenomId", ""),
						fmt.Sprintf("%s:%s:-:%s", "", "DenomId", "/chainmain.nft.v1.MsgIssueDenom"),
						fmt.Sprintf("%s:-:%s:%s", "", "", "/chainmain.nft.v1.MsgIssueDenom"),
						fmt.Sprintf("-:%s:%s:%s", "DenomId", "", "/chainmain.nft.v1.MsgIssueDenom"),
						fmt.Sprintf("%s:%s:%s:%s", "", "DenomId", "", "/chainmain.nft.v1.MsgIssueDenom"),
					}, int64(1)).
					Return(nil)

				nft.NewMessagesTotal = func(handle *rdb.Handle) view.MessagesTotal {
					return mockMessagesTotalView
				}

				mockMessagesView := &view.MockMessagesView{}
				mocks = append(mocks, &mockMessagesView.Mock)
				mockMessagesView.
					On("Insert", &view.MessageRow{
						BlockHeight:     1,
						BlockHash:       "",
						BlockTime:       utctime.UTCTime{},
						DenomId:         "DenomId",
						MaybeTokenId:    nil,
						MaybeDrop:       nil,
						TransactionHash: "TxHash",
						Success:         true,
						MessageIndex:    0,
						MessageType:     "/chainmain.nft.v1.MsgIssueDenom",
						Data:            typedEvent,
						Status:          constants.MINTED,
					}).
					Return(nil)

				nft.NewMessages = func(handle *rdb.Handle) view.Messages {
					return mockMessagesView
				}

				nft.UpdateLastHandledEventHeight = func(_ *nft.NFT, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgNFTMintNFT",
			Events: []entity_event.Event{
				&usecase_event.MsgNFTMintNFT{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_NFT_MINT_NFT,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					DenomId:   "DenomId",
					TokenId:   "TokenId",
					TokenName: "TokenName",
					URI:       "URI",
					Data:      "Data",
					Sender:    "Sender",
					Recipient: "Recipient",
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[0].(*usecase_event.MsgNFTMintNFT)

				mockTokensView := &view.MockTokensView{}
				mocks = append(mocks, &mockTokensView.Mock)
				mockTokensView.
					On("Insert", &view.TokenRow{
						TokenId:                      "TokenId",
						DenomId:                      "DenomId",
						MaybeDrop:                    nil,
						Name:                         "TokenName",
						URI:                          "URI",
						Data:                         "Data",
						Minter:                       "Sender",
						Owner:                        "Recipient",
						MintedAt:                     utctime.UTCTime{},
						MintedAtBlockHeight:          1,
						LastEditedAt:                 utctime.UTCTime{},
						LastEditedAtBlockHeight:      1,
						LastTransferredAt:            utctime.UTCTime{},
						LastTransferredAtBlockHeight: 1,
						Status:                       constants.MINTED,
					}).
					Return(nil)

				nft.NewTokens = func(handle *rdb.Handle) view.Tokens {
					return mockTokensView
				}

				mockTokensTotalView := &view.MockTokensTotalView{}
				mocks = append(mocks, &mockTokensTotalView.Mock)
				mockTokensTotalView.
					On("IncrementAll",
						[]string{
							"-:-:-:-",
							fmt.Sprintf("%s:-:-:-", "DenomId"),
							fmt.Sprintf("-:%s:-:-", ""),
							fmt.Sprintf("-:-:%s:-", "Sender"),
							fmt.Sprintf("-:-:-:%s", "Recipient"),
							fmt.Sprintf("%s:%s:-:-", "DenomId", ""),
							fmt.Sprintf("%s:-:%s:-", "DenomId", "Sender"),
							fmt.Sprintf("%s:-:-:%s", "DenomId", "Recipient"),
							fmt.Sprintf("-:%s:%s:-", "", "Sender"),
							fmt.Sprintf("-:%s:-:%s", "", "Recipient"),
							fmt.Sprintf("-:-:%s:%s", "Sender", "Recipient"),
							fmt.Sprintf("%s:%s:%s:-", "DenomId", "", "Sender"),
							fmt.Sprintf("%s:%s:-:%s", "DenomId", "", "Recipient"),
							fmt.Sprintf("%s:-:%s:%s", "DenomId", "Sender", "Recipient"),
							fmt.Sprintf("-:%s:%s:%s", "", "Sender", "Recipient"),
							fmt.Sprintf("%s:%s:%s:%s", "DenomId", "", "Sender", "Recipient"),
						}, int64(1)).
					Return(nil)

				nft.NewTokensTotal = func(handle *rdb.Handle) view.TokensTotal {
					return mockTokensTotalView
				}

				mockMessagesTotalView := &view.MockMessagesTotalView{}
				mocks = append(mocks, &mockMessagesTotalView.Mock)
				mockMessagesTotalView.
					On("IncrementAll", []string{
						"-:-:-:-",
						fmt.Sprintf("%s:-:-:-", ""),
						fmt.Sprintf("-:%s:-:-", "DenomId"),
						fmt.Sprintf("-:-:%s:-", "TokenId"),
						fmt.Sprintf("-:-:-:%s", "/chainmain.nft.v1.MsgMintNFT"),
						fmt.Sprintf("%s:%s:-:-", "", "DenomId"),
						fmt.Sprintf("%s:-:%s:-", "", "TokenId"),
						fmt.Sprintf("%s:-:-:%s", "", "/chainmain.nft.v1.MsgMintNFT"),
						fmt.Sprintf("-:%s:%s:-", "DenomId", "TokenId"),
						fmt.Sprintf("-:%s:-:%s", "DenomId", "/chainmain.nft.v1.MsgMintNFT"),
						fmt.Sprintf("-:-:%s:%s", "TokenId", "/chainmain.nft.v1.MsgMintNFT"),
						fmt.Sprintf("%s:%s:%s:-", "", "DenomId", "TokenId"),
						fmt.Sprintf("%s:%s:-:%s", "", "DenomId", "/chainmain.nft.v1.MsgMintNFT"),
						fmt.Sprintf("%s:-:%s:%s", "", "TokenId", "/chainmain.nft.v1.MsgMintNFT"),
						fmt.Sprintf("-:%s:%s:%s", "DenomId", "TokenId", "/chainmain.nft.v1.MsgMintNFT"),
						fmt.Sprintf("%s:%s:%s:%s", "", "DenomId", "TokenId", "/chainmain.nft.v1.MsgMintNFT"),
					}, int64(1)).
					Return(nil)

				nft.NewMessagesTotal = func(handle *rdb.Handle) view.MessagesTotal {
					return mockMessagesTotalView
				}

				mockMessagesView := &view.MockMessagesView{}
				mocks = append(mocks, &mockMessagesView.Mock)
				mockMessagesView.
					On("Insert", &view.MessageRow{
						BlockHeight:     1,
						BlockHash:       "",
						BlockTime:       utctime.UTCTime{},
						DenomId:         "DenomId",
						MaybeTokenId:    primptr.String("TokenId"),
						MaybeDrop:       nil,
						TransactionHash: "TxHash",
						Success:         true,
						MessageIndex:    0,
						MessageType:     "/chainmain.nft.v1.MsgMintNFT",
						Data:            typedEvent,
						Status:          constants.MINTED,
					}).
					Return(nil)

				nft.NewMessages = func(handle *rdb.Handle) view.Messages {
					return mockMessagesView
				}

				nft.UpdateLastHandledEventHeight = func(_ *nft.NFT, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgNFTEditNFT",
			Events: []entity_event.Event{
				&usecase_event.MsgNFTEditNFT{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_NFT_EDIT_NFT,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					DenomId:   "DenomId",
					TokenId:   "TokenId",
					TokenName: "TokenName",
					URI:       "URI",
					Data:      "Data",
					Sender:    "Sender",
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[0].(*usecase_event.MsgNFTEditNFT)

				mockTokensView := &view.MockTokensView{}
				mocks = append(mocks, &mockTokensView.Mock)

				mockTokensView.
					On("FindById", "DenomId", "TokenId").
					Return(&view.TokenRowWithDenomname{
						TokenRow: view.TokenRow{
							DenomId:                      "DenomId",
							TokenId:                      "TokenId",
							MaybeDrop:                    primptr.StringNil(),
							Name:                         "PrevName",
							URI:                          "PrevURI",
							Data:                         "PrevData",
							Minter:                       "PrevMinter",
							Owner:                        "PrevOwner",
							MintedAt:                     utctime.FromUnixNano(-1),
							MintedAtBlockHeight:          0,
							LastEditedAt:                 utctime.FromUnixNano(-1),
							LastEditedAtBlockHeight:      0,
							LastTransferredAt:            utctime.FromUnixNano(-1),
							LastTransferredAtBlockHeight: 0,
							Status:                       constants.MINTED,
						},
						DenomName:   "DenomName",
						DenomSchema: "DenomSchema",
					}, nil)

				mockTokensView.
					On("Update", view.TokenRow{
						TokenId:                      "TokenId",
						DenomId:                      "DenomId",
						MaybeDrop:                    nil,
						Name:                         "TokenName",
						URI:                          "URI",
						Data:                         "Data",
						Minter:                       "PrevMinter",
						Owner:                        "PrevOwner",
						MintedAt:                     utctime.FromUnixNano(-1),
						MintedAtBlockHeight:          0,
						LastEditedAt:                 utctime.UTCTime{},
						LastEditedAtBlockHeight:      1,
						LastTransferredAt:            utctime.FromUnixNano(-1),
						LastTransferredAtBlockHeight: 0,
					}).
					Return(nil)

				nft.NewTokens = func(handle *rdb.Handle) view.Tokens {
					return mockTokensView
				}

				mockTokensTotalView := &view.MockTokensTotalView{}
				mocks = append(mocks, &mockTokensTotalView.Mock)
				mockTokensTotalView.
					On("DecrementAll",
						[]string{
							fmt.Sprintf("-:-:-:%s", "PrevOwner"),
							fmt.Sprintf("%s:-:-:%s", "DenomId", "PrevOwner"),
							fmt.Sprintf("-:%s:-:%s", "", "PrevOwner"),
							fmt.Sprintf("-:-:%s:%s", "PrevMinter", "PrevOwner"),
							fmt.Sprintf("%s:%s:-:%s", "DenomId", "", "PrevOwner"),
							fmt.Sprintf("%s:-:%s:%s", "DenomId", "PrevMinter", "PrevOwner"),
							fmt.Sprintf("-:%s:%s:%s", "", "PrevMinter", "PrevOwner"),
							fmt.Sprintf("%s:%s:%s:%s", "DenomId", "", "PrevMinter", "PrevOwner"),
						}, int64(1)).
					Return(nil)
				mockTokensTotalView.
					On("IncrementAll",
						[]string{
							fmt.Sprintf("-:-:-:%s", "PrevOwner"),
							fmt.Sprintf("%s:-:-:%s", "DenomId", "PrevOwner"),
							fmt.Sprintf("-:%s:-:%s", "", "PrevOwner"),
							fmt.Sprintf("-:-:%s:%s", "PrevMinter", "PrevOwner"),
							fmt.Sprintf("%s:%s:-:%s", "DenomId", "", "PrevOwner"),
							fmt.Sprintf("%s:-:%s:%s", "DenomId", "PrevMinter", "PrevOwner"),
							fmt.Sprintf("-:%s:%s:%s", "", "PrevMinter", "PrevOwner"),
							fmt.Sprintf("%s:%s:%s:%s", "DenomId", "", "PrevMinter", "PrevOwner"),
						}, int64(1)).
					Return(nil)

				nft.NewTokensTotal = func(handle *rdb.Handle) view.TokensTotal {
					return mockTokensTotalView
				}

				mockMessagesTotalView := &view.MockMessagesTotalView{}
				mocks = append(mocks, &mockMessagesTotalView.Mock)
				mockMessagesTotalView.
					On("IncrementAll", []string{
						"-:-:-:-",
						fmt.Sprintf("%s:-:-:-", ""),
						fmt.Sprintf("-:%s:-:-", "DenomId"),
						fmt.Sprintf("-:-:%s:-", "TokenId"),
						fmt.Sprintf("-:-:-:%s", "/chainmain.nft.v1.MsgEditNFT"),
						fmt.Sprintf("%s:%s:-:-", "", "DenomId"),
						fmt.Sprintf("%s:-:%s:-", "", "TokenId"),
						fmt.Sprintf("%s:-:-:%s", "", "/chainmain.nft.v1.MsgEditNFT"),
						fmt.Sprintf("-:%s:%s:-", "DenomId", "TokenId"),
						fmt.Sprintf("-:%s:-:%s", "DenomId", "/chainmain.nft.v1.MsgEditNFT"),
						fmt.Sprintf("-:-:%s:%s", "TokenId", "/chainmain.nft.v1.MsgEditNFT"),
						fmt.Sprintf("%s:%s:%s:-", "", "DenomId", "TokenId"),
						fmt.Sprintf("%s:%s:-:%s", "", "DenomId", "/chainmain.nft.v1.MsgEditNFT"),
						fmt.Sprintf("%s:-:%s:%s", "", "TokenId", "/chainmain.nft.v1.MsgEditNFT"),
						fmt.Sprintf("-:%s:%s:%s", "DenomId", "TokenId", "/chainmain.nft.v1.MsgEditNFT"),
						fmt.Sprintf("%s:%s:%s:%s", "", "DenomId", "TokenId", "/chainmain.nft.v1.MsgEditNFT"),
					}, int64(1)).
					Return(nil)

				nft.NewMessagesTotal = func(handle *rdb.Handle) view.MessagesTotal {
					return mockMessagesTotalView
				}

				mockMessagesView := &view.MockMessagesView{}
				mocks = append(mocks, &mockMessagesView.Mock)
				mockMessagesView.
					On("Insert", &view.MessageRow{
						BlockHeight:     1,
						BlockHash:       "",
						BlockTime:       utctime.UTCTime{},
						DenomId:         "DenomId",
						MaybeTokenId:    primptr.String("TokenId"),
						MaybeDrop:       nil,
						TransactionHash: "TxHash",
						Success:         true,
						MessageIndex:    0,
						MessageType:     "/chainmain.nft.v1.MsgEditNFT",
						Data:            typedEvent,
						Status:          constants.MINTED,
					}).
					Return(nil)

				nft.NewMessages = func(handle *rdb.Handle) view.Messages {
					return mockMessagesView
				}

				nft.UpdateLastHandledEventHeight = func(_ *nft.NFT, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgNFTBurnNFT",
			Events: []entity_event.Event{
				&usecase_event.MsgNFTBurnNFT{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_NFT_BURN_NFT,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					DenomId: "DenomId",
					TokenId: "TokenId",
					Sender:  "Sender",
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[0].(*usecase_event.MsgNFTBurnNFT)

				mockTokensView := &view.MockTokensView{}
				mocks = append(mocks, &mockTokensView.Mock)

				mockTokensView.
					On("FindById", "DenomId", "TokenId").
					Return(&view.TokenRowWithDenomname{
						TokenRow: view.TokenRow{
							DenomId:                      "DenomId",
							TokenId:                      "TokenId",
							MaybeDrop:                    primptr.StringNil(),
							Name:                         "Name",
							URI:                          "URI",
							Data:                         "Data",
							Minter:                       "Minter",
							Owner:                        "Owner",
							MintedAt:                     utctime.FromUnixNano(-1),
							MintedAtBlockHeight:          0,
							LastEditedAt:                 utctime.FromUnixNano(-1),
							LastEditedAtBlockHeight:      0,
							LastTransferredAt:            utctime.FromUnixNano(-1),
							LastTransferredAtBlockHeight: 0,
							Status:                       constants.MINTED,
						},
						DenomName:   "DenomName",
						DenomSchema: "DenomSchema",
					}, nil)

				mockTokensView.
					On("SoftDelete", "DenomId", "TokenId").
					Return(nil)

				nft.NewTokens = func(handle *rdb.Handle) view.Tokens {
					return mockTokensView
				}

				mockMessagesView := &view.MockMessagesView{}
				mocks = append(mocks, &mockMessagesView.Mock)
				mockMessagesView.
					On("SoftDelete", "DenomId", "TokenId").
					Return(int64(1), nil)

				mocks = append(mocks, &mockMessagesView.Mock)
				mockMessagesView.
					On("Insert", &view.MessageRow{
						BlockHeight:     1,
						BlockHash:       "",
						BlockTime:       utctime.UTCTime{},
						DenomId:         "DenomId",
						MaybeTokenId:    primptr.String("TokenId"),
						MaybeDrop:       nil,
						TransactionHash: "TxHash",
						Success:         true,
						MessageIndex:    0,
						MessageType:     "/chainmain.nft.v1.MsgBurnNFT",
						Data:            typedEvent,
						Status:          constants.BURNED,
					}).
					Return(nil)

				nft.NewMessages = func(handle *rdb.Handle) view.Messages {
					return mockMessagesView
				}

				mockTokensTotalView := &view.MockTokensTotalView{}
				mocks = append(mocks, &mockTokensTotalView.Mock)
				mockTokensTotalView.
					On("DecrementAll",
						[]string{
							"-:-:-:-",
							fmt.Sprintf("%s:-:-:-", "DenomId"),
							fmt.Sprintf("-:%s:-:-", ""),
							fmt.Sprintf("-:-:%s:-", "Minter"),
							fmt.Sprintf("-:-:-:%s", "Owner"),
							fmt.Sprintf("%s:%s:-:-", "DenomId", ""),
							fmt.Sprintf("%s:-:%s:-", "DenomId", "Minter"),
							fmt.Sprintf("%s:-:-:%s", "DenomId", "Owner"),
							fmt.Sprintf("-:%s:%s:-", "", "Minter"),
							fmt.Sprintf("-:%s:-:%s", "", "Owner"),
							fmt.Sprintf("-:-:%s:%s", "Minter", "Owner"),
							fmt.Sprintf("%s:%s:%s:-", "DenomId", "", "Minter"),
							fmt.Sprintf("%s:%s:-:%s", "DenomId", "", "Owner"),
							fmt.Sprintf("%s:-:%s:%s", "DenomId", "Minter", "Owner"),
							fmt.Sprintf("-:%s:%s:%s", "", "Minter", "Owner"),
							fmt.Sprintf("%s:%s:%s:%s", "DenomId", "", "Minter", "Owner"),
						}, int64(1)).
					Return(nil)

				nft.NewTokensTotal = func(handle *rdb.Handle) view.TokensTotal {
					return mockTokensTotalView
				}

				nft.UpdateLastHandledEventHeight = func(_ *nft.NFT, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgNFTTransferNFT",
			Events: []entity_event.Event{
				&usecase_event.MsgNFTTransferNFT{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_NFT_TRANSFER_NFT,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					DenomId:   "DenomId",
					TokenId:   "TokenId",
					Sender:    "Sender",
					Recipient: "Recipient",
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[0].(*usecase_event.MsgNFTTransferNFT)

				mockTokensView := &view.MockTokensView{}
				mocks = append(mocks, &mockTokensView.Mock)

				mockTokensView.
					On("FindById", "DenomId", "TokenId").
					Return(&view.TokenRowWithDenomname{
						TokenRow: view.TokenRow{
							DenomId:                      "DenomId",
							TokenId:                      "TokenId",
							MaybeDrop:                    primptr.StringNil(),
							Name:                         "PrevName",
							URI:                          "PrevURI",
							Data:                         "PrevData",
							Minter:                       "PrevMinter",
							Owner:                        "PrevOwner",
							MintedAt:                     utctime.FromUnixNano(-1),
							MintedAtBlockHeight:          0,
							LastEditedAt:                 utctime.FromUnixNano(-1),
							LastEditedAtBlockHeight:      0,
							LastTransferredAt:            utctime.FromUnixNano(-1),
							LastTransferredAtBlockHeight: 0,
							Status:                       constants.MINTED,
						},
						DenomName:   "DenomName",
						DenomSchema: "DenomSchema",
					}, nil)

				mockTokensView.
					On("Update", view.TokenRow{
						TokenId:                      "TokenId",
						DenomId:                      "DenomId",
						MaybeDrop:                    nil,
						Name:                         "PrevName",
						URI:                          "PrevURI",
						Data:                         "PrevData",
						Minter:                       "PrevMinter",
						Owner:                        "Recipient",
						MintedAt:                     utctime.FromUnixNano(-1),
						MintedAtBlockHeight:          0,
						LastEditedAt:                 utctime.FromUnixNano(-1),
						LastEditedAtBlockHeight:      0,
						LastTransferredAt:            utctime.UTCTime{},
						LastTransferredAtBlockHeight: 1,
					}).
					Return(nil)

				nft.NewTokens = func(handle *rdb.Handle) view.Tokens {
					return mockTokensView
				}

				mockTokensTotalView := &view.MockTokensTotalView{}
				mocks = append(mocks, &mockTokensTotalView.Mock)
				mockTokensTotalView.
					On("DecrementAll",
						[]string{
							fmt.Sprintf("-:-:-:%s", "PrevOwner"),
							fmt.Sprintf("%s:-:-:%s", "DenomId", "PrevOwner"),
							fmt.Sprintf("-:%s:-:%s", "", "PrevOwner"),
							fmt.Sprintf("-:-:%s:%s", "PrevMinter", "PrevOwner"),
							fmt.Sprintf("%s:%s:-:%s", "DenomId", "", "PrevOwner"),
							fmt.Sprintf("%s:-:%s:%s", "DenomId", "PrevMinter", "PrevOwner"),
							fmt.Sprintf("-:%s:%s:%s", "", "PrevMinter", "PrevOwner"),
							fmt.Sprintf("%s:%s:%s:%s", "DenomId", "", "PrevMinter", "PrevOwner"),
						}, int64(1)).
					Return(nil)
				mockTokensTotalView.
					On("IncrementAll",
						[]string{
							fmt.Sprintf("-:-:-:%s", "Recipient"),
							fmt.Sprintf("%s:-:-:%s", "DenomId", "Recipient"),
							fmt.Sprintf("-:%s:-:%s", "", "Recipient"),
							fmt.Sprintf("-:-:%s:%s", "PrevMinter", "Recipient"),
							fmt.Sprintf("%s:%s:-:%s", "DenomId", "", "Recipient"),
							fmt.Sprintf("%s:-:%s:%s", "DenomId", "PrevMinter", "Recipient"),
							fmt.Sprintf("-:%s:%s:%s", "", "PrevMinter", "Recipient"),
							fmt.Sprintf("%s:%s:%s:%s", "DenomId", "", "PrevMinter", "Recipient"),
						}, int64(1)).
					Return(nil)

				nft.NewTokensTotal = func(handle *rdb.Handle) view.TokensTotal {
					return mockTokensTotalView
				}

				mockMessagesTotalView := &view.MockMessagesTotalView{}
				mocks = append(mocks, &mockMessagesTotalView.Mock)
				mockMessagesTotalView.
					On("IncrementAll", []string{
						"-:-:-:-",
						fmt.Sprintf("%s:-:-:-", ""),
						fmt.Sprintf("-:%s:-:-", "DenomId"),
						fmt.Sprintf("-:-:%s:-", "TokenId"),
						fmt.Sprintf("-:-:-:%s", "/chainmain.nft.v1.MsgTransferNFT"),
						fmt.Sprintf("%s:%s:-:-", "", "DenomId"),
						fmt.Sprintf("%s:-:%s:-", "", "TokenId"),
						fmt.Sprintf("%s:-:-:%s", "", "/chainmain.nft.v1.MsgTransferNFT"),
						fmt.Sprintf("-:%s:%s:-", "DenomId", "TokenId"),
						fmt.Sprintf("-:%s:-:%s", "DenomId", "/chainmain.nft.v1.MsgTransferNFT"),
						fmt.Sprintf("-:-:%s:%s", "TokenId", "/chainmain.nft.v1.MsgTransferNFT"),
						fmt.Sprintf("%s:%s:%s:-", "", "DenomId", "TokenId"),
						fmt.Sprintf("%s:%s:-:%s", "", "DenomId", "/chainmain.nft.v1.MsgTransferNFT"),
						fmt.Sprintf("%s:-:%s:%s", "", "TokenId", "/chainmain.nft.v1.MsgTransferNFT"),
						fmt.Sprintf("-:%s:%s:%s", "DenomId", "TokenId", "/chainmain.nft.v1.MsgTransferNFT"),
						fmt.Sprintf("%s:%s:%s:%s", "", "DenomId", "TokenId", "/chainmain.nft.v1.MsgTransferNFT"),
					}, int64(1)).
					Return(nil)

				nft.NewMessagesTotal = func(handle *rdb.Handle) view.MessagesTotal {
					return mockMessagesTotalView
				}

				mockMessagesView := &view.MockMessagesView{}
				mocks = append(mocks, &mockMessagesView.Mock)
				mockMessagesView.
					On("Insert", &view.MessageRow{
						BlockHeight:     1,
						BlockHash:       "",
						BlockTime:       utctime.UTCTime{},
						DenomId:         "DenomId",
						MaybeTokenId:    primptr.String("TokenId"),
						MaybeDrop:       nil,
						TransactionHash: "TxHash",
						Success:         true,
						MessageIndex:    0,
						MessageType:     "/chainmain.nft.v1.MsgTransferNFT",
						Data:            typedEvent,
						Status:          constants.MINTED,
					}).
					Return(nil)

				nft.NewMessages = func(handle *rdb.Handle) view.Messages {
					return mockMessagesView
				}

				nft.UpdateLastHandledEventHeight = func(_ *nft.NFT, _ *rdb.Handle, _ int64) error {
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
		mocks = append(mocks, &mockRDbConn.Mock)
		mocks = append(mocks, &mockTx.Mock)

		projection := NewNFTProjection(mockRDbConn)
		err := projection.HandleEvents(1, tc.Events)
		assert.NoError(t, err)

		for _, m := range mocks {
			m.AssertExpectations(t)
		}

		fmt.Println(tc.Name, "Passed")
	}
}
