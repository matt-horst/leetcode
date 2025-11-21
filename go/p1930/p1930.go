package p1930

func countPalindromicSubsequence(s string) int {
	n := len(s)
	v := []byte(s)

	// Build hashmap that indicates the right most instance of each char
	right := make(map[byte]int)

	for i := n-1; i >= 0; i-- {
		c := v[i]
		if _, ok := right[c]; !ok {
			right[c] = i
		}
	}

	// Build hashmap that indicates the next instance of each char
	next := make(map[byte][]int)

	for k := range right {
		next[k] = make([]int, n)
		next[k][n-1] = n

		curr := n
		for i := n-2; i >= 0; i-- {
			c := v[i]

			next[k][i] = curr

			if c == k {
				curr = i
			}
		}
	}

	set := make(map[string]any)

	for l, c := range v {
		if r, ok := right[c]; ok && r > l {
			for k := range next {
				if m, ok := next[k]; ok && m[l] > l && m[l] < r {
					set[string([]byte{c, k})] = struct{}{}
				}
			}
		}
	}

	return len(set)
}
