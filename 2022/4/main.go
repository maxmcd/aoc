package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/maxmcd/aoc"
)

func main() {
	containsCount := 0

	overlapAtAll := 0
	aoc.ReadInput(func(line string) {
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
			// fmt.Println(line)
		} else if sStart <= fStart && sEnd >= fEnd {
			containsCount++
			// fmt.Println(line)
		}
		if sStart <= fEnd && sEnd >= fStart {
			overlapAtAll++
		}
	})
	fmt.Println(containsCount, overlapAtAll)
}
