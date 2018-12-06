package day6

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestAbs(t *testing.T)  {
	assert.Equal(t, 3, abs(-3))
	assert.Equal(t, 3, abs(3))
	assert.Equal(t, 0, abs(0))
}

func TestDistanceTo(t *testing.T) {
	assert.Equal(t, 1, coordinate{0, 0}.distanceTo(coordinate{1, 0}))
	assert.Equal(t, 2, coordinate{0, 0}.distanceTo(coordinate{1, 1}))
	assert.Equal(t, 3, coordinate{0, 0}.distanceTo(coordinate{2, 1}))
	assert.Equal(t, 3, coordinate{0, 0}.distanceTo(coordinate{-2, 1}))
}

func TestArea(t *testing.T) {
	assert.Equal(t, 1, coordinate{0, 0}.getArea(coordinates{
		coordinate{1, 1},
		coordinate{1, -1},
		coordinate{-1, 1},
		coordinate{-1, -1},
		coordinate{0, 0},
	}))

	assert.Equal(t, 5, coordinate{0, 0}.getArea(coordinates{
		coordinate{2, 2},
		coordinate{2, -2},
		coordinate{-2, 2},
		coordinate{-2, -2},
		coordinate{0, 0},
	}))

	assert.Equal(t, 17, coordinate{5, 5}.getArea(coordinates{
		coordinate{1, 1},
		coordinate{1, 6},
		coordinate{8, 3},
		coordinate{3, 4},
		coordinate{5, 5},
		coordinate{8, 9},
	}))



	assert.Equal(t, 9, coordinate{3, 4}.getArea(coordinates{
		coordinate{1, 1},
		coordinate{1, 6},
		coordinate{8, 3},
		coordinate{3, 4},
		coordinate{5, 5},
		coordinate{8, 9},
	}))
}

func TestCorners(t *testing.T) {
	low, high := coordinates{
		coordinate{3, 1},
		coordinate{1, -5},
		coordinate{-1, 1},
		coordinate{-1, -1},
	}.outerBounds()

	assert.True(t, low.equals(coordinate{-1, -5}))
	assert.True(t, high.equals(coordinate{3, 1}))
}

func TestEquals(t *testing.T)  {
	assert.True(t, coordinate{1, 12}.equals(coordinate{1, 12}))
	assert.False(t, coordinate{1, 12}.equals(coordinate{3, 12}))
	assert.False(t, coordinate{1, 12}.equals(coordinate{1, 1}))
}

func TestBorder(t *testing.T) {
	cs := getAllSquareBounds(coordinate{-1, -1}, coordinate{1, 1})
	assert.Equal(t, 12, len(cs))
}

func TestClosest(t *testing.T) {
	cs := coordinates{
		coordinate{5, 0},
		coordinate{0, 15},
		coordinate{-5, 0},
	}
	c, i, err := cs.closesTo(coordinate{0, 0})
	assert.Error(t, err)
	assert.Equal(t, 5, i)

	c, i, err = cs.closesTo(coordinate{1, 0})
	assert.True(t, c.equals(coordinate{5, 0}))
	assert.NoError(t, err)
	assert.Equal(t, 4, i)

	c, i, err = cs.closesTo(coordinate{6, 0})
	assert.True(t, c.equals(coordinate{5, 0}))
	assert.NoError(t, err)
	assert.Equal(t, 1, i)

	c, i, err = cs.closesTo(coordinate{0, 14})
	assert.True(t, c.equals(coordinate{0, 15}))
	assert.NoError(t, err)
	assert.Equal(t, 1, i)
}

func TestInvalidPoints(t *testing.T) {
	invalid := coordinates{
		coordinate{1, 1},
		coordinate{1, -1},
		coordinate{-1, 1},
		coordinate{-1, -1},
		coordinate{0, 0},
	}.getInvalidPoints()

	assert.False(t, invalid.contains(coordinate{0, 0}))
	assert.Equal(t, 4, len(invalid))

}