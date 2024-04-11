package main

import (
	"bufio"
	"fmt"
	"os"
)

func StartRepl() {
	fmt.Println("welcom to the pokedex")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		err := controller(text)
		if err != nil {
			fmt.Println(err)
		}

	}

}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func controller(s string) error {
	commands := commands()

	switch {
	case s == "help":
		err := commands[s].callback()
		if err != nil {
			return err
		}
	case s == "exit":
		err := commands[s].callback()
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("invalid command")

	}
	return nil

}
