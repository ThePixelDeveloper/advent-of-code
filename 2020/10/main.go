package main

import (
	"fmt"
	"github.com/thepixeldeveloper/advent-of-code/2020/10/fixtures"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(fixtures.Input, "\n")

	var (
		voltages         map[string]struct{}
		maxVoltage       int
		maxVoltageBuffer = 3
	)

	voltages = make(map[string]struct{}, 0)

	for i := 0; i < len(lines); i++ {
		voltage := parseVoltage(lines[i])
		voltages[lines[i]] = struct{}{}

		if voltage > maxVoltage {
			maxVoltage = voltage
		}
	}

	maxVoltage += maxVoltageBuffer

	voltages[strconv.Itoa(maxVoltage)] = struct{}{}

	fmt.Printf("part 1 result: %+v", doStuff(voltages, 0))
}

func doStuff(voltages map[string]struct{}, start int) int {
	var (
		i           = start
		difference  int
		differences map[int]int
	)

	differences = make(map[int]int)

	for {
		_, ok := voltages[strconv.Itoa(i)]

		if !ok && difference > 3 {
			break
		}

		if ok {
			differences[difference]++
			difference = 0
		}

		difference++
		i++
	}

	return differences[1] * differences[3]
}

func parseVoltage(input string) int {
	voltage, _ := strconv.Atoi(input)
	return voltage
}
