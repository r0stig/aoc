package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func isNumber(s rune) bool {
	return s >= '0' && s <= '9'
}

func isSymbol(s rune) bool {
	return !isNumber(s) && s != '.'
}

func getPartNoAt(y, x int, grid [][]rune) int {
	start := grid[y][x]
	number := string(start)

	width := len(grid[y])
	// forwards
	i := 1
	for {
		if x+i <= width-1 && isNumber(grid[y][x+i]) {
			number += string(grid[y][x+i])
			i++
		} else {
			break
		}
	}

	// backwards
	i = 1
	for {
		if x-i >= 0 && isNumber(grid[y][x-i]) {
			number = string(grid[y][x-i]) + number
			i++
		} else {
			break
		}
	}

	nr, err := strconv.Atoi(number)
	if err != nil {
		fmt.Printf("Error converting %d\n", number)
	}
	return nr
}

func getNumberLength(y, x int, grid [][]rune) int {
	length := 1
	width := len(grid[y])
	// forwards
	i := 1
	for {
		if x+i <= width-1 && isNumber(grid[y][x+i]) {
			length++
			i++
		} else {
			break
		}
	}

	return length
}

func part2() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	var grid [][]rune

	for _, el := range inputArr {
		var line []rune
		for _, char := range el {
			line = append(line, char)
		}
		grid = append(grid, line)
	}

	var gearRatios []int
	for y, line := range grid {
		width := len(line)
		height := len(grid)

		for x, char := range line {
			isGearSymbol := char == '*'

			if isGearSymbol {

				var gearNumbers []int

				if y > 0 && isNumber(grid[y-1][x]) {
					gearNumbers = append(gearNumbers, getPartNoAt(y-1, x, grid))
				} else {
					if y > 0 && x > 0 && isNumber(grid[y-1][x-1]) {
						gearNumbers = append(gearNumbers, getPartNoAt(y-1, x-1, grid))
					}
					if y > 0 && x < width && isNumber(grid[y-1][x+1]) {
						gearNumbers = append(gearNumbers, getPartNoAt(y-1, x+1, grid))
					}
				}

				if x > 0 && isNumber(grid[y][x-1]) {
					gearNumbers = append(gearNumbers, getPartNoAt(y, x-1, grid))
				}
				if x < width && isNumber(grid[y][x+1]) {
					gearNumbers = append(gearNumbers, getPartNoAt(y, x+1, grid))
				}

				if y < height && isNumber(grid[y+1][x]) {
					gearNumbers = append(gearNumbers, getPartNoAt(y+1, x, grid))
				} else {
					if y < height && x > 0 && isNumber(grid[y+1][x-1]) {
						gearNumbers = append(gearNumbers, getPartNoAt(y+1, x-1, grid))
					}
					if y <= height && x < width && isNumber(grid[y+1][x+1]) {
						gearNumbers = append(gearNumbers, getPartNoAt(y+1, x+1, grid))
					}
				}

				if len(gearNumbers) == 2 {
					gearRatios = append(gearRatios, gearNumbers[0]*gearNumbers[1])
				}

			}
		}
	}

	sum := 0
	for _, partNo := range gearRatios {
		sum += partNo
	}

	fmt.Printf("Answer part2 %d\n", sum)
}

func part1() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	var grid [][]rune

	for _, el := range inputArr {
		var line []rune
		for _, char := range el {
			line = append(line, char)
		}
		grid = append(grid, line)
	}

	var partNumbers []int
	y := 0
	for y < len(grid) {
		width := len(grid[y])
		height := len(grid)

		x := 0
		for x < len(grid[y]) {
			char := grid[y][x]

			if isNumber(char) {
				length := getNumberLength(y, x, grid)

				for y0 := -1; y0 < 2; y0++ {
					for x0 := -1; x0 <= length; x0++ {

						checkY := y + y0
						checkX := x + x0

						if checkY < 0 {
							continue
						}
						if checkX < 0 {
							continue
						}
						if checkY >= height {
							continue
						}
						if checkX >= width {
							continue
						}

						if isSymbol(grid[checkY][checkX]) {
							p := getPartNoAt(y, x, grid)

							partNumbers = append(partNumbers, p)

						}
					}
				}
				// Step forward to make sure we don't check the same number again
				// use -1 since we always step one forward anyway
				x += length - 1
			}
			x++
		}

		y++
	}

	sum := 0
	for _, partNo := range partNumbers {
		sum += partNo
	}

	fmt.Printf("Answer part1 %d\n", sum)
}

func main() {
	part1()
	part2()
}
