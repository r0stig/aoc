package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func parseInput() []string {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")
	return inputArr
}

func closingChunkChar(r rune) rune {
	switch r {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	}
	return '.'
}

func part1() {
	lines := parseInput()

	var illegalChars []rune

	for _, line := range lines {
		var openChunks []rune
		for _, char := range line {
			if char == '(' || char == '[' || char == '{' || char == '<' {
				openChunks = append(openChunks, char)
			} else {
				lastChunk := openChunks[len(openChunks)-1]
				expectedClosingChar := closingChunkChar(lastChunk)
				if expectedClosingChar != char {
					illegalChars = append(illegalChars, char)
					break
				} else {
					openChunks = openChunks[:len(openChunks)-1]
				}
			}
		}
	}

	sum := 0
	for _, char := range illegalChars {
		switch char {
		case ')':
			sum += 3
		case ']':
			sum += 57
		case '}':
			sum += 1197
		case '>':
			sum += 25137
		}
	}
	fmt.Printf("Solution part 1: %d\n", sum)
}

type IncompleteLine struct {
	line       string
	openChunks []rune
}

func part2() {
	lines := parseInput()

	var incompleteLines []IncompleteLine

	for _, line := range lines {
		if line == "" {
			continue
		}
		var openChunks []rune
		isCorrupted := false
		for _, char := range line {
			if char == '(' || char == '[' || char == '{' || char == '<' {
				openChunks = append(openChunks, char)
			} else {
				lastChunk := openChunks[len(openChunks)-1]
				expectedClosingChar := closingChunkChar(lastChunk)
				if expectedClosingChar != char {
					isCorrupted = true
					break
				} else {
					openChunks = openChunks[:len(openChunks)-1]
				}
			}
		}
		if !isCorrupted {
			incompleteLines = append(incompleteLines, IncompleteLine{line, openChunks})
		}
	}

	var scores []int
	for _, line := range incompleteLines {
		closingChunks := ""
		sum := 0
		for i := len(line.openChunks) - 1; i >= 0; i-- {
			closingChar := closingChunkChar(line.openChunks[i])
			sum *= 5
			switch closingChar {
			case ')':
				sum += 1
			case ']':
				sum += 2
			case '}':
				sum += 3
			case '>':
				sum += 4
			}
			closingChunks += string(closingChar)
		}
		scores = append(scores, sum)
	}

	sort.Ints(scores)

	winningScore := scores[len(scores)/2]

	fmt.Printf("Solution part 2: %d\n", winningScore)
}

func main() {
	part1()
	part2()
}
