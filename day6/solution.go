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
		closest, _, err := cs.closesTo(current)
		if  err == nil && c.equals(closest) {
			*visited = append(*visited, current)
			area := 1
			area += c._getArea(cs, coordinate{current.x+1, current.y}, visited)
			area += c._getArea(cs, coordinate{current.x-1, current.y}, visited)
			area += c._getArea(cs, coordinate{current.x, current.y-1}, visited)
			area += c._getArea(cs, coordinate{current.x, current.y+1}, visited)
			return area
		} else {
			return 0
		}
	}
}

func (cs coordinates) findAllWithin(distance int) int {
	visited := make(coordinates, 0, 10000)
	a,  _ := cs._findAllWithin(distance, cs.findPointWithin(distance), visited)
	return a
}
func (cs coordinates) _findAllWithin(distance int, current coordinate, visited coordinates) (int, coordinates) {
	if visited.contains(current) {
		return 0, visited
	} else if cs.totalDistanceTo(current) < distance {
		visited = append(visited, current)
		area := 1
		var a int
		a, visited = cs._findAllWithin(distance, coordinate{current.x+1, current.y}, visited)
		area += a
		a, visited = cs._findAllWithin(distance, coordinate{current.x-1, current.y}, visited)
		area += a
		a, visited = cs._findAllWithin(distance, coordinate{current.x, current.y+1}, visited)
		area += a
		a, visited = cs._findAllWithin(distance, coordinate{current.x, current.y-1}, visited)
		area += a
		return area, visited
	} else {
		visited = append(visited, current)
		return 0, visited
	}
}

func (cs coordinates) findPointWithin(distance int) coordinate {
	lowBound, highBound := cs.outerBounds()
	visited := make(coordinates, 0, 100)
	c, _ := cs._findPointWithin(distance,
		coordinate{(lowBound.x + highBound.x) / 2, (lowBound.y + highBound.y) / 2},
		&visited,
	)
	return c
}
func (cs coordinates) _findPointWithin(distance int, current coordinate, visited *coordinates) (coordinate, error) {
	if visited.contains(current) {
		return coordinate{0, 0}, fmt.Errorf("not found")
	} else if cs.totalDistanceTo(current) < distance {
		return current, nil
	} else {
		return coordinate{0, 0}, fmt.Errorf("not implemnted")
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
func (cs coordinates) closesTo(point coordinate) (coordinate, int, error) {
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
		return coordinate{0, 0}, closestDistance, fmt.Errorf("middle ground")
	}
	return closest, closestDistance, nil
}

func getCoordinates() coordinates {
	c, err := aoc18.GetLines("day6/input")
	if err != nil {
		panic(err)
	}
	
	result := make([]coordinate, 0, 100)
	
	for _, line := range c {
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

func (cs coordinates) getInvalidPoints() coordinates  {
	lowBound, highBound := cs.outerBounds()
	border := getAllSquareBounds(lowBound, highBound)
	invalidPoints := make(coordinates, 0, 100)
	for _, borderPoint := range border {
		invalid, _, err:= cs.closesTo(borderPoint)
		if err == nil {
			if !invalidPoints.contains(invalid) {
				invalidPoints = append(invalidPoints, invalid)
			}
		}
	}
	return invalidPoints
}

func (cs coordinates) totalDistanceTo(point coordinate) int  {
	sum := 0
	for _, c := range cs {
		sum += c.distanceTo(point)
	}
	return sum
}

func a() interface{} {
	coords := getCoordinates()

	invalidPoints := coords.getInvalidPoints()

	biggestArea := 0
	for _, c := range coords {
		if !invalidPoints.contains(c) {
			area := c.getArea(coords)
			if area > biggestArea {
				biggestArea = area
			}
		}
	}
	return biggestArea
}

func b() interface{} {
	cs := getCoordinates()

	return cs.findAllWithin(10000)

}
