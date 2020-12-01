package main

import (
	"fmt"
	"github.com/thepixeldeveloper/advent-of-code/2020/01/fixtures"
)

func main() {
	for i := 0; i < len(fixtures.Input); i++ {
		for j := 0; j < len(fixtures.Input); j++ {
			if fixtures.Input[i]+fixtures.Input[j] == 2020 {
				fmt.Printf("part 1 result: %d\n", fixtures.Input[i]*fixtures.Input[j])
			}
			
			for k := 0; k < len(fixtures.Input); k++ {
				if fixtures.Input[i]+fixtures.Input[j]+fixtures.Input[k] == 2020 {
					fmt.Printf("part 2 result: %d\n", fixtures.Input[i]*fixtures.Input[j]*fixtures.Input[k])
				}	
			}
		}
	}
}
