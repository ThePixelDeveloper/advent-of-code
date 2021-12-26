package main

import (
	"github.com/ThePixelDeveloper/advent-of-code/2021/utils"
	"github.com/thoas/go-funk"
)

func partTwo() int {
	values := funk.Map(
		utils.SplitLines(file),
		utils.StringToInt,
	)

	var (
		window = 2
		i      = 0
		summed []float64
	)

	// Loop over every 3 values, sum them, and append them to a collection.
	for {
		v := values.([]int)

		if len(v[i:]) < 3 {
			break
		}

		summed = append(summed, funk.Sum(v[i:i+3]))

		i++
	}

	// Group values into pairs
	// 1,2,3,4 -> [1,2],[3,4]
	group1 := funk.Chunk(summed, window)

	// Group values into pairs offset by 1
	// 1,2,3,4 -> [2,3],[4]
	group2 := funk.Chunk(summed[1:], window)

	// Combined
	// [1,2],[3,4] + [2,3],[4] -> [[1,2],[2,3]], [[3,4][4]]
	combined := funk.Zip(group1, group2)

	// Flattened
	// [[1,2],[2,3]], [[3,4],[4]] -> [1,2],[2,3],[3,4],[4]
	flattened := funk.FlatMap(combined, func(x funk.Tuple) [][]float64 {
		var v [][]float64
		v = append(v, x.Element1.([]float64))
		v = append(v, x.Element2.([]float64))
		return v
	})

	// Filtered
	filtered := funk.Filter(flattened, func(x []float64) bool {
		if len(x) != 2 {
			return false
		}

		return x[1] > x[0]
	})

	return len(filtered.([][]float64))
}
