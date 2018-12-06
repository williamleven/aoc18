package day6

import (
	"fmt"
	"github.com/Gurgy/aoc18"
	"strings"
	"math"
)

var Solutions = aoc18.Day{
	A: a,
	B: b,
}

type coordinate struct {
	x int
	y int
}

func abs (a int) int {
	return int(math.Abs(float64(a)))
}

func (c coordinate) distanceTo(other coordinate) int {
	return abs(c.x - other.x) + abs(c.y - other.y)
}

func (c coordinate) getArea(cs coordinates) int {
	visited := make(coordinates, 0, 1000)
	return c._getArea(cs, c, &visited)
}
func (c coordinate) _getArea(cs coordinates, current coordinate, visited *coordinates) int {
	if visited.contains(current) {
		return 0
	} else {
		if closest, _ := cs.closesTo(current); closest != nil && c.equals(*closest) {
			*visited = append(*visited, current)
			area := 1
			area += c._getArea(cs, coordinate{current.x+1, current.y+1}, visited)
			area += c._getArea(cs, coordinate{current.x-1, current.y+1}, visited)
			area += c._getArea(cs, coordinate{current.x+1, current.y-1}, visited)
			area += c._getArea(cs, coordinate{current.x-1, current.y-1}, visited)
			return area
		} else {
			return 0
		}
	}
}

func (c coordinate) equals(other coordinate) bool {
	return c.x == other.x && c.y == other.y
}

type coordinates []coordinate

func (cs coordinates) outerBounds() (low, high coordinate) {

	low.x = cs[0].x
	low.y = cs[0].y

	high.x = cs[0].x
	high.y = cs[0].y

	for _, c := range cs[1:] {
		if low.x > c.x {
			low.x = c.x
		}
		if low.y > c.y {
			low.y = c.y
		}

		if high.x < c.x {
			high.x = c.x
		}
		if high.y < c.y {
			high.y = c.y
		}
	}

	return low, high
}

func (cs coordinates) contains(other coordinate) bool {
	for _, c := range cs {
		e := c.equals(other)
		if e {
			return true
		}
	}
	return false
}

// TODO some cache here
func (cs coordinates) closesTo(point coordinate) (*coordinate, int) {
	closest := cs[0]
	closestDistance := closest.distanceTo(point)
	similar := false
	for _,c := range cs[1:] {
		distance := c.distanceTo(point)
		if distance < closestDistance {
			closest = c
			closestDistance = distance
			similar = false
		}else if distance == closestDistance {
			similar = true
		}
	}
	if similar {
		return nil, closestDistance
	}
	return &closest, closestDistance
}

func getCoordinates() coordinates {
	c, err := aoc18.GetLines("day6/sample")
	if err != nil {
		panic(err)
	}
	
	result := make([]coordinate, 0, 100)
	
	for line := range c {
		coordinate := coordinate{}
		fmt.Fscanf(strings.NewReader(line), "%d, %d", &coordinate.x, &coordinate.y)
		result = append(result, coordinate)
	}
	return result
}

func getAllSquareBounds(low, high coordinate) (coordinates) {
	result := make(coordinates, 0, (high.x-low.x+high.y-low.y+8)*2)
	for i := low.x; i <= high.x; i++ {
		result = append(result, coordinate{i, low.y - 1}, coordinate{i, high.y + 1})
	}
	for i := low.y; i <= high.y; i++ {
		result = append(result, coordinate{low.x-1, i}, coordinate{high.x+1, i})
	}
	return result
}

func a() interface{} {
	coords := getCoordinates()
	lowBound, highBound := coords.outerBounds()
	border := getAllSquareBounds(lowBound, highBound)
	invalidPoints := make(coordinates, 0, 100)
	for _, borderPoint := range border {
		invalid, _ := coords.closesTo(borderPoint)
		if invalid != nil {
			if !invalidPoints.contains(*invalid) {
				invalidPoints = append(invalidPoints, *invalid)
			}
		}
	}


	fmt.Println(invalidPoints)
	biggestArea := 0
	for _, c := range coords {
		if !invalidPoints.contains(c) {
			fmt.Println(c)
			area := c.getArea(coords)
			fmt.Println(area)
			if area > biggestArea {
				biggestArea = area
			}
		}
	}
	return biggestArea
}

func b() interface{} {
	return fmt.Errorf("not implemented")
}
