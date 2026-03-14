package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	commands := getCommands()
	var config Config

	scanner := bufio.NewScanner(os.Stdin)
	for { //CLI loop
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		input := text[0]
		if len(text) > 1 {
			config.commandArg = text[1]
		} else {
			config.commandArg = ""
		}

		invalidCommand := true
		for key := range commands { //check if command is in registry
			if input == key {
				err := commands[input].callback(&config) //Run commands callback function
				if err != nil {
					fmt.Println("Error:", err)
				}
				invalidCommand = false
				break
			}
		}

		if invalidCommand {
			fmt.Println("Unknown command")
		}
	}
}
