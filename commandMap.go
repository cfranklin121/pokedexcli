package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(config *Config) error {
	/*
		body, ok := config.cache.Get(config.location.Next) //Get response from cache
		if ok {
			err := json.Unmarshal(body, &config.location)
			if err != nil {
				return fmt.Errorf("Error: %s", err)

			}
			results := config.location.Results
			for _, r := range results {
				fmt.Println(r.Name)
			}
			config.cache.Add(config.location.Next, body) //Add response to cache
			return nil
		}
	*/
	if len(config.location.Results) == 0 {
		res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
		if err != nil {
			return fmt.Errorf("Error: %s", err)
		}
		body, err := io.ReadAll((res.Body))
		res.Body.Close()
		if res.StatusCode > 299 {
			return fmt.Errorf("Response failed with status code %d\nMessage: %s", res.StatusCode, body)
		}
		if err != nil {
			return fmt.Errorf("Error: %s", err)
		}

		err = json.Unmarshal(body, &config.location)
		if err != nil {
			return fmt.Errorf("Error: %s", err)
		}
		config.cache.Add(config.location.Next, body) //Add response to cache
	} else {

		res, err := http.Get(config.location.Next)
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

		err = json.Unmarshal(body, &config.location)
		if err != nil {
			return err
		}
		config.cache.Add(config.location.Next, body) //Add response to cache
	}

	results := config.location.Results
	for _, r := range results {
		fmt.Println(r.Name)
	}

	return nil
}

func commandMapb(config *Config) error {
	if config.location.Previous == "" {
		return fmt.Errorf("No previous pages")
	}
	res, err := http.Get(config.location.Previous)
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

	err = json.Unmarshal(body, &config.location)
	if err != nil {
		return err
	}

	results := config.location.Results
	for _, r := range results {
		fmt.Println(r.Name)
	}

	return nil
}
