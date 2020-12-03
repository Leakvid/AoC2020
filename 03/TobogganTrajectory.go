package main

import (
	"bufio"
	"fmt"
	"os"
)

type Slope struct {
	x int
	y int
}

func getInputFromFile() []string {
	file, err := os.Open("input")
	check(err)

	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	check(scanner.Err())
	return input
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	input := getInputFromFile()

	slopes := []Slope{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	}

	product := 1

	for _, s := range slopes {
		trees := 0

		for i := 0; i*s.x < len(input); i++ {
			x := i * s.x
			y := i * s.y % len(input[x])
			if input[x][y] == '#' {
				trees++
			}
		}

		product *= trees
		fmt.Print(s.x, ":", s.y, " Count: ", trees, "\n")
	}

	fmt.Print("Product: ", product, "\n")
}
