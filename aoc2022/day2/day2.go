package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part1() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	totScore := 0
	for _, el := range inputArr {
		if el == "" {
			continue
		}

		picks := strings.Split(el, " ")

		score := 0

		// Draw
		if picks[0] == "A" && picks[1] == "X" ||
			picks[0] == "B" && picks[1] == "Y" ||
			picks[0] == "C" && picks[1] == "Z" {
			score += 3

			// Win
		} else if picks[0] == "A" && picks[1] == "Y" ||
			picks[0] == "B" && picks[1] == "Z" ||
			picks[0] == "C" && picks[1] == "X" {
			score += 6
		}

		if picks[1] == "X" {
			score += 1
		} else if picks[1] == "Y" {
			score += 2
		} else {
			score += 3
		}

		totScore += score
	}

	fmt.Printf("Part1: Totscore is: %d\n", totScore)
}

func part2() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	totScore := 0
	for _, el := range inputArr {
		if el == "" {
			continue
		}

		picks := strings.Split(el, " ")

		score := 0

		if picks[1] == "X" {
			if picks[0] == "A" { // Rock
				score += 3 // Sciccor 3 points
			} else if picks[0] == "B" { // Paper
				score += 1 // Rock 1 points
			} else if picks[0] == "C" { // Sciccor
				score += 2 // Paper 1 points
			}
		} else if picks[1] == "Y" {
			if picks[0] == "A" {
				score += 3 + 1 // Draw + Rock
			} else if picks[0] == "B" {
				score += 3 + 2 // Draw + Paper
			} else if picks[0] == "C" {
				score += 3 + 3 // Draw + Sciccor
			}
		} else if picks[1] == "Z" {
			if picks[0] == "A" {
				score += 6 + 2 // Win + Paper
			} else if picks[0] == "B" {
				score += 6 + 3 // Win + Sciccor
			} else if picks[0] == "C" {
				score += 6 + 1 // Win + Rock
			}
		}

		totScore += score
	}

	fmt.Printf("Part2: Totscore is: %d\n", totScore)
}

func main() {
	part1()
	part2()
}
