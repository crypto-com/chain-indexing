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

func GetCosmosAPIVersionFromMsgType(msgType string) string {
	cosmosAPIVersion := DefaultCosmosAPIVersion
	parts := strings.Split(msgType, ".")
	if len(parts) >= 2 {
		cosmosAPIVersion = parts[len(parts)-2]
		if _, exists := CosmosAPIVersionMap[cosmosAPIVersion]; !exists {
			cosmosAPIVersion = DefaultCosmosAPIVersion
		}
	}
	return cosmosAPIVersion
}
