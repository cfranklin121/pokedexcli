package main

import (
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func cleanInput(text string) []string {
	cleaned := strings.Fields(strings.ToLower(strings.TrimSpace(text)))
	return cleaned
}
