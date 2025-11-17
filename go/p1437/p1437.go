package p1437

func kLengthApart(nums []int, k int) bool {
	j := -1
	for i, v := range nums {
		if v == 1 {
			dist := i - j - 1
			if j >= 0 && dist < k {
				return false
			}

			j = i
		}
	}

	return true
}
