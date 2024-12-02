package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func PartOne(csvRows [][]string) {
	var totalDistance int
	leftList, rightList := ProcessCsvRows(csvRows)

	sort.Strings(leftList)
	sort.Strings(rightList)

	// Finds the difference between each sorted pair and increments total
	for i, _ := range leftList {
		leftInt, leftErr := strconv.Atoi(leftList[i])
		rightInt, rightErr := strconv.Atoi(rightList[i])

		if leftErr != nil || rightErr != nil {
			log.Fatal("There was an issue during string conversion.")
		}

		if leftInt > rightInt {
			totalDistance = totalDistance + (leftInt - rightInt)
		} else {
			totalDistance = totalDistance + (rightInt - leftInt)
		}
	}

	fmt.Printf("Prompt One Answer: %v ✅\n", totalDistance)
}
