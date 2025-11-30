package p3512

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
			nums: []int{3, 9, 7},
			k: 5,
			expected: 4,
		},
		{
			nums: []int{4, 1, 3},
			k: 4,
			expected: 0,
		},
		{
			nums: []int{3, 2},
			k: 6,
			expected: 5,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := minOperations(c.nums, c.k)
			assert.Equal(t, c.expected, result)
		})
	}
}
