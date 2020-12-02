package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func getInputFromFile() []int {
	file, err := os.Open("input")
	check(err)

	defer file.Close()

	var input []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i64, err := strconv.ParseInt(scanner.Text(), 10, 32)
		check(err)
		input = append(input, int(i64))
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
	var wg sync.WaitGroup

	input := getInputFromFile()

	for i, first := range input {
		wg.Add(1)
		go find2020InTwoNumbers(&wg, input[i+1:], first)
	}

	for i, first := range input {
		wg.Add(1)
		go find2020InThreeNumbers(&wg, input[i+1:], first)
	}

	wg.Wait()
}

func find2020InTwoNumbers(wg *sync.WaitGroup, numbers []int, first int) {
	defer wg.Done()

	for _, second := range numbers {
		if first+second == 2020 {
			product := first * second
			fmt.Print("\nFound Match ", first, ":", second)
			fmt.Print("\nProduct ", product)
		}
	}
}

func find2020InThreeNumbers(wg *sync.WaitGroup, numbers []int, first int) {
	defer wg.Done()

	for i, second := range numbers {
		if first+second < 2020 {
			wg.Add(1)
			go findThirdNumber(wg, numbers[i+1:], first, second)
		}
	}
}

func findThirdNumber(wg *sync.WaitGroup, numbers []int, first int, second int) {
	defer wg.Done()

	sum := first + second

	for _, third := range numbers {
		if sum+third == 2020 {
			product := first * second * third
			fmt.Print("\nFound Match ", first, ":", second, ":", third)
			fmt.Print("\nProduct ", product)
		}
	}
}
