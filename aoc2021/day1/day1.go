package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part1() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	increases := 0
	prevVal := -1
	for i := 0; i < len(inputArr); i++ {
		curNumber, err := strconv.Atoi(inputArr[i])
		if err != nil {
			continue
		}
		if prevVal == -1 {
			prevVal = curNumber
			continue
		}

		if curNumber > prevVal {
			increases++
		}

		prevVal = curNumber
	}

	fmt.Printf("Part 1, solution is %d\n", increases)
}

func part2() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	prevWindowSum := -1
	increases := 0

	for i := 2; i < len(inputArr); i++ {
		curNumber1, err := strconv.Atoi(inputArr[i-2])
		if err != nil {
			continue
		}
		curNumber2, err := strconv.Atoi(inputArr[i-1])
		if err != nil {
			continue
		}
		curNumber3, err := strconv.Atoi(inputArr[i])
		if err != nil {
			continue
		}
		curWindowSum := curNumber1 + curNumber2 + curNumber3
		if prevWindowSum == -1 {
			prevWindowSum = curWindowSum
		} else {
			if curWindowSum > prevWindowSum {
				increases++
			}
			prevWindowSum = curWindowSum
		}
	}

	fmt.Printf("Part 2, solution is %d\n", increases)
}

func main() {
	part1()
	part2()
}
