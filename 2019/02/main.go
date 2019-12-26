package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/thepixeldeveloper/advent-of-code/2019/02/computer"
	"github.com/thepixeldeveloper/advent-of-code/2019/02/types"
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

	var line int
	for scanner.Scan() {
		line++

		if line == 1 {
			output := computer.Calculate(types.Input(scanner.Text()))
			fmt.Printf("line 1 output: %s\n", output)
		}

		if line == 2 {
			for noun := 0; noun < 99; noun++ {
				for verb := 0; verb < 99; verb++ {
					output := computer.Calculate(types.Input(fmt.Sprintf(scanner.Text(), noun, verb)))

					if strings.HasPrefix(output, "19690720") {
						fmt.Printf("line 2 output: %d (verb=%d noun=%d)\n", 100*noun+verb, verb, noun)
					}
				}
			}
		}
	}

	return scanner.Err()
}
