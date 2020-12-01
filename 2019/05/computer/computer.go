package computer

import "strconv"

func MustOpcode(instruction int) int {
	s := strconv.Itoa(instruction)

	// Old opcode format
	if len(s) >= 1 && len(s) <= 2 {
		return instruction
	}

	// New computer format
	if len(s) > 2 {
		i, err := strconv.Atoi(s[len(s)-2:])
		if err != nil {
			panic(err)
		}

		return i
	}

	panic("unsupported opcode")
}

func MustReadValue(instruction int, computer [] int, index int) int {
	s := strconv.Itoa(instruction)

	// Old parameter mode (default)
	if len(s) >= 1 && len(s) <= 2 {
		return computer[computer[index]]
	}

	// New parameter mode (new)
	if len(s) > 2 {
		start := len(s) - (index + 3)
		finish := len(s) - (index + 2)
		
		var (
			i   = 0 // Default to the positional mode if the index isn't found.
			err error
		)

		if start >= 0 {
			i, err = strconv.Atoi(s[start:finish])
			if err != nil {
				panic(err)
			}
		}

		// Position mode
		if i == 0 {
			return computer[computer[index]]
		}

		// Immediate mode
		if i == 1 {
			return computer[index]
		}
	}

	panic("unsupported read")
}
