package main

import (
	"github.com/samber/lo"
	"strings"
)

func partOne() int {
	const (
		WIN      = 6
		DRAW     = 3
		LOSE     = 0
		ROCK     = 1
		PAPER    = 2
		SCISSORS = 3
	)

	scores := make(map[string]int)
	// Rock
	scores["A X"] = ROCK + DRAW
	scores["A Y"] = PAPER + WIN
	scores["A Z"] = SCISSORS + LOSE

	// Paper
	scores["B X"] = ROCK + LOSE
	scores["B Y"] = PAPER + DRAW
	scores["B Z"] = SCISSORS + WIN

	// Scissors
	scores["C X"] = ROCK + WIN
	scores["C Y"] = PAPER + LOSE
	scores["C Z"] = SCISSORS + DRAW

	return lo.Sum(
		lo.Map[string, int](strings.Split(file, "\n"), func(item string, _ int) int {
			return scores[item]
		}),
	)
}
