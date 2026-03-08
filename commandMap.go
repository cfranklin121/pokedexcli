package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(config *Config) error {
	if len(config.Results) == 0 {
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

		err = json.Unmarshal(body, &config)
		if err != nil {
			fmt.Println("Error:", err)
		}
	} else {

		res, err := http.Get(config.Next)
		if err != nil {
			return fmt.Errorf("Error: %s", err)
		}
		body, err := io.ReadAll((res.Body))
		res.Body.Close()
		if res.StatusCode > 299 {
			return fmt.Errorf("Response failed with status code %d\nMessage: %s", res.StatusCode, body)
		}
		if err != nil {
			return err
		}

		err = json.Unmarshal(body, &config)
		if err != nil {
			return err
		}
	}

	results := config.Results
	for _, r := range results {
		fmt.Println(r.Name)
	}

	return nil
}

func commandMapb(config *Config) error {
	res, err := http.Get(config.Previous)
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

	err = json.Unmarshal(body, &config)
	if err != nil {
		return err
	}

	results := config.Results
	for _, r := range results {
		fmt.Println(r.Name)
	}

	return nil
}
