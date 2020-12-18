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

// * 4 * * *
// * * * * 1
// * * 0 * *
// 3 * * * *
// * * * 2 *
func (i Instruction) execute(direction int, x int, y int, shipX int, shipY int) (int, int, int, int, int) {
	rotateLeft := i.rotateLeft
	rotateRight := i.rotateRight

	y += i.north + i.south
	x += i.east + i.west

	shipX += i.forward * x
	shipY += i.forward * y

	for {
		if rotateLeft > 0 {
			x, y = -y, x
			rotateLeft -= 90
		}

		if rotateRight > 0 {
			x, y = y, -x
			rotateRight -= 90
		}

		if rotateLeft == 0 && rotateRight == 0 {
			break
		}
	}

	return direction, x, y, shipX, shipY
}

func Simulate(instructions []Instruction, direction int, x int, y int, shipX int, shipY int) (int, int, int, int, int) {
	for _, i := range instructions {
		direction, x, y, shipX, shipY = i.execute(direction, x, y, shipX, shipY)
	}

	return direction, x, y, shipX, shipY
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
			i.rotateLeft = pace
			break
		case "R":
			i.rotateRight = pace
			break
		}

		instructions = append(instructions, i)
	}

	return instructions
}
