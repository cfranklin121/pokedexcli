package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func([]cliCommand) error
}

func commandExit(commands []cliCommand) error {
	fmt.Println("Closing Pokedex...Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(commands []cliCommand) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println("")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

// DEBUG: Test command
func commandTest(commands []cliCommand) error {
	fmt.Println("Testing... complete")
	return nil
}

// DEBUG: Error test
func commandError(commands []cliCommand) error {
	fmt.Println("This function should return an error message")
	err := fmt.Errorf("Error Message!")
	return err
}
