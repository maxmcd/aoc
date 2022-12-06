package main

import (
	"fmt"
	"os"

	"github.com/maxmcd/aoc"
)

func main() {
	aoc.ReadInput(func(line string) {
		// detectNDistinct(line, 4)
		detectNDistinct(line, 14)
	})
}

func detectNDistinct(line string, n int) {
	for i := 0; i < len(line)-n; i++ {
		track := make(map[string]struct{}, n)
		for j := 0; j < n; j++ {
			track[string(line[i+j])] = struct{}{}
			if len(track) < j+1 {
				// Early exit
				break
			}
		}
		if len(track) == n {
			fmt.Println(i+n, track)
			os.Exit(0)
		}
	}
}
