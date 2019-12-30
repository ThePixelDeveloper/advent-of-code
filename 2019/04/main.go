package main

import (
	"fmt"

	"github.com/thepixeldeveloper/advent-of-code/2019/04/password"
)

func main() {
	if err := run(347312, 805915); err != nil {
		panic(err)
	}
}

func run(start int, finish int) (err error) {
	validPt1 := 0
	validPt2 := 0

	for i := start; i < finish; i++ {
		if password.ValidPt1(i) {
			validPt1++
		}

		if password.ValidPt2(i) {
			validPt2++
		}
	}

	fmt.Printf("output: part_one=%d\n", validPt1)
	fmt.Printf("output: part_two=%d", validPt2)

	return nil
}
