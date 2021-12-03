package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func parseInput() []int16 {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	var codes []int16
	for _, row := range inputArr {
		if row == "" {
			continue
		}
		if number, err := strconv.ParseInt(row, 2, 16); err == nil {
			codes = append(codes, int16(number))

		} else {
			fmt.Printf("Err %v\n", err)
		}
	}
	return codes
}

func createBitMask(pos int, pad string) int16 {
	mask := "1"
	for i := 0; i < pos-1; i++ {
		mask += pad
	}
	res, err := strconv.ParseInt(mask, 2, 16)
	if err != nil {
		fmt.Printf("Error converting %v %v\n", mask, err)
		return 0
	}

	return int16(res)
}

func part1() {
	codes := parseInput()

	bits := 12
	var ones [12]int
	for _, code := range codes {
		for i := 0; i < bits; i++ {
			bit := (createBitMask(bits-i, "0") & code) >> (bits - (i + 1))
			ones[i] += int(bit)
		}
	}

	resultBinary := ""
	for _, one := range ones {
		if one > len(codes)/2 {
			resultBinary += "1"
		} else {
			resultBinary += "0"
		}
	}

	if gamma, err := strconv.ParseInt(resultBinary, 2, 16); err == nil {
		epsilon := int64(createBitMask(12, "1")) & (^gamma)

		fmt.Printf("Part 1 solution: %d\n", gamma*epsilon)
	} else {
		fmt.Printf("Error: %v\n", err)
	}
}

func findCommon(codes []int16, bitPos int, keepMostCommon bool) int16 {
	if len(codes) == 1 {
		return codes[0]
	}

	bits := 12
	ones := 0
	var oneNumbers []int16
	var zeroNumbers []int16
	for _, code := range codes {
		bit := (createBitMask(bits-bitPos, "0") & code) >> (bits - (bitPos + 1))
		ones += int(bit)

		if bit == 1 {
			oneNumbers = append(oneNumbers, code)
		} else {
			zeroNumbers = append(zeroNumbers, code)
		}
	}
	half := int(math.Ceil((float64(len(codes)) / 2)))

	if keepMostCommon {
		if ones >= half {
			return findCommon(oneNumbers, bitPos+1, keepMostCommon)
		}
		return findCommon(zeroNumbers, bitPos+1, keepMostCommon)
	} else {
		if ones < half {
			return findCommon(oneNumbers, bitPos+1, keepMostCommon)
		}
		return findCommon(zeroNumbers, bitPos+1, keepMostCommon)
	}
}

func part2() {
	codes := parseInput()

	oxy := findCommon(codes, 0, true)
	co2 := findCommon(codes, 0, false)

	fmt.Printf("Part 2 solution: %d\n", int64(oxy)*int64(co2))
}

func main() {
	part1()
	part2()
}
