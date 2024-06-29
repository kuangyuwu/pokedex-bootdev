package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kuangyuwu/pokedex-bootdev/internal/pokeapi"
)

type Status struct {
	pokeApiClient  *pokeapi.PokeApiClient
	nextLocAreaUrl *string
	prevLocAreaUrl *string
}

func startCli() {
	timeout := 5 * time.Second
	interval := 5 * time.Minute
	var LocAreaP1 string = "https://pokeapi.co/api/v2/location-area/"
	status := Status{
		pokeApiClient:  pokeapi.NewPokeApiClient(timeout, interval),
		nextLocAreaUrl: &LocAreaP1,
		prevLocAreaUrl: nil,
	}
	fmt.Println("-------------- Welcome to Pokedex!! ---------------")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		token := scanner.Text()

		words := tokenToWords(token)
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
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
		"exit": {
			name:        "exit",
			description: "exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: " display the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "display the previous 20 locations",
			callback:    commandMapB,
		},
	}
}
