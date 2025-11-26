package p2435

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		grid     [][]int
		k        int
		expected int
	}{
		{
			grid: [][]int{{5, 2, 4}, {3, 0, 5}, {0, 7, 2}},
			k: 3,
			expected: 2,
		},
		{
			grid: [][]int{{0, 0}},
			k: 5,
			expected: 1,
		},
		{
			grid: [][]int{{7,3,4,9},{2,3,6,2},{2,3,7,0}},
			k: 1,
			expected: 10,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := numberOfPaths(c.grid, c.k)
			assert.Equal(t, c.expected, result)
		})
	}
}
