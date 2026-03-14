package main

import (
	"encoding/json"
	"fmt"

	"github.com/cfranklin121/pokedexcli/internal/pokeapi"
)

func commandCatch(config *Config) error {
	baseUrl := "https://pokeapi.co/api/v2/pokemon/"
	if config.commandArg == "" {
		return fmt.Errorf("Please enter pokemon name")
	}

	data, err := pokeapi.GetApiData(baseUrl + config.commandArg)
	if err != nil {
		return err
	}

	var pokemon Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return fmt.Errorf("Error: %s", err)
	}

	fmt.Println(pokemon.Name)
	fmt.Println(pokemon.BaseExperience)
	return nil
}
