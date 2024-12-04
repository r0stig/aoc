package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func part1() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	var leftNumbers []int
	var rightNumbers []int
	for _, el := range inputArr {
		leftNumber := ""
		rightNumber := ""
		leftNumberCompleted := false
		for _, ch := range el {
			if ch != ' ' {
				if !leftNumberCompleted {
					leftNumber += string(ch)
				} else {
					rightNumber += string(ch)
				}
			} else {
				leftNumberCompleted = true
			}
		}

		left, err := strconv.Atoi(leftNumber)
		if err != nil {
			fmt.Printf("Error converting number %v\n", err)
			continue
		}

		right, err := strconv.Atoi(rightNumber)
		if err != nil {
			fmt.Printf("Error converting number %v\n", err)
			continue
		}

		leftNumbers = append(leftNumbers, left)
		rightNumbers = append(rightNumbers, right)

	}

	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	distanceSum := 0
	for i := 0; i < len(leftNumbers); i++ {
		distance := int(math.Abs(float64(leftNumbers[i]) - float64(rightNumbers[i])))
		distanceSum += distance
	}

	fmt.Printf("Distance sum is %v\n", distanceSum)
}

func part2() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	var leftNumbers []int
	var rightNumbers []int
	for _, el := range inputArr {
		leftNumber := ""
		rightNumber := ""
		leftNumberCompleted := false
		for _, ch := range el {
			if ch != ' ' {
				if !leftNumberCompleted {
					leftNumber += string(ch)
				} else {
					rightNumber += string(ch)
				}
			} else {
				leftNumberCompleted = true
			}
		}

		left, err := strconv.Atoi(leftNumber)
		if err != nil {
			fmt.Printf("Error converting number %v\n", err)
			continue
		}

		right, err := strconv.Atoi(rightNumber)
		if err != nil {
			fmt.Printf("Error converting number %v\n", err)
			continue
		}

		leftNumbers = append(leftNumbers, left)
		rightNumbers = append(rightNumbers, right)

	}

	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	duplicateScore := 0
	for _, l := range leftNumbers {
		duplicateCount := 0
		for _, r := range rightNumbers {
			if l == r {
				duplicateCount++
			}
		}
		duplicateScore += l * duplicateCount
	}

	fmt.Printf("Similarity score is  %v\n", duplicateScore)
}

func main() {
	part1()
	part2()
}
