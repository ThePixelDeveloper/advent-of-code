package main

import (
	"fmt"
	"github.com/thepixeldeveloper/advent-of-code/2020/11/fixtures"
	"strings"
)

func main() {
	grid := cols(rows(fixtures.Input))

	previousFrame := grid

	for {
		frame := simulate(previousFrame)

		if fmt.Sprintf("%+v", frame) == fmt.Sprintf("%+v", previousFrame) {
			break
		}

		previousFrame = frame
	}

	fmt.Printf("part 1 result: %d", count(previousFrame, []string{"#"}))
}

func simulate(grid [][]string) [][]string {
	var output [][]string

	for i, row := range grid {
		output = append(output, []string{})

		for j, col := range row {
			empty, emptyMax := countAdjacent(grid, i, j, []string{".", "L"})

			if col == "L" && empty == emptyMax {
				output[i] = append(output[i], "#")
				continue
			}

			occupied, _ := countAdjacent(grid, i, j, []string{"#"})

			if col == "#" && occupied >= 4 {
				output[i] = append(output[i], "L")
				continue
			}

			output[i] = append(output[i], col)
		}
	}

	return output
}

func count(grid [][]string, search []string) int {
	var count int

	for i, _ := range grid {
		for j, _ := range grid[i] {
			for _, s := range search {
				if grid[i][j] == s {
					count++
				}
			}
		}
	}

	return count
}

func countAdjacent(grid [][]string, row int, col int, search []string) (int, int) {
	var (
		count     int
		available int
	)

	for i := row - 1; i <= row+1; i++ {
		if i < 0 || i >= len(grid) {
			continue
		}

		for j := col - 1; j <= col+1; j++ {
			if i == row && j == col {
				continue
			}

			if j < 0 || j >= len(grid[i]) {
				continue
			}

			available++

			for _, s := range search {
				if grid[i][j] == s {
					count++
					break
				}
			}
		}
	}

	return count, available
}

func cols(input []string) [][]string {
	var output [][]string

	for i, s := range input {
		output = append(output, []string{})

		for _, j := range s {
			output[i] = append(output[i], string(j))
		}
	}

	return output
}

func rows(input string) []string {
	return strings.Split(input, "\n")
}
