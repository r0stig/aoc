package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Rule struct {
	Pair   string
	Result string
}

type Input struct {
	S               map[string]int
	StartingFormula string
	Rules           []Rule
}

func parseInput() Input {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	parts := strings.Split(string(data), "\n\n")
	var input Input

	input.S = make(map[string]int)
	for i := 1; i < len(parts[0]); i++ {
		input.S[fmt.Sprintf("%c%c", parts[0][i-1], parts[0][i])]++
	}

	input.StartingFormula = parts[0]
	var rules []Rule
	for _, line := range strings.Split(parts[1], "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " -> ")
		rules = append(rules, Rule{Pair: parts[0], Result: parts[1]})
	}

	input.Rules = rules

	return input

}

func step(formula map[string]int, rules []Rule, counts map[string]int) map[string]int {
	resultFormula := make(map[string]int)

	for formulaPair, v := range formula {
		for _, templateRule := range rules {
			if templateRule.Pair == formulaPair {
				counts[templateRule.Result] += v

				resultFormula[fmt.Sprintf("%c%s", formulaPair[0], templateRule.Result)] += v
				resultFormula[fmt.Sprintf("%s%c", templateRule.Result, formulaPair[1])] += v
			}
		}
	}
	return resultFormula
}

func part1() {
	input := parseInput()

	counts := make(map[string]int)

	for _, c := range input.StartingFormula {
		counts[fmt.Sprintf("%c", c)]++
	}

	formula := input.S
	for i := 0; i < 10; i++ {
		formula = step(formula, input.Rules, counts)
	}

	min := 99999999
	max := -1

	for _, v := range counts {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	fmt.Printf("Solution day 1: %d\n", max-min)
}

func part2() {
	input := parseInput()
	counts := make(map[string]int)

	for _, c := range input.StartingFormula {
		counts[fmt.Sprintf("%c", c)]++
	}

	formula := input.S
	for i := 0; i < 40; i++ {
		formula = step(formula, input.Rules, counts)
	}

	min := -1
	max := -1

	for _, v := range counts {
		if min == -1 {
			min = v
		}
		if max == -1 {
			max = v
		}
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	fmt.Printf("Solution day 2: %d\n", max-min)
}

func main() {
	part1()
	part2()
}
