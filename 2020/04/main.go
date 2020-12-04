package main

import (
	"fmt"
	"github.com/thepixeldeveloper/advent-of-code/2020/04/fixtures"
	"regexp"
)

func main() {
	// Formatting to a single line
	re := regexp.MustCompile("(?:(byr|iyr|eyr|hgt|hcl|ecl|pid|cid):([#a-z0-9]+)\\s?)")
	fmt.Printf("%s", re.ReplaceAllString(fixtures.Input, "$1"))

	// Now use a website like https://regex101.com/ and use
	//
	// ^(?=.*byr)(?=.*iyr)(?=.*eyr)(?=.*hgt)(?=.*hcl)(?=.*ecl)(?=.*pid)
	//
	// against the output.
	fmt.Printf("part 1 result: %d", 260)
}
