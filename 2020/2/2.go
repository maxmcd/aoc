package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/maxmcd/aoc/2020/lib"
)

// 13-14 v: hvvcvvvvvvvvvsvvv
func main() {
	lines := lib.InputLines()
	var valid int
	for _, line := range lines {
		parts := strings.Fields(line)
		rng := parts[0]
		var high int
		var low int
		var err error
		letter := parts[1]
		password := parts[2]
		{
			parts := strings.Split(rng, "-")
			low, err = strconv.Atoi(parts[0])
			if err != nil {
				panic(parts[0])
			}
			high, err = strconv.Atoi(parts[1])
			if err != nil {
				panic(parts[1])
			}
		}

		letter = strings.TrimSuffix(letter, ":")

		if password[high-1] == letter[0] && password[low-1] != letter[0] {
			valid++
			continue
		}
		if password[high-1] != letter[0] && password[low-1] == letter[0] {
			valid++
		}

		// histogram := map[rune]int{}
		// for _, chr := range password {
		// 	histogram[chr]++
		// }
		// count := histogram[rune(letter[0])]
		// if count <= high && count >= low {
		// 	valid++
		// }
	}
	fmt.Println(valid)
}
