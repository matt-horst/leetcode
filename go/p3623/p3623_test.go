package p3623

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		points   [][]int
		expected int
	}{
		{
			points:   [][]int{{1, 0}, {2, 0}, {3, 0}, {2, 2}, {3, 2}},
			expected: 3,
		},
		{
			points:   [][]int{{0, 0}, {1, 0}, {0, 1}, {2, 1}},
			expected: 1,
		},
		{
			points:   [][]int{{-94, -36}, {69, 35}, {39, -36}, {-17, 35}},
			expected: 1,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := countTrapezoids(c.points)
			assert.Equal(t, c.expected, result)
		})
	}
}
