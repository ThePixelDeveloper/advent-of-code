package fuel

func FromMass(mass int) int {
	fuel := mass/3 - 2

	if fuel < 0 {
		return 0
	}

	return fuel + FromMass(fuel)
}
