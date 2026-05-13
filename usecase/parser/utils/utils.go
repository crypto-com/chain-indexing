package utils

import "strconv"

func UnquoteOrRaw(s string) string {
	unquoted, err := strconv.Unquote(s)
	if err != nil {
		return s
	}
	return unquoted
}