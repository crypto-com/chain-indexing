package sanitizer

import (
	"strings"

	"github.com/crypto-com/chain-indexing/external/primptr"
)

func SanitizePostgresString(s string) string {
	return strings.ReplaceAll(
		strings.ReplaceAll(
			s,
			string([]byte{0x00}), "",
		),
		"\\u0000", "",
	)
}

func SanitizePostgresStringPtr(s *string) *string {
	if s == nil {
		return nil
	}
	return primptr.String(SanitizePostgresString(*s))
}
