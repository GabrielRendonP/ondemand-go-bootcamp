package controllers

import (
	"encoding/json"
	"net/http"
)

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
