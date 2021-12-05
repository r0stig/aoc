package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func parsePoint(point string) []int {
	parts := strings.Split(point, ",")
	x, err := strconv.Atoi(string(parts[0]))
	if err != nil {
		fmt.Printf("Error converting x %v\n", err)
	}
	y, err := strconv.Atoi(string(parts[1]))
	if err != nil {
		fmt.Printf("Error converting y %v\n", err)
	}

	return []int{x, y}
}

func drawLine(startPoint, endPoint []int, diagonal bool) [][]int {
	x1 := startPoint[0]
	y1 := startPoint[1]

	x2 := endPoint[0]
	y2 := endPoint[1]

	var line [][]int
	if x1 == x2 {
		lowest := int(math.Min(float64(y1), float64(y2)))
		largest := int(math.Max(float64(y1), float64(y2)))
		for i := lowest; i <= largest; i++ {
			line = append(line, []int{x1, i})
		}
	} else if y1 == y2 {
		lowest := int(math.Min(float64(x1), float64(x2)))
		largest := int(math.Max(float64(x1), float64(x2)))
		for i := lowest; i <= largest; i++ {
			line = append(line, []int{i, y2})
		}
	} else if diagonal {
		lowestX := int(math.Min(float64(x1), float64(x2)))
		largestX := int(math.Max(float64(x1), float64(x2)))

		for i := 0; i <= largestX-lowestX; i++ {
			newX := x1 + i
			if x1 > x2 {
				newX = x1 - i
			}
			newY := y1 + i
			if y1 > y2 {
				newY = y1 - i
			}

			line = append(line, []int{newX, newY})

		}
	}
	return line
}

func parseInput(diagonal bool) map[string]int {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")
	var points = make(map[string]int)

	for _, row := range inputArr {
		if row == "" {
			continue
		}
		lines := strings.Split(row, " -> ")
		startPoint := parsePoint(lines[0])
		endPoint := parsePoint(lines[1])

		lineCoordinates := drawLine(startPoint, endPoint, diagonal)
		for _, coordinate := range lineCoordinates {
			points[fmt.Sprintf("%d,%d", coordinate[0], coordinate[1])]++
		}
	}

	return points
}

func findNumberOfOverlappingPoints(points map[string]int) int {
	sum := 0
	for _, v := range points {
		if v > 1 {
			sum++
		}
	}
	return sum
}

func part1() {
	input := parseInput(false)
	overlappingPoints := findNumberOfOverlappingPoints(input)
	fmt.Printf("Solution part 1 %d\n", overlappingPoints)

}

func part2() {
	input := parseInput(true)
	overlappingPoints := findNumberOfOverlappingPoints(input)
	fmt.Printf("Solution part 2 %d\n", overlappingPoints)

}

func main() {
	part1()
	part2()
}
