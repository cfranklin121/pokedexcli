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

	for { //CLI loop
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		input := text[0]

		validCommand := false
		for key := range commands { //check if command is in registry
			if input == key {
				err := commands[input].callback()
				if input == "help" {
					for _, command := range commands {
						fmt.Printf("%s: %s\n", command.name, command.description)
					}
				}
				if err != nil {
					fmt.Println("Error:", err)
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
