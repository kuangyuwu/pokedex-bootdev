package main

import "fmt"

func commandHelp() error {
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
	for key, cmd := range getCommands() {
		fmt.Print(key+": ", cmd.description, "\n")
	}
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
	return nil
}
