package parsers

import "strings"

func DefaultLowercaseAddressHook(address string) (string, error) {
	return strings.ToLower(address), nil
}
