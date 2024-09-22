package pokecmd

import (
	"errors"
	"fmt"

	"github.com/poirei/pokedexcli/internal/pokecache"
)

func CommandMap(config *Config, cache *pokecache.Cache, _ string) error {
	locationAreas, err := fetchLocationAreas("map", config, cache)

	if err != nil {
		return err
	}

	for _, locationArea := range locationAreas {
		fmt.Println(locationArea.Name)
	}

	return nil
}

func CommandMapb(config *Config, cache *pokecache.Cache, _ string) error {
	if config.Previous == "" {
		return errors.New("cannot go back further. no previous page")
	}

	locationAreas, err := fetchLocationAreas("mapb", config, cache)

	if err != nil {
		return err
	}

	for _, locationArea := range locationAreas {
		fmt.Println(locationArea.Name)
	}

	return nil
}

func CommandExplore(_ *Config, cache *pokecache.Cache, locationArea string) error {
	fmt.Printf("\nExploring %s...\n", locationArea)

	availablePokemons, err := fetchAvailablePokemons(locationArea, cache)

	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, pokemonName := range availablePokemons {
		fmt.Println("- ", pokemonName)
	}
	return nil
}
