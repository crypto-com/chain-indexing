package tmcosmosutils

import (
	"fmt"

	"github.com/tendermint/tendermint/crypto"

	"github.com/btcsuite/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	ed255192 "github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func TmAddressFromTmPubKey(pubKey []byte) string {
	key := make(ed25519.PubKey, ed25519.PubKeySize)
	copy(key, pubKey)

	return key.Address().String()
}

func MustConsensusAddressFromTmPubKey(bech32Prefix string, pubKey []byte) string {
	address, err := ConsensusNodeAddressFromTmPubKey(bech32Prefix, pubKey)
	if err != nil {
		panic(err)
	}

	return address
}

func ConsensusNodeAddressFromTmPubKey(bech32Prefix string, pubKey []byte) (string, error) {
	cosmosPubKey := &ed255192.PubKey{
		Key: pubKey,
	}

	conv, err := bech32.ConvertBits(cosmosPubKey.Address().Bytes(), 8, 5, true)
	if err != nil {
		return "", fmt.Errorf("error converting tendermint public key to bech32 bits: %v", err)
	}
	address, err := bech32.Encode(bech32Prefix, conv)
	if err != nil {
		return "", fmt.Errorf("error encoding tendermint public key bits to consensus address: %v", err)
	}

	return address, nil
}

func MustConsensusNodePubKeyFromTmPubKey(bech32Prefix string, pubKey []byte) string {
	address, err := ConsensusNodePubKeyFromTmPubKey(bech32Prefix, pubKey)
	if err != nil {
		panic(err)
	}

	return address

}

func ConsensusNodePubKeyFromTmPubKey(bech32Prefix string, pubKey []byte) (string, error) {
	cosmosPubKey := &ed255192.PubKey{
		Key: pubKey,
	}
	pkToMarshal := cosmosPubKey.AsTmPubKey()

	conv, err := bech32.ConvertBits(legacy.Cdc.MustMarshalBinaryBare(pkToMarshal), 8, 5, true)
	if err != nil {
		return "", fmt.Errorf("error converting tendermint public key to bech32 bits: %v", err)
	}
	address, err := bech32.Encode(bech32Prefix, conv)
	if err != nil {
		return "", fmt.Errorf("error encoding tendermint public key bits to consensus public key: %v", err)
	}

	return address, nil
}

func ConsensusNodeAddressFromPubKey(bech32Prefix string, consensusNodePubKey string) (string, error) {
	_, conv, err := bech32.Decode(consensusNodePubKey)
	if err != nil {
		return "", fmt.Errorf("error converting consensus node pubkey to address")
	}

	pkToUnmarshal, err := bech32.ConvertBits(conv, 5, 8, false)
	if err != nil {
		return "", fmt.Errorf("error converting bech32 bits to tendermint public key: %v", err)
	}
	var pubKey crypto.PubKey
	legacy.Cdc.MustUnmarshalBinaryBare(pkToUnmarshal, &pubKey)

	conv, err = bech32.ConvertBits(pubKey.Address().Bytes(), 8, 5, true)
	if err != nil {
		return "", fmt.Errorf("error converting tendermint public key to bech32 bits: %v", err)
	}
	address, err := bech32.Encode(bech32Prefix, conv)
	if err != nil {
		return "", fmt.Errorf("error encoding tendermint public key bits to consensus address: %v", err)
	}

	return address, nil
}
