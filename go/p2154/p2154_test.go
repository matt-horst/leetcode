package p2154

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFinalValue(t *testing.T) {
	cases := []struct {
		nums     []int
		original int
		expected int
	}{
		{
			nums:     []int{5, 3, 6, 1, 12},
			original: 3,
			expected: 24,
		},
		{
			nums:     []int{2, 7, 9},
			original: 4,
			expected: 4,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := findFinalValue(c.nums, c.original)
			assert.Equal(t, result, c.expected)
		})
	}
}
