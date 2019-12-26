package fuel

import "testing"

func TestFromMass(t *testing.T) {
	tests := []struct {
		name string
		mass int
		fuel int
	}{
		{
			name: "example one",
			mass: 12,
			fuel: 2,
		},
		{
			name: "example two",
			mass: 14,
			fuel: 2,
		},
		{
			name: "example five",
			mass: 7,
			fuel: 0,
		},
		{
			name: "example six",
			mass: 100756,
			fuel: 50346,
		},
		{
			name: "negative fuel",
			mass: 2,
			fuel: 0,
		},
		{
			name: "recursive",
			mass: 1969,
			fuel: 966,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := FromMass(test.mass)

			if got != test.fuel {
				t.Fatalf("expected: %d, got: %d", test.fuel, got)
			}
		})
	}
}
