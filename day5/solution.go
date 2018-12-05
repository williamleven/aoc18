package day5

import (
	"github.com/Gurgy/aoc18"

	"strings"
)

var Solutions = aoc18.Day{
	A: a,
	B: b,
}


func reduce(polymer string) (string) {
	result := make([]uint8, 0, len(polymer))
	for i := 0; i < len(polymer); i++ {
		if len(result) <= 0 {
			result = append(result, polymer[i])
		} else if result[len(result) - 1] ^ polymer[i] == 32 { // matches A a, B b...
			result = result[:len(result)-1]
		} else {
			result = append(result, polymer[i])
		}
	}
	return string(result)
}

func a() interface{} {
	lines, err := aoc18.GetLines("day5/input")
	if err != nil {
		return err
	}
	return len(reduce(lines[0]))
}

func removeUnit(polymer string, unit rune) string {
	// Replace lower case char
	polymer = strings.Replace(polymer, string(unit), "", -1)
	// Replace upper case char
	polymer = strings.Replace(polymer, string(unit - 32), "", -1)
	return polymer
}

func b() interface{} {
	lines, err := aoc18.GetLines("day5/input")
	if err != nil {
		return err
	}

	var polymerLengths = make(chan int, 100)
	answers := 0
	for letter := 'a'; letter <= 'z'; letter++  {
		go func(l rune) {
			// Send back length on chanel
			polymerLengths <- len(reduce(removeUnit(lines[0], l)))
		}(letter)
		answers++
	}
	shortest := <-polymerLengths
	for answers--; answers > 0; answers-- {
		l := <-polymerLengths
		if l < shortest {
			shortest = l
		}
	}
	return shortest
}
