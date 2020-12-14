package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := getInputFromFile()
	Part1(input)
	Part2(input)
}

func Part1(input []string) {
	earliest, err := strconv.Atoi(input[0])
	check(err)

	busses := map[int]int{}
	for _, busStr := range strings.Split(input[1], ",") {
		if busStr == "x" {
			continue
		}
		id, err := strconv.Atoi(busStr)
		check(err)
		busses[id] = id - (earliest % id)
	}

	min := -1
	for k, v := range busses {
		if min == -1 || busses[min] > v {
			min = k
		}
	}

	fmt.Print("Part 1: ", busses[min]*min, "\n")
}

func Part2(input []string) {
	var busses []int
	for _, busStr := range strings.Split(input[1], ",") {
		if busStr == "x" {
			busses = append(busses, 1)
			continue
		}
		id, err := strconv.Atoi(busStr)
		check(err)
		busses = append(busses, id)
	}

	isNumberValid := func(number int) int {

		for i, bus := range busses {
			if (number+i)%bus != 0 {
				return i
			}
		}

		return -1
	}

	number := 1
	for {
		index := isNumberValid(number)
		if index == -1 {
			break
		}
		fmt.Print(index, "\n")
		number *= busses[index]
	}

	fmt.Print("Part 2: ", number, "\n")
}
