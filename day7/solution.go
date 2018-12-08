package day7

import (
	"fmt"
	"github.com/Gurgy/aoc18"
	"strings"
	"sort"
)

var Solutions = aoc18.Day{
	A: a,
	B: b,
}



type pieces map[uint8][]uint8

func getPieces() pieces  {
	lines, err := aoc18.GetLinesStream("day7/input")
	if err != nil {
		panic(err)
	}

	ps := make(pieces)

	for line := range lines {
		var step string
		var req string
		fmt.Fscanf(strings.NewReader(line), "Step %s must be finished before step %s can begin.", &req, &step)

		ps[step[0]] = append(ps[step[0]], req[0])
		ps[req[0]] = append(ps[req[0]])
	}


	return ps
}

func (ps pieces) available() []uint8 {
	av := make([]uint8, 0, len(ps))
	for key := range ps {
		if len(ps[key]) == 0 {
			av = append(av, key)
		}
	}

	sort.Slice(av, func(i, j int) bool {
		return av[i] < av[j]
	})

	return av
}

func (ps pieces) complete(e uint8) (pieces) {
	for key := range ps {
		ps[key] = remove(
			ps[key],
			e,
		)
	}
	delete(ps, e)
	return ps
}

func remove(es []uint8, e uint8) []uint8{
	result := make([]uint8, 0, len(es))
	for _,value := range es {
		if value != e {
			result = append(result, value)
		}
	}
	return result
}

func a() interface{} {
	ps := getPieces()

	var solution []uint8

	for len(ps) > 0 {
		av := ps.available()
		ps = ps.complete(av[0])
		solution = append(solution, av[0])
	}
	return string(solution)
}

func b() interface{} {
	ps := getPieces()

	var solution []uint8
	workers := make(map[uint8]int)
	i := 0
	for ; len(ps) > 0; i++ {
		// clear upp workers
		for key := range workers {
			if workers[key] == i {
				delete(workers, key)
				ps = ps.complete(key)
				solution = append(solution, key)
			}
		}
		av := ps.available()
		// assign workers
		for j := 0; j < len(av) && len(workers) < 5; j++ {
			_, working := workers[av[j]]
			if !working {
				done := i + (int(av[j]) - 64) + 60 // 60 time units delat
				workers[av[j]] = done
			}
		}
	}
	return i-1 // -1 because nothing is done in the last step of the loop
}
