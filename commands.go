package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing Pokedex...Goodbye!")
	os.Exit(0)
	return fmt.Errorf("")
}
