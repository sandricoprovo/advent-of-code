package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./rows.csv")

	if err != nil {
		log.Fatal("Could not open input file.")
	}

	reader := csv.NewReader(file)

	// Removes column titles row from reader
	if _, err := reader.Read(); err != nil {
		log.Fatal("Could read first line of csv.")
	}

	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal("Could not read input csv.")
	}

	PartOne(records)
	PartTwo(records)
}
