package tendermint

import "encoding/base64"

type RawBlockResultEventAttributeDecoder struct {
}

func (decoder RawBlockResultEventAttributeDecoder) DecodeKey(s string) (string, error) {
	return s, nil
}

func (decoder RawBlockResultEventAttributeDecoder) DecodeValue(s string) (string, error) {
	return s, nil
}

type Base64BlockResultEventAttributeDecoder struct {
}

func (decoder Base64BlockResultEventAttributeDecoder) DecodeKey(s string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}

func (decoder Base64BlockResultEventAttributeDecoder) DecodeValue(s string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
