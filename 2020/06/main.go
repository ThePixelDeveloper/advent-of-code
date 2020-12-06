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
		answers := strings.Split(group, "\n")

		// If the group size is one, short circuit
		if len(answers) == 1 {
			total += len(answers[0])
			continue
		}

		var (
			count = map[uint8]int{}
		)

		for _, s := range answers {
			answerUnique(count, s)
		}

		for _, i := range count {
			if i == len(answers) {
				total += 1
			}
		}
	}

	fmt.Printf("part 2 result: %d", total)
}

func answerUnique(count map[uint8]int, input string) {
	for i, _ := range input {
		count[input[i]]++
	}
}
