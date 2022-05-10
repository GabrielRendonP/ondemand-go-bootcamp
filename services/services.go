package services

import (
	"fmt"
	"log"

	"github.com/GabrielRendonP/ondemand-go-bootcamp/entities"
)

type service struct {
	r repository
}

// Implements repository interface
type repository interface {
	ReadCSVData() ([][]string, error)
}

// Instantiates a new service object
func NewService(r repository) service {
	return service{r}
}

func (s service) GetAllPokemons() ([]entities.Pokemon, error) {
	data, err := s.r.ReadCSVData()

	if err != nil {
		log.Panic("Could not read csv data")
		return nil, err
	}

	var pokeList []entities.Pokemon
	for i, line := range data {
		if i > 0 {
			var pokemon entities.Pokemon
			for j, pokeAtt := range line {
				if j == 0 {
					pokemon.Number = pokeAtt
				} else if j == 1 {
					pokemon.Name = pokeAtt
				} else if j == 2 {
					pokemon.PokeType = pokeAtt
				}
			}
			pokeList = append(pokeList, pokemon)
		}
	}
	return pokeList, nil
}

func (s service) GetPokemon(id string) (entities.Pokemon, error) {
	pokeList, err := s.GetAllPokemons()

	if err != nil {
		return entities.Pokemon{}, fmt.Errorf("error in data list")
	}

	for _, poke := range pokeList {
		if poke.Number == id {
			log.Println("Found!", poke)
			return poke, nil
		}
	}

	return entities.Pokemon{}, fmt.Errorf("pokemon with id %s not found", id)
}
