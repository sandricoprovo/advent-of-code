package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
)

func PartTwo(csvRows [][]string) {
	similarityScore := 0
	leftList, rightList := ProcessCsvRows(csvRows)

	for _, item := range rightList {
		if slices.Contains(leftList, item) {
			itemInt, err := strconv.Atoi(item)

			if err != nil {
				log.Fatal("Could not convert string to int.")
			}

			similarityScore += itemInt
		}
	}

	fmt.Printf("Prompt Two Answer: %v ✅\n", similarityScore)
}
