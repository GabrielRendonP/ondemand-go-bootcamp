package repo

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/GabrielRendonP/ondemand-go-bootcamp/entities"
	"github.com/GabrielRendonP/ondemand-go-bootcamp/helpers"
	"github.com/GabrielRendonP/ondemand-go-bootcamp/worker"
)

type localData struct{}

// Defines interface for local repo
type LocalDataInterface interface {
	ReadCSVData() ([][]string, error)
	GetAllPokemonsApi() []entities.Pokemon
	SaveToCsv([]entities.Pokemon) error
	ConcurrentRead(int, int, string) (*entities.DataResponse, error)
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

// ConcurrentRead reads a csv file based on max items, items per worker and even or odd type param
func (r localData) ConcurrentRead(ipw int, items int, types string) (*entities.DataResponse, error) {
	var resList [][]string

	list, _ := helpers.ReadCSV()
	filtered := helpers.Filter(list, types)

	ctx, cancel := context.WithCancel(context.Background())
	jobsNum := helpers.CapMaxItems(len(filtered), items)

	if !worker.CanBeDone(ipw, jobsNum) {
		cancel()
		return nil, fmt.Errorf("not enough items per worker for total amount of work")
	}

	jobs := make(chan []string, jobsNum)
	results := make(chan []string)

	var w worker.Worker
	w.WorkerPool(ipw, jobs, results, ctx, cancel)
	w.CreateJobs(jobs, filtered, jobsNum)

	for i := 0; i < jobsNum; i++ {
		res := <-results
		resList = append(resList, res)
	}

	close(results)
	response := entities.DataResponse{
		ResponseSize: len(resList),
		Ipw:          ipw,
		Types:        types,
		Data:         resList,
	}
	return &response, nil
}

func (r localData) getApiResponse(urlPath string) (*http.Response, error) {
	return http.Get(urlPath)
}
