package main

import (
	"fmt"
	"time"
	"github.com/Gurgy/aoc18"
	"github.com/Gurgy/aoc18/day1"
	"github.com/Gurgy/aoc18/day2"
	"github.com/Gurgy/aoc18/day3"
	"github.com/Gurgy/aoc18/day4"
)

var days = map[int]aoc18.Day{
	1: day1.Solutions,
	2: day2.Solutions,
	3: day3.Solutions,
	4: day4.Solutions,
}

func main() {
	for id, day := range days {
		printDay(id, day)
	}
}

func printDay(id int, day aoc18.Day) {
	printTask(fmt.Sprintf("%dA", id), day.A)
	printTask(fmt.Sprintf("%dB", id), day.B)
}

func printTask(id string, task func() interface{}) {
	start := time.Now()
	answer := task()
	time := time.Since(start)
	fmt.Printf("%s: %35v (%s) \n", id, answer, time)
}
