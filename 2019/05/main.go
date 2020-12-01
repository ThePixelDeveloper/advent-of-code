package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/thepixeldeveloper/advent-of-code/2019/02/computer"
	"github.com/thepixeldeveloper/advent-of-code/2019/02/types"
	types2 "github.com/thepixeldeveloper/advent-of-code/2019/05/types"
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

	for scanner.Scan() {
		output := computer.Calculate(types.Computer(scanner.Text()), types2.IntPtr(100))
		fmt.Printf("line 1 output: %s\n", output)
	}

	return scanner.Err()
}
