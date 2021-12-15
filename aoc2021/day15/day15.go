package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

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

func findMin(distances map[string]int, visiteds [][]int) []int {
	minDistance := 999999999999999999
	var minDistancePos []int

	for d, distance := range distances {
		found := false
		for _, visited := range visiteds {
			if d == fmt.Sprintf("%d,%d", visited[1], visited[0]) {
				found = true
			}
		}
		if !found && distance < minDistance {
			minDistance = distance
			parts := strings.Split(d, ",")
			y, _ := strconv.Atoi(parts[0])
			x, _ := strconv.Atoi(parts[1])
			minDistancePos = []int{y, x}
		}
	}
	return minDistancePos
}

func findMinStack(grid [][]int, stack [][]int, distances map[string]int) int {
	minDistance := 999999999999999999
	var minDistancePosIndex int

	for i, s := range stack {
		distance := distances[fmt.Sprintf("%d,%d", s[1], s[0])]
		heuristic := calcHeuristic(grid, s)
		if distance+heuristic < minDistance {
			minDistance = distance + heuristic
			minDistancePosIndex = i
		}
	}

	return minDistancePosIndex
}

func traverse(grid [][]int) {
	var visited [][]int

	var stack [][]int
	stack = append(stack, []int{0, 0})
	came_from := make(map[string]string)

	distances := make(map[string]int)
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			distances[fmt.Sprintf("%d,%d", col, row)] = 999999999999999999
		}
	}
	distances["0,0"] = 0

	for len(stack) > 0 {
		headIndex := findMinStack(grid, stack, distances)
		head := stack[headIndex]
		stack = append(stack[:headIndex], stack[headIndex+1:]...)

		if head[0] == len(grid)-1 && head[1] == len(grid)-1 {
			fmt.Printf("Goal reached!! \n")
			break
		} else {
			visited = append(visited, head)
		}

		neighboars := getNeighboars(grid, head[0], head[1])
		for _, neighboar := range neighboars {
			found := false
			for _, visited := range visited {
				if neighboar[0] == visited[0] && neighboar[1] == visited[1] {
					found = true
					break
				}
			}
			if !found {
				tempDistance := distances[fmt.Sprintf("%d,%d", head[1], head[0])] + grid[neighboar[0]][neighboar[1]]

				if tempDistance < distances[fmt.Sprintf("%d,%d", neighboar[1], neighboar[0])] {
					distances[fmt.Sprintf("%d,%d", neighboar[1], neighboar[0])] = tempDistance
					came_from[fmt.Sprintf("%d,%d", neighboar[1], neighboar[0])] = fmt.Sprintf("%d,%d", head[1], head[0])
				}

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

	fmt.Printf("Distances %v\n", distances[fmt.Sprintf("%d,%d", len(grid)-1, len(grid)-1)])
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
