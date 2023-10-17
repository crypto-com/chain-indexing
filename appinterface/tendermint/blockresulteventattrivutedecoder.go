package tendermint

type BlockResultEventAttributeDecoder interface {
	DecodeKey(string) (string, error)
	DecodeValue(string) (string, error)
}
