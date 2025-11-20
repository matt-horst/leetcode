package p757

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersectionSizeTwo(t *testing.T) {
	cases := []struct {
		intervals [][]int
		expected  int
	}{
		{
			intervals: [][]int{{1, 3}, {3, 7}, {8, 9}},
			expected: 5,
		},
		{
			intervals: [][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}},
			expected: 3,
		},
		{
			intervals: [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}},
			expected: 5,
		},
		{
			intervals: [][]int{{2,10},{3,7},{3,15},{4,11},{6,12},{6,16},{7,8},{7,11},{7,15},{11,12}},
			expected: 5,
		},
		{
			intervals: [][]int{{1, 3}, {3, 7}, {5, 7}, {7, 8}},
			expected: 5,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := intersectionSizeTwo(c.intervals)
			assert.Equal(t, c.expected, result)
		})
	}
}
