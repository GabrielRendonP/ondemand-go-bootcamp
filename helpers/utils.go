package helpers

import (
	"strconv"
)

// Filters by odd or even
func Filter(list [][]string, filterBy string) [][]string {
	var filteredList [][]string

	for _, row := range list {
		num, _ := strconv.Atoi(row[0])
		if checkFilter(num, filterBy) {
			filteredList = append(filteredList, row)
		}
	}

	return filteredList
}

// Caps max amount of items based on items and jobs numbers
func CapMaxItems(maxJobs int, maxItems int) int {
	if maxItems > maxJobs {
		return maxJobs
	}
	return maxItems
}

func checkFilter(num int, filter string) bool {
	if filter == "even" {
		return num%2 == 0
	}
	if filter == "odd" {
		return num%2 != 0
	}
	return false
}
