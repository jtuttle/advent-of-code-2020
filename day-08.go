package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	ACC = "acc"
	JMP = "jmp"
	NOP = "nop"
)

type Computer struct {
	ptr int
	accumulator int
	history map[int]bool
}

type Instruction struct {
	op string
	arg int
}

func NewComputer() (Computer) {
	return Computer{
		history: make(map[int]bool),
	}
}

func (c *Computer) Execute(program []Instruction) {
	for {
		_, exist := c.history[c.ptr]

		if exist {
			break
		}

		c.history[c.ptr] = true

		instruct := program[c.ptr]

		switch instruct.op {
		case ACC:
			c.accumulator += instruct.arg
			c.ptr++
		case JMP:
			c.ptr += instruct.arg
		case NOP:
			c.ptr++
		}

		if c.ptr >= len(program) {
			break
		}
	}
}

func ParseProgram(program []string) []Instruction {
	var instructs []Instruction

	for _, line := range program {
		splitLine := strings.Split(line, " ")
		op, argStr := splitLine[0], splitLine[1]

		arg, err := strconv.Atoi(argStr)

		if err != nil {
			panic(err)
		}

		instruct := Instruction{
			op: op,
			arg: arg,
		}

		instructs = append(instructs, instruct)
	}

	return instructs
}

func main() {
	lines := ReadInput("day-08-input.txt")

	program := ParseProgram(lines)

	cpu := NewComputer()
	cpu.Execute(program)

	fmt.Println("Broken exit accumulator (part 1):", cpu.accumulator)

	for i := 0; i < len(program); i++ {
		line := program[i]

		if line.op == JMP || line.op == NOP {
			// swap operation
			if line.op == JMP {
				line.op = NOP
			} else if line.op == NOP {
				line.op = JMP
			}

			program[i] = line

			// run modified program
			cpu = NewComputer()
			cpu.Execute(program)

			if cpu.ptr >= len(program) {
				fmt.Println("Fixed exit accumulator (part 2):", cpu.accumulator)
			}

			// restore originl operation before moving to next line
			if line.op == JMP {
				line.op = NOP
			} else if line.op == NOP {
				line.op = JMP
			}

			program[i] = line
		}
	}

}
