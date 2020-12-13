package main

import (
	"fmt"
	"github.com/thepixeldeveloper/advent-of-code/2020/09/fixtures"
	"strconv"
	"strings"
)

func main() {
	var (
		target = 85848519
		values []int
		input  = fixtures.Input
		lines  = strings.Split(input, "\n")
	)

	values = make([]int, len(lines))

	for i, line := range lines {
		atoi, _ := strconv.Atoi(line)
		values[i] = atoi
	}

	result := EqualsTarget(target, 0, values)

	fmt.Printf("part 2 result: %d", result)
}

func EqualsTarget(target int, start int, values []int) int {
	var (
		smallest int
		largest  int
		count    int
	)

	smallest = values[start]
	largest = values[start]

	for i := start; i < len(values); i++ {
		count += values[i]

		if values[i] > largest {
			largest = values[i]
		}

		if values[i] < smallest {
			smallest = values[i]
		}

		if count > target {
			return EqualsTarget(target, start+1, values)
		}

		if count == target {
			return smallest + largest
		}
	}

	return 0
}
