package pokecmd

type EncounterMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VersionDetails struct {
	MaxLevel int `json:"max_level"`
	MinLevel int `json:"min_level"`
	Version  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"version"`
}

type EncounterMethodRate struct {
	EncounterMethod EncounterMethod  `json:"encounter_method"`
	VersionDetails  []VersionDetails `json:"version_details"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Name struct {
	Language Language `json:"language"`
	Name     string   `json:"name"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EncounterDetail struct {
	Chance          int   `json:"chance"`
	ConditionValues []any `json:"condition_values"`
	MaxLevel        int   `json:"max_level"`
	Method          struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"method"`
	MinLevel int `json:"min_level"`
}

type VersionDetail struct {
	EncounterDetails []EncounterDetail `json:"encounter_details"`
	MaxChance        int               `json:"max_chance"`
	Version          struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"version"`
}

type PokemonEncounter struct {
	Pokemon        Pokemon         `json:"pokemon"`
	VersionDetails []VersionDetail `json:"version_details"`
}

type ExploreData struct {
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	GameIndex            int                   `json:"game_index"`
	ID                   int                   `json:"id"`
	Location             Location              `json:"location"`
	Name                 string                `json:"name"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}
