package p2536

import (
	"fmt"
	"testing"
)

func TestRangeAddQueries(t *testing.T) {
	cases := []struct {
		n        int
		queries  [][]int
		expected [][]int
	}{
		{
			n:        3,
			queries:  [][]int{{1, 1, 2, 2}, {0, 0, 1, 1}},
			expected: [][]int{{1, 1, 0}, {1, 2, 1}, {0, 1, 1}},
		},
		{
			n:        2,
			queries:  [][]int{{0, 0, 1, 1}},
			expected: [][]int{{1, 1}, {1, 1}},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			result := rangeAddQueries(c.n, c.queries)

			if len(result) != c.n {
				t.Errorf("Expected %v; got %v", c.expected, result)
			}

			for i, row := range result {
				if len(row) != c.n {
					t.Errorf("Expected %v; got %v", c.expected, result)
				}

				for j, v := range row {
					if v != c.expected[i][j] {
						t.Errorf("Expected %v; got %v", c.expected, result)
					}
				}
			}
		})
	}
}
