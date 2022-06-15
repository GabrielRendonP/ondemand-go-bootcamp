package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/GabrielRendonP/ondemand-go-bootcamp/helpers"
)

// GetPokemon finds and returns a pokemon by id
func (c controller) GetPokemon(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	id := query.Get("id")
	pokemon, err := c.s.GetPokemon(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"Invalid pokemon id or missing csv file"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(pokemon)
	w.Write(response)
}

// GetPokemons returns all pokemons from csv file
func (c controller) GetPokemons(w http.ResponseWriter, r *http.Request) {

	pokeList, err := c.s.GetAllPokemons()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"No list found"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(pokeList)
	w.Write(response)
}

// PokemonIndex shows all pokemons found on external poke api
func (c controller) PokemonIndex(w http.ResponseWriter, r *http.Request) {
	list, _ := c.s.GetAllPokemonsFromApi()

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(list)
	w.Write(response)
}

// SavePokeApi Stores a pokemon list from external poke api
func (c controller) SavePokeApi(w http.ResponseWriter, r *http.Request) {
	list, _ := c.s.GetAllPokemonsFromApi()
	err := c.s.SaveToCsv(list)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal("unable to save data")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Data Saved to CSV"}`))
}

// ConcurrentRead displays read results from csv or errors
func (c controller) ConcurrentRead(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var err error
	err = helpers.ValidateParams(query)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := `{%s: %s}`
		w.Write([]byte(fmt.Sprintf(msg, "error", err.Error())))
		return
	}
	ipw, _ := strconv.Atoi(query.Get("ipw"))
	items, _ := strconv.Atoi(query.Get("items"))
	types := query.Get("type")

	list, err := c.s.ConcurrentRead(ipw, items, types)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := `{%s: %s}`
		w.Write([]byte(fmt.Sprintf(msg, "error", err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(list)
	w.Write(response)

}
