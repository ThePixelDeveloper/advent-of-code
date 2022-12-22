package main

import (
	"github.com/ThePixelDeveloper/advent-of-code/2022/utils"
	"github.com/samber/lo"
	"sort"
	"strings"
)

func partTwo() int {
	rs := lo.Map[string, int](strings.Split(file, "\n\n"), func(item string, _ int) int {
		return lo.Sum(
			lo.Map[string, int](strings.Split(item, "\n"), utils.StringToInt))
	})

	sort.Slice(rs, func(i, j int) bool {
		return rs[i] > rs[j]
	})

	return rs[0] + rs[1] + rs[2]
}
