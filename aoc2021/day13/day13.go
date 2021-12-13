package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type FoldInstruction struct {
	FoldType   string
	Coordinate int
}

type Input struct {
	Coordinates  []Coordinate // map[string]Coordinate
	Instructions []FoldInstruction
}
type Grid [][]int
type Coordinate []int

func parseInput() Input {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	parts := strings.Split(string(data), "\n\n")
	var input Input

	//coords := make(map[string]Coordinate)
	var coords []Coordinate
	for _, line := range strings.Split(parts[0], "\n") {
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Error converting x %v\n", err)
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("Error converting y %v\n", err)
		}

		coords = append(coords, Coordinate{x, y})
	}

	var instructions []FoldInstruction
	for _, line := range strings.Split(parts[1], "\n") {
		if line == "" {
			continue
		}
		instruction := line[11:]
		parts := strings.Split(instruction, "=")
		coord, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("Error converting fold instruction coord: %v \n", err)
		}
		instructions = append(instructions, FoldInstruction{FoldType: parts[0], Coordinate: coord})
	}

	input.Coordinates = coords
	input.Instructions = instructions

	return input
}

func part1() {
	input := parseInput()

	instruction := input.Instructions[0]
	for i := 0; i < len(input.Coordinates); i++ {
		x := input.Coordinates[i][0]
		y := input.Coordinates[i][1]
		if instruction.FoldType == "y" {
			if y > instruction.Coordinate {
				newY := y - ((y - instruction.Coordinate) * 2)
				input.Coordinates[i] = Coordinate{x, newY}
			}
		} else {
			if x > instruction.Coordinate {
				newX := x - ((x - instruction.Coordinate) * 2)
				input.Coordinates[i] = Coordinate{newX, y}
			}
		}
	}

	uniqueCoords := make(map[string]Coordinate)
	for _, coord := range input.Coordinates {
		uniqueCoords[fmt.Sprintf("%d,%d", coord[0], coord[1])] = coord
	}

	fmt.Printf("Solution part 1: %d\n", len(uniqueCoords))
}

func printGrid(coordinates []Coordinate) {
	maxX := 0
	maxY := 0
	for _, coord := range coordinates {
		if coord[0] > maxX {
			maxX = coord[0]
		}
		if coord[1] > maxY {
			maxY = coord[1]
		}
	}

	for i := 0; i <= maxY; i++ {
		//var xForThisRow []int
		xForThisRow := make(map[int]bool)
		for _, coord := range coordinates {
			if coord[1] == i {
				//xForThisRow = append(xForThisRow, coord[0])
				xForThisRow[coord[0]] = true
			}
		}

		for j := 0; j <= maxX; j++ {
			if xForThisRow[j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func part2() {
	input := parseInput()

	for _, instruction := range input.Instructions {
		for i := 0; i < len(input.Coordinates); i++ {
			x := input.Coordinates[i][0]
			y := input.Coordinates[i][1]
			if instruction.FoldType == "y" {
				if y > instruction.Coordinate {
					newY := y - ((y - instruction.Coordinate) * 2)
					input.Coordinates[i] = Coordinate{x, newY}
				}
			} else {
				if x > instruction.Coordinate {
					newX := x - ((x - instruction.Coordinate) * 2)
					input.Coordinates[i] = Coordinate{newX, y}
				}
			}
		}
	}

	uniqueCoords := make(map[string]Coordinate)
	for _, coord := range input.Coordinates {
		uniqueCoords[fmt.Sprintf("%d,%d", coord[0], coord[1])] = coord
	}

	fmt.Println("Solution part 2 is:")
	printGrid(input.Coordinates)

}

func main() {
	part1()
	part2()
}
