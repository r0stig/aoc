package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func getHandStrenghtWithJoker(hand []string) int {
	typesOfCards := make(map[string]int)
	for _, card := range hand {
		typesOfCards[card] += 1
	}

	nrOfJokers := typesOfCards["J"]
	if nrOfJokers == 0 {
		return getHandStrenght(hand)
	}

	// Replace jokers with the card which has the most of the same..
	highestAmount := 0
	highestCard := ""
	for key, value := range typesOfCards {
		if key != "J" && value > highestAmount {
			highestAmount = value
			highestCard = key
		}
	}

	var newHand []string
	for i := 0; i < len(hand); i++ {
		if hand[i] == "J" {
			newHand = append(newHand, highestCard)
		} else {
			newHand = append(newHand, hand[i])
		}
	}

	return getHandStrenght(newHand)
}

func getHandStrenght(hand []string) int {
	typesOfCards := make(map[string]int)
	for _, card := range hand {
		typesOfCards[card] += 1
	}

	mapSize := len(typesOfCards)

	if mapSize == 1 {
		// Five of a kind
		return 7
	} else if mapSize == 2 {
		// Four of a kind
		// or full house
		for _, value := range typesOfCards {
			if value == 4 {
				// Four of a kind
				return 6
			}
		}
		// Full house
		return 5
	} else {
		pairs := 0
		for _, value := range typesOfCards {
			if value == 3 {
				// Three of a kind
				return 4
			} else if value == 2 {
				pairs++
			}
		}
		if pairs == 2 {
			// Two pair
			return 3
		} else if pairs == 1 {
			// One pair
			return 2
		}
		// High card
		return 1
	}
}

type HandWithStrenght struct {
	hand     []string
	strenght int
	bid      int
}

type ByHandRanking []HandWithStrenght

func (a ByHandRanking) Len() int      { return len(a) }
func (a ByHandRanking) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByHandRanking) Less(i, j int) bool {
	var cardToSymbol = map[string]string{
		"A": "E",
		"K": "D",
		"Q": "C",
		"J": "B",
		"T": "A",
		"9": "9",
		"8": "8",
		"7": "7",
		"6": "6",
		"5": "5",
		"4": "4",
		"3": "3",
		"2": "2",
	}
	if a[i].strenght == a[j].strenght {
		// Go through cards until we find one that is higher than the other
		for i, card := range a[i].hand {
			cardA := cardToSymbol[card]
			cardB := cardToSymbol[a[j].hand[i]]
			if cardA == cardB {
				continue
			}
			return cardA > cardB
		}
	}

	return a[i].strenght > a[j].strenght
}

type ByHandRankingWithJoker []HandWithStrenght

func (a ByHandRankingWithJoker) Len() int      { return len(a) }
func (a ByHandRankingWithJoker) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByHandRankingWithJoker) Less(i, j int) bool {
	var cardToSymbol = map[string]string{
		"A": "E",
		"K": "D",
		"Q": "C",
		"J": "1",
		"T": "A",
		"9": "9",
		"8": "8",
		"7": "7",
		"6": "6",
		"5": "5",
		"4": "4",
		"3": "3",
		"2": "2",
	}
	if a[i].strenght == a[j].strenght {
		// Go through cards until we find one that is higher than the other
		for i, card := range a[i].hand {
			cardA := cardToSymbol[card]
			cardB := cardToSymbol[a[j].hand[i]]
			if cardA == cardB {
				continue
			}
			return cardA > cardB
		}
	}

	return a[i].strenght > a[j].strenght
}

func part1() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	lines := strings.Split(string(data), "\n")

	var list []HandWithStrenght
	for _, line := range lines {
		parts := strings.Split(line, " ")
		handStr := strings.Split(parts[0], "")
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("Failed to convert bid %s\n", parts[1])
		}

		handStrenght := getHandStrenght(handStr)

		handWithStrength := HandWithStrenght{
			hand:     handStr,
			strenght: handStrenght,
			bid:      bid,
		}

		list = append(list, handWithStrength)
	}

	sort.Sort(ByHandRanking(list))

	sum := 0
	totalHands := len(list)
	for i, hand := range list {
		sum += (totalHands - i) * hand.bid
	}
	fmt.Printf("Answer part1: %d\n", sum)
}

func part2() {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	lines := strings.Split(string(data), "\n")

	var list []HandWithStrenght
	for _, line := range lines {
		parts := strings.Split(line, " ")
		handStr := strings.Split(parts[0], "")
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("Failed to convert bid %s\n", parts[1])
		}

		handStrenght := getHandStrenghtWithJoker(handStr)

		handWithStrength := HandWithStrenght{
			hand:     handStr,
			strenght: handStrenght,
			bid:      bid,
		}

		list = append(list, handWithStrength)
	}

	sort.Sort(ByHandRankingWithJoker(list))

	sum := 0
	totalHands := len(list)
	for i, hand := range list {
		sum += (totalHands - i) * hand.bid
	}
	fmt.Printf("Answer part2: %d\n", sum)
}

func main() {
	part1()
	part2()
}
