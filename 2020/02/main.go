package main

import (
	"fmt"
	"github.com/thepixeldeveloper/advent-of-code/2020/02/fixtures"
	"regexp"
	"strconv"
)

func main() {
	var (
		compliant = 0
		regex     = regexp.MustCompile("^([0-9]+)\\-([0-9]+)\\s([a-z]):\\s(.*)$")
	)

	for i := 0; i < len(fixtures.Input); i++ {
		matches := regex.FindStringSubmatch(fixtures.Input[i])

		var (
			min, _   = strconv.Atoi(matches[1])
			max, _   = strconv.Atoi(matches[2])
			letter   = matches[3]
			password = matches[4]
			count    = 0
		)

		for _, c := range password {
			if string(c) == letter {
				count++
			}
		}

		if count >= min && count <= max {
			compliant++
		}
	}

	fmt.Printf("part 1 result: %d", compliant)
}
