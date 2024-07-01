package pokeapi

import (
	"encoding/json"
	"errors"
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

func getJsonObj(p *PokeApiClient, url string) ([]byte, error) {
	var jsonObj []byte
	if val, ok := p.cache.Get(url); ok {
		jsonObj = val
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		resp, err := p.client.Do(req)
		if err != nil {
			return nil, err
		}
		jsonObj, err = io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return nil, err
		}
		p.cache.Add(url, jsonObj)
	}
	return jsonObj, nil
}

func GetPage(p *PokeApiClient, url string) (*Page, error) {
	jsonObj, err := getJsonObj(p, url)
	if err != nil {
		return nil, err
	}
	var result Page
	err = json.Unmarshal(jsonObj, &result)
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
	jsonObj, err := getJsonObj(p, url)
	if err != nil {
		return nil, err
	}
	var result LocArea
	err = json.Unmarshal(jsonObj, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type prePokemon struct {
	Name    string `json:"name"`
	BaseExp int    `json:"base_experience"`
	Height  int    `json:"height"`
	Weight  int    `json:"weight"`
	Stats   []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

type Pokemon struct {
	Name    string
	BaseExp int
	Height  int
	Weight  int
	Stats   [6]int
	Types   [2]string
}

func GetPokemon(p *PokeApiClient, url string) (*Pokemon, error) {
	jsonObj, err := getJsonObj(p, url)
	if err != nil {
		return nil, err
	}
	var preResult prePokemon
	err = json.Unmarshal(jsonObj, &preResult)
	if err != nil {
		return nil, err
	}
	result, err := preResult.toPokemon()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p prePokemon) toPokemon() (*Pokemon, error) {
	result := Pokemon{
		Name:    p.Name,
		BaseExp: p.BaseExp,
		Height:  p.Height,
		Weight:  p.Weight,
		Stats:   [6]int{},
		Types:   [2]string{},
	}
	for _, s := range p.Stats {
		i, err := statToI(s.Stat.Name)
		if err != nil {
			return nil, err
		}
		result.Stats[i] = s.BaseStat
	}
	for _, t := range p.Types {
		if t.Slot == 1 {
			result.Types[0] = t.Type.Name
		} else if t.Slot == 2 {
			result.Types[1] = t.Type.Name
		} else {
			return nil, errors.New("invalid slot in types")
		}
	}
	return &result, nil
}

func statToI(stat string) (int, error) {
	switch stat {
	case "hp":
		return 0, nil
	case "attack":
		return 1, nil
	case "defense":
		return 2, nil
	case "special-attack":
		return 3, nil
	case "special-defense":
		return 4, nil
	case "speed":
		return 5, nil
	default:
		return 0, errors.New("invalid stat")
	}
}
