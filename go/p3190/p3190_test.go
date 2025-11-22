package p3190

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{
			nums: []int{1, 2, 3, 4},
			expected: 3,
		},
		{
			nums: []int{3, 6, 9},
			expected: 0,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := minimumOperations(c.nums)
			assert.Equal(t, c.expected, result)
		})
	}
}
