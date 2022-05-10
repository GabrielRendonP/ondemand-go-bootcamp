package controllers

import (
	"encoding/json"
	"net/http"
)

func (c controller) GetPokemon(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	id := query.Get("id")
	pokemon, err := c.s.GetPokemon(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"Invalid pokemon id"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(pokemon)
	w.Write(response)
}
