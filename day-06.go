package main

import (
	"fmt"
)

func CountYesAnswers(group []string) map[rune]int {
	set := make(map[rune]int)

	for _, answers := range group {
		for _, letter := range answers {
			set[letter]++
		}
	}

	return set
}

func CountAllYes(answers map[rune]int, groupSize int) int {
	count := 0

	for _, v := range answers {
		if v == groupSize {
			count++
		}
	}

	return count
}

func main() {
	lines := ReadInput("day-06-input.txt")

	i := 0
	anySum := 0
	allSum := 0

	for i < len(lines) {
		var group []string

		line := lines[i]

		for line != "" {
			group = append(group, line)

			i++

			if i >= len(lines) {
				line = ""
			} else {
				line = lines[i]
			}
		}

		yesCount := CountYesAnswers(group)

		anySum += len(yesCount)
		allSum += CountAllYes(yesCount, len(group))

		i++
	}

	fmt.Println("Any yes sum (part 1):", anySum)
	fmt.Println("All yes sum (part 2):", allSum)
}
