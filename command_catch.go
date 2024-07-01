package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/kuangyuwu/pokedex-bootdev/internal/pokeapi"
)

func commandCatch(s *Status) error {
	if s.extraArgs == nil || len(s.extraArgs) > 1 {
		return errors.New("incorrect number of arg: expect 1")
	}
	name := s.extraArgs[0]
	url, ok := s.currPkms[name]
	if !ok {
		return errors.New(name + " is not found in this area")
	}
	pkm, err := pokeapi.GetPokemon(s.pokeApiClient, url)
	if err != nil {
		return nil
	}
	exp := pkm.BaseExp
	pct := catchPercentage(exp)
	num := rand.Intn(100)
	fmt.Printf("Throwing a Pokeball at %v", name)
	for i := 0; i < 3; i++ {
		fmt.Printf(".")
		time.Sleep(time.Second)
	}
	fmt.Println("")
	if num < pct {
		fmt.Println("Congrats,", name, "was caught!")
		fmt.Println("You may now inspect it with the inspect command.")
		if _, ok := s.pkmCaught[name]; !ok {
			s.pkmCaught[name] = pkm
		}
	} else {
		fmt.Println("Oh no,", name, "escaped!")
	}
	return nil
}

func catchPercentage(exp int) int {
	offset := float64(300-exp) / 3.0
	if offset < 0.0 {
		offset = 0.0
	}
	return int(0.8*offset + 10.0)
}
