package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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
		return fmt.Errorf("http.GET error: %s", err)
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

	*config = Config(location)
	return nil
}

func commandMapb(config *Config) error {
	type Location struct {
		Count    int    `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"results"`
	}

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

	location := Location{}
	err = json.Unmarshal(body, &location)
	if err != nil {
		return err
	}

	*config = Config(location)

	results := config.Results
	for _, r := range results {
		fmt.Println(r.Name)
	}

	return nil
}
