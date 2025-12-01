package p2141

import (
	"math"
)

func maxRunTime(n int, batteries []int) int64 {
	total := int64(0)
	for _, v := range batteries {
		if int64(v) > math.MaxInt64-total {
			total = math.MaxInt64
			break
		}

		total += int64(v)
	}

	l := int64(0)
	r := total/int64(n) + 1

	for l <= r {
		m := (r-l)/2 + l

		surplus := int64(0)
		for _, b := range batteries {
			surplus += min(m, int64(b))
		}

		if surplus/int64(n) >= m {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return l - 1
}
