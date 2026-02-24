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

		validCommand := false
		for key := range commands {
			if input[0] == key {
				err := commands[input[0]].callback()
				if input[0] == "help" {
					for _, command := range commands {
						fmt.Printf("%s: %s\n", command.name, command.description)
					}
				}
				if err != nil {
					fmt.Println(err)
				}
				validCommand = true
				break
			}
		}

		if validCommand == false {
			fmt.Println("Unknown command")
		}
	}
}

func commandExit() error {
	fmt.Println("Closing Pokedex...Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	return nil
}
