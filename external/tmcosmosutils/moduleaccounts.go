package tmcosmosutils

import "crypto/sha256"

type ModuleAccounts struct {
	FeeCollector        string
	Mint                string
	Distribution        string
	Gov                 string
	BondedTokensPool    string
	NotBondedTokensPool string
	IBCTransfer         string
}

func NewModuleAccounts(addressPrefix string) ModuleAccounts {
	return ModuleAccounts{
		FeeCollector:        getModuleAccount(addressPrefix, "fee_collector"),
		Mint:                getModuleAccount(addressPrefix, "mint"),
		Distribution:        getModuleAccount(addressPrefix, "distribution"),
		Gov:                 getModuleAccount(addressPrefix, "gov"),
		BondedTokensPool:    getModuleAccount(addressPrefix, "bonded_tokens_pool"),
		NotBondedTokensPool: getModuleAccount(addressPrefix, "not_bonded_tokens_pool"),
		IBCTransfer:         getModuleAccount(addressPrefix, "transfer"),
	}
}

func getModuleAccount(addressPrefix string, module string) string {
	b := sha256.Sum256([]byte(module))
	return MustModuleAccountFromBytes(addressPrefix, b[:20])
}
