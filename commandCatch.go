package main

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/cfranklin121/pokedexcli/internal/pokeapi"
)

func commandCatch(config *Config) error {
	baseUrl := "https://pokeapi.co/api/v2/pokemon/"
	if config.commandArg == "" {
		return fmt.Errorf("Please enter pokemon name")
	}

	data, err := pokeapi.GetApiData(baseUrl + config.commandArg)
	if err != nil {
		return fmt.Errorf("Pokemon not found")
	}

	var pokemon Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return fmt.Errorf("Error: %s", err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	catchChance := rand.Intn(100) + 150
	if catchChance < pokemon.BaseExperience {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it wit the inspect command")
	if config.pokedex == nil {
		config.pokedex = make(map[string]Pokemon)
	}
	config.pokedex[pokemon.Name] = pokemon
	return nil
}
