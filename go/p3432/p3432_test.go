package p3432

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
			nums:     []int{10, 10, 3, 7, 6},
			expected: 4,
		},
		{
			nums:     []int{1, 2, 2},
			expected: 0,
		},
		{
			nums:     []int{2, 4, 6, 8},
			expected: 3,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := countPartitions(c.nums)
			assert.Equal(t, c.expected, result)
		})
	}
}
