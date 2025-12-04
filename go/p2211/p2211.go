package p2211

func countCollisions(directions string) int {
	n := len(directions)
	total := 0

	prev := directions[0] != 'L'
	for i := 1; i < n; i++ {
		hit := false

		if directions[i] == 'L' && prev {
			total++
			hit = true
		}

		prev = hit || directions[i] != 'L'
	}

	prev = directions[n-1] != 'R'
	for i := n - 2; i >= 0; i-- {
		hit := false
		if directions[i] == 'R' && prev {
			total++
			hit = true
		}

		prev = hit || directions[i] != 'R'
	}
	return total
}
