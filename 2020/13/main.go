package main

import (
	"fmt"
	"github.com/thepixeldeveloper/advent-of-code/2020/13/fixtures"
	"strconv"
	"strings"
)

type Estimate struct {
	DepatureTime int
	Buses        []int
}

func (e Estimate) ClosestBus() (int, int) {
	var (
		i = e.DepatureTime
	)

	for {
		for _, bus := range e.Buses {
			if i%bus == 0 {
				return bus, i
			}
		}

		i++
	}
}

func main() {
	estimate := parse(fixtures.Input)

	bus, i := estimate.ClosestBus()

	fmt.Printf("part 1 result: %d", (i-estimate.DepatureTime)*bus)
}

func parse(input string) *Estimate {
	lines := strings.Split(input, "\n")

	var (
		depatureTimeInput = lines[0]
		busesInput        = lines[1]

		depatureTime int
		bus          int
		buses        []int
	)

	depatureTime, _ = strconv.Atoi(depatureTimeInput)

	for _, s := range strings.Split(busesInput, ",") {
		if s != "x" {
			bus, _ = strconv.Atoi(s)
			buses = append(buses, bus)
		}
	}

	return &Estimate{
		DepatureTime: depatureTime,
		Buses:        buses,
	}
}
