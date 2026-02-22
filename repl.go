package main

import (
	"strings"
)

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lower := strings.ToLower(trimmed)
	lst := strings.Fields(lower)
	return lst
}
