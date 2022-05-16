package main

import (
	"net/http"

	"github.com/GabrielRendonP/ondemand-go-bootcamp/controllers"
	"github.com/GabrielRendonP/ondemand-go-bootcamp/repo"
	"github.com/GabrielRendonP/ondemand-go-bootcamp/services"
)

func main() {
	lr := repo.NewLocalData()
	ns := services.NewService(lr)
	nc := controllers.NewController(ns)

	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/pokemons", nc.GetPokemons)
	http.HandleFunc("/pokemon", nc.GetPokemon)

	_ = http.ListenAndServe(":8080", nil)
}
