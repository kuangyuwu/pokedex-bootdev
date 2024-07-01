package main

import (
	"errors"
	"fmt"
)

func commandPokemon(s *Status) error {
	if s.extraArgs != nil {
		return errors.New("incorrect number of arg: expect 0")
	}
	fmt.Println("Your Pokedex:")
	for p := range s.pkmCaught {
		fmt.Println(" -", p)
	}
	return nil
}
