package main

import (
	"fmt"

	"github.com/maxmcd/aoc/2020/lib"
)

func main() {
	input := lib.InputIntLines()
	for _, x := range input {
		for _, y := range input {
			for _, z := range input {
				if x+y+z == 2020 {
					fmt.Println(x, y, z, x*y*z)
					return
				}
			}
		}
	}
}
