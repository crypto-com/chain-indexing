package txdecoder

import (
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/model/icaauth"
)

type TxDecoder interface {
	DecodeBase64(string) (*model.CosmosTx, error)
	DeserializeCosmosTx([]byte) (icaauth.CosmosTx, []map[string]interface{}, error)
}
