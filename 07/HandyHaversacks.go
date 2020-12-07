package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := getInputFromFile()

	bags := make(map[string]map[string]int)
	for _, line := range input {
		bag := strings.Split(strings.TrimSuffix(line, "."), " contain ")
		subbags := make(map[string]int)

		if bag[1] == "no other bags" {
			continue
		}

		subbagsStr := strings.Split(bag[1], ", ")
		for _, subbag := range subbagsStr {
			amount, err := strconv.Atoi(string(subbag[0]))
			check(err)
			subbags[strings.TrimSuffix(subbag[2:], "s")] = amount
		}
		bags[strings.TrimSuffix(bag[0], "s")] = subbags
	}

	count := 0
	for bag := range bags {
		if ContainsBag(bag, "shiny gold bag", bags) {
			count++
		}
	}

	fmt.Print("Part 1: ", count, "\n")
	fmt.Print("Part 2: ", CountBags("shiny gold bag", bags))

}

func ContainsBag(start string, target string, bags map[string]map[string]int) bool {
	subbags := bags[start]

	for subbag := range subbags {
		if subbag == target {
			return true
		}
		if ContainsBag(subbag, target, bags) {
			return true
		}
	}

	return false
}

func CountBags(start string, bags map[string]map[string]int) int {
	subbags := bags[start]

	amountAll := 0
	for subbag, amount := range subbags {
		amountAll += amount
		amountAll += amount * CountBags(subbag, bags)
	}

	return amountAll
}
