package main

import (
	"fmt"
	"github.com/thepixeldeveloper/advent-of-code/2020/10/fixtures"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		lines    = strings.Split(fixtures.Input, "\n")
		voltages = make([]int, 0, len(lines))
	)

	for i := 0; i < len(lines); i++ {
		voltages = append(voltages, parseVoltage(lines[i]))
	}

	sort.Ints(voltages)

	fmt.Printf("part 2 result: %+v", possibilities(deltas(voltages)))
}
func possibilities(deltas []int) int {
	var (
		recurse func(deltas []int) int
		cache   = make(map[string]int)
	)

	recurse = func(deltas []int) int {
		deltasAsString := fmt.Sprintf("%+v", deltas)

		if val, ok := cache[deltasAsString]; ok {
			return val
		}

		if len(deltas) <= 1 {
			return 1
		}

		if deltas[0]+deltas[1] > 3 {
			return recurse(deltas[1:])
		}

		meh := make([]int, len(deltas))

		copy(meh, deltas)

		meh[1] = meh[0] + meh[1]

		result := recurse(deltas[1:]) + recurse(meh[1:])

		cache[deltasAsString] = result

		return result
	}

	return recurse(deltas)
}

func deltas(voltages []int) []int {
	var (
		previous int
		deltas   []int
	)

	for _, voltage := range voltages {
		deltas = append(deltas, voltage-previous)
		previous = voltage
	}

	deltas = append(deltas, 3)

	return deltas
}

func parseVoltage(input string) int {
	voltage, _ := strconv.Atoi(input)
	return voltage
}
