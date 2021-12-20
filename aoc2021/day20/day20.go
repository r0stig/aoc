package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Coordinate struct {
	row, col int
}

type Input struct {
	algoritm []int
	grid     [][]int
	image    map[Coordinate]int
}

func parseInput() Input {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	parts := strings.Split(string(data), "\n\n")

	var input Input
	input.image = make(map[Coordinate]int)
	for _, letter := range strings.Split(parts[0], "") {
		if letter == "" {
			continue
		}
		number := 0
		if letter == "#" {
			number = 1
		}
		input.algoritm = append(input.algoritm, number)
	}

	for iRow, row := range strings.Split(parts[1], "\n") {
		var inputRow []int
		for iLetter, letter := range strings.Split(row, "") {
			number := 0
			if letter == "#" {
				number = 1
			}
			inputRow = append(inputRow, number)
			if number > 0 {
				input.image[Coordinate{row: iRow, col: iLetter}] = number
			}
		}
		input.grid = append(input.grid, inputRow)
	}

	return input
}

type Bounds struct {
	minCol, maxCol, minRow, maxRow int
}

func getBounds(image map[Coordinate]int, expandWith int) Bounds {
	minCol, minRow := 0, 0
	maxCol, maxRow := 0, 0
	for k := range image {
		if k.col < minCol {
			minCol = k.col
		}
		if k.row < minRow {
			minRow = k.col
		}
		if k.col > maxCol {
			maxCol = k.col
		}
		if k.row > maxRow {
			maxRow = k.row
		}
	}

	if expandWith > 0 {
		return Bounds{
			minCol: minCol - expandWith,
			maxCol: maxCol + expandWith,
			minRow: minRow - expandWith,
			maxRow: maxRow + expandWith,
		}
	} else {
		return Bounds{
			minCol: minCol,
			maxCol: maxCol,
			minRow: minRow,
			maxRow: maxRow,
		}
	}
}

func getIndex(image map[Coordinate]int, bounds Bounds, fillOnOutOfBounds bool, isOdd bool, searchRow, searchCol int) int64 {
	binaryIndex := ""

	for row := 1; row >= -1; row-- {
		for col := 1; col >= -1; col-- {
			checkRow := searchRow - row
			checkCol := searchCol - col
			if fillOnOutOfBounds {
				// In bounds?
				if checkCol >= bounds.minCol && checkCol <= bounds.maxCol &&
					checkRow >= bounds.minRow && checkRow <= bounds.maxRow {
					if _, ok := image[Coordinate{row: checkRow, col: checkCol}]; ok {
						binaryIndex += "1"
					} else {
						binaryIndex += "0"
					}
				} else {
					if isOdd {
						binaryIndex += "1"
					} else {
						binaryIndex += "0"
					}
				}
			} else {
				if _, ok := image[Coordinate{row: checkRow, col: checkCol}]; ok {
					binaryIndex += "1"
				} else {
					binaryIndex += "0"
				}
			}
		}
	}
	index, err := strconv.ParseInt(binaryIndex, 2, 64)
	if err != nil {
		fmt.Printf("Error parsing %s: %v\n", binaryIndex, err)
	}

	return index
}

func enhance(image map[Coordinate]int, algoritm []int, isOdd bool, bounds Bounds) map[Coordinate]int {
	newPixels := make(map[Coordinate]int)

	// In the AoC input of the algoritm the number at index 0 is 1, which means
	// that the "void" will all be flipped to 1 on iteration 1. Then look at the
	// last number in the algoritm (0), which means that on next iteration all of the void
	// wil be flipped to 0.
	// To make this easier just assume the pixels out of bounds is 0 on even iterations
	// and 1 on uneven iterations.

	// This just checks if the input should use the "flipping void" or not by
	// looking at the algoritm.
	fillOnOutOfBounds := true
	if algoritm[0] == 0 {
		fillOnOutOfBounds = false
	}

	for row := bounds.minRow; row <= bounds.maxRow; row++ {
		for col := bounds.minCol; col <= bounds.maxCol; col++ {
			index := getIndex(image, bounds, fillOnOutOfBounds, isOdd, row, col)
			newData := algoritm[index]
			if newData > 0 {
				newPixels[Coordinate{row, col}] = newData
			}
		}
	}
	return newPixels
}

func countLitPixels(image map[Coordinate]int, enhancedBounds Bounds) int {
	sum := 0
	for k := range image {
		if k.col >= enhancedBounds.minCol && k.col <= enhancedBounds.maxCol &&
			k.row >= enhancedBounds.minRow && k.row <= enhancedBounds.maxRow {
			sum++
		}
	}
	return sum
}

func printImage(image map[Coordinate]int) {
	bounds := getBounds(image, 0)

	for row := bounds.minRow; row <= bounds.maxRow; row++ {
		for col := bounds.minCol; col <= bounds.maxCol; col++ {
			if _, ok := image[Coordinate{row, col}]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println("")
	}
}

func part1() {
	input := parseInput()

	steps := 2
	enhancedBounds := getBounds(input.image, steps)

	newImage := input.image
	for i := 0; i < steps; i++ {
		fillOnOutOfBounds := i%2 != 0
		if input.algoritm[0] == 0 {
			fillOnOutOfBounds = i%2 == 0
		}
		newImage = enhance(newImage, input.algoritm, fillOnOutOfBounds, enhancedBounds)
	}

	fmt.Printf("Solution part1: %d\n", countLitPixels(newImage, enhancedBounds))
}

func part2() {
	input := parseInput()

	steps := 50
	enhancedBounds := getBounds(input.image, steps)

	newImage := input.image
	for i := 0; i < steps; i++ {
		newImage = enhance(newImage, input.algoritm, i%2 != 0, enhancedBounds)
	}

	fmt.Printf("Solution part2: %d\n", countLitPixels(newImage, enhancedBounds))
}

func main() {
	part1()
	part2()
}
