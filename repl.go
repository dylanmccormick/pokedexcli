package main

import (
	"bufio"
	"fmt"
	"os"
)

func StartRepl() {
	fmt.Println("welcom to the pokedex")
	scanner := bufio.NewScanner(os.Stdin)
	c := &config{
		next:     "https://pokeapi.co/api/v2/location-area/",
		previous: "",
	}
	for {
		fmt.Printf("pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		err := controller(text, c)
		if err != nil {
			fmt.Println(err)
		}

	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

type config struct {
	previous string
	next     string
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
	}
}

func controller(s string, c *config) error {
	commands := commands()

	if _, ok := commands[s]; !ok {
		return fmt.Errorf("invalid command")
	}

	err := commands[s].callback(c)
	if err != nil {
		return err
	}

	return nil

}
