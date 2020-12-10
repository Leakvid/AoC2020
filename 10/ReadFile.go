package main

import (
	"bufio"
	"os"
	"strconv"
)

func getInputFromFile() []int {
	file, err := os.Open("input")
	check(err)

	defer file.Close()

	var input []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		check(err)
		input = append(input, value)
	}
	check(scanner.Err())
	return input
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
