package main

import (
	"fmt"
	"strings"
)

func main() {
	input := getInputFromFile()

	var answers []map[rune]bool
	answers = append(answers, make(map[rune]bool))
	for _, line := range input {
		if line == "" {
			answers = append(answers, make(map[rune]bool))
		} else {
			for _, char := range line {
				answers[len(answers)-1][char] = true
			}
		}
	}

	sum := 0
	for _, a := range answers {
		sum += len(a)
	}
	fmt.Print("Part 1: ", sum)

	index := 0
	for _, line := range input {
		if line == "" {
			index++
		} else {
			for k := range answers[index] {
				if !strings.Contains(line, string(k)) {
					delete(answers[index], k)
				}
			}
		}
	}

	sum = 0
	for _, a := range answers {
		sum += len(a)
	}
	fmt.Print("\nPart 2: ", sum)
}
