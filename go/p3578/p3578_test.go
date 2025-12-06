package p3578

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			nums:     []int{9, 4, 1, 3, 7},
			k:        4,
			expected: 6,
		},
		{
			nums:     []int{3, 3, 4},
			k:        0,
			expected: 2,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := countPartitions(c.nums, c.k)
			assert.Equal(t, c.expected, result)
		})
	}
}
