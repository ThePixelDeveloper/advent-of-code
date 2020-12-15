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
			if col == "L" {
				occupied := countVisible(grid, i, j, []string{"#"})
				if occupied == 0 {
					output[i] = append(output[i], "#")
					continue
				}
			}

			if col == "#" {
				occupied := countVisible(grid, i, j, []string{"#"})
				if occupied >= 5 {
					output[i] = append(output[i], "L")
					continue
				}
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

func countVisible(grid [][]string, row int, col int, search []string) int {
	var (
		count int
		steps = [][]int{
			{0, -1},  // North
			{1, -1},  // North East
			{1, 0},   // East
			{1, 1},   // South East
			{0, 1},   // South
			{-1, 1},  // South West
			{-1, 0},  // West
			{-1, -1}, // North West
		}
	)

	var (
		step = 0
		x    = col
		y    = row
	)

	for {
		if step >= len(steps) {
			break
		}

		x += steps[step][0]
		y += steps[step][1]

		if (x < 0 || y < 0) || (y >= len(grid) || x >= len(grid[y])) && step < len(steps) {
			step++
			x = col
			y = row
			continue
		}

		for _, s := range search {
			if grid[y][x] == "." {
				continue
			}

			if grid[y][x] == s {
				count++
			}

			step++
			x = col
			y = row
			continue
		}
	}

	return count
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
