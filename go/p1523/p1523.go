package p1523

func countOdds(low, high int) int {
	count := (high - low) / 2

	if low % 2 != 0 || high % 2 != 0 {
		count += 1
	}
	
	return count
} 

