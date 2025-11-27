package p3381

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		nums     []int
		k        int
		expected int64
	}{
		{
			nums:     []int{1, 2},
			k:        1,
			expected: 3,
		},
		{
			nums:     []int{-1, -2, -3, -4, -5},
			k:        4,
			expected: -10,
		},
		{
			nums:     []int{-5, 1, 2, -3, 4},
			k:        2,
			expected: 4,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := maxSubarraySum(c.nums, c.k)
			assert.Equal(t, c.expected, result)
		})
	}
}
