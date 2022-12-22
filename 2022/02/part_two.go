package main

import (
	"github.com/samber/lo"
	"strings"
)

func partTwo() int {
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
	scores["A X"] = SCISSORS + LOSE // Lose
	scores["A Y"] = ROCK + DRAW     // Draw
	scores["A Z"] = PAPER + WIN     // Win

	// Paper
	scores["B X"] = ROCK + LOSE    // Lose
	scores["B Y"] = PAPER + DRAW   // Draw
	scores["B Z"] = SCISSORS + WIN // Win

	// Scissors
	scores["C X"] = PAPER + LOSE    // Lose
	scores["C Y"] = SCISSORS + DRAW // Draw
	scores["C Z"] = ROCK + WIN      // Win

	return lo.Sum(
		lo.Map[string, int](strings.Split(file, "\n"), func(item string, _ int) int {
			return scores[item]
		}),
	)
}
