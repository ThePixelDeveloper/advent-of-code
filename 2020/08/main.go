package main

import (
	"fmt"
	"github.com/thepixeldeveloper/advent-of-code/2020/08/fixtures"
	"strconv"
	"strings"
)

func main() {
	accumulator := 0

	lines := strings.Split(fixtures.Input, "\n")

	currentLine := lines[0]
	currentIndex := 0
	direction := 1
	visited := make(map[int]struct{})

	for {
		switch true {
		case strings.HasPrefix(currentLine, "nop"):
			direction = 1
			break

		case strings.HasPrefix(currentLine, "acc"):
			acc, _ := strconv.Atoi(strings.Split(currentLine, " ")[1])
			accumulator += acc
			direction = 1
			break

		case strings.HasPrefix(currentLine, "jmp"):
			direction, _ = strconv.Atoi(strings.Split(currentLine, " ")[1])
			break
		}

		currentIndex += direction
		currentLine = lines[currentIndex]

		if _, ok := visited[currentIndex]; ok {
			fmt.Printf("part 1 result: %d", accumulator)
			break
		}

		visited[currentIndex] = struct{}{}
	}
}
