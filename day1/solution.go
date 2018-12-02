package day1

import (
	"github.com/Gurgy/aoc18"
	"fmt"
	"strconv"
)

var Solutions = aoc18.Day{
	A:  a,
	B:  b,
}


func a() interface{} {
	f := 0
	c, err := aoc18.GetLines("day1/input")
	if err != nil {
		return err
	}
	for line := range c {
		value, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		f += value
	}
	if err != nil {
		return err
	}
	return f
}

func b() interface{} {
	f := 0
	fs := map[int]bool{0: true}
	for {
		c, err := aoc18.GetLines("day1/input")
		if err != nil {
			return err
		}
		for line := range c {
			value, _ := strconv.Atoi(line)

			f += value

			if fs[f] == true {
				return f
			} else {
				fs[f] = true
			}

		}
	}
	return fmt.Errorf("unreachable")
}