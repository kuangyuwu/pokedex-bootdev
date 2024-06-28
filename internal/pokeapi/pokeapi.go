package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Client = http.Client

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

func NewClient() *Client {
	client := new(Client)
	client.Timeout = 5 * time.Second
	return client
}

func GetPage(c *Client, url string) (*Page, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
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
