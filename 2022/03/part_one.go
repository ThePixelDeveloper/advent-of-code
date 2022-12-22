package main

import (
	"github.com/samber/lo"
	"strings"
)

func partOne() int {
	a := make(map[rune]int, 52)

	for i, l := range []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		a[l] = i + 1
	}

	return lo.Sum(
		lo.Map[[]rune, int](
			lo.Map[[]string, []rune](
				lo.Map[string, []string](strings.Split(file, "\n"),
					func(item string, index int) []string {
						return lo.ChunkString(item, len(item)/2)
					},
				),
				func(item []string, index int) []rune {
					return lo.Uniq[rune](
						lo.Intersect[rune]([]rune(item[0]), []rune(item[1])),
					)
				}),
			func(item []rune, index int) int {
				return a[item[0]]
			}),
	)
}
