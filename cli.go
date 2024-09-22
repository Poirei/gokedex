package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/poirei/pokedexcli/internal/pokecache"
	"github.com/poirei/pokedexcli/internal/pokecmd"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *pokecmd.Config, cache *pokecache.Cache, locationArea string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations, and so on.",
			callback:    pokecmd.CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back.",
			callback:    pokecmd.CommandMapb,
		},
		"explore": {
			name:        "explore <location_area>",
			description: "Displays the names of all the Pokémon that can be encountered in the given location area.",
			callback:    pokecmd.CommandExplore,
		},
	}
}

func commandHelp(_ *pokecmd.Config, _ *pokecache.Cache, _ string) error {
	fmt.Println("\nWelcome to the Pokedex CLI!\nUsage:")
	fmt.Print("\n")

	for _, val := range getCommands() {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}

	fmt.Print("\n")

	return nil
}

func commandExit(_ *pokecmd.Config, _ *pokecache.Cache, _ string) error {
	os.Exit(0)

	return errors.New("Error exiting Pokedex.")
}
