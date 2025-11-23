package p1262

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{
			nums: []int{3, 6, 5, 1, 8},
			expected: 18,
		},
		{
			nums: []int{4},
			expected: 0,
		},
		{
			nums: []int{1, 2, 3, 4, 4},
			expected: 12,
		},
		{
			nums: []int{2, 6, 2, 2, 7},
			expected: 15,
		},
		{
			nums: []int{4, 1, 5, 3, 1},
			expected: 12,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := maxSumDivThree(c.nums)
			assert.Equal(t, c.expected, result)
		})
	}
}
