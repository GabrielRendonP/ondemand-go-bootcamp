package entities

type Pokemon struct {
	Number string
	Name   string
}
type Response struct {
	PokeList []Poke `json:"pokemon_entries"`
}
type Poke struct {
	EntryNumber int            `json:"entry_number"`
	Species     PokemonSpecies `json:"pokemon_species"`
}
type PokemonSpecies struct {
	Name string `json:"name"`
}
