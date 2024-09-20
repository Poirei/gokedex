package pokecmd

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/poirei/pokedexcli/internal/pokecache"
)

func TestFetchLocationAreas(t *testing.T) {
	cache := pokecache.NewCache(5 * time.Second)

	config := &Config{}

	t.Run("fetches data from cache successfully", func(t *testing.T) {
		url := "http://example.com/next"
		config.Next = url

		mockLocationAreas := []LocationArea{
			{
				Name: "Area1",
				URL:  "http://example.com/next/area1",
			},
			{
				Name: "Area2",
				URL:  "http://example.com/next/area2",
			},
		}

		data, err := json.Marshal(mockLocationAreas)

		if err != nil {
			t.Fatalf("error marshaling mock data: %v", err)
		}

		cache.Add(url, data)

		locationAreas, err := fetchLocationAreas("map", config, cache)

		if err != nil {
			t.Fatalf("error fetching location areas: %v", err)
		}

		if len(locationAreas) != len(mockLocationAreas) {
			t.Fatalf("expected %d location areas, got %d", len(mockLocationAreas), len(locationAreas))
		}

		if locationAreas[0].Name != mockLocationAreas[0].Name {
			t.Fatalf("expected location area name to be %s, got %s", mockLocationAreas[0].Name, locationAreas[0].Name)
		}
	})

	t.Run("fetches data from API successfully", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			responseData := ResponseData{
				Next:     "http://example.com/next",
				Previous: nil,
				Results: []LocationArea{
					{
						Name: "Area1",
						URL:  "http://example.com/next/area1",
					},
					{
						Name: "Area2",
						URL:  "http://example.com/next/area2",
					},
				},
			}

			json.NewEncoder(w).Encode(responseData)
		}))

		defer server.Close()

		config.Next = server.URL

		locationAreas, err := fetchLocationAreas("map", config, cache)

		if err != nil {
			t.Fatalf("error fetching location areas: %v", err)
		}

		if len(locationAreas) != 2 {
			t.Fatalf("expected 2 location areas, got %d", len(locationAreas))
		}

		if locationAreas[0].Name != "Area1" {
			t.Fatalf("expected location area name to be Area1, got %s", locationAreas[0].Name)
		}

		cachedData, isPresent := cache.Get(config.Next)

		if !isPresent {
			t.Fatalf("expected data to be cached, but it wasn't")
		}

		cachedLocationAreas := []LocationArea{}

		if err := json.Unmarshal(cachedData, &cachedLocationAreas); err != nil {
			t.Fatalf("error unmarshaling cached data: %v", err)
		}

		if len(cachedLocationAreas) != 2 {
			t.Fatalf("expected 2 location areas in cache, got %d", len(cachedLocationAreas))
		}
	})
}
