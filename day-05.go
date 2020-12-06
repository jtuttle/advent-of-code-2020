package main

import (
	"fmt"
	"math"
	"sort"
)

func BinSearch(letters string) int {
	index := 0
	elements := int(math.Pow(2 ,float64(len(letters))))

	for _, letter := range letters {
		elements /= 2

		if letter == 'B' || letter == 'R' {
			index += elements
		}
	}

	return index
}

func main() {
	lines := ReadInput("day-05-input.txt")

	var seatIds []int

	for _, seat := range lines {
		row := BinSearch(seat[0:7])
		col := BinSearch(seat[7:10])

		seatIds = append(seatIds, row * 8 + col)
	}

	sort.Ints(seatIds)

	fmt.Println("Highest Seat ID (part 1):", seatIds[len(seatIds) - 1])

	for i := 1; i < len(seatIds) - 1; i++ {
		if seatIds[i + 1] != seatIds[i] + 1 {
			fmt.Println("Missing Seat ID (part 2):", seatIds[i] + 1)
		}
	}
}
