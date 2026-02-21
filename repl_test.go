package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " Hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Pikachu",
			expected: []string{"pikachu"},
		},
		{
			input:    "Gotta catch them all!!",
			expected: []string{"gotta", "catch", "them", "all!!"},
		},
		{
			input:    "more    pokeMon",
			expected: []string{"more", "pokemon"},
		},
	}

	for _, c := range cases {
		fmt.Println("Run test...")
		fmt.Println("Case:", c.input)
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Length doesn't match")
		}
		fmt.Println("")
		for i := range actual {
			word := actual[i]
			fmt.Println("Input:", word)
			if word == c.expected[i] {
				fmt.Printf("Expected: %s\n", c.expected[i])
				fmt.Printf("Actual:   %s\n", word)
				fmt.Println("Pass")
			} else {
				fmt.Printf("Expected: %s\n", c.expected[i])
				fmt.Printf("Actual:   %s\n", word)
				fmt.Println("Fail")
				t.Errorf("Output does not match")
			}
			fmt.Println("======================")
		}
		fmt.Println("----------------------")
	}
}
