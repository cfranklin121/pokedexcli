package main

import "fmt"

func commandInspect(config *Config) error {
	for _, pokemon := range config.pokedex {
		if config.commandArg == pokemon.Name {
			fmt.Printf("Name: %s\n", pokemon.Name)
			fmt.Printf("Weight: %d\n", pokemon.Weight)
			fmt.Printf("Stats\n")
			for i := 0; i < len(pokemon.Stats); i++ {
				fmt.Printf(" - %s: %d\n", pokemon.Stats[i].Stat.Name, pokemon.Stats[i].BaseStat)
			}
			fmt.Printf("Types:\n")
			for i := 0; i < len(pokemon.Types); i++ {
				fmt.Printf(" - %s\n", pokemon.Types[i].Type.Name)
			}
			return nil
		}
	}
	return fmt.Errorf("You have not caught this pokemon.")
}
