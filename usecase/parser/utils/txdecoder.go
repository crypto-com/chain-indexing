package utils

import (
	"fmt"

	"github.com/calvinlauyh/cosmosutils"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibctransfertypes "github.com/cosmos/ibc-go/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	ibcconnectiontypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	ibccommitmenttypes "github.com/cosmos/ibc-go/modules/core/23-commitment/types"
	ibctypes "github.com/cosmos/ibc-go/modules/core/types"
	"github.com/crypto-com/chain-indexing/usecase/model"
	nfttypes "github.com/crypto-org-chain/chain-main/v3/x/nft/types"
	cronostypes "github.com/crypto-org-chain/cronos/x/cronos/types"
	liquiditytypes "github.com/gravity-devs/liquidity/x/liquidity/types"
	jsoniter "github.com/json-iterator/go"
	gravitytypes "github.com/peggyjv/gravity-bridge/module/x/gravity/types"
	ethermintcryptotypes "github.com/tharsis/ethermint/crypto/codec"
	ethereminttypes "github.com/tharsis/ethermint/types"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"

	"github.com/crypto-com/chain-indexing/usecase/coin"
)

type TxDecoder struct {
	decoder *cosmosutils.Decoder
}

func NewTxDecoder() *TxDecoder {
	return &TxDecoder{
		cosmosutils.NewDecoder().RegisterInterfaces(RegisterDecoderInterfaces),
	}
}

func RegisterDecoderInterfaces(interfaceRegistry types.InterfaceRegistry) {
	std.RegisterInterfaces(interfaceRegistry)

	authtypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	crisistypes.RegisterInterfaces(interfaceRegistry)
	distributiontypes.RegisterInterfaces(interfaceRegistry)
	evidencetypes.RegisterInterfaces(interfaceRegistry)
	govtypes.RegisterInterfaces(interfaceRegistry)
	proposal.RegisterInterfaces(interfaceRegistry)
	slashingtypes.RegisterInterfaces(interfaceRegistry)
	stakingtypes.RegisterInterfaces(interfaceRegistry)
	upgradetypes.RegisterInterfaces(interfaceRegistry)
	vestingtypes.RegisterInterfaces(interfaceRegistry)
	ibctypes.RegisterInterfaces(interfaceRegistry)
	ibcclienttypes.RegisterInterfaces(interfaceRegistry)
	ibctransfertypes.RegisterInterfaces(interfaceRegistry)
	ibcconnectiontypes.RegisterInterfaces(interfaceRegistry)
	ibcchanneltypes.RegisterInterfaces(interfaceRegistry)
	ibccommitmenttypes.RegisterInterfaces(interfaceRegistry)
	nfttypes.RegisterInterfaces(interfaceRegistry)
	authztypes.RegisterInterfaces(interfaceRegistry)
	feegranttypes.RegisterInterfaces(interfaceRegistry)
	ethereminttypes.RegisterInterfaces(interfaceRegistry)
	ethermintcryptotypes.RegisterInterfaces(interfaceRegistry)
	evmtypes.RegisterInterfaces(interfaceRegistry)
	gravitytypes.RegisterInterfaces(interfaceRegistry)
	cronostypes.RegisterInterfaces(interfaceRegistry)
	liquiditytypes.RegisterInterfaces(interfaceRegistry)

	// FIXME
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil),
		&cronostypes.MsgConvertVouchers{},
		&cronostypes.MsgTransferTokens{},
	)
}

func (decoder *TxDecoder) Decode(base64Tx string) (*CosmosTx, error) {
	rawTx, err := decoder.decoder.DecodeBase64(base64Tx)
	if err != nil {
		return nil, fmt.Errorf("error decoding transaction: %v", err)
	}

	txJSONBytes, err := rawTx.MarshalToJSON()
	if err != nil {
		return nil, fmt.Errorf("error encoding decoded transaction to JSON: %v", err)
	}

	var tx *CosmosTx
	if err := jsoniter.Unmarshal(txJSONBytes, &tx); err != nil {
		return nil, fmt.Errorf("error decoding transaction JSON: %v", err)
	}

	return tx, nil
}

func (decoder *TxDecoder) GetFee(base64Tx string) (coin.Coins, error) {
	tx, err := decoder.Decode(base64Tx)
	if err != nil {
		return nil, fmt.Errorf("error decoding transaction: %v", err)
	}

	return decoder.sumAmount(tx.AuthInfo.Fee.Amount)
}

func (decoder *TxDecoder) sumAmount(amounts []Amount) (coin.Coins, error) {
	var err error

	coins := coin.NewEmptyCoins()
	for _, amount := range amounts {
		var amountCoin coin.Coin
		amountCoin, err = coin.NewCoinFromString(amount.Denom, amount.Amount)
		if err != nil {
			return nil, fmt.Errorf("error parsing amount %s to Coin: %v", amount.Amount, err)
		}
		coins = coins.Add(amountCoin)
	}

	return coins, nil
}

type CosmosTx struct {
	Body       Body     `json:"body"`
	AuthInfo   AuthInfo `json:"auth_info"`
	Signatures []string `json:"signatures"`
}

type Body struct {
	Messages                    []map[string]interface{} `json:"messages"`
	Memo                        string                   `json:"memo"`
	TimeoutHeight               string                   `json:"timeout_height"`
	ExtensionOptions            []interface{}            `json:"extension_options"`
	NonCriticalExtensionOptions []interface{}            `json:"non_critical_extension_options"`
}

type AuthInfo struct {
	SignerInfos []SignerInfo `json:"signer_infos"`
	Fee         Fee          `json:"fee"`
}

type Fee struct {
	Amount   []Amount `json:"amount"`
	GasLimit string   `json:"gas_limit"`
	Payer    string   `json:"payer"`
	Granter  string   `json:"granter"`
}

type Amount struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type SignerInfo struct {
	MaybePublicKey *SignerInfoPublicKey `json:"public_key"`
	ModeInfo       ModeInfo             `json:"mode_info"`
	Sequence       string               `json:"sequence"`
}

type SignerInfoPublicKey struct {
	Type            string      `json:"@type"`
	MaybeThreshold  *int        `json:"threshold,omitempty"`
	MaybePublicKeys []PublicKey `json:"public_keys,omitempty"`
	MaybeKey        *string     `json:"key,omitempty"`
}

type PublicKey struct {
	Type string `json:"@type"`
	Key  string `json:"key"`
}

type ModeInfo struct {
	MaybeSingle *Single `json:"single,omitempty"`
	MaybeMulti  *Multi  `json:"multi,omitempty"`
}

type Single struct {
	Mode string `json:"mode"`
}

type Multi struct {
	Bitarray  Bitarray         `json:"bitarray"`
	ModeInfos []SingleModeInfo `json:"mode_infos"`
}

type SingleModeInfo struct {
	Single Single `json:"single"`
}

type Bitarray struct {
	ExtraBitsStored int64  `json:"extra_bits_stored"`
	Elems           string `json:"elems"`
}

func SumAmount(amounts []model.CosmosTxAuthInfoFeeAmount) (coin.Coins, error) {
	var err error

	coins := coin.NewEmptyCoins()
	for _, amount := range amounts {
		var amountCoin coin.Coin
		amountCoin, err = coin.NewCoinFromString(amount.Denom, amount.Amount)
		if err != nil {
			return nil, fmt.Errorf("error parsing amount %s to Coin: %v", amount.Amount, err)
		}
		coins = coins.Add(amountCoin)
	}

	return coins, nil
}
