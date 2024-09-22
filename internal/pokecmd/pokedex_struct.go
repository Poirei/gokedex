package pokecmd

type Pokedex map[string]PokemonInfo

func CreatePokedex() *Pokedex {
	pokedex := &Pokedex{}

	return pokedex
}
