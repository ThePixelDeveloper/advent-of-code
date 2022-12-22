package main

import (
	_ "embed"
	"fmt"
)

//go:embed _input
var file string

func main() {
	fmt.Printf("answer part 1: %+v\n", partOne())
	fmt.Printf("answer part 2: %+v\n", partTwo())
}
