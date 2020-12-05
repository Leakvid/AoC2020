package main

import (
	"fmt"
	"math"
)

func main() {
	input := getInputFromFile()

	var seats []int
	for _, bsp := range input {
		seats = append(seats, convertBSP(bsp))
	}

	maxId := -1
	for _, id := range seats {
		if id > maxId {
			maxId = id
		}
	}

	fmt.Print("Max Id: ", maxId)

	for _, id := range seats {
		for _, other := range seats {
			if other-id == 2 && !contains(id+1, seats) {
				fmt.Print("My Id: ", id+1)
			}
		}

	}
}

func contains(value int, values []int) bool {
	for _, v := range values {
		if value == v {
			return true
		}
	}
	return false
}

func convertBSP(str string) int {
	row := getPosition(str[0:7], 'B', 'F')
	col := getPosition(str[7:10], 'R', 'L')

	return row*8 + col
}

func getPosition(str string, upper byte, lower byte) int {
	min := 0
	max := int(math.Pow(2, float64(len(str))))

	for i := 0; i < len(str); i++ {
		switch str[i] {
		case upper:
			min += (max - min) / 2
		case lower:
			max -= (max - min) / 2
		default:
			panic("something went wrong")
		}
	}

	return max - 1
}
