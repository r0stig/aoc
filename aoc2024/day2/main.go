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

	var reports [][]int
	for _, el := range inputArr {
		var numbers []int
		number := ""
		for index, ch := range el {
			if ch != ' ' {
				number += string(ch)
			}
			if ch == ' ' || index == len(el)-1 {
				n, err := strconv.Atoi(number)
				if err != nil {
					fmt.Printf("Error converting number %v %v\n", number, err)
				}
				numbers = append(numbers, n)
				number = ""
			}
		}
		reports = append(reports, numbers)
	}

	safeCount := 0
	for _, report := range reports {
		if isReportValid(report) {
			safeCount++
		}
	}

	fmt.Printf("Safe reports: %d\n", safeCount)
}

func isReportValid(report []int) bool {
	isSafe := true
	prevDiff := 0
	for i := 0; i < len(report)-1; i++ {
		diff := report[i] - report[i+1]

		if diff > 3 || diff < -3 || diff == 0 {
			isSafe = false
		}
		if i > 0 {
			if diff > 0 && prevDiff < 0 {
				isSafe = false
			} else if diff < 0 && prevDiff > 0 {
				isSafe = false
			}
		}

		prevDiff = diff
	}

	return isSafe
}

func part2() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	var reports [][]int
	for _, el := range inputArr {
		var numbers []int
		number := ""
		for index, ch := range el {
			if ch != ' ' {
				number += string(ch)
			}
			if ch == ' ' || index == len(el)-1 {
				n, err := strconv.Atoi(number)
				if err != nil {
					fmt.Printf("Error converting number %v %v\n", number, err)
				}
				numbers = append(numbers, n)
				number = ""
			}
		}
		reports = append(reports, numbers)
	}

	safeCount := 0
	for _, report := range reports {
		isSafe := isReportValid(report)

		if !isSafe {
			for i := 0; i < len(report); i++ {
				dup := append([]int{}, report...)
				newReport := append(dup[:i], dup[i+1:]...)
				if isReportValid(newReport) {
					isSafe = true
					break
				}
			}
		}

		if isSafe {
			safeCount++
		}
	}

	fmt.Printf("Safe reports: %d\n", safeCount)
}

func main() {
	part1()
	part2()
}
