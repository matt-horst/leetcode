package p1925

func sqrt(v int) (int, bool) {
	l := 1
	r := 250 // From problem constraints

	for l <= r {
		m := l + (r-l)/2

		sqr := m * m
		if sqr == v {
			return m, true
		}

		if sqr > v {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return 0, false
}

func countTriples(n int) int {
	count := 0
	for a := 1; a <= n; a++ {
		for b := a + 1; b <= n; b++ {
			cSqr := a*a + b*b
			if cSqr > n*n {
				break
			}
			if _, ok := sqrt(cSqr); ok {
				count += 2
			}
		}
	}
	return count
}
