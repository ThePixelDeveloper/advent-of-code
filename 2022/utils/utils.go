package utils

import (
	"bufio"
	"strconv"
	"strings"
)

func StringToInt(s string, _ int) int {
	atoi, _ := strconv.Atoi(s)
	return atoi
}

func SplitLines(s string) []string {
	var lines []string

	sc := bufio.NewScanner(strings.NewReader(s))

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines
}
