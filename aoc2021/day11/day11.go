package main

import (
	"fmt"
	"io/ioutil"
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

	if row > 0 && col > 0 {
		neighboars = append(neighboars, []int{row - 1, col - 1})
	}
	if row > 0 {
		neighboars = append(neighboars, []int{row - 1, col})
	}
	if row > 0 && col < len(grid[row])-1 {
		neighboars = append(neighboars, []int{row - 1, col + 1})
	}

	if col > 0 {
		neighboars = append(neighboars, []int{row, col - 1})
	}
	if col < len(grid[row])-1 {
		neighboars = append(neighboars, []int{row, col + 1})
	}

	if row < len(grid)-1 && col > 0 {
		neighboars = append(neighboars, []int{row + 1, col - 1})
	}
	if row < len(grid)-1 {
		neighboars = append(neighboars, []int{row + 1, col})
	}

	if row < len(grid)-1 && col < len(grid[row])-1 {
		neighboars = append(neighboars, []int{row + 1, col + 1})
	}

	return neighboars
}

func hasFlashed(flashed [][]int, row, col int) bool {
	for _, f := range flashed {
		if f[0] == row && f[1] == col {
			return true
		}
	}
	return false
}

func flash(grid [][]int, flashed *[][]int, row, col int) int {
	flashes := 1
	*flashed = append(*flashed, []int{row, col})
	neighboars := getNeighboars(grid, row, col)

	for _, neighboar := range neighboars {
		nRow := neighboar[0]
		nCol := neighboar[1]
		grid[nRow][nCol]++

		if grid[nRow][nCol] > 9 && !hasFlashed(*flashed, nRow, nCol) {
			flashes += flash(grid, flashed, nRow, nCol)
		}

	}
	return flashes
}

func tick(grid [][]int) int {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			grid[row][col]++
		}
	}

	var flashed [][]int
	flashes := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			level := grid[row][col]
			if level > 9 && !hasFlashed(flashed, row, col) {
				flashes += flash(grid, &flashed, row, col)
			}
		}
	}

	for _, fl := range flashed {
		grid[fl[0]][fl[1]] = 0
	}
	return flashes
}

func printGrid(grid [][]int) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			fmt.Printf("%2d ", grid[row][col])
		}
		fmt.Println("")
	}
}

func part1() {
	grid := parseInput()

	sum := 0
	for i := 0; i < 100; i++ {
		sum += tick(grid)
	}

	fmt.Printf("Solution part 1: %d\n", sum)
}

func part2() {
	grid := parseInput()

	octopuses := len(grid) * len(grid[0])
	i := 0
	for {
		flashes := tick(grid)
		if flashes == octopuses {
			break
		}
		i++
	}

	fmt.Printf("Solution part 2: %d\n", i+1)
}

func main() {
	part1()
	part2()
}
