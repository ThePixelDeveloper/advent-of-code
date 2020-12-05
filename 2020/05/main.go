package main

import (
	"fmt"
	"github.com/thepixeldeveloper/advent-of-code/2020/05/fixtures"
)

func main() {
	passes := fixtures.Input

	var (
		highest = 0
		id      = 0
	)

	for _, pass := range passes {
		id = seatIdentifier(pass)

		if id > highest {
			highest = id
		}
	}

	fmt.Printf("part 1 result: %d", highest)
}

func seatIdentifier(boardingPass []string) int {
	var (
		seatRowLower = 0
		seatRowMid   = 0
		seatRowUpper = 128

		seatColumnLower = 0
		seatColumnMid   = 0
		seatColumnUpper = 8
	)

	for _, pass := range boardingPass {
		seatRowMid = seatRowLower + (seatRowUpper-seatRowLower)/2

		if pass == "F" {
			seatRowUpper = seatRowMid
		}

		if pass == "B" {
			seatRowLower = seatRowMid
		}

		if pass == "L" || pass == "R" {
			seatColumnMid = seatColumnLower + (seatColumnUpper-seatColumnLower)/2

			if pass == "R" {
				seatColumnLower = seatColumnMid
			}

			if pass == "L" {
				seatColumnUpper = seatColumnMid
			}
		}
	}

	return seatRowLower*8 + seatColumnLower
}
