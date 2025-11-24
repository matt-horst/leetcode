package p1018

func prefixesDivBy5(nums []int) []bool {
	ans := make([]bool, len(nums))
	
	v := 0
	for i, b := range nums {
		v += b
		ans[i] = v % 5 == 0
		v <<= 1
		v %= 5
	}

	return ans
}
