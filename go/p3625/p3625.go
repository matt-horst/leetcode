package p3625

type Fraction struct {
	Num   int
	Denom int
}

func NewFraction(num, denom int) Fraction {
	if denom == 0 {
		return Fraction{Num: 1, Denom: 0}
	}

	if num == 0 {
		return Fraction{Num: 0, Denom: 1}
	}

	sign := 1
	if (num < 0) != (denom < 0) {
		sign = -1
	}

	gcd := gcd(num, denom)

	return Fraction{
		Num:   sign * abs(num/gcd),
		Denom: abs(denom / gcd),
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

type Point struct {
	x int
	y int
}

type PointFraction struct {
	x Fraction
	y Fraction
}

func calcSlope(a, b Point) Fraction {
	dx := a.x - b.x
	dy := a.y - b.y

	return NewFraction(dy, dx)
}

func calcIntercept(a, b Point) Fraction {
	dx := a.x - b.x
	dy := a.y - b.y

	if dx == 0 {
		// return float64(a.x)
		return Fraction{Num: a.x, Denom: 1}
	}

	// return float64(a.y*dx - a.x*dy) / float64(dx)
	return NewFraction(a.y*dx-a.x*dy, dx)
}

func nChoose2(n int) int {
	if n == 2 {
		return 1
	} else if n < 2 {
		return 0
	}

	return n * (n - 1) / 2
}

func countTrapezoids(points [][]int) int {
	n := len(points)

	hm := make(map[Fraction]map[Fraction]map[Point]struct{})
	midPointToSlope := make(map[PointFraction]map[Fraction]int)

	for i := 0; i < n-1; i++ {
		a := Point{x: points[i][0], y: points[i][1]}

		for j := i + 1; j < n; j++ {
			b := Point{x: points[j][0], y: points[j][1]}

			slope := calcSlope(a, b)
			midPoint := PointFraction{x: NewFraction(a.x+b.x, 2), y: NewFraction(a.y+b.y, 2)}
			if _, ok := midPointToSlope[midPoint]; !ok {
				midPointToSlope[midPoint] = make(map[Fraction]int)
			}
			midPointToSlope[midPoint][slope] += 1
			intercept := calcIntercept(a, b)

			if _, ok := hm[slope]; !ok {
				hm[slope] = make(map[Fraction]map[Point]struct{})
			}

			if _, ok := hm[slope][intercept]; !ok {
				hm[slope][intercept] = make(map[Point]struct{})
			}

			hm[slope][intercept][a] = struct{}{}
			hm[slope][intercept][b] = struct{}{}
		}
	}

	total := 0
	for _, ps := range hm {
		others := 0
		for _, vs := range ps {
			pairs := nChoose2(len(vs))
			total += others * pairs
			others += pairs
		}
	}

	for _, slopes := range midPointToSlope {
		others := 0
		for _, cnt := range slopes {
			total -= others * cnt
			others += cnt
		}
	}

	return total
}
