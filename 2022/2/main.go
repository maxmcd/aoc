package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func first() {
	var (
		rockL    = "A"
		paperL   = "B"
		scissorL = "C"

		rockR    = "X"
		paperR   = "Y"
		scissorR = "Z"
	)
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		first, second, _ := strings.Cut(line, " ")

		startScore := map[string]int{rockR: 1, paperR: 2, scissorR: 3}[second]

		switch first + second {
		case rockL + rockR, paperL + paperR, scissorL + scissorR:
			// tie
			startScore += 3
		case rockL + paperR, paperL + scissorR, scissorL + rockR:
			// win
			startScore += 6
		case rockL + scissorR, paperL + rockR, scissorL + paperR:
			// loss
		default:
			panic(line)
		}
		total += startScore
	}

	fmt.Println(total)
}

var (
	rockL    = "A"
	paperL   = "B"
	scissorL = "C"

	lose = "X"
	draw = "Y"
	win  = "Z"

	rock    = 1
	paper   = 2
	scissor = 3
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		first, second, _ := strings.Cut(line, " ")

		startScore := map[string]int{lose: 0, draw: 3, win: 6}[second]

		switch first + second {
		case scissorL + draw, rockL + lose, paperL + win:
			startScore += scissor
		case paperL + draw, rockL + win, scissorL + lose:
			startScore += paper
		case rockL + draw, paperL + lose, scissorL + win:
			startScore += rock
		default:
			panic(line)
		}
		total += startScore
	}

	fmt.Println(total)
}
