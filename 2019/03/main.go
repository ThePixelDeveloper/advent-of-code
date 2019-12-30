package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/thepixeldeveloper/advent-of-code/2019/03/circuit"
	"github.com/thepixeldeveloper/advent-of-code/2019/03/wire"
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

	c := circuit.New()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		c.AddWire(wire.Wire(scanner.Text()))
	}

	closest, distance := c.Closest()
	fastest := c.Fastest()

	fmt.Printf("output (closest): x=%d y=%d distance=%d\n", closest.X, closest.Y, distance)
	fmt.Printf("output (fastest): x=%d y=%d steps=%d", fastest.X, fastest.Y, fastest.Steps)

	return scanner.Err()
}
