package main

import (
	"fmt"
	"github.com/thepixeldeveloper/advent-of-code/2020/08/computer"
	"github.com/thepixeldeveloper/advent-of-code/2020/08/fixtures"
	"strings"
)

func main() {
	lines := strings.Split(fixtures.Input, "\n")

	var (
		instructions []*computer.Instruction
		permutations [][]*computer.Instruction
	)

	for _, line := range lines {
		instructions = append(instructions, computer.NewInstruction(line))
	}

	permutations = make([][]*computer.Instruction, len(instructions))

	for i := 0; i < len(instructions); i++ {
		for j := 0; j < len(instructions); j++ {
			if i == j {
				if instructions[i].IsJmp() {
					permutations[i] = append(permutations[i], instructions[i].ToNop())
					continue
				}

				if instructions[i].IsNop() {
					permutations[i] = append(permutations[i], instructions[i].ToJmp())
					continue
				}
			}

			permutations[i] = append(permutations[i], instructions[j])
		}
	}

	for _, instructions := range permutations {
		execute, err := computer.Execute(instructions)
		if err != nil && err.Error() == "execution finished" {
			fmt.Printf("part 2 result: %d - %s\n", execute, err)
		}
	}
}
