package main

import (
	"fmt"
	"os"
)

func CommandExit(c *config, _ ...string) error {
	fmt.Println("exiting the program")
	os.Exit(0)
	return nil

}
