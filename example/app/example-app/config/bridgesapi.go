package config

import (
	"github.com/crypto-com/chain-indexing/external/bridges/parsers"
	"github.com/crypto-com/chain-indexing/external/bridges/parsers/cronos"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi/handlers"
)

type BridgesAPIConfig struct {
	Networks []BridgeNetworkConfig        `yaml:"networks"`
	Chains   []handlers.BridgeChainConfig `yaml:"chains"`
}

type BridgeNetworkConfig struct {
	ChainName string `yaml:"chain_name" mapstructure:"chain_name"`
	// Chain network unique abbreviation, used in URL query params
	Abbreviation                    handlers.NetworkAbbreviation `yaml:"abbreviation" mapstructure:"abbreviation"`
	MaybeAddressHookKey             *string                      `yaml:"maybe_address_hook_key"`
	MaybeCronosAccountAddressPrefix *string                      `yaml:"maybe_cronos_account_address_prefix"`
}

var BridgeAddressHooks = (func() map[string]func(
	networkConfig *BridgeNetworkConfig,
) handlers.AddressHook {
	hooks := make(map[string]func(config *BridgeNetworkConfig) handlers.AddressHook, 0)

	hooks["DefaultLowercaseAddressHook"] = func(_ *BridgeNetworkConfig) handlers.AddressHook {
		return parsers.DefaultLowercaseAddressHook
	}
	hooks["DefaultCronosEVMAddressHookGenerator"] = func(config *BridgeNetworkConfig) handlers.AddressHook {
		return cronos.DefaultCronosEVMAddressHookGenerator(*config.MaybeCronosAccountAddressPrefix)
	}

	return hooks
})()

func ParseBridgesConfig(rawConfig *BridgesAPIConfig) handlers.BridgesConfig {
	config := handlers.BridgesConfig{
		Networks: make([]handlers.BridgeNetworkConfig, 0, len(rawConfig.Networks)),
		Chains:   rawConfig.Chains,
	}

	for i, network := range rawConfig.Networks {
		if network.MaybeAddressHookKey == nil {
			config.Networks = append(config.Networks, handlers.BridgeNetworkConfig{
				ChainName:    network.ChainName,
				Abbreviation: network.Abbreviation,
			})
		} else {
			hook := BridgeAddressHooks[*network.MaybeAddressHookKey](&rawConfig.Networks[i])
			config.Networks = append(config.Networks, handlers.BridgeNetworkConfig{
				ChainName:        network.ChainName,
				Abbreviation:     network.Abbreviation,
				MaybeAddressHook: &hook,
			})
		}
	}

	return config
}
