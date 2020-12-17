package main

import (
	"fmt"
	"github.com/thepixeldeveloper/advent-of-code/2020/12/fixtures"
	"github.com/thepixeldeveloper/advent-of-code/2020/12/ship"
	"strings"
)

func main() {
	instructions := ship.Parse(strings.Split(fixtures.Input, "\n"))

	_, x, y := ship.Simulate(instructions, 90, 0, 0)

	fmt.Printf("part 1 result: %d", abs(x)+abs(y))
}

func abs(input int) int {
	if input < 0 {
		return -input
	}

	return input
}
