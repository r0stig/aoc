package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type FishBundle struct {
	Age      int
	ResetVal int
	Count    int
}

func parseInput() []FishBundle {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), ",")

	var fishes []FishBundle
	for _, nr := range inputArr {
		if nr == "" {
			continue
		}
		age, err := strconv.Atoi(strings.Trim(nr, "\n"))
		if err != nil {
			fmt.Printf("Error converting %v\n", err)
		}
		foundFishBundle := false
		for i := 0; i < len(fishes); i++ {
			if fishes[i].Age == age {
				fishes[i].Count++
				foundFishBundle = true
			}
		}
		if !foundFishBundle {
			fishes = append(fishes, FishBundle{Age: age, ResetVal: 6, Count: 1})
		}
	}

	return fishes
}

func tick(fishes []FishBundle) []FishBundle {
	fishesToAdd := 0
	for i := 0; i < len(fishes); i++ {
		if fishes[i].Age == 0 {
			fishes[i].Age = fishes[i].ResetVal
			fishesToAdd += fishes[i].Count
		} else {
			fishes[i].Age--
		}
	}
	fishes = append(fishes, FishBundle{Age: 8, ResetVal: 6, Count: fishesToAdd})
	return fishes
}

func part1() {
	fishes := parseInput()

	for i := 0; i < 80; i++ {
		fishes = tick(fishes)
	}

	sum := 0
	for _, bundles := range fishes {
		sum += bundles.Count
	}
	fmt.Printf("Solution part 1: %d\n", sum)
}

func part2() {
	fishes := parseInput()

	for i := 0; i < 256; i++ {
		fishes = tick(fishes)
	}

	sum := 0
	for _, bundles := range fishes {
		sum += bundles.Count
	}
	fmt.Printf("Solution part 2: %d\n", sum)
}

func main() {
	part1()
	part2()
}
