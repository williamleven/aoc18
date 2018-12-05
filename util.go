package aoc18

import (
	"bufio"
	"os"
)

func GetLinesStream(file string) (chan string, error) {
	r, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	c := make(chan string)
	go func() {
		reader := bufio.NewReader(r)
		var line []byte
		line, err := readLine(reader)
		for err == nil {
			c <- string(line)
			line, err = readLine(reader)
		}
		close(c)
	}()
	return c, nil
}

func readLine(reader *bufio.Reader) ([]byte, error) {
	var line []byte
	for hasMore := true; hasMore;  {
		var additional []byte
		var err error
		additional, hasMore, err = reader.ReadLine()
		if err != nil {
			return nil, err
		}
		line = append(line, additional...)
	}
	return line, nil
}

func GetLines(file string) ([]string, error) {
	c, err := GetLinesStream(file)
	if err != nil {
		return nil, err
	}
	lines := make([]string, 0, 1000)
	for line := range c {
		lines = append(lines, line)
	}
	return lines, nil
}
