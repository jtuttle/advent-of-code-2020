package main

import (
	"fmt"
	"math"
)

func CountTrees(treeMap []string, right int, down int) int {
	pos := 0
	treeCount := 0

	for _, line := range treeMap {
		if line[pos] == byte('#') {
			treeCount++
		}

		pos = int(math.Mod(float64(pos + right), float64(len(line))))
	}

	return treeCount
}

func main() {
	lines := ReadInput("day-03-input.txt")

	treeCount := CountTrees(lines, 3, 1)

	fmt.Println("Tree count (part 1):", treeCount)

	treeCount *= CountTrees(lines, 1, 1)
	treeCount *= CountTrees(lines, 5, 1)
	treeCount *= CountTrees(lines, 7, 1)
	treeCount *= CountTrees(lines, 1, 2)

	fmt.Println("Tree count product (part 2):", treeCount)
}
