package parser

import (
	"fmt"

	"github.com/calvinlauco/cosmostxdecoder"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	jsoniter "github.com/json-iterator/go"
)

type TxDecoder struct {
	decoder *cosmostxdecoder.Decoder

	baseDenom string
}

func NewTxDecoder(baseDenom string) *TxDecoder {
	return &TxDecoder{
		cosmostxdecoder.DefaultDecoder,

		baseDenom,
	}
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

func (decoder *TxDecoder) GetFee(base64Tx string) (coin.Coin, error) {
	tx, err := decoder.Decode(base64Tx)
	if err != nil {
		return coin.Coin{}, fmt.Errorf("error decoding transaction: %v", err)
	}

	return decoder.sumAmount(tx.AuthInfo.Fee.Amount)
}

func (decoder *TxDecoder) sumAmount(amounts []Amount) (coin.Coin, error) {
	var err error

	sum := coin.Zero()
	for _, amount := range amounts {
		if amount.Denom != decoder.baseDenom {
			return coin.Coin{}, fmt.Errorf("unrecognized denom %s when parsing amount", amount.Denom)
		}

		var amountCoin coin.Coin
		amountCoin, err = coin.NewCoinFromString(amount.Amount)
		if err != nil {
			return coin.Coin{}, fmt.Errorf("error parsing amount %s to coin: %v", amount.Amount, err)
		}
		sum, err = sum.Add(amountCoin)
		if err != nil {
			return coin.Coin{}, fmt.Errorf("error adding amount togetehr: %v", err)
		}
	}

	return sum, nil
}

type CosmosTx struct {
	Body       Body     `json:"body"`
	AuthInfo   AuthInfo `json:"auth_info"`
	Signatures []string `json:"signatures"`
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
	PublicKey PublicKey `json:"public_key"`
	ModeInfo  ModeInfo  `json:"mode_info"`
	Sequence  string    `json:"sequence"`
}

type ModeInfo struct {
	Single Single `json:"single"`
}

type Single struct {
	Mode string `json:"mode"`
}

type PublicKey struct {
	Type string `json:"@type"`
	Key  string `json:"key"`
}

type Body struct {
	Messages                    []map[string]interface{} `json:"messages"`
	Memo                        string                   `json:"memo"`
	TimeoutHeight               string                   `json:"timeout_height"`
	ExtensionOptions            []interface{}            `json:"extension_options"`
	NonCriticalExtensionOptions []interface{}            `json:"non_critical_extension_options"`
}

type Message struct {
	Type string `json:"@type"`
}

type MsgSend struct {
	Type        string   `json:"@type"`
	FromAddress string   `json:"from_address"`
	ToAddress   string   `json:"to_address"`
	Amount      []Amount `json:"amount"`
}
