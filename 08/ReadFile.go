package main

import (
	"bufio"
	"os"
)

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
