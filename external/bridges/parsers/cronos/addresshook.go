package cronos

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/crypto-com/chain-indexing/external/tmcosmosutils"
)

func DefaultCronosEVMAddressHookGenerator(accountAddressPrefix string) func(string) (string, error) {
	return func(hexAddress string) (string, error) {
		var addr []byte
		if strings.HasPrefix(hexAddress, "0x") {
			addr = common.HexToAddress(hexAddress[2:]).Bytes()
		} else {
			addr = common.HexToAddress(hexAddress).Bytes()
		}
		accountAddr, accountAddrErr := tmcosmosutils.AccountAddressFromBytes(accountAddressPrefix, addr)
		if accountAddrErr != nil {
			return "", fmt.Errorf("error converting cronosevm address to account address: %v", accountAddrErr)
		}

		return strings.ToLower(accountAddr), nil
	}
}
