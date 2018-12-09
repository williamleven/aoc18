package day9

import (
	"fmt"
	"github.com/Gurgy/aoc18"
	"strings"
)

var Solutions = aoc18.Day{
	A: a,
	B: b,
}

type marble struct {
	value int
	previous *marble
	next *marble
}

func a() interface{} {
	players, marbles := input("day9/input")

	score := playGame(players, marbles)

	return max(score)
}

func b() interface{} {
	players, marbles := input("day9/input")

	marbles *= 100

	score := playGame(players, marbles)

	return max(score)
}

func playGame(players int, marbles int) ([]int) {
	score := make([]int, players)

	// setup
	ring := &marble{value:0}
	ring.next = ring
	ring.previous = ring

	for i := 1; i <= marbles ; i++  {
		if i % 23 == 0 {
			player := (i - 1) % len(score)
			score[player] += i
			// back 7 steps
			for j := 0; j < 7; j++ {
				ring = ring.previous
			}
			score[player] += ring.value
			// remove marble
			ring.previous.next = ring.next
			ring.next.previous = ring.previous
			ring = ring.next
		} else {
			ring = ring.next
			// Insert new marble
			nw := &marble{
				value: i,
				previous: ring,
				next: ring.next,
			}
			ring.next.previous = nw
			ring.next = nw
			ring = nw
		}
	}
	return score
}

func input(file string) (players, marbles int) {
	lines, err := aoc18.GetLines(file)
	if err != nil {
		panic(err)
	}

	fmt.Fscanf(
		strings.NewReader(lines[0]),
		"%d players; last marble is worth %d points",
		&players,
		&marbles,
	)

	return players, marbles
}

func max(as []int) int {
	max := 0
	for _, a := range as  {
		if a > max {
			max = a
		}
	}

	return max
}