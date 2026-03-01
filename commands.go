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
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println("")
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

// DEBUG: Test command
func commandTest() error {
	fmt.Println("Testing... complete")
	return nil
}

// DEBUG: Error test
func commandError() error {
	fmt.Println("This function should return an error message")
	err := fmt.Errorf("Error Message!")
	return err
}
