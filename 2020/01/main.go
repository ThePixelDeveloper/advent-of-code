package main

import (
	"fmt"
	"github.com/thepixeldeveloper/advent-of-code/2020/01/fixtures"
)

func main() {
	for i := 0; i < len(fixtures.Input); i++ {
		for j := 0; j < len(fixtures.Input); j++ {
			if fixtures.Input[i]+fixtures.Input[j] == 2020 {
				fmt.Printf("result: %d", fixtures.Input[i]*fixtures.Input[j])
				return
			}
		}
	}
}
