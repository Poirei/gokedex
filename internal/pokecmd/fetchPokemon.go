package pokecmd

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/poirei/pokedexcli/internal/pokecache"
)

func fetchPokemon(pokemonName string, cache *pokecache.Cache) (PokemonInfo, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	cachedData, isPresent := cache.Get(url)
	var pokemonInfo PokemonInfo

	if isPresent {
		if err := json.Unmarshal(cachedData, &pokemonInfo); err != nil {
			return PokemonInfo{}, err
		}

		return pokemonInfo, nil
	}

	res, err := http.Get(url)

	if err != nil {
		return PokemonInfo{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return PokemonInfo{}, err
	}

	resData, err := io.ReadAll(res.Body)

	if err != nil {
		return PokemonInfo{}, err
	}

	cache.Add(url, resData)

	if err := json.Unmarshal(resData, &pokemonInfo); err != nil {
		return PokemonInfo{}, err
	}

	return pokemonInfo, nil
}
