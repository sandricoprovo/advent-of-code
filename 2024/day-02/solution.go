package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func isDiffInRange(left int, right int) bool {
	d := math.Abs(float64(left - right))

	return d <= float64(3) && d >= float64(1)
}

func checkLineSafety(reportLine *string, safetyCount *int) {
	var isIncreasing bool
	var isDecreasing bool
	var isDiffValid bool

	lineSlice := strings.Split(*reportLine, " ")

	for i := 0; i <= len(lineSlice); i++ {
		if i+1 >= len(lineSlice) {
			break
		}

		leftInt, leftErr := strconv.Atoi(lineSlice[i])
		rightInt, rightErr := strconv.Atoi(lineSlice[i+1])

		if leftErr != nil || rightErr != nil {
			log.Fatal("Could not convert string to number")
		}

		if leftInt < rightInt {
			isIncreasing = true
		} else {
			isDecreasing = true
		}

		isDiffValid = isDiffInRange(leftInt, rightInt)

		if isDecreasing == isIncreasing || !isDiffValid {
			break
		}

	}

	if isDiffValid && (isDecreasing != isIncreasing) && (isDecreasing || isIncreasing) {
		*safetyCount = *safetyCount + 1
	}

	fmt.Printf("line: %v, isDiffValid: %v, isGradual: %v, safetyCount: %v \n", *reportLine, isDiffValid, isDecreasing != isIncreasing, *safetyCount)

}

func Solution(path string) {
	var numOfSafeReports int

	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		log.Fatal("Could not open input file.")
	}

	linesScanner := bufio.NewScanner(file)

	for linesScanner.Scan() {
		line := linesScanner.Text()
		checkLineSafety(&line, &numOfSafeReports)
	}

	if scannerErr := linesScanner.Err(); scannerErr != nil {
		log.Fatal("Could not open input file.")
	}

	fmt.Printf("Prompt One Answer: %v ✅\n", numOfSafeReports)
}
