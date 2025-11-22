package p3190

func minimumOperations(nums []int) int {
	total := 0

	for _, v := range nums {
		if v % 3 != 0 {
			total++
		}
	}

	return total
}
