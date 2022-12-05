package aoc

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"
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

func Atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

func Sscanf(str string, format string, a ...any) {
	num, err := fmt.Sscanf(str, format, a...)
	if err != nil {
		panic(errors.Wrapf(err, "sscanf err with str %q and format %q", str, format))
	}
	if num != len(a) {
		panic(fmt.Errorf("expected %d args in scan but found %d in string %q", len(a), num, str))
	}
}
