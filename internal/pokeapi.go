package internal

import (
	"encoding/json"
	"io"
	"net/http"
)

type Page struct {
	Next    *string `json:"next"`
	Prev    *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

var LocationAreaPage1 string = "https://pokeapi.co/api/v2/location-area/"

var LocationAreaPage = Page{
	Next:    &LocationAreaPage1,
	Prev:    nil,
	Results: nil,
}

func GetPage(url string) (*Page, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Page
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
