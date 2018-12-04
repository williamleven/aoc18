package day3

import (
	"fmt"
	"github.com/Gurgy/aoc18"
	"strings"
)

var Solutions = aoc18.Day{
	A: a,
	B: b,
}

type claim struct {
	id int
	x  int
	y  int
	xs int
	ys int
}

func getClaims(lines chan string) []claim {
	claims := make([]claim, 0, 1357)

	for line := range lines {
		claims = append(claims, parseClaim(line))

	}
	return claims
}

func parseClaim(string string) claim {
	var claim claim
	fmt.Fscanf(strings.NewReader(string), "#%d @ %d,%d: %dx%d", &claim.id, &claim.x, &claim.y, &claim.xs, &claim.ys)
	return claim
}

func plotClaims(claims []claim) [1000][1000]int {
	fabric := [1000][1000]int{}

	for _, claim := range claims {
		for i := claim.x; i < (claim.x + claim.xs); i++ {
			for j := claim.y; j < (claim.y + claim.ys); j++ {
				fabric[i][j]++
			}
		}
	}

	return fabric
}

func a() interface{} {
	c, err := aoc18.GetLines("day3/input")
	if err != nil {
		panic(err)
	}
	claims := getClaims(c)

	fabric := plotClaims(claims)

	counter := 0
	for _, row := range fabric {
		for _, cell := range row {
			if cell > 1 {
				counter++
			}
		}
	}
	return counter
}

func b() interface{} {
	c, err := aoc18.GetLines("day3/input")
	if err != nil {
		panic(err)
	}

	claims := getClaims(c)

	fabric := plotClaims(claims)

	for _, claim := range claims {
		free := true
		for i := claim.x; i < (claim.x + claim.xs); i++ {
			for j := claim.y; j < (claim.y + claim.ys); j++ {
				if fabric[i][j] > 1 {
					free = false
					break
				}
			}
			if !free {
				break
			}
		}
		if free {
			return claim.id
		}

	}
	return fmt.Errorf("no none-overlapping")
}
