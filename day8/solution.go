package day8

import (
	"github.com/Gurgy/aoc18"
	"strings"
	"strconv"
)

var Solutions = aoc18.Day{
	A: a,
	B: b,
}
type node struct {
	children []node
	metadata []int
}
func parseTree(input []string) (node, []string, error) {
	nNodes, err := strconv.Atoi(input[0])
	if err != nil {
		return node{}, nil, err
	}
	nData, err := strconv.Atoi(input[1])
	if err != nil {
		return node{}, nil, err
	}
	head := node{
		children: make([]node, nNodes),
		metadata: make([]int, nData),
	}
	input = input[2:]
	for i := 0; i < len(head.children); i++   {
		head.children[i], input, err = parseTree(input)
		if err != nil {
			return node{}, nil, err
		}
	}
	for i := 0; i < len(head.metadata); i++   {
		data, err := strconv.Atoi(input[0])
		if err != nil {
			return node{}, nil, err
		}
		input = input[1:]
		head.metadata[i] = data
	}

	return head, input, nil
}

func (n node) sumOfMetaValues() int  {
	sum := 0
	for _, value := range n.metadata {
		sum += value
	}
	return sum
}

func (n node) deepSumOfMetaValues() int  {
	sum := n.sumOfMetaValues()
	for _, child := range n.children  {
		sum += child.deepSumOfMetaValues()
	}
	return sum
}

func (n node) valueOf() int {
	if len(n.children) == 0 {
		return n.sumOfMetaValues()
	} else {
		sum := 0
		for _, value := range n.metadata {
			if value > 0 && value <= len(n.children) {
				sum += n.children[value-1].valueOf()
			}
		}
		return sum
	}
}

func a() interface{} {
	lines, err := aoc18.GetLines("day8/input")
	if err != nil {
		panic(err)
	}
	line := lines[0]
	entries := strings.Split(line, " ")
	head, _, err := parseTree(entries)
	return head.deepSumOfMetaValues()
}

func b() interface{} {
	lines, err := aoc18.GetLines("day8/input")
	if err != nil {
		panic(err)
	}
	line := lines[0]
	entries := strings.Split(line, " ")
	head, _, err := parseTree(entries)
	return head.valueOf()
}
