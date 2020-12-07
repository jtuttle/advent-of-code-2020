package main

import (
	"fmt"
	"strconv"
	"strings"
)

// count number of colors that can be parents by tracing parents through map
func CountParents(child string, parentMap map[string]map[string]bool) int {
	count := 0
	visited := make(map[string]bool)

	queue := []string{ child }

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for k, _ := range(parentMap[current]) {
			_, ok := visited[k]

			// follow (and count) parent link if not yet visited
			if !ok {
				queue = append(queue, k)
				visited[k] = true
				count += 1
			}
		}
	}

	return count
}

// count number of bags that must be contained by a given parent
func CountChildren(parent string, childMap map[string][]string) int {
	count := 0

	queue := []string{ parent }

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, child := range childMap[current] {
			childSplit := strings.Split(child, " ")

			childCount, err := strconv.Atoi(childSplit[0])

			if err != nil {
				panic(err)
			}

			count += childCount

			childColor := strings.Join(childSplit[1:3], " ")

			for i := 0; i < childCount; i++ {
				queue = append(queue, childColor)
			}
		}
	}

	return count
}

func main() {
	lines := ReadInput("day-07-input.txt")

	parentMap := make(map[string]map[string]bool)
	childMap := make(map[string][]string)

	for _, line := range lines {
		rule := strings.Trim(line, ".")
		ruleSplit := strings.Split(rule, " bags contain ")

		parent := ruleSplit[0]
		children := strings.Split(ruleSplit[1], ", ")

		for _, child := range children {
			// split child string into useful parts
			childSplit := strings.Split(child, " ")
			count := childSplit[0]
			color := strings.Join(childSplit[1:3], " ")

			// init nested map if this is first instance of key
			_, ok := parentMap[color]

			if !ok {
				parentMap[color] = make(map[string]bool)
			}

			// record parent relationship
			parentMap[color][parent] = true

			// record child relationship(s)
			if child == "no other bags" {
				continue
			}

			childMap[parent] = append(childMap[parent], count + " " + color)
		}
	}

	parentCount := CountParents("shiny gold", parentMap)

	fmt.Println("Count of colors that hold 'shiny gold' (part 1):", parentCount)

	childCount := CountChildren("shiny gold", childMap)

	fmt.Println("Count of bags required in 'shiny gold' (part 2):", childCount)
}
