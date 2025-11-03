package service

import (
	"go-sprint6/pkg/morse"
	"strings"
)

func isMorse(text string) bool {
	empty := strings.ReplaceAll(text, ".", "")
	empty = strings.ReplaceAll(empty, "-", "")
	empty = strings.ReplaceAll(empty, " ", "")

	return len(empty) == 0
}

func ConvertText(text string) string {
	if isMorse(text) {
		return morse.ToText(text)
	} else {
		return morse.ToMorse(text)
	}
}
