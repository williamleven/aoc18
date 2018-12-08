package day8

import (
	"fmt"
	"github.com/Gurgy/aoc18"
	"strings"
	"strconv"
	"encoding/json"
)

var Solutions = aoc18.Day{
	A: a,
	B: b,
}
type node struct {
	children []node `json:"children"`
	metadata []int `json:"metadata"`
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
	for i := 0; i < len(head.children); i++   {
		head.children[i], input, err = parseTree(input[2:])
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

func a() interface{} {
	lines, err := aoc18.GetLines("day8/sample")
	if err != nil {
		panic(err)
	}
	line := lines[0]
	entries := strings.Split(line, " ")
	head, remaining, err := parseTree(entries)
	if err != nil {
		panic(err)
	}
	json, err := json.Marshal(head)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json))
	return remaining
}

func b() interface{} {
	return fmt.Errorf("not implemented")
}
