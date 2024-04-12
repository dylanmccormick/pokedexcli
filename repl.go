package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dylanmccormick/pokedexcli/internal/pokeapi"
)

func StartRepl(cfg *config) {
	fmt.Println("welcom to the pokedex")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()

		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}

		cmd := text[0]
		args := []string{}
		if len(text) > 1 {
			args = text[1:]
		}

		err := controller(cmd, cfg, args...)
		if err != nil {
			fmt.Println(err)
		}

	}

}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient        pokeapi.Client
	previousLocationsURL *string
	nextLocationsUrl     *string
}

func commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    CommandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the program",
			callback:    CommandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next list of 20 regions",
			callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the next list of 20 regions",
			callback:    CommandMapB,
		},
		"explore": {
			name:        "explore <location-name>",
			description: "Explores a given location",
			callback:    CommandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Attempts to catch a given pokemon",
			callback:    CommandCatch,
		},
	}
}

func controller(s string, c *config, args ...string) error {
	commands := commands()

	if _, ok := commands[s]; !ok {
		return fmt.Errorf("invalid command")
	}

	err := commands[s].callback(c, args...)
	if err != nil {
		return err
	}

	return nil

}
