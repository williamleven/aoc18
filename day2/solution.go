package day2

import (
	"fmt"
	"github.com/Gurgy/aoc18"
)

var Solutions = aoc18.Day{
	A: a,
	B: b,
}

func a() interface{} {
	c, err := aoc18.GetLines("day2/input")
	if err != nil {
		return err
	}

	counter2 := 0
	counter3 := 0
	for line := range c {
		counter := map[rune]int{}

		for _, symbol := range line {
			counter[symbol]++
		}

		has3Count := false
		has2Count := false
		for _, counts := range counter {
			if counts == 2 {
				has2Count = true
			} else if counts == 3 {
				has3Count = true
			}
		}
		if has3Count {
			counter3++
		}
		if has2Count {
			counter2++
		}

	}
	return counter2 * counter3
}

func b() interface{} {
	c, err := aoc18.GetLines("day2/input")
	if err != nil {
		return err
	}

	lines := make([]string, 0, 50)
	for line := range c {
		for _, l := range lines {
			similar := -1
			for i, r := range line {
				if r != int32(l[i]) {
					if similar == -1 {
						similar = i
					} else {
						similar = -2
						break
					}
				}
			}
			if similar >= 0 {
				return line[:similar] + line[similar+1:]
			}
		}
		lines = append(lines, line)
	}
	return fmt.Errorf("could not find any similar lines")
}
