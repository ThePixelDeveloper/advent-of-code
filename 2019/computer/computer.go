package computer

import (
	"strings"
)

func Run(input string) {
	instructions := strings.Split(input, ",")

	for i := 0; i < len(instructions); i++ {
		switch Opcode(instructions[i]) {
		case "02":
			i += Add(Args(instructions[i], i+1, i+2))
		}
	}

}

func Add(args string) int {
	
}

func Args(instruction string) string {
	var args string
 
	for i := len(instruction) - 3; i >= 0; i-- {
		args += string(instruction[i])
	}

	return args
}

func Opcode(instruction string) string {
	return instruction[len(instruction)-2:]
}
