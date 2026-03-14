package main

import "fmt"

func commandPokedex(config *Config) error {
	fmt.Println("Your Pokedex:")
	if config.pokedex == nil {
		fmt.Println("You have not caught any pokemon yet.")
		return nil
	}
	for _, pokemon := range config.pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
