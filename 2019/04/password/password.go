package password

import (
	"strconv"
)

func ValidPt1(password int) bool {
	p := strconv.Itoa(password)

	double := false

	for i := range p {
		if i > 0 && p[i] < p[i-1] {
			return false
		}

		if i > 0 && i < len(p) && !double {
			double = p[i] == p[i-1]
		}
	}

	return double
}

func ValidPt2(password int) bool {
	p := strconv.Itoa(password)

	buckets := make(map[uint8]int, 10)

	for i := range p {
		if i > 0 && p[i] < p[i-1] {
			return false
		}

		buckets[p[i]]++
	}

	for _, c := range buckets {
		if c == 2 {
			return true
		}
	}

	return false
}
