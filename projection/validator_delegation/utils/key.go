package utils

import (
	"encoding/base64"
	"fmt"

	"github.com/crypto-com/chain-indexing/external/tmcosmosutils"
)

func GetConsensusNodeAddress(
	tendermintPubKeyInString string,
	conNodeAddressPrefix string,
) (
	string,
	error,
) {
	tendermintPubKey, err := base64.StdEncoding.DecodeString(tendermintPubKeyInString)
	if err != nil {
		return "", fmt.Errorf("error base64 decoding Tendermint node pubkey: %v", err)
	}

	consensusNodeAddress, err := tmcosmosutils.ConsensusNodeAddressFromTmPubKey(
		conNodeAddressPrefix, tendermintPubKey,
	)
	if err != nil {
		return "", fmt.Errorf("error converting Tendermint node pubkey to consensus node address: %v", err)
	}

	return consensusNodeAddress, nil
}
