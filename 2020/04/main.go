package main

import (
	"fmt"
	"github.com/thepixeldeveloper/advent-of-code/2020/04/fixtures"
	"regexp"
)

func main() {
	// Formatting to a single line
	re := regexp.MustCompile("(?:(byr|iyr|eyr|hgt|hcl|ecl|pid|cid):([#a-z0-9]+)\\s?)")
	fmt.Printf("%s", re.ReplaceAllString(fixtures.Input, "$2"))

	// Now use a website like https://regex101.com/ and use
	//
	// ^(?=.*byr)(?=.*iyr)(?=.*eyr)(?=.*hgt)(?=.*hcl)(?=.*ecl)(?=.*pid)
	//
	// against the output.
	fmt.Printf("part 1 result: %d", 260)

	// Part two (use a website like https://regex101.com/)
	//
	// ^(?=.*byr:(?:19[2-9][0-9]|200[0-2])\s)(?=.*iyr:(?:20(?:1[0-9]|20))\s)(?=.*eyr:(?:20(?:2[0-9]|30))\s)(?=.*hgt:(?:1(?:[5-8][0-9]|9[0-3])cm|(?:59|6[0-9]|7[0-6])in)\s)(?=.*hcl:#[a-f0-9]{6}\s)(?=.*ecl:(?:(amb|blu|brn|gry|grn|hzl|oth))\s)(?=.*pid:[0-9]{9}\s)
	//
	// Birth year match
	//  ^(?=.*byr:(?:19[2-9][0-9]|200[0-2])\s)
	//
	// Issue year match
	//  ^(?=.*iyr:(?:20(?:1[0-9]|20))\s)
	//
	// Expiration year
	//  ^(?=.*eyr:(?:20(?:2[0-9]|30))\s)
	//
	// Height
	//  ^(?=.*hgt:(?:1(?:[5-8][0-9]|9[0-3])cm|(?:59|6[0-9]|7[0-6])in)\s)
	//
	// Hair colour
	//  ^(?=.*hcl:#[a-f0-9]{6}\s)
	//
	// Eye colour
	//  ^(?=.*ecl:(?:(amb|blu|brn|gry|grn|hzl|oth))\s)
	//
	// Passport ID
	//  ^(?=.*pid:[0-9]{9}\s)
	//
	// against the output.
	fmt.Printf("part 2 result: %d", 153)
}
