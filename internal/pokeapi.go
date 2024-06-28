package internal

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Page struct {
	Next    *string `json:"next"`
	Prev    *string `json:"previous"`
	Results []Item  `json:"results"`
}

type Item struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

var LocationAreaPage1 string = "https://pokeapi.co/api/v2/location-area/"

var LocationAreaPage = Page{
	Next:    &LocationAreaPage1,
	Prev:    nil,
	Results: nil,
}

func GetNextLocationAreaPage() ([]Item, error) {
	if LocationAreaPage.Next == nil {
		return nil, errors.New("invalid command - already on the last page")
	}
	resp, err := http.Get(*LocationAreaPage.Next)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &LocationAreaPage)
	if err != nil {
		return nil, err
	}
	return LocationAreaPage.Results, nil
}

func GetPrevLocationAreaPage() ([]Item, error) {
	if LocationAreaPage.Prev == nil {
		return nil, errors.New("invalid command - already on the first page")
	}
	resp, err := http.Get(*LocationAreaPage.Prev)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &LocationAreaPage)
	if err != nil {
		return nil, err
	}
	return LocationAreaPage.Results, nil
}

// type LocationArea struct {
// 	Name    string `json:"name"`
// 	Id      int    `json:"id"`
// 	PkmEncs []Item `json:"pokemon_encounters"`
// }

// type Item struct {
// 	Pkm Pokemon `json:"pokemon"`
// }

// type Pokemon struct {
// 	Name string `json:"name"`
// 	Url  string `json:"url"`
// }

// func GetLocationArea(Id int) (*LocationArea, error) {
// 	resp, err := http.Get("https://pokeapi.co/api/v2/location-area/" + strconv.Itoa(Id))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	loc := new(LocationArea)
// 	err = json.Unmarshal(body, loc)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return loc, nil
// }

// func GetLocation() error {
// 	fmt.Println("abc")
// 	resp, err := http.Get("https://pokeapi.co/api/v2/location-area/2")
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return err
// 	}
// 	loc := LocationArea{}
// 	err = json.Unmarshal(body, &loc)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	fmt.Println(loc.Name, loc.Id)
// 	for _, p := range loc.PkmEncs {
// 		fmt.Println(p.Pkm.Name, p.Pkm.Url)
// 	}
// 	return nil
// }
