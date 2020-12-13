package main

import (
	"io/ioutil"
	"math"
	"strings"
	"strconv"
)

func ReadInput(filename string) []string {
	input, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	inputStr := string(input)
	inputStr = strings.TrimSuffix(inputStr, "\n")

	return strings.Split(inputStr, "\n")
}

func ConvertToInts(strings []string) []int {
	var vals []int

	for _, str := range strings {
		iVal, err := strconv.Atoi(str)

		if err != nil {
			panic(err)
		}

		vals = append(vals, iVal)
	}

	return vals
}

func ManhattanDistance(x1 int, y1 int, x2 int, y2 int) int {
	return int(math.Abs(float64(x2 - x1)) + math.Abs(float64(y2 - y1)))
}
