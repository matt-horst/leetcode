package p757

import (
	"slices"
)

func intersectionSizeTwo(intervals [][]int) int {
	// Sort the intervals by ending point
	slices.SortFunc(intervals, func(a, b []int) int { return a[1] - b[1] })

	total := 2

	a := intervals[0][1] // current highest number in set
	b := a - 1           // current second highest number in set
	for i := 1; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]

		if start > a {
			total += 2
			a = end
			b = a - 1
		} else if start > b {
			total += 1
			if a != end {
				b = a
				a = end
			} else {
				a = end
				b = a - 1
			}
		}
	}

	return total
}
