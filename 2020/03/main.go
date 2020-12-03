package main

import (
	"fmt"
	"github.com/thepixeldeveloper/advent-of-code/2020/03/fixtures"
)

func main() {
	forest := fixtures.Input

	var (
		i     = 0
		j     = 1
		trees = 0
	)

	for {
		// Part 1
		i = i + 1
		j = j + 3

		// Part 2 (Lazy!)
		//i = i + 1
		//j = j + 1

		//i = i + 1
		//j = j + 3

		//i = i + 1
		//j = j + 5

		//i = i + 1
		//j = j + 6

		//i = i + 2
		//j = j + 1

		if i >= len(forest) {
			break
		}

		location := forest[i][(j-1)%len(forest[i])]

		if location == "#" {
			trees++
		}
	}

	fmt.Printf("part 1 result: %d", trees)
}
