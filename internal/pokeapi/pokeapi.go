package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/kuangyuwu/pokedex-bootdev/internal/pokecache"
)

type PokeApiClient struct {
	client *http.Client
	cache  *pokecache.Cache
}

type Page struct {
	Next    *string `json:"next"`
	Prev    *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func NewPokeApiClient(timeout, interval time.Duration) *PokeApiClient {
	pokeApiClient := PokeApiClient{
		client: &http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(interval),
	}
	return &pokeApiClient
}

func GetPage(p *PokeApiClient, url string) (*Page, error) {
	var body []byte
	if val, ok := p.cache.Get(url); ok {
		body = val
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		resp, err := p.client.Do(req)
		if err != nil {
			return nil, err
		}
		body, err = io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return nil, err
		}
	}

	var result Page
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type LocArea struct {
	PkmEncs []struct {
		Pkm struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func GetLocArea(p *PokeApiClient, url string) (*LocArea, error) {
	var body []byte
	if val, ok := p.cache.Get(url); ok {
		body = val
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		resp, err := p.client.Do(req)
		if err != nil {
			return nil, err
		}
		body, err = io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return nil, err
		}
	}

	var result LocArea
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
