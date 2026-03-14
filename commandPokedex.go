package main

import "fmt"

func commandPokedex(config *Config) error {
	for _, pokemon := range config.pokedex {
		fmt.Println(pokemon.Name)
	}
	return nil
}
