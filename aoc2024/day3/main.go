package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func part1() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	matches := re.FindAllStringSubmatch(string(data), -1)

	sum := 0
	for _, match := range matches {
		if len(match) > 2 {
			number1, _ := strconv.Atoi(match[1])
			number2, _ := strconv.Atoi(match[2])

			sum += number1 * number2
		}
	}

	fmt.Printf("Sum is %d\n", sum)
}

func part2() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}

	re := regexp.MustCompile(`(mul\((\d+),(\d+)\))|don\'t\(\)|do\(\)`)

	matches := re.FindAllStringSubmatch(string(data), -1)

	sum := 0
	isEnabled := true
	for _, match := range matches {
		if match[0] == "do()" {
			isEnabled = true
		} else if match[0] == "don't()" {
			isEnabled = false
		} else if len(match) > 2 && isEnabled {
			number1, _ := strconv.Atoi(match[2])
			number2, _ := strconv.Atoi(match[3])

			sum += number1 * number2
		}
	}

	fmt.Printf("Sum is %d\n", sum)
}

func main() {
	part1()
	part2()
}
