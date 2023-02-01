package txdecoder

import (
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type TxDecoder interface {
	DecodeBase64(string) (*model.CosmosTx, error)
	DeserializeCosmosTx([]byte) (icatypes.CosmosTx, []map[string]interface{}, error)
}
