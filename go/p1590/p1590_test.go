package p1590

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		nums     []int
		p        int
		expected int
	}{
		{
			nums:     []int{3, 1, 4, 2},
			p:        6,
			expected: 1,
		},
		{
			nums:     []int{6, 3, 5, 2},
			p:        9,
			expected: 2,
		},
		{
			nums:     []int{1, 2, 3},
			p:        3,
			expected: 0,
		},
		{
			nums:     []int{1, 2, 3},
			p:        7,
			expected: -1,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := minSubarray(c.nums, c.p)
			assert.Equal(t, c.expected, result)
		})
	}
}
