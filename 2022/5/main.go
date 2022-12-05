package main

import (
	"fmt"
	"strings"

	"github.com/maxmcd/aoc"
)

func main() {
	run(false)
	run(true)

}

func run(is9001 bool) {

	initialized := false
	stacks := [9][]string{}
	stackLines := []string{}

	initializeStacks := func() {
		for i := len(stackLines) - 1; i >= 0; i-- {
			for x, letter := range stackLines[i] {
				if x%4 == 1 && string(letter) != " " {
					stacks[x/4] = append(stacks[x/4], string(letter))
				}
			}
		}
	}

	popFromStack := func(i int) (x string) {
		x, stacks[i] = stacks[i][len(stacks[i])-1], stacks[i][:len(stacks[i])-1]
		return x
	}

	aoc.ReadInput(func(line string) {
		if strings.Contains(line, "[") {
			stackLines = append(stackLines, line)
		}
		if strings.HasPrefix(line, "move") {
			if !initialized {
				initializeStacks()
				initialized = true
			}
			parts := strings.Split(line, " ")
			amount := aoc.Atoi(parts[1])
			start := aoc.Atoi(parts[3])
			end := aoc.Atoi(parts[5])
			if is9001 {
				values := []string{}
				for i := 0; i < amount; i++ {
					values = append(values, popFromStack(start-1))
				}
				for i := len(values) - 1; i >= 0; i-- {
					stacks[end-1] = append(stacks[end-1], values[i])
				}
			} else {
				for i := 0; i < amount; i++ {
					stacks[end-1] = append(stacks[end-1], popFromStack(start-1))
				}
			}
		}
	})
	fmt.Println(stacks)
	var out string
	for _, stack := range stacks {
		out += stack[len(stack)-1]
	}
	fmt.Println(out)
}
