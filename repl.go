package main

import (
	"strings"

	"github.com/cfranklin121/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Config struct {
	location     Location
	cache        pokecache.Cache
	exploreInput string
}

func cleanInput(text string) []string {
	cleaned := strings.Fields(strings.ToLower(strings.TrimSpace(text)))
	return cleaned
}
