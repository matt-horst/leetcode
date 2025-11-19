package p2154

func findFinalValue(nums []int, original int) int {
	hm := make(map[int]any, len(nums))

	for _, v := range nums {
		hm[v] = struct {}{}
	}

	for {
		if _, ok := hm[original]; ok {
			original *= 2
		} else {
			return original
		}
	}
}
