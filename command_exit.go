package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(s *Status) error {
	if s.extraArgs != nil {
		return errors.New("incorrect number of arg: expect 0")
	}
	fmt.Println("---------- Thank you for using Pokedex!! ----------")
	os.Exit(0)
	return nil
}
