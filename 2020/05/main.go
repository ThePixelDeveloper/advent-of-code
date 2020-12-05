package main

import (
	"fmt"
	"github.com/thepixeldeveloper/advent-of-code/2020/05/fixtures"
)

func main() {
	passes := fixtures.Input

	var (
		occupied = map[int]bool{}
		empty    = []int{}
	)

	for _, pass := range passes {
		occupied[seatIdentifier(pass)] = true
	}

	for i := 0; i < 842; i++ {
		if _, ok := occupied[i]; ok {
			continue
		}

		_, previous := occupied[i-1]
		_, next := occupied[i+1]

		if previous && next {
			empty = append(empty, i)
		}
	}

	fmt.Printf("part 2 result: %+v", empty[0])
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
