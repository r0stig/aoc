package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Heightmap [][]int

func parseInput() Heightmap {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	var heightmap Heightmap
	for _, row := range inputArr {
		if row == "" {
			continue
		}
		var columns []int
		for _, column := range strings.Split(row, "") {
			height, err := strconv.Atoi(column)
			if err != nil {
				fmt.Printf("Error converting column %v\n", err)
			}
			columns = append(columns, height)
		}
		heightmap = append(heightmap, columns)
	}

	return heightmap
}

func part1() {
	heightmap := parseInput()

	var lowPoints []int
	for row := 0; row < len(heightmap); row++ {
		for col := 0; col < len(heightmap[row]); col++ {
			current := heightmap[row][col]
			higherNeigbours := 0
			if row == 0 || heightmap[row-1][col] > current {
				higherNeigbours++
			}
			if row == len(heightmap)-1 || heightmap[row+1][col] > current {
				higherNeigbours++
			}
			if col == 0 || heightmap[row][col-1] > current {
				higherNeigbours++
			}
			if col == len(heightmap[row])-1 || heightmap[row][col+1] > current {
				higherNeigbours++
			}
			if higherNeigbours == 4 {
				lowPoints = append(lowPoints, current)
			}
		}
	}

	sum := 0
	for _, p := range lowPoints {
		sum += 1 + p
	}

	fmt.Printf("Solution part 1 is %d\n", sum)
}

func isInBasin(row, col int, heighmap Heightmap, curBasin *[][]int) bool {
	for _, c := range *curBasin {
		if c[0] == row && c[1] == col {
			return true
		}
	}
	return false
}

func getSizeOfBasin(row, col int, heightmap Heightmap, curBasin *[][]int) int {
	// Is the coordinate already in the current basin?
	for _, c := range *curBasin {
		if c[0] == row && c[1] == col {
			return 0
		}
	}

	current := heightmap[row][col]
	higherNeigbours := 0
	var checkNeighboards [][]int
	if row == 0 || heightmap[row-1][col] > current || isInBasin(row-1, col, heightmap, curBasin) {
		higherNeigbours++
		if row != 0 && heightmap[row-1][col] != 9 && !isInBasin(row-1, col, heightmap, curBasin) {
			checkNeighboards = append(checkNeighboards, []int{row - 1, col})
		}
	}
	if row == len(heightmap)-1 || heightmap[row+1][col] > current || isInBasin(row+1, col, heightmap, curBasin) {
		higherNeigbours++
		if row != len(heightmap)-1 && heightmap[row+1][col] != 9 && !isInBasin(row+1, col, heightmap, curBasin) {
			checkNeighboards = append(checkNeighboards, []int{row + 1, col})
		}
	}
	if col == 0 || heightmap[row][col-1] > current || isInBasin(row, col-1, heightmap, curBasin) {
		higherNeigbours++
		if col != 0 && heightmap[row][col-1] != 9 && !isInBasin(row, col-1, heightmap, curBasin) {
			checkNeighboards = append(checkNeighboards, []int{row, col - 1})
		}
	}
	if col == len(heightmap[row])-1 || heightmap[row][col+1] > current || isInBasin(row, col+1, heightmap, curBasin) {
		higherNeigbours++
		if col != len(heightmap[row])-1 && heightmap[row][col+1] != 9 && !isInBasin(row, col+1, heightmap, curBasin) {
			checkNeighboards = append(checkNeighboards, []int{row, col + 1})
		}
	}

	*curBasin = append(*curBasin, []int{row, col})
	if higherNeigbours >= 2 {
		sum := 1
		for _, neighboar := range checkNeighboards {
			sum += getSizeOfBasin(neighboar[0], neighboar[1], heightmap, curBasin)
		}
		return sum
	}
	return 1
}

func part2() {
	heightmap := parseInput()

	var lowPoints [][]int
	for row := 0; row < len(heightmap); row++ {
		for col := 0; col < len(heightmap[row]); col++ {
			current := heightmap[row][col]
			higherNeigbours := 0
			if row == 0 || heightmap[row-1][col] > current {
				higherNeigbours++
			}
			if row == len(heightmap)-1 || heightmap[row+1][col] > current {
				higherNeigbours++
			}
			if col == 0 || heightmap[row][col-1] > current {
				higherNeigbours++
			}
			if col == len(heightmap[row])-1 || heightmap[row][col+1] > current {
				higherNeigbours++
			}
			if higherNeigbours == 4 {
				lowPoints = append(lowPoints, []int{row, col})
			}
		}
	}

	var basinSizes []int
	for _, lowPoint := range lowPoints {
		var curBasin [][]int
		size := getSizeOfBasin(lowPoint[0], lowPoint[1], heightmap, &curBasin)
		basinSizes = append(basinSizes, size)
	}

	sort.Ints(basinSizes)

	sum := basinSizes[len(basinSizes)-1] * basinSizes[len(basinSizes)-2] * basinSizes[len(basinSizes)-3]

	fmt.Printf("Solution part 2 is %d\n", sum)

}

func main() {
	part1()
	part2()
}
