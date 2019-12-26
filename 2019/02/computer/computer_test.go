package computer

import (
	"testing"

	"github.com/thepixeldeveloper/advent-of-code/2019/02/types"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "example one",
			input:  "1,0,0,0,99",
			output: "2,0,0,0,99",
		},
		{
			name:   "example two",
			input:  "2,3,0,3,99",
			output: "2,3,0,6,99",
		},
		{
			name:   "example three",
			input:  "2,4,4,5,99,0",
			output: "2,4,4,5,99,9801",
		},
		{
			name:   "example four",
			input:  "1,1,1,4,99,5,6,0,99",
			output: "30,1,1,4,2,5,6,0,99",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Calculate(types.Input(test.input))

			if got != test.output {
				t.Fatalf("expected: %s, got: %s", test.output, got)
			}
		})
	}
}
