package main

import "fmt"

func main() {
	input := getInputFromFile()
	preambleSize := 25

	for i, number := range input[preambleSize:] {
		preambleEnd := preambleSize + i
		preambleStart := preambleEnd - preambleSize
		if preambleStart < 0 {
			preambleStart = 0
		}

		if !FindCombination(input[preambleStart:preambleEnd], number) {
			fmt.Print("Part 1: ", number, "\n")
			set := FindContiguousSet(input, number)
			fmt.Print("Part 2: ", GetSumOfMinMax(set), "\n")
		}
	}
}

func FindCombination(numbers []int, target int) bool {

	for i, first := range numbers {
		for _, second := range numbers[i:] {
			if first+second == target {
				return true
			}
		}
	}

	return false
}

func FindContiguousSet(numbers []int, target int) []int {

	for i, _ := range numbers {
		sum := 0
		for j, next := range numbers[i:] {
			sum += next

			if sum > target {
				break
			}

			if sum == target {
				return numbers[i : i+j+1]
			}
		}
	}

	return nil
}

func GetSumOfMinMax(numbers []int) int {
	min := numbers[0]
	max := numbers[0]
	for _, number := range numbers {
		if min > number {
			min = number
		}
		if max < number {
			max = number
		}
	}

	return min + max
}
