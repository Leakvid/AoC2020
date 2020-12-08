package main

import (
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	funcType       string
	param          int
	executionCount int
	loopThreshhold int
}

type executable interface {
	execute() int
}

func (i *instruction) execute() int {
	i.executionCount++
	if i.executionCount >= i.loopThreshhold {
		panic("Threshold exceeded")
	}

	funcs := map[string]func(int) int{
		"jmp": func(val int) int { return val },
		"acc": func(val int) int { accumalator += val; return 1 },
		"nop": func(val int) int { return 1 },
	}

	return funcs[i.funcType](i.param)
}

var accumalator int

func main() {
	var instructions []instruction
	for _, line := range getInputFromFile() {
		args := strings.Split(line, " ")

		value, err := strconv.Atoi(string(args[1]))
		check(err)

		instructions = append(instructions, instruction{args[0], value, 0, 2})
	}

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Print("Part 1 : ", accumalator, "\n")
			}
		}()
		accumalator = 0
		for i := 0; i < len(instructions); i += instructions[i].execute() {
		}
	}()

	runInstructions := func() {
		defer func() {
			if r := recover(); r == nil {
				fmt.Print("Part 2 : ", accumalator, "\n")
			}
		}()
		accumalator = 0
		for i := 0; i < len(instructions); i += instructions[i].execute() {
		}
	}

	swapJmpNop := func(index int) {
		if instructions[index].funcType == "jmp" {
			instructions[index].funcType = "nop"
		} else if instructions[index].funcType == "nop" {
			instructions[index].funcType = "jmp"
		}
	}

	for i := range instructions {
		for j := range instructions {
			instructions[j].executionCount = 0
		}
		swapJmpNop(i)
		runInstructions()
		swapJmpNop(i)
	}
}
