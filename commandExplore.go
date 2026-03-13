package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandExplore(config *Config) error {
	if config.location.Results == nil {
		err := commandMap(config)
		if err != nil {
			return fmt.Errorf("Error: %s", err)
		}
	}

	if config.exploreInput == "" {
		return fmt.Errorf("Please enter city name")
	}
	var url string

	for _, result := range config.location.Results {
		if result.Name == config.exploreInput {
			fmt.Printf("Exploring %s...\n", result.Name)
			url = result.URL
			break
		}
	}
	if url == "" {
		return fmt.Errorf("Location not found")
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error: %s", err)
	}
	body, err := io.ReadAll((res.Body))
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code %d\nMessage: %s", res.StatusCode, body)
	}
	if err != nil {
		return fmt.Errorf("Error: %s", err)
	}

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
