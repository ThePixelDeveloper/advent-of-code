package main

import (
	"fmt"
	"github.com/thepixeldeveloper/advent-of-code/2020/09/fixtures"
	"strconv"
	"strings"
)

func main() {
	var (
		preamble       []int
		preambleSums   map[int]struct{}
		preambleLength = 25
		input          = fixtures.Input
		lines          = strings.Split(input, "\n")
	)

	preamble = make([]int, len(lines))

	for i, line := range lines {
		atoi, _ := strconv.Atoi(line)
		preamble[i] = atoi
	}

	for i := preambleLength; i < len(preamble); i++ {
		preambleSums = CalculateSums(preamble, i-preambleLength, preambleLength)
		if _, ok := preambleSums[preamble[i]]; !ok {
			fmt.Printf("part 1 result: %d\n", preamble[i])
		}
	}
}

func CalculateSums(preamble []int, start int, preambleLength int) map[int]struct{} {
	var (
		preambleSums map[int]struct{}
	)

	preambleSums = make(map[int]struct{}, preambleLength*preambleLength)

	for _, i := range preamble[start : start+preambleLength] {
		for _, j := range preamble[start : start+preambleLength] {
			if i == j {
				continue
			}
			preambleSums[i+j] = struct{}{}
		}
	}

	return preambleSums
}
