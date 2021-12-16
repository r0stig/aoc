package main

import "fmt"

type Target struct {
	fromX, fromY, toX, toY int
}

func step(x, y, velX, velY, stepNr int) (int, int) {
	newX := x
	stepVelX := velX - (stepNr - 1)
	if stepVelX > 0 {
		newX += stepVelX
	}
	newY := y
	stepVelY := velY - (stepNr - 1)
	newY += stepVelY
	return newX, newY
}

func isInTarget(x, y int, target Target) bool {
	return x >= target.fromX && x <= target.toX &&
		y >= target.fromY && y <= target.toY
}

func part1() {
	/*
		// example input
		target := Target{
			fromX: 20,
			toX:   30,
			fromY: -10,
			toY:   -5,
		}
	*/

	// AOC input
	target := Target{
		fromX: 34,
		toX:   67,
		fromY: -215,
		toY:   -186,
	}

	var maxYVals []int

	for testVelY := 0; testVelY < 10000; testVelY++ {
		beenInTarget := false
		newX, newY := 0, 0
		velX, velY := 0, testVelY
		maxY := 0

		stepNr := 1
		for newY >= target.toY {
			newX, newY = step(newX, newY, velX, velY, stepNr)
			if newY > maxY {
				maxY = newY
			}
			if newY >= target.fromY && newY <= target.toY {
				beenInTarget = true
				break
			}
			stepNr++
		}
		if beenInTarget {
			maxYVals = append(maxYVals, maxY)
		}
	}

	fmt.Printf("Solution part 1: %d\n", maxYVals[len(maxYVals)-1])
}

func part2() {
	target := Target{
		fromX: 34,
		toX:   67,
		fromY: -215,
		toY:   -186,
	}
	var yVals []int

	for i := -10000; i < 10000; i++ {
		beenInTarget := false
		newX, newY := 0, 0
		velX, velY := 0, i
		stepNr := 1
		for newY >= target.toY {
			newX, newY = step(newX, newY, velX, velY, stepNr)
			if newY >= target.fromY && newY <= target.toY {
				beenInTarget = true
				break
			}
			stepNr++
		}
		if beenInTarget {
			yVals = append(yVals, i)
		}
	}

	targetXYCounts := 0

	for i := 0; i < len(yVals); i++ {
		for j := 0; j < 10000; j++ {
			newX, newY := 0, 0
			velX, velY := j, yVals[i]
			maxY := 0
			for stepNr := 1; stepNr < 10000; stepNr++ {
				newX, newY = step(newX, newY, velX, velY, stepNr)
				if newY > maxY {
					maxY = newY
				}
				if isInTarget(newX, newY, target) {
					targetXYCounts++
					break
				}
			}
		}
	}

	fmt.Printf("Solution part 2: %d\n", targetXYCounts)
}

func main() {
	part1()
	part2()
}

// 8,47 thats not right
// 8,214 That's not the right answer.
