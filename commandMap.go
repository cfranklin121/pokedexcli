package main

import (
	"encoding/json"
	"fmt"

	"github.com/cfranklin121/pokedexcli/internal/pokeapi"
)

func commandMap(config *Config) error {
	/*
		body, ok := config.cache.Get(config.location.Next) //Get response from cache
		if ok {
			err := json.Unmarshal(body, &config.location)
			if err != nil {
				return fmt.Errorf("Error: %s", err)

			}
			results := config.location.Results
			for _, r := range results {
				fmt.Println(r.Name)
			}
			config.cache.Add(config.location.Next, body) //Add response to cache
			return nil
		}
	*/
	if len(config.location.Results) == 0 {
		body, err := pokeapi.GetApiData("https://pokeapi.co/api/v2/location-area/")
		if err != nil {
			return err
		}
		err = json.Unmarshal(body, &config.location)
		if err != nil {
			return fmt.Errorf("Error: %s", err)
		}
		config.cache.Add(config.location.Next, body) //Add response to cache
	} else {
		body, err := pokeapi.GetApiData(config.location.Next)
		if err != nil {
			return err
		}
		err = json.Unmarshal(body, &config.location)
		if err != nil {
			return err
		}
		config.cache.Add(config.location.Next, body) //Add response to cache
	}

	results := config.location.Results
	for _, r := range results {
		fmt.Println(r.Name)
	}

	return nil
}

func commandMapb(config *Config) error {
	if config.location.Previous == "" {
		return fmt.Errorf("No previous pages")
	}

	body, err := pokeapi.GetApiData(config.location.Previous)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &config.location)
	if err != nil {
		return err
	}
	config.cache.Add(config.location.Previous, body) //Add response to cache

	results := config.location.Results
	for _, r := range results {
		fmt.Println(r.Name)
	}

	return nil
}
