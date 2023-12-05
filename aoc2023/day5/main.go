package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type MapEntry struct {
	destRangeStart   int64
	sourceRangeStart int64
	rangeLength      int64
}

func createMap(start int, lines []string) ([]MapEntry, int) {
	var entries []MapEntry

	lastI := 0
	for i := start; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			lastI = i
			break
		}
		var entry MapEntry
		parts := strings.Split(line, " ")
		destRangeStart, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			fmt.Printf("Failed to convert %s\n", parts[0])
		}
		sourceRangeStart, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			fmt.Printf("Failed to convert %s\n", parts[1])
		}
		rangeLength, err := strconv.ParseInt(parts[2], 10, 64)
		if err != nil {
			fmt.Printf("Failed to convert %s\n", parts[2])
		}

		entry.destRangeStart = destRangeStart
		entry.sourceRangeStart = sourceRangeStart
		entry.rangeLength = rangeLength

		entries = append(entries, entry)
	}

	return entries, lastI
}

var cache = make(map[int]int)

func getMapValue(number int64, seedMaps []MapEntry) int64 {
	for _, seedMap := range seedMaps {
		if number >= seedMap.sourceRangeStart && number < seedMap.sourceRangeStart+seedMap.rangeLength {
			var diff int64
			diff = number - seedMap.sourceRangeStart
			return seedMap.destRangeStart + diff
		}
	}

	return number
}

func part() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	lines := strings.Split(string(data), "\n")

	var seeds []int64
	seedRow := strings.Split(lines[0], ":")
	seedsTable := strings.Split(strings.Trim(seedRow[1], " "), " ")
	for _, e := range seedsTable {
		nr, err := strconv.ParseInt(e, 10, 64)
		if err != nil {
			fmt.Printf("Failed to convert %s\n", e)
		}
		seeds = append(seeds, nr)
	}

	var seedMappings [][]MapEntry

	lineNr := 1
	for lineNr < len(lines) {
		if lineNr == 0 {
			break
		}

		fmt.Printf("Try to convert from line %d\n", lineNr)
		seedMap, lastI := createMap(lineNr+2, lines)
		seedMappings = append(seedMappings, seedMap)
		lineNr = lastI
	}

	var lowest int64
	lowest = -1
	for _, seed := range seeds {

		lastResult := seed
		for _, seedMap := range seedMappings {
			lastResult = getMapValue(lastResult, seedMap)
		}
		if lowest == -1 || lastResult < lowest {
			lowest = lastResult
		}
	}
	fmt.Printf("Answer part 1: %d\n", lowest)

	lowest = -1
	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		amount := seeds[i+1]

		fmt.Printf("Seeds from %d %d\n", start, amount)

		for x := start; x < start+amount; x++ {
			var lastResult int64
			lastResult = x

			for _, seedMap := range seedMappings {
				lastResult = getMapValue(lastResult, seedMap)
			}
			if lowest == -1 || lastResult < lowest {
				lowest = lastResult
				fmt.Printf("lowest is now %d seedNr: %d\n", lastResult, x)
			}
		}
	}

	fmt.Printf("Answer part 2: %d\n", lowest)
}

func main() {
	part()
}
