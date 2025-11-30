package p1590

import (
	"math"
)

func minSubarray(nums []int, p int) int {
	n := len(nums)

	preSum := 0
	for _, v := range nums {
		preSum += v
		preSum %= p
	}

	if preSum == 0 {
		return 0
	}

	hm := make(map[int]int, p)
	hm[0] = -1

	length := math.MaxInt
	rem := 0
	for i, v := range nums {
		rem += v
		rem %= p
		needed := (rem - preSum + p) % p
		if j, ok := hm[needed]; ok {
			length = min(length, i-j)
		}

		hm[rem] = i
	}

	if length >= n {
		return -1
	}

	return length
}
