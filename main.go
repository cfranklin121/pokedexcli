package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	commands := getCommands()
	var lst []cliCommand
	for _, val := range commands {
		lst = append(lst, val)
	}

	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		fmt.Println("Error:", err)
	}
	body, err := io.ReadAll((res.Body))
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code %d\nMessage: %s", res.StatusCode, body)
	}
	if err != nil {
		fmt.Println("Error:", err)
	}

	config := Config{}
	err = json.Unmarshal(body, &config)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for { //CLI loop
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		input := text[0]

		//fmt.Println("DEBUG:")
		//fmt.Println(config.Next)

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
