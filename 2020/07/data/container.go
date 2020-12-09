package data

type Container map[string][]string

func (c Container) Count(colour string) int {
	count := 0

	for container := range c {
		if c.Has(container, colour) {
			count++
		}
	}

	return count
}

func (c Container) Has(container string, colour string) bool {
	for _, bag := range c[container] {
		if bag == colour {
			return true
		}
	}

	for _, bag := range c[container] {
		if c.Has(bag, colour) {
			return true
		}
	}

	return false
}
