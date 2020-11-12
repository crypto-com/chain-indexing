package parser

import "strings"

func TrimAmountDenom(s string) string {
	return strings.TrimRight(strings.TrimRight(s, "basetcro"), "basecro")
}
