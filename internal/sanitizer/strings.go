package sanitizer

import (
	"github.com/crypto-com/chain-indexing/external/primptr"
	"strings"
)

func SanitizePostgresString(s string) string {
	return strings.ReplaceAll(s, "\\u0000", "")
}

func SanitizePostgresStringPtr(s *string) *string {
	if s == nil {
		return nil
	}
	return primptr.String(SanitizePostgresString(*s))
}
