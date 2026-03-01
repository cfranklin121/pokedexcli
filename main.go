package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	commands := getCommands()
	var lst []cliCommand
	for _, val := range commands {
		lst = append(lst, val)
	}

	for { //CLI loop
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		input := text[0]

		invalidCommand := true
		for key := range commands { //check if command is in registry
			if input == key {
				err := commands[input].callback() //Run commands callback function
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
