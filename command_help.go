package main

import "fmt"

func commandHelp(s *Status) error {
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
	for key, cmd := range getCommands() {
		fmt.Print(key+": ", cmd.description, "\n")
	}
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
	return nil
}
