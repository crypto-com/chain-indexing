package evminnermsgdecoder

type EvmInnerMsgDecoder interface {
	DecodeCosmosMsgFromTxInput([]byte, string) (map[string]interface{}, error)
}
