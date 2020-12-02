package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Password struct {
	rule  Rule
	value string
}

type Rule struct {
	min   int
	max   int
	value string
}

func getInputFromFile() []Password {
	file, err := os.Open("input")
	check(err)

	defer file.Close()

	var input []Password
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, parsePassword(scanner.Text()))
	}
	check(scanner.Err())
	return input
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parsePassword(value string) Password {
	splitted := strings.Split(value, ": ")
	pw := Password{value: splitted[1]}
	//rule
	splitted = strings.Split(splitted[0], " ")
	character := splitted[1]
	splitted = strings.Split(splitted[0], "-")
	min, err := strconv.ParseInt(splitted[0], 10, 32)
	check(err)
	max, err := strconv.ParseInt(splitted[1], 10, 32)
	check(err)
	pw.rule = Rule{min: int(min), max: int(max), value: character}
	return pw
}

func main() {
	var wg sync.WaitGroup

	input := getInputFromFile()

	isValidChannel := make(chan Password, len(input))
	for _, pw := range input {
		wg.Add(1)
		go checkPasswordB(&wg, pw, isValidChannel)
	}

	wg.Wait()

	close(isValidChannel)

	fmt.Print(" Count: ", len(isValidChannel))
}

func checkPasswordA(wg *sync.WaitGroup, pw Password, valid chan Password) {

	defer wg.Done()

	amount := strings.Count(pw.value, pw.rule.value)

	if amount >= pw.rule.min && amount <= pw.rule.max {
		valid <- pw
	}
}

func checkPasswordB(wg *sync.WaitGroup, pw Password, valid chan Password) {

	defer wg.Done()

	posA := string(pw.value[pw.rule.min-1]) == pw.rule.value
	posB := string(pw.value[pw.rule.max-1]) == pw.rule.value
	if (posA || posB) && (posA != posB) {
		valid <- pw
	}
}
