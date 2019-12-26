package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/thepixeldeveloper/advent-of-code/2019/01/fuel"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() (err error) {
	file, err := os.Open("input")
	if err != nil {
		return err
	}

	defer func() {
		err = file.Close()
	}()

	scanner := bufio.NewScanner(file)

	var total int

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}

		total += fuel.FromMass(i)
	}

	fmt.Printf("output: %d", total)

	return scanner.Err()
}
