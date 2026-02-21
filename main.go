package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lower := strings.ToLower(trimmed)
	lst := strings.Fields(lower)
	return lst
}
