package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Line struct {
	Codes   []string
	Display []string
}

func parseInput() []Line {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	var lines []Line
	for _, row := range inputArr {
		if row == "" {
			continue
		}
		parts := strings.Split(row, " | ")

		var line Line
		codes := strings.Split(parts[0], " ")
		for _, code := range codes {
			line.Codes = append(line.Codes, code)
		}

		displays := strings.Split(parts[1], " ")
		for _, display := range displays {
			line.Display = append(line.Display, display)
		}
		lines = append(lines, line)
	}

	return lines
}

func part1() {
	lines := parseInput()

	uniqueNumberSegments := 0
	for _, line := range lines {
		for _, display := range line.Display {
			codeLength := len(display)
			if codeLength == 2 || codeLength == 4 || codeLength == 3 || codeLength == 7 {
				uniqueNumberSegments++
			}
		}
	}

	fmt.Printf("Solution day 1: %d\n", uniqueNumberSegments)
}

func countCommon(codeA, codeB string) int {
	sum := 0
	for _, a := range codeA {
		for _, b := range codeB {
			if a == b {
				sum++
			}
		}
	}
	return sum
}

func buildMapping(codes []string) map[int]string {

	mappings := make(map[int]string)

	for len(mappings) != 10 {
		for _, code := range codes {
			l := len(code)

			if l == 2 {
				mappings[1] = sortString(code)
			} else if l == 4 {
				mappings[4] = sortString(code)
			} else if l == 3 {
				mappings[7] = sortString(code)
			} else if l == 7 {
				mappings[8] = sortString(code)
			} else if l == 5 {
				// 2, 3, 5

				one, ok := mappings[1]
				if !ok {
					continue
				}

				// if this number contains all segments of a 1 it's 3
				if countCommon(code, one) == 2 {
					mappings[3] = sortString(code)
					continue
				}
				four, ok := mappings[4]
				if !ok {
					continue
				}
				if countCommon(code, four) == 2 {
					mappings[2] = sortString(code)
				} else {
					mappings[5] = sortString(code)
				}
			} else if l == 6 {
				// 0, 6, 9
				four, ok := mappings[4]
				if !ok {
					continue
				}

				if countCommon(code, four) == 4 {
					mappings[9] = sortString(code)
					continue
				}

				one, ok := mappings[1]
				if !ok {
					continue
				}
				if countCommon(code, one) == 2 {
					mappings[0] = sortString(code)
				} else {
					mappings[6] = sortString(code)
				}
			}
		}
	}

	return mappings
}

func getFromMapping(mappings map[int]string, code string) string {
	for k, v := range mappings {
		if v == sortString(code) {
			return fmt.Sprintf("%d", k)
		}
	}
	return "0"
}

func part2() {
	lines := parseInput()

	var numbers []int
	for _, line := range lines {
		mappings := buildMapping(line.Codes)
		number := ""
		for _, display := range line.Display {
			number += getFromMapping(mappings, display)
		}
		nr, err := strconv.Atoi(number)
		if err != nil {
			fmt.Printf("Error converting number %s\n", number)
		}
		numbers = append(numbers, nr)
	}

	sum := 0
	for _, n := range numbers {
		sum += n
	}

	fmt.Printf("Solution part 2: %d\n", sum)
}

func main() {
	part1()
	part2()

}
