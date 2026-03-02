package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Config struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println("")
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(config *Config) error {
	type Location struct {
		Count    int    `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"results"`
	}

	results := config.Results
	for _, r := range results {
		fmt.Println(r.Name)
	}

	res, err := http.Get(config.Next)
	if err != nil {
		return err
	}
	body, err := io.ReadAll((res.Body))
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code %d\nMessage: %s", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	location := Location{}
	err = json.Unmarshal(body, &location)
	if err != nil {
		return err
	}
	//results := location.Results
	//for _, r := range results {
	//	fmt.Println(r.Name)
	//}

	*config = Config{
		Next:     location.Next,
		Previous: location.Previous,
		Results:  location.Results,
	}
	return nil
}

// DEBUG: Test command
func commandTest(config *Config) error {
	fmt.Println("Testing... complete")
	return nil
}

// DEBUG: Error test
func commandError(config *Config) error {
	fmt.Println("This function should return an error message")
	err := fmt.Errorf("Error Message!")
	return err
}
