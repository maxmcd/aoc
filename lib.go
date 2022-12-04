package aoc

import (
	"bufio"
	"os"
)

func ReadInput(cb func(line string)) {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		cb(line)
	}
}
