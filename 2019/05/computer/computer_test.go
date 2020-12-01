package computer

import (
	"testing"

	"github.com/thepixeldeveloper/advent-of-code/2019/02/opcode"
)

func TestMustOpcode(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		opcode int
	}{
		{
			name:   "new format (add)",
			input:  1001,
			opcode: opcode.ADD,
		},
		{
			name:   "new format (multiply)",
			input:  1002,
			opcode: opcode.MULTIPLY,
		},
		{
			name:   "add",
			input:  1,
			opcode: opcode.ADD,
		},
		{
			name:   "multiply",
			input:  2,
			opcode: opcode.MULTIPLY,
		},
		{
			name:   "input",
			input:  3,
			opcode: opcode.INPUT,
		},
		{
			name:   "output",
			input:  4,
			opcode: opcode.OUTPUT,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opcode := MustOpcode(test.input)

			if opcode != test.opcode {
				t.Fatalf("expected: %d, got: %d", test.opcode, opcode)
			}
		})
	}
}

func TestMustReadValue(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		computer []int
		index    int
		output   int
	}{
		{
			name:     "positional mode",
			input:    1002,
			computer: []int{1, 2, 3, 4},
			index:    0,
			output:   2,
		},
		{
			name:     "immediate mode",
			input:    1102,
			computer: []int{1, 2, 3, 4},
			index:    0,
			output:   1,
		},
		{
			name:     "default option (positional)",
			input:    1102,
			computer: []int{1, 2, 3, 4},
			index:    2,
			output:   4,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value := MustReadValue(test.input, test.computer, test.index)

			if value != test.output {
				t.Fatalf("expected: %d, got: %d", test.output, value)
			}
		})
	}
}
