package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startCli() {
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
			command.callback()
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
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit the pokedex",
			callback:    commandExit,
		},
	}
}
