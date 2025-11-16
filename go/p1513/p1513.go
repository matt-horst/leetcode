package p1513

const modulo int = 1e9 + 7

func numSub(s string) int {
	total := 0

	cnt := 0
	for _, c := range s {
		if c == '0' {
			total = (total + (cnt+1)*(cnt)/2) % modulo
			cnt = 0
		} else {
			cnt++
		}
	}

	total = (total + (cnt+1)*(cnt)/2) % modulo

	return total
}
