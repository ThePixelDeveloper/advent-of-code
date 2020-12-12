package computer

import (
	"strconv"
	"strings"
)

type Instruction struct {
	key   string
	value int
}

func (i Instruction) IsJmp() bool {
	return i.key == "jmp"
}

func (i Instruction) IsNop() bool {
	return i.key == "nop"
}

func (i Instruction) ToJmp() *Instruction {
	return &Instruction{
		key:   "jmp",
		value: i.value,
	}
}

func (i *Instruction) ToNop() *Instruction {
	return &Instruction{
		key:   "nop",
		value: i.value,
	}
}

func NewInstruction(input string) *Instruction {
	split := strings.Split(input, " ")
	value, _ := strconv.Atoi(split[1])

	return &Instruction{
		key:   split[0],
		value: value,
	}
}
