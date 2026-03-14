package main

import (
	"encoding/json"
	"fmt"

	"github.com/cfranklin121/pokedexcli/internal/pokeapi"
)

func commandExplore(config *Config) error {
	if config.location.Results == nil {
		err := commandMap(config)
		if err != nil {
			return fmt.Errorf("Error: %s", err)
		}
	}

	if config.commandArg == "" {
		return fmt.Errorf("Please enter city name")
	}
	var url string

	for _, result := range config.location.Results {
		if result.Name == config.commandArg {
			fmt.Printf("Exploring %s...\n", result.Name)
			url = result.URL
			break
		}
	}
	if url == "" {
		return fmt.Errorf("Location not found")
	}

	body, err := pokeapi.GetApiData(url)

	var area Area
	err = json.Unmarshal(body, &area)
	if err != nil {
		return fmt.Errorf("Error: %s", err)
	}

	results := area.PokemonEncounters
	for _, r := range results {
		fmt.Printf(" - %s\n", r.Pokemon.Name)
	}

	return nil
}
