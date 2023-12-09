package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func predict(input []int) (int, int) {
	var extrapolates [][]int
	extrapolates = append(extrapolates, input)
	for {
		lastEntry := extrapolates[len(extrapolates)-1]
		allZeroes := true
		for _, value := range lastEntry {
			if value != 0 {
				allZeroes = false
			}
		}
		if allZeroes {
			break
		}

		var newExtrapolare []int
		for i := 0; i < len(lastEntry)-1; i++ {
			val1 := lastEntry[i]
			val2 := lastEntry[i+1]

			newExtrapolare = append(newExtrapolare, val2-val1)
		}

		extrapolates = append(extrapolates, newExtrapolare)
	}

	for i := len(extrapolates) - 1; i >= 0; i-- {
		hasBelow := i < len(extrapolates)-1
		increaseWith := -1
		if !hasBelow {
			increaseWith = 0
		} else {
			belowEntry := extrapolates[i+1]
			increaseWith = belowEntry[len(belowEntry)-1]
		}
		currentEntry := extrapolates[i]
		extrapolates[i] = append(currentEntry, currentEntry[len(currentEntry)-1]+increaseWith)
	}

	for i := len(extrapolates) - 1; i >= 0; i-- {
		hasBelow := i < len(extrapolates)-1
		increaseWith := -1
		if !hasBelow {
			increaseWith = 0
		} else {
			belowEntry := extrapolates[i+1]
			increaseWith = belowEntry[0]
		}
		currentEntry := extrapolates[i]
		extrapolates[i] = append([]int{currentEntry[0] - increaseWith}, extrapolates[i]...)
	}

	return extrapolates[0][len(extrapolates[0])-1], extrapolates[0][0]
}

func part1() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	lines := strings.Split(string(data), "\n")

	var input [][]int

	for _, line := range lines {
		parts := strings.Split(line, " ")
		var dataPoints []int
		for _, part := range parts {
			nr, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("Could not convert %s\n", part)
			}
			dataPoints = append(dataPoints, nr)
		}
		input = append(input, dataPoints)
	}

	sumPart1 := 0
	sumPart2 := 0
	for _, value := range input {
		p1, p2 := predict(value)
		sumPart1 += p1
		sumPart2 += p2
	}

	fmt.Printf("Answer part 1: %d\n", sumPart1)
	fmt.Printf("Answer part 2: %d\n", sumPart2)

}

func main() {
	part1()
}
