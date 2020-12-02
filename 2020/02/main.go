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
			one, _   = strconv.Atoi(matches[1])
			two, _   = strconv.Atoi(matches[2])
			letter   = matches[3]
			password = matches[4]
		)

		if (string(password[one-1]) == letter && string(password[two-1]) != letter) || (string(password[two-1]) == letter && string(password[one-1]) != letter) {
			compliant++
		}
	}

	fmt.Printf("part 2 result: %d", compliant)
}
