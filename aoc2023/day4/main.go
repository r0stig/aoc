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

	var cardPoints []int
	for _, el := range inputArr {
		cardParts := strings.Split(el, ":")
		numberParts := strings.Split(cardParts[1], "|")

		winningNumbersStr := strings.Split(strings.Trim(numberParts[0], " "), " ")
		elfNumbersStr := strings.Split(strings.Trim(numberParts[1], " "), " ")

		var winningNumbers []int
		for _, str := range winningNumbersStr {
			i, err := strconv.Atoi(str)
			if err != nil {
				continue
			}
			winningNumbers = append(winningNumbers, i)
		}

		var elfNumbers []int
		for _, str := range elfNumbersStr {
			i, err := strconv.Atoi(str)
			if err != nil {
				continue
			}
			elfNumbers = append(elfNumbers, i)
		}

		points := 0
		for _, winningNumber := range winningNumbers {
			for _, elfNumber := range elfNumbers {
				if winningNumber == elfNumber {
					if points == 0 {
						points = 1
					} else {
						points = points * 2
					}
				}
			}
		}

		if points > 0 {
			cardPoints = append(cardPoints, points)
		}

	}

	sum := 0
	for _, nr := range cardPoints {
		sum += nr
	}

	fmt.Printf("Answer part1: %d\n", sum)
}

type Entry struct {
	winningNumbers []int
	elfNumbers     []int
}

func check(index int, entries []Entry) int {
	matchingNumbers := 0
	for _, winningNumber := range entries[index].winningNumbers {
		for _, elfNumber := range entries[index].elfNumbers {
			if winningNumber == elfNumber {
				matchingNumbers++
			}
		}
	}

	sum := 0
	for i := 0; i < matchingNumbers; i++ {
		sum += 1 + check(index+i+1, entries)
	}
	return sum
}

func part2() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	var entries []Entry
	for _, el := range inputArr {
		cardParts := strings.Split(el, ":")
		numberParts := strings.Split(cardParts[1], "|")

		winningNumbersStr := strings.Split(strings.Trim(numberParts[0], " "), " ")
		elfNumbersStr := strings.Split(strings.Trim(numberParts[1], " "), " ")

		var winningNumbers []int
		for _, str := range winningNumbersStr {
			i, err := strconv.Atoi(str)
			if err != nil {
				continue
			}
			winningNumbers = append(winningNumbers, i)
		}

		var elfNumbers []int
		for _, str := range elfNumbersStr {
			i, err := strconv.Atoi(str)
			if err != nil {
				continue
			}
			elfNumbers = append(elfNumbers, i)
		}
		entries = append(entries, Entry{
			winningNumbers,
			elfNumbers,
		})
	}

	sum := 0
	for i := 0; i < len(entries); i++ {
		sum += 1 + check(i, entries)
	}

	fmt.Printf("Answer part2: %d\n", sum)
}

func main() {
	part1()
	part2()
}
