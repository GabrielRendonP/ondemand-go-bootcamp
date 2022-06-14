package repo

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/GabrielRendonP/ondemand-go-bootcamp/entities"
	"github.com/GabrielRendonP/ondemand-go-bootcamp/helpers"
)

type localData struct{}

// Defines interface for local repo
type LocalDataInterface interface {
	ReadCSVData() ([][]string, error)
	GetAllPokemonsApi() []entities.Pokemon
	SaveToCsv([]entities.Pokemon) error
}

// Instantiates new localdata struct
func NewLocalData() LocalDataInterface {
	return &localData{}
}

// ReadCsvData calls helper method that access csv file
func (r localData) ReadCSVData() ([][]string, error) {
	results, err := helpers.ReadCSV()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return results, nil
}

// GetAllPokemonsApi makes exteneral api request and returns a pokemon list
func (r localData) GetAllPokemonsApi() []entities.Pokemon {
	response, _ := r.getApiResponse("https://pokeapi.co/api/v2/pokedex/1")

	defer response.Body.Close()

	var pokeResponse entities.Response

	rb, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal("Error")
	}

	json.Unmarshal(rb, &pokeResponse)
	var pokeList []entities.Pokemon

	for _, poke := range pokeResponse.PokeList {
		var pokemon entities.Pokemon
		pokemon.Name = poke.Species.Name
		pokemon.Number = strconv.Itoa(poke.EntryNumber)
		pokeList = append(pokeList, pokemon)
	}

	return pokeList
}

// SaveToCsv creates a new csv file to store pokemon list
func (r localData) SaveToCsv(list []entities.Pokemon) error {
	csvFile, err := os.Create("./lib/pokemonsFromApi.csv")

	if err != nil {
		return err
	}

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	cWriter := csv.NewWriter(csvFile)

	for _, pok := range list {
		row := []string{pok.Number, pok.Name}
		cWriter.Write((row))
	}

	cWriter.Flush()
	csvFile.Close()

	return nil
}

func (r localData) getApiResponse(urlPath string) (*http.Response, error) {
	return http.Get(urlPath)
}
