package lib

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func InputLines() []string {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(f)
	var out []string
	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		out = append(out, strings.TrimSuffix(line, "\n"))
	}
	return out
}

func InputIntLines() []int {
	var out []int
	for _, line := range InputLines() {
		if line == "" {
			continue
		}
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(line)
		}
		out = append(out, i)
	}
	return out
}
