package computer

import "errors"

func Execute(instructions []*Instruction) (int, error) {
	var (
		visited     map[int]struct{}
		instruction *Instruction
		acc         = 0
		i           = 0
	)

	visited = make(map[int]struct{})

	for {
		if i >= len(instructions) {
			return acc, errors.New("execution finished")
		}

		instruction = instructions[i]

		if _, ok := visited[i]; ok {
			return acc, errors.New("infinite loop detected")
		}

		visited[i] = struct{}{}

		switch instruction.key {
		case "nop":
			break

		case "acc":
			acc += instruction.value
			break

		case "jmp":
			i += instruction.value
			continue
		}

		i++
	}
}
