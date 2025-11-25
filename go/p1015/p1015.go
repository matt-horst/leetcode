package p1015

func smallestRepunitDivByK(k int) int {
	set := make(map[int]any)
	v := 1
	cnt := 1
	for {
		if _, ok := set[v]; ok {
			break
		}
		set[v] = struct{}{}

		if v % k == 0 {
			return cnt
		}

		cnt++
		v *= 10
		v += 1
		v %= k
	}

	return -1
}
