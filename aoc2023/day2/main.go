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

	constraintRed := 12
	constraintGreen := 13
	constraintBlue := 14

	var gamesWithinConstraints []int
	for _, el := range inputArr {
		parts := strings.Split(el, ":")
		gameParts := strings.Split(parts[0], " ")
		gameId, err := strconv.Atoi(gameParts[1])
		if err != nil {
			fmt.Printf("Error converting %d\n", gameParts[1])
		}

		bagConfigs := strings.Split(parts[1], ";")
		isWithinConstraints := true

		for _, config := range bagConfigs {
			entries := strings.Split(config, ",")

			for _, entry := range entries {
				color := strings.Split(strings.Trim(entry, " "), " ")
				number, err := strconv.Atoi(color[0])
				if err != nil {
					fmt.Printf("Error converting color %d\n", color[0])
				}

				if color[1] == "red" {
					if number > constraintRed {
						isWithinConstraints = false
					}
				} else if color[1] == "green" {
					if number > constraintGreen {
						isWithinConstraints = false
					}
				} else if color[1] == "blue" {
					if number > constraintBlue {
						isWithinConstraints = false
					}
				}
			}
		}

		if isWithinConstraints {
			gamesWithinConstraints = append(gamesWithinConstraints, gameId)
		}
	}

	sum := 0

	for _, gameId := range gamesWithinConstraints {
		sum += gameId
	}

	fmt.Printf("Answer part1: %d\n", sum)

}

func part2() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	var gamesPowers []int
	for _, el := range inputArr {
		parts := strings.Split(el, ":")

		bagConfigs := strings.Split(parts[1], ";")

		maxRed := 0
		maxBlue := 0
		maxGreen := 0

		for _, config := range bagConfigs {
			entries := strings.Split(config, ",")

			for _, entry := range entries {
				color := strings.Split(strings.Trim(entry, " "), " ")
				number, err := strconv.Atoi(color[0])
				if err != nil {
					fmt.Printf("Error converting color %d\n", color[0])
				}

				if color[1] == "red" {
					if number > maxRed {
						maxRed = number
					}
				} else if color[1] == "green" {
					if number > maxGreen {
						maxGreen = number
					}
				} else if color[1] == "blue" {
					if number > maxBlue {
						maxBlue = number
					}
				}
			}
		}

		gamesPowers = append(gamesPowers, maxRed*maxBlue*maxGreen)

	}

	sum := 0

	for _, nr := range gamesPowers {
		sum += nr
	}

	fmt.Printf("Answer part2: %d\n", sum)

}

func main() {
	part1()
	part2()
}
