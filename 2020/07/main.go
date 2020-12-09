package main

import (
	"fmt"
	"github.com/thepixeldeveloper/advent-of-code/2020/07/data"
	"github.com/thepixeldeveloper/advent-of-code/2020/07/fixtures"
	"regexp"
	"strings"
)

func main() {
	container := parse(fixtures.Input)

	fmt.Printf("part 1 result: %d", container.Count("shiny gold"))
}

func parse(input string) data.Container {
	re := regexp.MustCompile("(?:^|(\\d) )(\\w+ \\w+) bags?")

	c := data.Container{}

	for _, line := range strings.Split(input, "\n") {
		matches := re.FindAllStringSubmatch(line, -1)

		// Container
		container := matches[0]
		c[container[2]] = []string{}

		// Bags
		for _, i := range matches[1:] {
			c[container[2]] = append(c[container[2]], i[2])
		}
	}

	return c
}
