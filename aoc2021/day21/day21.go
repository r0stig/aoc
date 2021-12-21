package main

import "fmt"

type DeterministicDice struct {
	diceMax   int
	diceStart int
	rolls     int
}

func (d *DeterministicDice) Roll(count int) int {
	sum := 0
	for i := 0; i < count; i++ {
		if d.diceStart == d.diceMax+1 {
			sum += 1
			d.diceStart = 2
		} else {
			sum += d.diceStart
			d.diceStart += 1
		}
	}
	d.rolls += 3
	return sum
}

type Player struct {
	pos, score   int
	winningScore int
	hasWon       bool
}

func (p Player) Step(steps int) Player {
	p.pos += steps % 10
	if p.pos > 10 {
		p.pos = p.pos % 10
	}
	p.score += p.pos

	if p.score >= p.winningScore {
		p.hasWon = true
	}
	return p
}

type GameState struct {
	dice       DeterministicDice
	player1Pos int
	player2Pos int

	player1Score int
	player2Score int

	player1Win bool
	player2Win bool

	player1, player2 Player
}

func turn(state GameState) GameState {
	steps := state.dice.Roll(3)
	state.player1 = state.player1.Step(steps)
	if state.player1.hasWon {
		return state
	}
	steps = state.dice.Roll(3)
	state.player2 = state.player2.Step(steps)
	return state
}

var cache = make(map[string][]int64)

func turnPart2(player1, player2 Player, p1Turn bool) []int64 {
	// Manual memorization with cacheKey of the method arguments
	cacheKey := fmt.Sprintf("%d%d%d%d%t", player1.pos, player1.score, player2.pos, player2.score, p1Turn)
	if player1.hasWon {
		return []int64{1, 0}
	} else if player2.hasWon {
		return []int64{0, 1}
	} else if v, ok := cache[cacheKey]; ok {
		return v
	} else {
		newPlayer1 := player1
		newPlayer2 := player2

		p1Wins := int64(0)
		p2Wins := int64(0)

		for d1 := 1; d1 <= 3; d1++ {
			for d2 := 1; d2 <= 3; d2++ {
				for d3 := 1; d3 <= 3; d3++ {
					if p1Turn {
						newPlayer1 = player1.Step(d1 + d2 + d3)
					} else {
						newPlayer2 = player2.Step(d1 + d2 + d3)
					}
					res := turnPart2(newPlayer1, newPlayer2, !p1Turn)
					p1Wins += res[0]
					p2Wins += res[1]
				}
			}
		}
		cache[cacheKey] = []int64{p1Wins, p2Wins}
		return []int64{p1Wins, p2Wins}
	}
}
func part1() {
	state := GameState{
		dice:         DeterministicDice{diceMax: 100, diceStart: 1},
		player1Pos:   4,
		player2Pos:   5,
		player1Score: 0,
		player2Score: 0,
		player1:      Player{pos: 4, score: 0, winningScore: 1000},
		player2:      Player{pos: 5, score: 0, winningScore: 1000},
	}

	for !state.player1.hasWon && !state.player2.hasWon {
		state = turn(state)
	}

	losingScore := state.player2.score
	if state.player2.hasWon {
		losingScore = state.player1.score
	}

	fmt.Printf("Solution part 1: %d\n", state.dice.rolls*losingScore)
}

func part2() {
	res := turnPart2(
		Player{pos: 4, score: 0, winningScore: 21},
		Player{pos: 5, score: 0, winningScore: 21},
		true,
	)

	max := res[0]
	if res[1] > max {
		max = res[1]
	}
	fmt.Printf("Solution part 2: %d \n", max)
}

func main() {
	part1()
	part2()
}
