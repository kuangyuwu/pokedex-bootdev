package main

import (
	"errors"
	"fmt"

	"github.com/kuangyuwu/pokedex-bootdev/internal/pokeapi"
)

func commandExplore(s *Status) error {
	if s.extraArgs == nil || len(s.extraArgs) > 1 {
		return errors.New("incorrect number of arg: expect 1")
	}
	name := s.extraArgs[0]
	url, ok := s.currLocAreas[name]
	if !ok {
		return errors.New("location area not found")
	}
	locArea, err := pokeapi.GetLocArea(s.pokeApiClient, url)
	if err != nil {
		return err
	}
	fmt.Println("Exploring", name, "...")
	for key := range s.currPkms {
		delete(s.currPkms, key)
	}
	for _, item := range locArea.PkmEncs {
		s.currPkms[item.Pkm.Name] = item.Pkm.Url
		fmt.Println(" -", item.Pkm.Name)
	}
	return nil
}
