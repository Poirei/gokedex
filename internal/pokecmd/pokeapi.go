package pokecmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/poirei/pokedexcli/internal/pokecache"
)

type Config struct {
	Next     string
	Previous string
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ResponseData struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

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

func CommandMap(config *Config, cache *pokecache.Cache) error {
	locationAreas, err := fetchLocationAreas("map", config, cache)

	if err != nil {
		return err
	}

	for _, locationArea := range locationAreas {
		fmt.Println(locationArea.Name)
	}

	return nil
}

func CommandMapb(config *Config, cache *pokecache.Cache) error {
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
