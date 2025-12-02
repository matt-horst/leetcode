package p3623

const MOD int = 1e9 + 7

func nChoose2(n int) int {
	if n == 2 {
		return 1
	} else if n < 2 {
		return 0
	}

	return n * (n - 1) / 2
}

func countTrapezoids(points [][]int) int {
	hm := make(map[int]int)
	for _, p := range points {
		y := p[1]
		hm[y] += 1
	}

	total := 0
	others := 0
	for _, v := range hm {
		pairs := nChoose2(v) % MOD
		total += others * pairs
		total %= MOD

		others += pairs
		others %= MOD
	}

	return total
}
