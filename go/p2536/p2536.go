package p2536

func rangeAddQueries(n int, queries [][]int) [][]int {
	s := make([][]int, n)
	for i := range n {
		s[i] = make([]int, n)
	}

	for _, q := range queries {
		for i := q[0]; i <= q[2]; i++ {
			s[i][q[1]]++

			if q[3] + 1 < n {
				s[i][q[3] + 1]--
			}
		}
	}

	for i := range n {
		for j := 1; j < n; j++ {
			s[i][j] += s[i][j - 1]
		}
	}

	return s
}
