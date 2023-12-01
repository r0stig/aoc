package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func isNumberWord(s string) int {
	switch true {
	case strings.HasSuffix(s, "one"):
		return 1
	case strings.HasSuffix(s, "two"):
		return 2
	case strings.HasSuffix(s, "three"):
		return 3
	case strings.HasSuffix(s, "four"):
		return 4
	case strings.HasSuffix(s, "five"):
		return 5
	case strings.HasSuffix(s, "six"):
		return 6
	case strings.HasSuffix(s, "seven"):
		return 7
	case strings.HasSuffix(s, "eight"):
		return 8
	case strings.HasSuffix(s, "nine"):
		return 9
	}
	return 0
}

func part() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	var lines [][]rune
	for _, el := range inputArr {
		var numbers []rune
		curWord := ""
		for _, char := range el {
			nr := char - 48
			if int(nr) < 10 {
				numbers = append(numbers, nr)
				curWord = ""
			} else {
				curWord += string(char)
				wordNumber := isNumberWord(curWord)
				if wordNumber != 0 {
					numbers = append(numbers, rune(wordNumber))
				}
			}
		}
		lines = append(lines, numbers)
	}

	sum := 0
	for i, line := range lines {
		nr, err := strconv.Atoi(strconv.Itoa(int(line[0])) + strconv.Itoa(int(line[len(line)-1])))
		if err != nil {
			fmt.Printf("Failed to convert\n")
			break
		}
		sum += nr
		fmt.Printf("%d: Adding %d %d %d\n", i+1, int(line[0]), int(line[len(line)-1]), nr)
	}

	fmt.Printf("Answer part2: %d\n", sum)
	fmt.Printf("nr lines %d\n", len(lines))
}

func main() {
	part()
}
