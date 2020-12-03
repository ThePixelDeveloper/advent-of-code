package main

import (
	"fmt"
	"github.com/thepixeldeveloper/advent-of-code/2020/03/fixtures"
)

func main() {
	forest := fixtures.Input

	fmt.Printf("part 1 result: %d\n",
		calculate(forest, 1, 3))

	fmt.Printf(
		"part 2 result: %d\n",
		calculate(forest, 1, 1)*
			calculate(forest, 1, 3)*
			calculate(forest, 1, 5)*
			calculate(forest, 1, 7)*
			calculate(forest, 2, 1),
	)
}

func calculate(forest [][]string, moveDown int, moveRight int) int {
	var (
		i     = 0
		j     = 1
		trees = 0
	)

	for {
		i = i + moveDown
		j = j + moveRight

		if i >= len(forest) {
			break
		}

		if forest[i][(j-1)%len(forest[i])] == "#" {
			trees++
		}
	}

	return trees
}
