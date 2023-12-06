package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part1() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	timeLineStr := strings.Split(strings.TrimSpace(string(inputArr[0][9:])), " ")
	distanceLineStr := strings.Split(strings.TrimSpace(string(inputArr[1][9:])), " ")

	var timeLine []int
	for _, t := range timeLineStr {
		c, err := strconv.Atoi(t)
		if err != nil {
			fmt.Printf("Failed to convert %s\n", t)
		} else {
			timeLine = append(timeLine, c)
		}

	}

	var distanceLine []int
	for _, t := range distanceLineStr {
		c, err := strconv.Atoi(t)
		if err != nil {
			fmt.Printf("Failed to convert %s\n", t)
		} else {
			distanceLine = append(distanceLine, c)
		}
	}

	var winningHoldTimes []int
	for i, totalTime := range timeLine {
		nrWinningHoldTimes := 0
		for holdTime := 0; holdTime < totalTime; holdTime++ {
			timeLeftAfterHold := totalTime - holdTime
			speed := holdTime
			distance := speed * timeLeftAfterHold
			if distance > distanceLine[i] {
				nrWinningHoldTimes++
			}
		}
		if nrWinningHoldTimes > 0 {
			winningHoldTimes = append(winningHoldTimes, nrWinningHoldTimes)
		}
	}

	sum := 1
	for _, w := range winningHoldTimes {
		sum *= w
	}

	fmt.Printf("Answer part1: %d\n", sum)
}

func part2() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	timeLineStr := strings.ReplaceAll(string(inputArr[0][9:]), " ", "")
	distanceLineStr := strings.ReplaceAll(string(inputArr[1][9:]), " ", "")

	var errConvert error
	var totalTime int
	totalTime, errConvert = strconv.Atoi(timeLineStr)
	if errConvert != nil {
		fmt.Printf("Failed to convert %s\n", timeLineStr)
	}

	var totalDistance int
	totalDistance, errConvert = strconv.Atoi(distanceLineStr)
	if errConvert != nil {
		fmt.Printf("Failed to convert %s\n", distanceLineStr)
	}

	fmt.Printf("Time: %d Distance %d\n", totalTime, totalDistance)

	nrWinningHoldTimes := 0
	for holdTime := 0; holdTime < totalTime; holdTime++ {
		timeLeftAfterHold := totalTime - holdTime
		speed := holdTime
		distance := speed * timeLeftAfterHold
		if distance > totalDistance {
			nrWinningHoldTimes++
		}
	}

	fmt.Printf("Answer part2: %d\n", nrWinningHoldTimes)
}

func main() {
	part1()
	part2()
}
