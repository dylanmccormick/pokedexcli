package main

import (
	"fmt"
	"os"
)

func CommandExit(c *config) error {
	fmt.Println("exiting the program")
	os.Exit(0)
	return nil

}
