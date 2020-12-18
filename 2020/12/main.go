package main

import (
	"fmt"
	"github.com/thepixeldeveloper/advent-of-code/2020/12/fixtures"
	"github.com/thepixeldeveloper/advent-of-code/2020/12/ship"
	"strings"
)

func main() {
	instructions := ship.Parse(strings.Split(fixtures.Input, "\n"))

	_, _, _, shipX, shipY := ship.Simulate(instructions, 90, 10, 1, 0, 0)

	fmt.Printf("part 1 result: %d", abs(shipX)+abs(shipY))
}

func abs(input int) int {
	if input < 0 {
		return -input
	}

	return input
}
