package main

import "fmt"

func CommandHelp(c *config, _ ...string) error {
	fmt.Println("why do you need help")
	return nil

}
