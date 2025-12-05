package p3432

func countPartitions(nums []int) int {
	count := 0

	total := 0
	for _, v := range nums {
		total += v
	}

	sum := 0
	for _, v := range nums[:len(nums)-1] {
		sum += v
		diff := total - 2 * sum
		if diff % 2 == 0{
			count++
		}
	}


	return count
}
