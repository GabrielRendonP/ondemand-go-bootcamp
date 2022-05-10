package controllers

import "github.com/GabrielRendonP/ondemand-go-bootcamp/entities"

type controller struct {
	s bigController
}

// Here I can define getters, setters, deleter and all other methods
type bigController interface {
	getter
}

type getter interface {
	GetAllPokemons() ([]entities.Pokemon, error)
	GetPokemon(id string) (entities.Pokemon, error)
}

func NewController(s bigController) controller {
	return controller{s}
}
