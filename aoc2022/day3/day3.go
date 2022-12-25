package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part1() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	//priorities := make(map[rune]struct{})
	var priorities []map[rune]struct{}
	for _, el := range inputArr {
		if el == "" {
			continue
		}

		part1 := el[0 : len(el)/2]
		part2 := el[len(el)/2:]

		CODE_POINT_UPPER_CASE_A := rune('A')
		CODE_POINT_LOWER_CASE_A := rune('a')

		// var common []rune
		var exists struct{}
		common := make(map[rune]struct{})
		for _, p1 := range part1 {
			for _, p2 := range part2 {
				if p1 == p2 {
					// Is it lowercase?
					if p1 >= CODE_POINT_LOWER_CASE_A {
						common[p1-CODE_POINT_LOWER_CASE_A+1] = exists
					} else {
						common[p1-CODE_POINT_UPPER_CASE_A+27] = exists
					}
					// Only unique is needed
					break
				}
			}
		}
		priorities = append(priorities, common)
	}

	sum := 0
	for _, p := range priorities {
		for key, _ := range p {
			sum += int(key)
		}
	}

	fmt.Printf("Sum is %d\n", sum)
}

func part2() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	var priorities []map[rune]struct{}
	for i := 2; i < len(inputArr); i += 3 {

		part1 := inputArr[i-2]
		part2 := inputArr[i-1]
		part3 := inputArr[i]

		CODE_POINT_UPPER_CASE_A := rune('A')
		CODE_POINT_LOWER_CASE_A := rune('a')

		var exists struct{}
		common := make(map[rune]struct{})
		for _, p1 := range part1 {
			for _, p2 := range part2 {
				if p1 == p2 {
					common[p1] = exists
					// Only unique is needed
					break
				}
			}
		}

		allThreeCommon := make(map[rune]struct{})
		for p, _ := range common {
			for _, p3 := range part3 {
				if p == p3 {
					if p3 >= CODE_POINT_LOWER_CASE_A {
						allThreeCommon[p3-CODE_POINT_LOWER_CASE_A+1] = exists
					} else {
						allThreeCommon[p3-CODE_POINT_UPPER_CASE_A+27] = exists
					}
					// Only unique is needed
					break
				}
			}
		}

		priorities = append(priorities, allThreeCommon)
	}

	sum := 0
	for _, p := range priorities {
		for key, _ := range p {
			sum += int(key)
		}
	}

	fmt.Printf("Sum is %d\n", sum)
}

func main() {
	fmt.Println("Part1")
	part1()
	fmt.Println("Part1")
	part2()
}
