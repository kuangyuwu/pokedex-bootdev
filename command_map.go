package main

import (
	"errors"
	"fmt"

	"github.com/kuangyuwu/pokedex-bootdev/internal"
)

func commandMap(s *Status) error {
	if s.nextLocAreaUrl == nil {
		return errors.New("invalid command - already on the last page")
	}
	page, err := internal.GetPage(*s.nextLocAreaUrl)
	if err != nil {
		return nil
	}
	for _, item := range page.Results {
		fmt.Println(item.Name)
	}
	s.nextLocAreaUrl = page.Next
	s.prevLocAreaUrl = page.Prev
	return nil
}

func commandMapB(s *Status) error {
	if s.prevLocAreaUrl == nil {
		return errors.New("invalid command - already on the first page")
	}
	page, err := internal.GetPage(*s.prevLocAreaUrl)
	if err != nil {
		return nil
	}
	for _, item := range page.Results {
		fmt.Println(item.Name)
	}
	s.nextLocAreaUrl = page.Next
	s.prevLocAreaUrl = page.Prev
	return nil
}
