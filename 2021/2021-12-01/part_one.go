package main

import (
	"github.com/ThePixelDeveloper/advent-of-code/2021/utils"
	"github.com/thoas/go-funk"
)

func partOne() int {
	values := funk.Map(
		utils.SplitLines(file),
		utils.StringToInt,
	)

	var window = 2

	// Group values into pairs
	// 1,2,3,4 -> [1,2],[3,4]
	group1 := funk.Chunk(values, window)

	// Group values into pairs offset by 1
	// 1,2,3,4 -> [2,3],[4]
	group2 := funk.Chunk(values.([]int)[1:], window)

	// Combined
	// [1,2],[3,4] + [2,3],[4] -> [[1,2],[2,3]], [[3,4][4]]
	combined := funk.Zip(group1, group2)

	// Flattened
	// [[1,2],[2,3]], [[3,4],[4]] -> [1,2],[2,3],[3,4],[4]
	flattened := funk.FlatMap(combined, func(x funk.Tuple) [][]int {
		var v [][]int
		v = append(v, x.Element1.([]int))
		v = append(v, x.Element2.([]int))
		return v
	})

	// Filtered
	filtered := funk.Filter(flattened, func(x []int) bool {
		if len(x) != 2 {
			return false
		}

		return x[1] > x[0]
	})

	return len(filtered.([][]int))
}
