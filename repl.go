package main

import (
	"strings"
)

func cleanInput(text string) []string {
	cleaned := strings.Fields(strings.ToLower(strings.TrimSpace(text)))
	return cleaned
}
