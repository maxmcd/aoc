package main

import (
	"fmt"

	"github.com/maxmcd/aoc/2020/lib"
)

func main() {
	lines := lib.InputLines()
	width := len(lines[0])
	fmt.Println(len(lines), width)

	// get constants from the output above
	grid := [323][31]bool{}

	for x := range grid {
		for y := range grid[x] {
			grid[x][y] = lines[x][y] == '#'
		}
	}

	calc := func(yinc, xinc int) int {
		var x int
		var y int
		var bang int
		for {
			y += yinc
			x += xinc
			if x+1 > len(grid) {
				break
			}
			if grid[x][y%31] {
				bang++
			}
		}
		fmt.Println(yinc, xinc, bang)
		return bang
	}
	fmt.Println(calc(1, 1) *
		calc(3, 1) *
		calc(5, 1) *
		calc(7, 1) *
		calc(1, 2))

}
