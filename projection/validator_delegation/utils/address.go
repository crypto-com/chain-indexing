package utils

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/bech32"

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

func GetTendermintAddress(tendermintPubKeyInString string) (string, error) {
	tendermintPubKey, err := base64.StdEncoding.DecodeString(tendermintPubKeyInString)
	if err != nil {
		return "", fmt.Errorf("error base64 decoding Tendermint node pubkey: %v", err)
	}

	return tmcosmosutils.TmAddressFromTmPubKey(tendermintPubKey), nil
}

func IsValAddrEqualsDelAddr(
	validatorAddress string,
	delegatorAddress string,
	validatorAddressPrefix string,
	accountAddressPrefix string,
) (bool, error) {

	valAddr, err := ValAddressFromBech32(validatorAddress, validatorAddressPrefix) // OperatorAddress
	if err != nil {
		return false, fmt.Errorf("error in ValAddressFromBech32: %v", err)
	}
	delAddr, err := AccAddressFromBech32(delegatorAddress, accountAddressPrefix)
	if err != nil {
		return false, fmt.Errorf("error in AccAddressFromBech32: %v", err)
	}

	return bytes.Equal(valAddr, delAddr), nil
}

// AccAddressFromBech32 creates an AccAddress from a Bech32 string.
func AccAddressFromBech32(address, bech32PrefixAccAddr string) (addr []byte, err error) {
	if len(strings.TrimSpace(address)) == 0 {
		return nil, errors.New("empty address string is not allowed")
	}

	bz, err := GetFromBech32(address, bech32PrefixAccAddr)
	if err != nil {
		return nil, err
	}

	err = VerifyAddressFormat(bz)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

// ValAddressFromBech32 creates a ValAddress from a Bech32 string.
func ValAddressFromBech32(address, bech32PrefixValAddr string) (addr []byte, err error) {
	if len(strings.TrimSpace(address)) == 0 {
		return nil, errors.New("empty address string is not allowed")
	}

	bz, err := GetFromBech32(address, bech32PrefixValAddr)
	if err != nil {
		return nil, err
	}

	err = VerifyAddressFormat(bz)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

// GetFromBech32 decodes a bytestring from a Bech32 encoded string.
func GetFromBech32(bech32str, prefix string) ([]byte, error) {
	if len(bech32str) == 0 {
		return nil, errors.New("decoding Bech32 address failed: must provide a non empty address")
	}

	hrp, bz, err := bech32.DecodeAndConvert(bech32str)
	if err != nil {
		return nil, err
	}

	if hrp != prefix {
		return nil, fmt.Errorf("invalid Bech32 prefix; expected %s, got %s", prefix, hrp)
	}

	return bz, nil
}

// VerifyAddressFormat verifies that the provided bytes form a valid address
// according to the default address rules or a custom address verifier set by
// GetConfig().SetAddressVerifier()
func VerifyAddressFormat(bz []byte) error {
	// verifier := GetConfig().GetAddressVerifier()
	// if verifier != nil {
	// 	return verifier(bz)
	// }

	if len(bz) == 0 {
		return errors.New("addresses cannot be empty")
	}

	if len(bz) > 255 {
		return fmt.Errorf("address max length is %d, got %d", 255, len(bz))
	}

	return nil
}
