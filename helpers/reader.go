package helpers

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// Probably should go in a driver
func ReadCSV() ([][]string, error) {
	var data [][]string

	file, err := os.Open("./lib/pokemon.csv")
	if err != nil {
		fmt.Println("file not found")
		return nil, err
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err = csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return data, nil
}
