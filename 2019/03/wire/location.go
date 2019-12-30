package wire

import (
	"fmt"
	"sort"
)

type Location struct {
	X     int
	Y     int
	Steps int
}

type By func(p1, p2 *Location) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(locations []Location) {
	ps := &locationSorter{
		locations: locations,
		by:        by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

type locationSorter struct {
	locations []Location
	by        func(p1, p2 *Location) bool
}

// Len is part of sort.Interface.
func (s *locationSorter) Len() int {
	return len(s.locations)
}

// Swap is part of sort.Interface.
func (s *locationSorter) Swap(i, j int) {
	s.locations[i], s.locations[j] = s.locations[j], s.locations[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *locationSorter) Less(i, j int) bool {
	return s.by(&s.locations[i], &s.locations[j])
}

func (l Location) ManhattanDistance(to Location) int {
	x := l.X - to.X
	y := l.Y - to.Y

	if x < 0 {
		x = -x
	}

	if y < 0 {
		y = -y
	}

	return x + y
}

func (l Location) Hash() string {
	return fmt.Sprintf("%d%d", l.X, l.Y)
}
