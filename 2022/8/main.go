package main

import (
	"fmt"

	"github.com/maxmcd/aoc"
)

type Tree struct {
	Height  int
	Visible bool
	Score   int
}

func one() {
	trees := [][]*Tree{}
	aoc.ReadInput(func(line string) {
		row := []*Tree{}
		for _, t := range line {
			row = append(row, &Tree{Height: aoc.Atoi(string(t))})
		}
		trees = append(trees, row)
	})

	// Mark outside trees as visible
	for _, tree := range append(trees[0], trees[len(trees)-1]...) {
		tree.Visible = true
	}
	for _, row := range trees {
		row[0].Visible = true
		row[len(row)-1].Visible = true
	}

	for _, row := range trees {
		// Left -> right rows
		height := row[0].Height
		for _, tree := range row {
			if tree.Height > height {
				tree.Visible = true
				height = tree.Height
			}
		}
		height = row[len(trees)-1].Height
		// Right -> left rows
		for i := len(row) - 1; i >= 0; i-- {
			tree := row[i]
			if tree.Height > height {
				tree.Visible = true
				height = tree.Height
			}
		}
	}

	// columns to top bottom
	for i := 0; i < len(trees[0]); i++ {
		height := trees[0][i].Height
		for j := 0; j < len(trees); j++ {
			tree := trees[j][i]
			if tree.Height > height {
				tree.Visible = true
				height = tree.Height
			}
		}

		height = trees[len(trees)-1][i].Height
		for j := len(trees) - 1; j >= 0; j-- {
			tree := trees[j][i]
			if tree.Height > height {
				tree.Visible = true
				height = tree.Height
			}
		}
	}

	count := 0
	for _, row := range trees {
		for _, tree := range row {
			if tree.Visible {
				count++
			}
		}
	}
	fmt.Println(count)
}

func main() {
	trees := [][]*Tree{}
	aoc.ReadInput(func(line string) {
		row := []*Tree{}
		for _, t := range line {
			row = append(row, &Tree{Height: aoc.Atoi(string(t))})
		}
		trees = append(trees, row)
	})

	safeGet := func(x, y int) *Tree {
		if y < 0 {
			return nil
		}
		if x < 0 {
			return nil
		}
		if y > len(trees)-1 {
			return nil
		}
		if x > len(trees[0])-1 {
			return nil
		}
		return trees[y][x]
	}

	for y, row := range trees {
		for x, tree := range row {
			score := 1
			for _, mutator := range []func(int, int) (int, int){
				func(x, y int) (int, int) {
					return x, y + 1
				},
				func(x, y int) (int, int) {
					return x, y - 1
				},
				func(x, y int) (int, int) {
					return x + 1, y
				},
				func(x, y int) (int, int) {
					return x - 1, y
				},
			} {
				thisScore := 0
				nx, ny := mutator(x, y)
				for {
					ot := safeGet(nx, ny)
					if ot == nil {
						break
					}
					thisScore++
					if ot.Height >= tree.Height {
						break
					}
					nx, ny = mutator(nx, ny)
				}
				score = score * thisScore
			}

			tree.Score = score
		}
	}

	score := 0
	for _, row := range trees {
		for _, tree := range row {
			if tree.Score > score {
				score = tree.Score
			}
		}
	}
	fmt.Println(score)
}
