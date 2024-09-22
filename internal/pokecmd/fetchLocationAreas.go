package pokecmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/poirei/pokedexcli/internal/pokecache"
)

func fetchLocationAreas(cmd string, config *Config, cache *pokecache.Cache) ([]LocationArea, error) {
	url := ""

	if cmd == "map" {
		url = config.Next
	} else {
		url = config.Previous
	}

	// Check if data already present in cache, if so, return it
	val, isPresent := cache.Get(url)

	if isPresent {
		locationAreas := []LocationArea{}

		if err := json.Unmarshal(val, &locationAreas); err != nil {
			return nil, fmt.Errorf("error unmarshaling data from cache: %w", err)
		}

		return locationAreas, nil
	}

	res, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("error fetching location areas: %w", err)
	}

	if res.StatusCode > http.StatusOK {
		return nil, fmt.Errorf("unable to fetch location areas: %v", res.StatusCode)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Add new data to cache
	cache.Add(url, body)

	var responseData ResponseData

	if err := json.Unmarshal(body, &responseData); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	config.Next = responseData.Next

	if responseData.Previous != nil {
		config.Previous = *responseData.Previous
	} else {
		config.Previous = ""
	}

	locationAreas := responseData.Results

	return locationAreas, nil

}
