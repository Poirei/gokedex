package pokecmd

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/poirei/pokedexcli/internal/pokecache"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CommandMap(config *Config, cache *pokecache.Cache, _ string, _ string, _ *Pokedex) error {
	locationAreas, err := fetchLocationAreas("map", config, cache)

	if err != nil {
		return err
	}

	for _, locationArea := range locationAreas {
		fmt.Println(locationArea.Name)
	}

	return nil
}

func CommandMapb(config *Config, cache *pokecache.Cache, _ string, _ string, _ *Pokedex) error {
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

func CommandExplore(_ *Config, cache *pokecache.Cache, locationArea string, _ string, _ *Pokedex) error {
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

func CommandCatch(_ *Config, cache *pokecache.Cache, _ string, pokemonName string, pokedex *Pokedex) error {
	if _, ok := (*pokedex)[pokemonName]; ok {
		fmt.Printf("\n%s is already in your Pokedex!\n", cases.Title(language.English).String(pokemonName))
		return nil
	}

	fmt.Printf("\nThrowing a Pokeball at %s...\n", pokemonName)

	pokemonInfo, err := fetchPokemon(pokemonName, cache)

	if err != nil {
		return fmt.Errorf("error fetching pokemon info: %w", err)
	}

	if pokemonInfo.BaseExperience == 0 {
		return fmt.Errorf("invalid pokemon: %s", pokemonName)
	}

	generatedBaseXp := rand.Intn(pokemonInfo.BaseExperience)

	fmt.Printf("\nGenerated base XP: %d\nActual base XP: %d\nRatio: %.2f\n", generatedBaseXp, pokemonInfo.BaseExperience, float64(generatedBaseXp)/float64(pokemonInfo.BaseExperience))

	if float64(generatedBaseXp)/float64(pokemonInfo.BaseExperience) < .8 {
		fmt.Printf("\n%s escaped!\n", cases.Title(language.English).String(pokemonName))

	} else {
		(*pokedex)[pokemonName] = pokemonInfo
		fmt.Printf("\n%s was caught!\n", cases.Title(language.English).String(pokemonName))
	}

	return nil
}
