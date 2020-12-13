package main

import (
	"fmt"
	"math"
)

const (
	EMPTY = 'L'
	OCCUPIED = '#'
	FLOOR = '.'
)

type Vec2 struct {
	x int
	y int
}

type Grid struct {
	width int
	height int
	data [][]rune
}

func NewGrid(width, height int) (Grid) {
	data := make([][]rune, height)

	for i := 0; i < height; i++ {
		data[i] = make([]rune, width)
	}

	return Grid{
		width: width,
		height: height,
		data: data,
	}
}

func NeighborOccupied(grid Grid, coord Vec2, direction Vec2, distance int) bool {
	for i := 1; i <= distance; i++ {
		nx := coord.x + (direction.x * i)
		ny := coord.y + (direction.y * i)

		// out of bounds
		if nx < 0 || nx >= grid.width || ny < 0 || ny >= grid.height {
			break
		}

		if grid.data[ny][nx] == EMPTY {
			return false
		}

		if grid.data[ny][nx] == OCCUPIED {
			return true
		}
	}

	return false
}

func CountOccupiedNeighbors(grid Grid, coord Vec2, occupiedDist int) int {
	count := 0

	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if y == 0 && x == 0 {
				continue
			}

			direction := Vec2{
				x: x,
				y: y,
			}

			if NeighborOccupied(grid, coord, direction, occupiedDist) {
				count++
			}
		}
	}

	return count
}

func Step(grid Grid, occupiedDist int, occupiedCount int) Grid {
	newGrid := NewGrid(grid.width, grid.height)

	for y, row := range grid.data {
		for x, _ := range row {
			state := grid.data[y][x]

			if state == FLOOR {
				newGrid.data[y][x] = FLOOR
				continue
			}

			coord := Vec2{
				x: x,
				y: y,
			}

			occupiedNeighbors := CountOccupiedNeighbors(grid, coord, occupiedDist)

			if state == EMPTY && occupiedNeighbors == 0 {
				newGrid.data[y][x] = OCCUPIED
			} else if state == OCCUPIED && occupiedNeighbors >= occupiedCount {
				newGrid.data[y][x] = EMPTY
			} else {
				newGrid.data[y][x] = state
			}
		}
	}

	return newGrid
}

func StepUntilStable(grid Grid, occupiedDist int, occupiedCount int) Grid {
	nextGrid := NewGrid(grid.width, grid.height)

	for {
		nextGrid = Step(grid, occupiedDist, occupiedCount)

		if GridsEqual(grid, nextGrid) {
			break
		}

		grid = nextGrid
	}

	return grid
}

func GridsEqual(gridA Grid, gridB Grid) bool {
	for y := 0; y < len(gridA.data); y++ {
		for x := 0; x < len(gridA.data[y]); x++ {
			if gridA.data[y][x] != gridB.data[y][x] {
				return false
			}
		}
	}

	return true
}

func PrintGrid(grid Grid) {
	for _, row := range grid.data {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Print("\n")
	}
}

func CountOccupied(grid Grid) int {
	count := 0

	for _, row := range grid.data {
		for _, col := range row {
			if col == OCCUPIED {
				count++
			}
		}
	}

	return count
}

func main() {
	lines := ReadInput("day-11-input.txt")

	inputGrid := NewGrid(len(lines[0]), len(lines))

	for i, line := range lines {
		for j, char := range line {
			inputGrid.data[i][j] = char
		}
	}

	adjacentGrid := StepUntilStable(inputGrid, 1, 4)

	fmt.Println("Occupied seats when stable (part 1):", CountOccupied(adjacentGrid))

	lineOfSightGrid := StepUntilStable(inputGrid, math.MaxInt32, 5)
	
	fmt.Println("Occupied seats when stable (part 2):", CountOccupied(lineOfSightGrid))
}
