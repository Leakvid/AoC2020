package main

import "fmt"

func main() {
	next := getInputFromFile()
	var curr []string

	for !Equal(next, curr) {
		curr = next
		next = ApplyRules1(curr)
	}

	fmt.Print("Part 1: ", CountOccupied(curr), "\n")

	next = getInputFromFile()
	curr = nil
	for !Equal(next, curr) {
		curr = next
		next = ApplyRules2(curr)
	}

	fmt.Print("Part 2: ", CountOccupied(curr), "\n")
}

func ApplyRules1(layout []string) []string {
	var output []string
	for i := 0; i < len(layout); i++ {
		outputString := ""
		for j := 0; j < len(layout[i]); j++ {
			adj := GetAdjacened(layout, i, j)
			outputByte := layout[i][j]
			switch outputByte {
			case '#':
				if adj == 0 {
					outputByte = 'L'
				}
			case 'L':
				if adj >= 4 {
					outputByte = '#'

				}
			}
			outputString = outputString + string(outputByte)
		}
		output = append(output, outputString)
	}
	return output
}

func ApplyRules2(layout []string) []string {
	var output []string
	for i := 0; i < len(layout); i++ {
		outputString := ""
		for j := 0; j < len(layout[i]); j++ {
			adj := GetInSight(layout, i, j)
			outputByte := layout[i][j]
			switch outputByte {
			case '#':
				if adj == 0 {
					outputByte = 'L'
				}
			case 'L':
				if adj >= 5 {
					outputByte = '#'

				}
			}
			outputString = outputString + string(outputByte)
		}
		output = append(output, outputString)
	}
	return output
}

func GetAdjacened(layout []string, x, y int) int {
	adj := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) {
				if x+i < len(layout) && x+i >= 0 {
					if y+j < len(layout[x+i]) && y+j >= 0 {
						if layout[x+i][y+j] == 'L' {
							adj++
						}
					}
				}
			}
		}
	}

	return adj
}

func GetInSight(layout []string, x, y int) int {
	adj := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) {
				for offset := 1; offset < len(layout) && offset < len(layout[x]); offset++ {
					iOffset := i * offset
					jOffset := j * offset
					if x+iOffset >= len(layout) || x+iOffset < 0 {
						break
					}
					if y+jOffset >= len(layout[x+iOffset]) || y+jOffset < 0 {
						break
					}
					if layout[x+iOffset][y+jOffset] == 'L' {
						adj++
						break
					}
					if layout[x+iOffset][y+jOffset] == '#' {
						break
					}
				}
			}
		}
	}

	return adj
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func CountOccupied(input []string) int {
	occ := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 'L' {
				occ++
			}
		}
	}

	return occ
}
