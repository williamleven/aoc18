package day4


import (
	"fmt"
	"github.com/Gurgy/aoc18"
	"strings"
	"time"
	"sort"
)
var Solutions = aoc18.Day{
	A: a,
	B: b,
}


/*
	TODO: Refactor this answer
 */

type logEntry struct {
	time time.Time
	verb string
	id int
}

type bed struct {
	from time.Time
	to time.Time
}

type beds []bed

func (bs beds) duration() int  {
	sum := 0
	for _, bed := range bs {
		sum += bed.duration()
	}
	return sum
}

func (bs beds) asleep(minute int) int  {
	sum := 0
	for _, bed := range bs {
		if int(bed.from.Minute()) <= minute && int(bed.to.Minute()) > minute {
			sum++
		}
	}
	return sum
}

func (bs beds) mostActiveMinute() int  {
	sum := [60]int{}
	for _, bed := range bs {
		for i := int(bed.from.Minute()); i < int(bed.to.Minute()); i++ {
			sum[i]++
		}
	}
	highest := 0
	highestM := 0
	for m, s := range sum {
		if s > highest {
			highest = s
			highestM = m
		}
	}
	return highestM
}

func (b bed) duration() int  {
	return int(b.to.Sub(b.from).Minutes())
}

type guards map[int]beds

func (g guards) sleptMost() int {
	highestTime := 0
	highestTimeId := 0
	for id, beds := range g {
		if beds.duration() >= highestTime {
			highestTime = beds.duration()
			highestTimeId = id
		}
	}
	return highestTimeId
}

func (g guards) sleptMostAt(minute int) (int, int) {
	highestTime := 0
	highestTimeId := 0
	for id, beds := range g {
		if beds.asleep(minute) >= highestTime {
			highestTime = beds.asleep(minute)
			highestTimeId = id
		}
	}
	return highestTimeId, highestTime
}

func getGuards() guards {
	c, err := aoc18.GetLines("day4/input")
	if err != nil {
		panic(err)
	}

	entries := make([]logEntry, 0, 1500)
	for line := range c {
		entry := logEntry{}
		var (
			timeString string
			dateString string
		)
		fmt.Fscanf(strings.NewReader(line), "[%s %5s] %s #%d", &dateString, &timeString, &entry.verb, &entry.id)

		t, _ := time.Parse("2006-01-02 15:04", dateString  + " " + timeString)
		entry.time = t
		entries = append(entries, entry)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].time.Before(entries[j].time)
	})

	var latestId int
	var fell time.Time
	guards := make(map[int]beds)
	for _, entry := range entries {
		if entry.verb == "Guard" {
			latestId = entry.id
		} else if entry.verb == "falls"{
			fell = entry.time
		} else if entry.verb == "wakes"{
			guards[latestId] = append(guards[latestId], bed{
				from: fell,
				to:   entry.time,
			})
		}
	}
	return guards
}

func a() interface{} {
	guards := getGuards()

	id := guards.sleptMost()

	return guards[id].mostActiveMinute() * id
}

func b() interface{} {
	guards := getGuards()

	highestDuration := 0
	highestDurationID := 0
	highestDurationMinute := 0
	for minute := 0; minute < 60; minute++ {
		masterId, duration := guards.sleptMostAt(minute)
		if duration > highestDuration {
			highestDurationID = masterId
			highestDuration = duration
			highestDurationMinute = minute
		}
	}
	return highestDurationID * highestDurationMinute
}
