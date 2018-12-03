package day1

import (
	"fmt"
	"github.com/Gurgy/aoc18"
	"strconv"
)

var Solutions = aoc18.Day{
	A: a,
	B: b,
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

func getValues() ([]int, error) {
	c, err := aoc18.GetLines("day1/input")
	if err != nil {
		return nil, err
	}
	values := make([]int, 0, 1000)
	for line := range c {
		value, _ := strconv.Atoi(line)
		values = append(values, value)
	}
	return values, nil
}

func b() interface{} {
	f := 0
	fs := map[int]bool{0: true}
	values, err := getValues()
	if err != nil {
		return err
	}
	for {
		for _, value := range values {
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
