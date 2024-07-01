package main

import (
	"errors"
	"fmt"

	"github.com/kuangyuwu/pokedex-bootdev/internal/pokeapi"
)

func commandMap(s *Status) error {
	if s.nextLocAreaUrl == nil {
		return errors.New("invalid command - already on the last page")
	}
	page, err := pokeapi.GetPage(s.pokeApiClient, *s.nextLocAreaUrl)
	if err != nil {
		return nil
	}
	for key := range s.currLocAreas {
		delete(s.currLocAreas, key)
	}
	for _, item := range page.Results {
		s.currLocAreas[item.Name] = item.Url
		fmt.Println(item.Name)
	}
	s.nextLocAreaUrl = page.Next
	s.prevLocAreaUrl = page.Prev
	for key := range s.currPkms {
		delete(s.currPkms, key)
	}
	return nil
}

func commandMapB(s *Status) error {
	if s.prevLocAreaUrl == nil {
		return errors.New("invalid command - already on the first page")
	}
	page, err := pokeapi.GetPage(s.pokeApiClient, *s.prevLocAreaUrl)
	if err != nil {
		return nil
	}
	for key := range s.currLocAreas {
		delete(s.currLocAreas, key)
	}
	for _, item := range page.Results {
		s.currLocAreas[item.Name] = item.Url
		fmt.Println(item.Name)
	}
	s.nextLocAreaUrl = page.Next
	s.prevLocAreaUrl = page.Prev
	for key := range s.currPkms {
		delete(s.currPkms, key)
	}
	return nil
}
