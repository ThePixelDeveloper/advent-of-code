package circuit

import "github.com/thepixeldeveloper/advent-of-code/2019/03/wire"

type Circuit struct {
	wireOne map[string]wire.Location
	wireTwo map[string]wire.Location
}

func New() *Circuit {
	return &Circuit{
		wireOne: make(map[string]wire.Location),
		wireTwo: make(map[string]wire.Location),
	}
}

func (c *Circuit) AddWire(w wire.Wire) {
	if len(c.wireOne) == 0 {
		for _, s := range w.Steps() {
			c.wireOne[s.Hash()] = s
		}
		return
	}

	if len(c.wireTwo) == 0 {
		for _, s := range w.Steps() {
			c.wireTwo[s.Hash()] = s
		}
		return
	}
}

func (c *Circuit) Intersections() []wire.Location {
	var i []wire.Location

	for h, one := range c.wireOne {
		if two, ok := c.wireTwo[h]; ok {
			i = append(i, wire.Location{
				X:     one.X,
				Y:     one.Y,
				Steps: one.Steps + two.Steps,
			})
		}
	}

	return i
}

func (c *Circuit) Closest() (wire.Location, int) {
	zero := wire.Location{
		X: 0,
		Y: 0,
	}

	distance := func(p1, p2 *wire.Location) bool {
		return p1.ManhattanDistance(zero) < p2.ManhattanDistance(zero)
	}

	intersections := c.Intersections()

	wire.By(distance).Sort(intersections)

	return intersections[0], intersections[0].ManhattanDistance(zero)
}

func (c *Circuit) Fastest() wire.Location {
	steps := func(p1, p2 *wire.Location) bool {
		return p1.Steps < p2.Steps
	}

	intersections := c.Intersections()

	wire.By(steps).Sort(intersections)

	return intersections[0]
}
