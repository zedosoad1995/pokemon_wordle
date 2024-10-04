package utils

import "strings"

func ExtractLabelAndValue(input string) (string, string) {
	parts := strings.SplitN(input, ":", 2)
	label := parts[0]
	value := ""
	if len(parts) > 1 {
		value = parts[1]
	}
	return label, value
}
