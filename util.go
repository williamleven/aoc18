package aoc18

import (
	"bufio"
	"os"
)

func GetLines(file string) (chan string, error) {
	r, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	c := make(chan string)
	go func() {
		reader := bufio.NewReader(r)
		var line []byte
		line, _, err := reader.ReadLine()
		for err == nil {
			c <- string(line)
			line, _, err = reader.ReadLine()
		}
		close(c)
	}()
	return c, nil
}
