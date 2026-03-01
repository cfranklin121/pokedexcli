package main

import (
	"strings"
)

func getCommands() map[string]cliCommand {
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
		"test": {
			name:        "test",
			description: "DEBUG: This is a test command",
			callback:    commandTest,
		},
		"error": {
			name:        "error",
			description: "This command tests errors",
			callback:    commandError,
		},
	}

	return commands
}

func cleanInput(text string) []string {
	cleaned := strings.Fields(strings.ToLower(strings.TrimSpace(text)))
	return cleaned
}
