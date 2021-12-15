package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

// PriorityQueue from docs
// https://golang.google.cn/pkg/container/heap/#example__priorityQueue
type Item struct {
	point    Point
	priority int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

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

func getNeighboars(grid [][]int, row, col int) [4][2]int {
	var neighboars [4][2]int

	if row > 0 {
		neighboars[0][0] = row - 1
		neighboars[0][1] = col
	}

	if col > 0 {
		neighboars[1][0] = row
		neighboars[1][1] = col - 1
	}
	if col < len(grid[row])-1 {
		neighboars[2][0] = row
		neighboars[2][1] = col + 1
	}
	if row < len(grid)-1 {
		neighboars[3][0] = row + 1
		neighboars[3][1] = col
	}

	return neighboars
}

func calcHeuristic(grid [][]int, pos []int) int {
	goal := len(grid) - 1

	dX := int(math.Abs(float64(pos[0]) - float64(goal)))
	dY := int(math.Abs(float64(pos[1]) - float64(goal)))

	return dX + dY
}

func traverse(grid [][]int) int {
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

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{
		point:    Point{row: 0, col: 0},
		priority: grid[0][0],
	})

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		head := cur.point
		if head.col == len(grid)-1 && head.row == len(grid)-1 {
			return distances[Point{row: len(grid) - 1, col: len(grid) - 1}]
		}

		neighboars := getNeighboars(grid, head.row, head.col)
		for _, neighboar := range neighboars {
			tempDistance := distances[Point{row: head.row, col: head.col}] + grid[neighboar[0]][neighboar[1]]

			if tempDistance < distances[Point{row: neighboar[0], col: neighboar[1]}] {
				distances[Point{row: neighboar[0], col: neighboar[1]}] = tempDistance
				distancesWithHeuristics[Point{row: neighboar[0], col: neighboar[1]}] = tempDistance // + calcHeuristic(grid, neighboar) // Quicker without this, why?
				came_from[Point{row: neighboar[0], col: neighboar[1]}] = fmt.Sprintf("%d,%d", head.col, head.row)

				heap.Push(&pq, &Item{
					point:    Point{row: neighboar[0], col: neighboar[1]},
					priority: tempDistance,
				})
			}
		}
	}

	return -1
}

func part1() {
	input := parseInput()
	sum := traverse(input)
	fmt.Printf("Solution part 1: %d\n", sum)
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

	sum := traverse(newGrid)
	fmt.Printf("Solution part 2: %d\n", sum)
}

func main() {
	part1()
	part2()
}
