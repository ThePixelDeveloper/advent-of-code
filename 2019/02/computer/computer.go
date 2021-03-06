package computer

import (
	"github.com/thepixeldeveloper/advent-of-code/2019/02/opcode"
	"github.com/thepixeldeveloper/advent-of-code/2019/02/types"
	"github.com/thepixeldeveloper/advent-of-code/2019/05/computer"
)

func Calculate(input types.Input) string {
	instructions := input.SplitInt(",")

loop:
	for i := 0; i < len(instructions); {
		switch computer.MustOpcode(instructions[i]) {
		case opcode.ADD:
			i += add(i, instructions)
			break
		case opcode.MULTIPLY:
			i += multiply(i, instructions)
			break
		case opcode.TERMINATE:
			break loop
		}
	}

	return types.Output(instructions).JoinString(",")
}

func multiply(i int, instructions []int) int {
	left := computer.MustReadValue(instructions[i], instructions, 0)
	right :=  computer.MustReadValue(instructions[i], instructions, 1)
	
	output := instructions[i+3]

	instructions[output] = left * right

	return 4
}

func add(i int, instructions []int) int {
	left := instructions[instructions[i+1]]
	right := instructions[instructions[i+2]]
	output := instructions[i+3]

	instructions[output] = left + right

	return 4
}
