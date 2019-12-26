package types

import (
	"strconv"
	"strings"
)

type Input string

// SplitInt allows us to split a string into a slice of integers.
func (i Input) SplitInt(delimiter string) []int {
	var instructions []int

	for _, s := range strings.Split(string(i), delimiter) {
		if i, err := strconv.Atoi(s); err == nil {
			instructions = append(instructions, i)
		}
	}

	return instructions
}

type Output []int

// JoinString allows us to join a slice of integers into a string, separated by the delimiter.
func (o Output) JoinString(delimiter string) string {
	var output []string

	for _, i := range o {
		output = append(output, strconv.Itoa(i))
	}

	return strings.Join(output, delimiter)
}
