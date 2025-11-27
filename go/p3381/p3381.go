package p3381

import (
	"math"
)

func maxSubarraySum(nums []int, k int) int64 {
	prefix := make(map[int]int64, k)

	ans := int64(math.MinInt64)

	tot := int64(0)
	prefix[0] = 0
	for i, num := range nums {
		tot += int64(num)

		if pre, ok := prefix[(i+1)%k]; ok {
			ans = max(ans, tot-pre)
		}

		if pre, ok := prefix[(i+1)%k]; ok {
			prefix[(i+1)%k] = min(pre, tot)
		} else {
			prefix[(i+1)%k] = tot
		}
	}

	return ans
}
