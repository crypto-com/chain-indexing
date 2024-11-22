package ethereumtxinnermsgdecoder

type EthereumTxInnerMsgDecoder interface {
	DecodeCosmosMsgFromTxInput([]byte, string) (map[string]interface{}, error)
}
