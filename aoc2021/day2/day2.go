package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Command struct {
	Name  string
	Count int
}

func parseInput() []Command {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	var commands []Command
	for _, row := range inputArr {
		if row == "" {
			continue
		}
		parts := strings.SplitN(row, " ", 2)
		cmd := parts[0]
		number := parts[1]
		convertedNumber, err := strconv.Atoi(number)
		if err != nil {
			continue
		}
		commands = append(commands, Command{Name: cmd, Count: convertedNumber})
	}
	return commands
}

func part1() {
	commands := parseInput()

	horPos := 0
	depth := 0
	for _, command := range commands {
		if command.Name == "forward" {
			horPos += command.Count
		} else if command.Name == "down" {
			depth += command.Count
		} else if command.Name == "up" {
			depth -= command.Count
		}
	}

	fmt.Printf("Part 1, solution is %d\n", horPos*depth)
}

func part2() {
	commands := parseInput()

	aim := 0
	horPos := 0
	depth := 0
	for _, command := range commands {
		if command.Name == "forward" {
			horPos += command.Count
			depth += aim * command.Count
		} else if command.Name == "down" {
			aim += command.Count
		} else if command.Name == "up" {
			aim -= command.Count
		}
	}

	fmt.Printf("Part 2, solution is %d\n", horPos*depth)
}

func main() {
	part1()
	part2()
}
