package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func part() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	var elfs []int

	curElfVal := 0
	for _, el := range inputArr {
		if el == "" {
			elfs = append(elfs, curElfVal)
			curElfVal = 0
		} else {
			intVal, err := strconv.Atoi(el)
			if err != nil {
				fmt.Printf("Erro converting %s\n", el)
			}
			curElfVal += intVal
		}
	}

	sort.Slice(elfs, func(i, j int) bool {
		return elfs[i] > elfs[j]
	})

	fmt.Printf("Answer part1: %d\n", elfs[0])
	fmt.Printf("Answer part2: %d\n", elfs[0]+elfs[1]+elfs[2])

}

func main() {
	part()
}
