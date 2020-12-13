package main

import (
	"fmt"
	"strconv"
)

const (
	NORTH = "N"
	SOUTH = "S"
	EAST = "E"
	WEST = "W"
	LEFT = "L"
	RIGHT = "R"
	FORWARD = "F"
)

type Instruction struct {
	action string
	value int
}

type Vec2 struct {
	x int
	y int
}

func ParseInstructions(lines []string) []Instruction {
	var instructs []Instruction

	for _, line := range lines {
		action := string(line[0])

		valStr := line[1:len(line)]

		val, err := strconv.Atoi(valStr)

		if err != nil {
			panic(err)
		}

		instruct := Instruction{
			action: action,
			value: val,
		}

		instructs = append(instructs, instruct)
	}

	return instructs
}

func Move(pos Vec2, direction string, amount int) Vec2 {
	newPos := Vec2{
		x: pos.x,
		y: pos.y,
	}

	switch direction {
	case NORTH:
		newPos.y += amount
	case SOUTH:
		newPos.y -= amount
	case EAST:
		newPos.x += amount
	case WEST:
		newPos.x -= amount
	}

	return newPos
}

func Rotate(waypoint Vec2, clockwise bool, amount int) Vec2 {
	for i := 0; i < amount / 90; i++ {
		// perform 90-degree rotation trick
		if clockwise {
			oldX := waypoint.x
			waypoint.x = waypoint.y
			waypoint.y = -oldX
		} else {
			oldX := waypoint.x
			waypoint.x = -waypoint.y
			waypoint.y = oldX
		}
	}

	return waypoint
}

func FollowInstructions(instructs []Instruction, waypointStart Vec2, moveShip bool) Vec2 {
	ship := Vec2{
		x: 0,
		y: 0,
	}

	// waypoint coords are relative to ship
	waypoint := waypointStart

	for _, instruct := range instructs {
		switch instruct.action {
		case LEFT:
			waypoint = Rotate(waypoint, false, instruct.value)
		case RIGHT:
			waypoint = Rotate(waypoint, true, instruct.value)
		case FORWARD:
			for i := 0; i < instruct.value; i++ {
				ship.x = ship.x + waypoint.x
				ship.y = ship.y + waypoint.y
			}
		default:
			if moveShip {
				ship = Move(ship, instruct.action, instruct.value)
			} else {
				waypoint = Move(waypoint, instruct.action, instruct.value)
			}
		}
	}

	return ship
}

func main() {
	lines := ReadInput("day-12-input.txt")

	instructs := ParseInstructions(lines)

	ship := FollowInstructions(instructs, Vec2{ x: 1, y: 0 }, true)
	md := ManhattanDistance(0, 0, ship.x, ship.y)

	fmt.Println("Destination Manhattan distance (part 1):", md)

	ship = FollowInstructions(instructs, Vec2{ x: 10, y: 1 }, false)
	md = ManhattanDistance(0, 0, ship.x, ship.y)

	fmt.Println("Destination Manhattan distance (part 2):", md)
}
