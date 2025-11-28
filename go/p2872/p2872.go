package p2872

type component struct {
	root int
	sum  int
}

func dfs(adjList map[int][]int, seen map[int]struct{}, values []int, k int, par, node int, components []*component, total int) int {
	seen[node] = struct{}{}

	val := values[node]
	components[node] = &component{root: node, sum: val}

	for _, next := range adjList[node] {
		if _, ok := seen[next]; !ok {
			total = dfs(adjList, seen, values, k, node, next, components, total)
		}
	}

	if components[node].sum%k != 0 && par >= 0 {
		components[par].sum += components[node].sum
		components[node] = components[par]

		total--
	}

	return total
}

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	adjList := make(map[int][]int, n)
	for _, e := range edges {
		adjList[e[0]] = append(adjList[e[0]], e[1])
		adjList[e[1]] = append(adjList[e[1]], e[0])
	}

	components := make([]*component, n)

	seen := make(map[int]struct{}, n)
	total := dfs(adjList, seen, values, k, -1, 0, components, n)

	return total
}
