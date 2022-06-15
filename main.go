package main

import (
	"log"
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
	http.HandleFunc("/index", nc.PokemonIndex)
	http.HandleFunc("/save", nc.SavePokeApi)
	http.HandleFunc("/concurrentRead", nc.ConcurrentRead)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("error!!!")
		return
	}

}
