package day5

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/stvp/assert"
)

type myTestSuite struct {
	suite.Suite
}

func (suite *myTestSuite) SetupSuite() {
}

func (suite *myTestSuite) TearDownSuite() {
}

func (suite *myTestSuite) SetupTest() {
}

func (suite *myTestSuite) TearDownTest() {
}

func (suite *myTestSuite) TestReduce() {
	assert.Equal(suite.T(), reduce("aA"), "")
	assert.Equal(suite.T(), reduce("abBA"), "")
	assert.Equal(suite.T(), reduce("abAB"), "abAB")
	assert.Equal(suite.T(), reduce("aabAAB"), "aabAAB")
	assert.Equal(suite.T(), reduce("dabAcCaCBAcCcaDA"), "dabCBAcaDA")
}

func (suite *myTestSuite) TestRemoveUnit() {
	assert.Equal(suite.T(), removeUnit("aA", 'a'), "")
	assert.Equal(suite.T(), removeUnit("aA", 'b'), "aA")
	assert.Equal(suite.T(), removeUnit("dabAcCaCBAcCcaDA", 'c'), "dabAaBAaDA")
}


func TestMyTestSuite(t *testing.T) {
	tests := new(myTestSuite)
	suite.Run(t, tests)
}
