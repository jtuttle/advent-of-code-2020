package main

import (
	"fmt"
	"sort"
)

var multipliers = []int{ 1, 1, 1, 2, 4, 7 }

func main() {
	lines := ReadInput("day-10-input.txt")

	vals := ConvertToInts(lines)
	sort.Ints(vals)

	vals = append([]int{0}, vals...)
	vals = append(vals, vals[len(vals) - 1] + 3)

	var diffs []int
	diffCounts := make(map[int]int)

	length := 0
	combos := 1

	for i := 1; i < len(vals); i++ {
		diff := vals[i] - vals[i - 1]
		diffs = append(diffs, diff)
		diffCounts[diff]++

		length++

		if diff == 3 {
			combos *= multipliers[length]
			length = 0
		}
	}

	diffCountProduct := diffCounts[1] * diffCounts[3]
	fmt.Println("Product of 1 diffs and 3 diffs (part 1):", diffCountProduct)

	fmt.Println("Possible combinations (part 2):", combos)
}
