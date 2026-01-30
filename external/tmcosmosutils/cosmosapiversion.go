package tmcosmosutils

import (
	"strings"
)

const (
	DefaultCosmosAPIVersion = "v1beta1"
	CosmosAPIVersionV1      = "v1"
)

var CosmosAPIVersionMap = map[string]bool{
	DefaultCosmosAPIVersion: true,
	CosmosAPIVersionV1:      true,
}

// MustNewCoinFromAmountInterface returns a Coin from the amount map in the form of
// map[string]interface{}. It panics when the amount is invalid.
func GetCosmosAPIVersionByMsgType(msgType string) string {
	cosmosAPIVersion := "v1beta1"
	parts := strings.Split(msgType, ".")
	if len(parts) >= 2 {
		cosmosAPIVersion = parts[len(parts)-2]
		if _, exists := CosmosAPIVersionMap[cosmosAPIVersion]; !exists {
			cosmosAPIVersion = DefaultCosmosAPIVersion
		}
	}
	return cosmosAPIVersion
}
