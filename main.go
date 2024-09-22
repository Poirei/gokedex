package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/poirei/pokedexcli/internal/pokecache"
	"github.com/poirei/pokedexcli/internal/pokecmd"
)

func main() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(5 * time.Second)
	pokedex := pokecmd.CreatePokedex()

	config := pokecmd.Config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
	}

	for {
		fmt.Print(">Pokedex ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading input", err)
			}

			break
		}

		input := scanner.Text()
		input = strings.TrimSpace(input)

		fields := strings.Fields(input)
		input = fields[0]

		locationArea := ""
		pokemonName := ""

		if input == "explore" || input == "catch" {
			if len(fields) != 2 {
				fmt.Println("\nMissing arg. Please type 'help' for available commands.")
				fmt.Print("\n")

				continue
			} else {
				switch input {
				case "explore":
					locationArea = fields[1]
				case "catch":
					pokemonName = fields[1]
				default:
					fmt.Println("\nInvalid command. Please type 'help' for available commands.")
				}
			}
		}

		command, ok := commands[input]

		if !ok {
			fmt.Println("\nInvalid command. Please type 'help' for available commands.")
			fmt.Print("\n")

			continue
		}

		if err := command.callback(&config, cache, locationArea, pokemonName, pokedex); err != nil {
			fmt.Println("\nError executing command:\n", err)
			fmt.Println()

			continue
		}
	}
}
