package main

import (
	"errors"
	"fmt"
)

func commandHelp(s *Status) error {
	if s.extraArgs != nil {
		return errors.New("incorrect number of arg: expect 0")
	}
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
	for key, cmd := range getCommands() {
		fmt.Print(key+": ", cmd.description, "\n")
	}
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
	return nil
}
