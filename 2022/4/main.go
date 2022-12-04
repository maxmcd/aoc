package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	containsCount := 0

	overlapAtAll := 0
	for scanner.Scan() {
		line := scanner.Text()
		// 1-2,2-96
		first, second, _ := strings.Cut(line, ",")
		fStartS, fEndS, _ := strings.Cut(first, "-")
		sStartS, sEndS, _ := strings.Cut(second, "-")

		fStart, _ := strconv.Atoi(fStartS)
		fEnd, _ := strconv.Atoi(fEndS)
		sStart, _ := strconv.Atoi(sStartS)
		sEnd, _ := strconv.Atoi(sEndS)

		if fStart <= sStart && fEnd >= sEnd {
			containsCount++
			fmt.Println(line)
		} else if sStart <= fStart && sEnd >= fEnd {
			containsCount++
			fmt.Println(line)
		}
		if sStart <= fEnd && sEnd >= fStart {
			overlapAtAll++
		}
	}
	fmt.Println(containsCount, overlapAtAll)
}
