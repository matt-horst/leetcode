package p3512

func minOperations(nums []int, k int) int {
	rem := 0
	for _, v := range nums {
		rem += v
		rem %= k
	}

	return rem
}
