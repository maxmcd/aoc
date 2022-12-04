package main

import (
	"bufio"
	"fmt"
	"os"
)

func findOneCommonLetter(a, b string) rune {
	letters := make(map[rune]bool)

	for _, r := range a {
		letters[r] = true
	}
	for _, r := range b {
		if _, found := letters[r]; found {
			return r
		}
	}
	panic(fmt.Sprintf("None found: %s %s", a, b))
}

func findOneCommonLetterOfThree(a, b, c string) rune {
	letters := make(map[rune][2]bool)

	for _, r := range a {
		v := letters[r]
		v[0] = true
		letters[r] = v
	}
	for _, r := range b {
		v := letters[r]
		v[1] = true
		letters[r] = v
	}

	for _, r := range c {
		if val, found := letters[r]; found && val[0] && val[1] {
			return r
		}
	}
	panic(fmt.Sprintf("None found: %s %s %s", a, b, c))
}

func one() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		first, second := line[:len(line)/2], line[len(line)/2:]
		commonLetter := findOneCommonLetter(first, second)
		score := letterScore(commonLetter)
		fmt.Println(first, second, commonLetter, score)
		total += score
	}

	fmt.Println(total)
}

func letterScore(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c) - 96
	}
	if c >= 'A' && c <= 'Z' {
		return int(c) - 38
	}
	panic(fmt.Sprintf("%s not found", string(c)))
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	group := []string{}
	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		group = append(group, line)
		if len(group) == 3 {
			commonLetter := findOneCommonLetterOfThree(group[0], group[1], group[2])
			score := letterScore(commonLetter)
			total += score
			group = nil
		}
	}

	fmt.Println(total)
}
