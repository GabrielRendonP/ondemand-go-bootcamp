package services

import (
	"errors"
	"fmt"
	"log"

	"github.com/GabrielRendonP/ondemand-go-bootcamp/entities"
	"github.com/GabrielRendonP/ondemand-go-bootcamp/repo"
)

type ServiceInterface interface {
	GetAllPokemons() ([]entities.Pokemon, error)
	GetPokemon(string) (entities.Pokemon, error)
	GetAllPokemonsFromApi() ([]entities.Pokemon, error)
	SaveToCsv([]entities.Pokemon) error
}
type service struct {
	r repo.LocalDataInterface
}

// Instantiates a new service object
func NewService(r repo.LocalDataInterface) ServiceInterface {
	return &service{r}
}

func (s service) GetAllPokemons() ([]entities.Pokemon, error) {
	data, err := s.r.ReadCSVData()
	if err != nil {
		newError := errors.New("missing csv file")
		log.Println(newError)
		return nil, newError
	}

	var pokeList []entities.Pokemon
	for _, line := range data {
		var pokemon entities.Pokemon
		for j, pokeAtt := range line {
			if j == 0 {
				pokemon.Number = pokeAtt
			} else if j == 1 {
				pokemon.Name = pokeAtt
			}
		}
		pokeList = append(pokeList, pokemon)
	}
	return pokeList, nil
}

func (s service) GetPokemon(id string) (entities.Pokemon, error) {
	pokeList, err := s.GetAllPokemons()

	if err != nil {
		return entities.Pokemon{}, err
	}

	for _, poke := range pokeList {
		if poke.Number == id {
			log.Println("Found!", poke)
			return poke, nil
		}
	}

	return entities.Pokemon{}, fmt.Errorf("pokemon with id %s not found", id)
}

func (s service) GetAllPokemonsFromApi() ([]entities.Pokemon, error) {
	list := s.r.GetAllPokemonsApi()
	var newError error
	return list, newError
}

func (s service) SaveToCsv(list []entities.Pokemon) error {
	err := s.r.SaveToCsv(list)

	if err != nil {
		return err
	}
	return nil
}
