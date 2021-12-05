package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Brick struct {
	Number int
	Marked bool
}

type Board struct {
	Rows             [][]Brick
	RightMarkings    [5]int
	DownMarkings     [5]int
	HasWon           bool
	WinningIteration int
}

type Input struct {
	Numbers []int
	Boards  []Board
}

func parseInput() Input {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	var input Input
	numbers := strings.Split(inputArr[0], ",")

	for _, number := range numbers {
		if nr, err := strconv.Atoi(number); err == nil {
			input.Numbers = append(input.Numbers, nr)
		}
	}

	// fmt.Printf("length is %v\n", len(inputArr))
	for i := 2; i < len(inputArr); i += 6 {
		var board Board

		for j := i; j < i+5; j++ {
			// fmt.Printf("Parsing row %v -- %v\n", inputArr[j], i)
			line := strings.Split(inputArr[j], " ")
			var columns []Brick
			for _, l := range line {
				if l == " " {
					continue
				}
				if c, err := strconv.Atoi(l); err == nil {
					columns = append(columns, Brick{Number: c})
				}
			}
			board.Rows = append(board.Rows, columns)
		}
		input.Boards = append(input.Boards, board)
	}
	return input
}

func markNumbers(input Input, number int) Input {
	for i := 0; i < len(input.Boards); i++ {
		for j := 0; j < len(input.Boards[i].Rows); j++ {
			for k := 0; k < len(input.Boards[i].Rows[j]); k++ {
				if input.Boards[i].Rows[j][k].Number == number {
					input.Boards[i].Rows[j][k].Marked = true

					input.Boards[i].RightMarkings[j]++
					input.Boards[i].DownMarkings[k]++
				}
			}
		}
	}
	return input
}

type WinResult struct{}

func checkForWin(input Input) int {
	for i := 0; i < len(input.Boards); i++ {
		for _, rightMarking := range input.Boards[i].RightMarkings {
			if rightMarking == 5 {
				return i
			}
		}
		for _, downMarking := range input.Boards[i].RightMarkings {
			if downMarking == 5 {
				return i
			}
		}
	}
	return -1
}

func markWinningBoards(input Input, iteration int) Input {
	for i := 0; i < len(input.Boards); i++ {
		if input.Boards[i].HasWon {
			continue
		}
		for _, rightMarking := range input.Boards[i].RightMarkings {
			if rightMarking == 5 {
				input.Boards[i].HasWon = true
				input.Boards[i].WinningIteration = iteration
			}
		}
		for _, downMarking := range input.Boards[i].DownMarkings {
			if downMarking == 5 {
				input.Boards[i].HasWon = true
				input.Boards[i].WinningIteration = iteration
			}
		}
	}
	return input
}

func calcNumberOfWinningBoard(input Input) int {
	numberOfWinningBoards := 0
	for i := 0; i < len(input.Boards); i++ {
		if input.Boards[i].HasWon {
			numberOfWinningBoards++
		}
	}
	return numberOfWinningBoards
}

func sumOfUnmarked(board Board) int {
	sum := 0
	for _, row := range board.Rows {
		for _, brick := range row {
			if !brick.Marked {
				sum += brick.Number
			}
		}
	}
	return sum
}

func getLastWinningBoard(input Input) Board {
	lastBoard := input.Boards[0]
	for i := 1; i < len(input.Boards); i++ {
		if input.Boards[i].WinningIteration > lastBoard.WinningIteration {
			lastBoard = input.Boards[i]
		}
	}
	return lastBoard
}

func prettyPrintBoards(input Input) {
	for _, board := range input.Boards {
		for _, row := range board.Rows {
			for _, brick := range row {
				if brick.Marked {
					fmt.Printf("*%1d  ", brick.Number)
				} else {
					fmt.Printf(" %1d  ", brick.Number)
				}
			}
			fmt.Println("")
		}
		fmt.Println("")
		fmt.Println("")
	}
}

func part1() {
	input := parseInput()

	for _, number := range input.Numbers {
		input := markNumbers(input, number)
		winningBoard := checkForWin(input)
		if winningBoard != -1 {
			fmt.Printf("We got a win! last number %d\n", number)
			sum := sumOfUnmarked(input.Boards[winningBoard])
			fmt.Printf("Sum of unmarked %d\n", sum)
			fmt.Printf("Solution part 1 is %d\n", sum*number)
			break
		}
	}
}

func part2() {
	input := parseInput()

	for i, number := range input.Numbers {
		input = markNumbers(input, number)
		input = markWinningBoards(input, i)

		if calcNumberOfWinningBoard(input) == len(input.Boards) {
			lastBoard := getLastWinningBoard(input)
			fmt.Printf("All boards has won..\n")
			sum := sumOfUnmarked(lastBoard)
			fmt.Printf("Sum of unmarked %d\n", sum)
			fmt.Printf("Solution part 2 is %d\n", sum*number)
			break
		}
	}
}

func main() {
	part1()
	part2()
}
