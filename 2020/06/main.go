package main

import (
	"fmt"
	"github.com/thepixeldeveloper/advent-of-code/2020/06/fixtures"
	"strings"
)

func main() {
	landingCards := fixtures.Input

	groups := strings.Split(landingCards, "\n\n")

	var (
		total = 0
	)

	for _, group := range groups {
		var (
			count = map[uint8]int{}
		)

		for _, s := range strings.Split(group, "\n") {
			answerUnique(count, s)
		}

		total += len(count)
	}

	fmt.Printf("part 1 result: %d", total)
}

func answerUnique(count map[uint8]int, input string) {
	for i, _ := range input {
		count[input[i]]++
	}
}
