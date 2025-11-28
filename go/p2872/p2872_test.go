package p2872

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		n        int
		edges    [][]int
		values   []int
		k        int
		expected int
	}{
		{
			n:        5,
			edges:    [][]int{{0, 2}, {1, 2}, {1, 3}, {2, 4}},
			values:   []int{1, 8, 1, 4, 4},
			k:        6,
			expected: 2,
		},
		{
			n:        7,
			edges:    [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}},
			values:   []int{3, 0, 6, 1, 5, 2, 1},
			k:        3,
			expected: 3,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := maxKDivisibleComponents(c.n, c.edges, c.values, c.k)
			assert.Equal(t, c.expected, result)
		})
	}
}
