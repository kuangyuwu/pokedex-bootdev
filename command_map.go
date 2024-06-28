package main

import (
	"fmt"

	"github.com/kuangyuwu/pokedex-bootdev/internal"
)

func commandMap() error {
	results, err := internal.GetNextLocationAreaPage()
	for _, item := range results {
		fmt.Println(item.Name)
	}
	return err
}

func commandMapB() error {
	results, err := internal.GetPrevLocationAreaPage()
	for _, item := range results {
		fmt.Println(item.Name)
	}
	return err
}
