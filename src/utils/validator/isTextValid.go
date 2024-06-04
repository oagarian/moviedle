package validator

import "strings"

func IsTextEmpty(text string) bool {
	text = strings.TrimSpace(text)
	return len(text) <= 0
}