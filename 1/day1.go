package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	// split out column 1 & column 2
	csvReader := csv.NewReader(file)
	csvReader.Comma = '\t'

	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	var l1 []int
	var l2 []int

	for _, record := range records {
		res := strings.Fields(record[0])

		val1, err := strconv.Atoi(res[0])
		if err != nil {
			panic(err)
		}

		val2, err := strconv.Atoi(res[1])
		if err != nil {
			panic(err)
		}

		l1 = append(l1, val1)
		l2 = append(l2, val2)
	}

	sort.Ints(l1)
	sort.Ints(l2)

	distance := 0

	for i := 0; i < len(l1); i++ {
		distance += abs(l1[i] - l2[i])
	}

	fmt.Println(distance)
}
