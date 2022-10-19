package txdecoder

import (
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type TxDecoder interface {
	DecodeBase64(string) (*model.CosmosTx, error)
}
