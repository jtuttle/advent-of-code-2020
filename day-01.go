package main

import (
	"fmt"
)

func main() {
	lines := ReadInput("day-01-input.txt")

	vals := ConvertToInts(lines)

	// find pair that sums to 2020
	for i := 0; i < len(vals); i++ {
		for j := i + 1; j < len(vals); j++ {
			if vals[i] + vals[j] == 2020 {
				fmt.Println("2020 pair product (part 1):", vals[i] * vals[j])
			}
		}
	}

	// find triplet that sums to 2020
	for i := 0; i < len(vals); i++ {
		for j := i + 1; j < len(vals); j++ {
			for k := j + 1; k < len(vals); k++ {
				if vals[i] + vals[j] + vals[k] == 2020 {
					fmt.Println("2020 triple product (part 2):", vals[i] * vals[j] * vals[k])
				}
			}
		}
	}
}
