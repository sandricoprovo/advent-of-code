package main

func ProcessCsvRows(csvRows [][]string) ([]string, []string) {
	leftList := []string{}
	rightList := []string{}

	for _, row := range csvRows {
		leftInt := row[0]
		rightInt := row[1]

		leftList = append(leftList, leftInt)
		rightList = append(rightList, rightInt)
	}

	return leftList, rightList
}
