package services

import (
	"github.com/GabrielRendonP/ondemand-go-bootcamp/entities"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLocalData struct {
	mock.Mock
}

// Mock methods
func (r MockLocalData) ReadCSVData() ([][]string, error) {
	var result = [][]string{
		{"1", "PokeMock"},
		{"2", "FakeMon"},
	}
	return result, nil
}

func (r MockLocalData) SaveToCsv([]entities.Pokemon) error {
	return nil
}

func (r MockLocalData) GetAllPokemonsApi() []entities.Pokemon {
	pokemons := []entities.Pokemon{
		{Number: "1", Name: "PokeMock"},
	}
	return pokemons
}

func (r MockLocalData) ConcurrentRead(int, int, string) (*entities.DataResponse, error) {
	return nil, nil
}

// Test methods
func TestGetAllPokemons(t *testing.T) {
	repo := MockLocalData{}
	repo.On("ReadCSVData").Return([][]string{})

	service := NewService(repo)
	pokeList, _ := service.GetAllPokemons()
	assert.NotEmpty(t, pokeList)
	assert.GreaterOrEqual(t, len(pokeList), 1)
}

func TestGetPokemon(t *testing.T) {
	repo := MockLocalData{}
	repo.On("ReadCSVData").Return([][]string{})

	service := NewService(repo)
	poke, _ := service.GetPokemon("2")
	assert.Equal(t, poke.Name, "FakeMon", "Pokemon with id 2 should exist")
	assert.NotEmpty(t, poke)
}

func TestGetAllPokemonsFromApi(t *testing.T) {
	repo := MockLocalData{}
	repo.On("GetAllPokemonsApi").Return([]entities.Pokemon{})

	service := NewService(repo)
	pokeList, _ := service.GetAllPokemonsFromApi()
	assert.Equal(t, pokeList[0].Name, "PokeMock", "PokeMocks should be the first item in the list")
	assert.NotEmpty(t, pokeList)
}
