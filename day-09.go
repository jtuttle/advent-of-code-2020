package main

import (
	"fmt"
	"sort"
)

func IsValid(num int, window []int) bool {
	for i := 0; i < len(window); i++ {
		for j := i + 1; j < len(window); j++ {
			if window[i] + window[j] == num {
				return true
			}
		}
	}

	return false
}

// find set of contigious numbers in list that sums to target
func FindAddends(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		sum := nums[i]

		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]

			if sum == target {
				return nums[i:j + 1]
			} else if sum > target {
				break
			}
		}
	}

	return nil
}

func main() {
	lines := ReadInput("day-09-input.txt")

	nums := ConvertToInts(lines)

	for i := 25; i < len(nums); i++ {
		window := nums[i - 25:i]
		nextNum := nums[i]

		if !IsValid(nextNum, window) {
			fmt.Println("First invalid number (part 1):", nextNum)

			addends := FindAddends(nums, nextNum)

			sort.Ints(addends)
			weakness := addends[0] + addends[len(addends) - 1]

			fmt.Println("Sum of min and max addend (part 2):", weakness)

			break
		}
	}
}
