package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	row int
	col int
}

func parseInput() [][]int {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")
	var grid [][]int
	for _, line := range inputArr {
		if line == "" {
			continue
		}
		var row []int
		for _, char := range strings.Split(line, "") {
			nr, err := strconv.Atoi(char)
			if err != nil {
				fmt.Printf("Error converting %s\n", err)
			}
			row = append(row, nr)
		}
		grid = append(grid, row)
	}
	return grid
}

func getNeighboars(grid [][]int, row, col int) [][]int {
	var neighboars [][]int

	if row > 0 {
		neighboars = append(neighboars, []int{row - 1, col})
	}

	if col > 0 {
		neighboars = append(neighboars, []int{row, col - 1})
	}
	if col < len(grid[row])-1 {
		neighboars = append(neighboars, []int{row, col + 1})
	}
	if row < len(grid)-1 {
		neighboars = append(neighboars, []int{row + 1, col})
	}

	return neighboars
}

func calcHeuristic(grid [][]int, pos []int) int {
	goal := len(grid) - 1

	dX := int(math.Abs(float64(pos[0]) - float64(goal)))
	dY := int(math.Abs(float64(pos[1]) - float64(goal)))

	return dX + dY
}

func findMinStack(stack [][]int, distances map[Point]int) int {
	minDistance := 999999999999999999
	var minDistancePosIndex int

	for i, s := range stack {
		distance := distances[Point{row: s[0], col: s[1]}]
		if distance < minDistance {
			minDistance = distance
			minDistancePosIndex = i
		}
	}

	return minDistancePosIndex
}

func traverse(grid [][]int) {
	var stack [][]int
	stack = append(stack, []int{0, 0})
	came_from := make(map[Point]string)

	distancesWithHeuristics := make(map[Point]int)
	distances := make(map[Point]int)
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			distances[Point{row, col}] = 999999999999999999
			distancesWithHeuristics[Point{row, col}] = 999999999999999999
		}
	}
	distances[Point{row: 0, col: 0}] = 0
	distancesWithHeuristics[Point{row: 0, col: 0}] = 0

	for len(stack) > 0 {
		headIndex := findMinStack(stack, distancesWithHeuristics)
		head := stack[headIndex]
		stack = append(stack[:headIndex], stack[headIndex+1:]...)

		if head[0] == len(grid)-1 && head[1] == len(grid)-1 {
			fmt.Printf("Goal reached!! \n")
			break
		}

		neighboars := getNeighboars(grid, head[0], head[1])
		for _, neighboar := range neighboars {
			tempDistance := distances[Point{row: head[0], col: head[1]}] + grid[neighboar[0]][neighboar[1]]

			if tempDistance < distances[Point{row: neighboar[0], col: neighboar[1]}] {
				distances[Point{row: neighboar[0], col: neighboar[1]}] = tempDistance
				distancesWithHeuristics[Point{row: neighboar[0], col: neighboar[1]}] = tempDistance + calcHeuristic(grid, neighboar)
				came_from[Point{row: neighboar[0], col: neighboar[1]}] = fmt.Sprintf("%d,%d", head[1], head[0])

				alreadyInStack := false
				for _, s := range stack {
					if neighboar[0] == s[0] && neighboar[1] == s[1] {
						alreadyInStack = true
						break
					}
				}
				if !alreadyInStack {
					stack = append(stack, neighboar)
				}
			}
		}
	}

	fmt.Printf("Distances %v\n", distances[Point{row: len(grid) - 1, col: len(grid) - 1}])
}

func part1() {
	input := parseInput()
	traverse(input)
}

func part2() {
	input := parseInput()

	originalGridSize := len(input)
	gridSize := len(input) * 5
	newGrid := make([][]int, gridSize)
	for i := range newGrid {
		newGrid[i] = make([]int, gridSize)
	}

	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			addition := row/originalGridSize + col/originalGridSize
			newVal := (input[row%(originalGridSize)][col%(originalGridSize)] + addition)
			if newVal > 9 {
				newVal = (newVal % 10) + 1
			}
			newGrid[row][col] = newVal
		}
	}

	traverse(newGrid)
}

func main() {
	fmt.Println("Solution part 1:")
	part1()
	fmt.Println("Solution part 2:")
	part2()
}
