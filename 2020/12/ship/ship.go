package ship

import "strconv"

type Instruction struct {
	forward     int
	rotateLeft  int
	rotateRight int
	north       int
	east        int
	south       int
	west        int
}

func (i Instruction) execute(direction int, x int, y int) (int, int, int) {
	direction += i.rotateLeft + i.rotateRight

	if direction < 0 {
		direction += 360
	}

	direction %= 360

	y += i.north + i.south
	x += i.east + i.west

	switch direction {
	case 0:
		y += i.forward
	case 90:
		x += i.forward
	case 180:
		y -= i.forward
	case 270:
		x -= i.forward
	}

	return direction, x, y
}

func Simulate(instructions []Instruction, direction int, x int, y int) (int, int, int) {
	for _, i := range instructions {
		direction, x, y = i.execute(direction, x, y)
	}

	return direction, x, y
}

func Parse(input []string) []Instruction {
	var instructions []Instruction

	for _, s := range input {
		i := Instruction{}

		pace, _ := strconv.Atoi(s[1:])

		switch s[:1] {
		case "F":
			i.forward = pace
			break
		case "N":
			i.north = pace
			break
		case "E":
			i.east = pace
			break
		case "S":
			i.south = pace * -1
			break
		case "W":
			i.west = pace * -1
			break
		case "L":
			i.rotateLeft = pace * -1
			break
		case "R":
			i.rotateRight = pace
			break
		}

		instructions = append(instructions, i)
	}

	return instructions
}
