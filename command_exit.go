package main

import (
	"fmt"
	"os"
)

func CommandExit() error {
	fmt.Println("exiting the program")
	os.Exit(0)
	return nil

}
