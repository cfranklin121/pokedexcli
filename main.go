package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		input := cleanInput(text)

		for key := range commands {
			if input[0] == key {
				fmt.Println(commands[input[0]].description)
				err := commands[input[0]].callback()
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func commandExit() error {
	fmt.Println("Closing Pokedex...Goodbye!")
	os.Exit(0)
	return fmt.Errorf("")
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	return nil
}
