package controllers

import (
	"net/http"

	"github.com/GabrielRendonP/ondemand-go-bootcamp/services"
)

type GetterInterface interface {
	GetPokemon(http.ResponseWriter, *http.Request)
	GetPokemons(http.ResponseWriter, *http.Request)
	PokemonIndex(http.ResponseWriter, *http.Request)
	SavePokeApi(http.ResponseWriter, *http.Request)
}
type controller struct {
	s services.ServiceInterface
}

func NewController(s services.ServiceInterface) GetterInterface {
	return &controller{s}
}
