package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/kuangyuwu/pokedex-bootdev/internal/pokeapi"
)

type Status struct {
	pokeApiClient  *pokeapi.PokeApiClient
	nextLocAreaUrl *string
	prevLocAreaUrl *string
	currLocAreas   map[string]string
	currPkms       map[string]string
	pkmCaught      map[string]*pokeapi.Pokemon
	extraArgs      []string
}

func startCli() {
	timeout := 5 * time.Second
	interval := 5 * time.Minute
	var LocAreaP1 string = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=10"
	status := Status{
		pokeApiClient:  pokeapi.NewPokeApiClient(timeout, interval),
		nextLocAreaUrl: &LocAreaP1,
		prevLocAreaUrl: nil,
		currLocAreas:   make(map[string]string),
		currPkms:       make(map[string]string),
		pkmCaught:      make(map[string]*pokeapi.Pokemon),
		extraArgs:      nil,
	}
	fmt.Println("-------------- Welcome to Pokedex!! ---------------")
	scanner := bufio.NewScanner(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		token := scanner.Text()

		words := tokenToWords(token)
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		if len(words) > 1 {
			status.extraArgs = words[1:]
		} else {
			status.extraArgs = nil
		}
		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback(&status)
			if err != nil {
				fmt.Println("pokedex:", err)
			}
		} else {
			fmt.Println("pokedex: command not found: " + commandName)
		}
	}
}

func tokenToWords(token string) []string {
	line := strings.ToLower(token)
	return strings.Fields(line)
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Status) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch",
			description: "  [+pokemon] throw a pokeball to the pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "   exit the pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "[+loc area] explore a location area",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "   display a help message",
			callback:    commandHelp,
		},
		"inspect": {
			name:        "inspect",
			description: "[+pokemon] inspect a caught pokemon",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "    display the next 10 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "   display the previous 10 locations",
			callback:    commandMapB,
		},
		"pokedex": {
			name:        "pokedex",
			description: "list out all caught pokemons",
			callback:    commandPokemon,
		},
	}
}
