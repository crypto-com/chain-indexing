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

func GetMsgVersionFromMsgType(msgType string) string {
	msgVersion := DefaultCosmosAPIVersion
	parts := strings.Split(msgType, ".")
	if len(parts) >= 2 {
		msgVersion = parts[len(parts)-2]
		if _, exists := CosmosAPIVersionMap[msgVersion]; !exists {
			msgVersion = DefaultCosmosAPIVersion
		}
	}
	return msgVersion
}
