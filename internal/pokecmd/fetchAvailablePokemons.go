package pokecmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/poirei/pokedexcli/internal/pokecache"
)

func fetchAvailablePokemons(locationArea string, cache *pokecache.Cache) ([]string, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + locationArea
	availablePokemons := []string{}

	val, isPresent := cache.Get(url)

	if isPresent {

		if err := json.Unmarshal(val, &availablePokemons); err != nil {
			return nil, fmt.Errorf("error unmarshaling data from cache: %w", err)
		}

		return availablePokemons, nil
	}

	res, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("error fetching available pokemons: %w", err)
	}

	if res.StatusCode > http.StatusOK {
		return nil, fmt.Errorf("unable to fetch available pokemons: %v", res.StatusCode)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	cache.Add(url, body)

	var resData ExploreData

	if err := json.Unmarshal(body, &resData); err != nil {
		return nil, fmt.Errorf("error unmarshaling data from response body: %w", err)
	}

	for _, pokemon := range resData.PokemonEncounters {
		availablePokemons = append(availablePokemons, pokemon.Pokemon.Name)
	}

	return availablePokemons, nil
}
