package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/pkg/errors"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	totals := []int{}
	val := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			totals = append(totals, val)
			val = 0
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				panic(errors.Wrap(err, line))
			}
			val += num
		}
	}
	sort.Slice(totals, func(i, j int) bool { return totals[i] < totals[j] })
	top3 := totals[len(totals)-3:]
	fmt.Println(top3)
	total := 0
	for _, v := range top3 {
		total += v
	}
	fmt.Println(total)
}
