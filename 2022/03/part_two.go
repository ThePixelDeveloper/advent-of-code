package main

import (
	"github.com/samber/lo"
	"strings"
)

func partTwo() int {
	a := make(map[rune]int, 52)

	for i, l := range []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		a[l] = i + 1
	}

	return lo.Sum(lo.Map[[]rune, int](
		lo.Map[[]string, []rune](
			lo.Chunk[string](strings.Split(file, "\n"), 3), func(item []string, index int) []rune {
				return lo.Uniq[rune](
					lo.Intersect[rune](
						lo.Intersect[rune](
							[]rune(item[0]),
							[]rune(item[1]),
						),
						[]rune(item[2]),
					),
				)
			}), func(item []rune, index int) int {
			return a[item[0]]
		},
	))
}
