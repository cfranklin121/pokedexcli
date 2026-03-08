package main

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
		"map": {
			name:        "map",
			description: "List all the locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous locations",
			callback:    commandMapb,
		},
	}

	return commands
}
