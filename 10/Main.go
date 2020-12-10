package main

import (
	"fmt"
	"sort"
)

func main() {
	input := getInputFromFile()

	sort.Ints(input)

	Part1(input)
	Part2(input)
}

func Part1(input []int) {

	jolts := make(map[int]int)

	for i := 1; i < len(input); i++ {
		jolts[input[i]-input[i-1]]++
	}

	jolts[1]++
	jolts[3]++

	fmt.Print("Part 1: ", jolts[1]*jolts[3], "\n")
}

func Part2(input []int) {
	jolts := map[int]int{0: 1}

	for _, i := range input {
		jolts[i] = jolts[i-1] + jolts[i-2] + jolts[i-3]
	}

	fmt.Println("Part 2: ", jolts[input[len(input)-1]])
}
