package p1015

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		k        int
		expected int
	}{
		{
			k:        1,
			expected: 1,
		},
		{
			k:        2,
			expected: -1,
		},
		{
			k:        3,
			expected: 3,
		},
		{
			k:        17,
			expected: 16,
		},
		{
			k:        4,
			expected: -1,
		},
		{
			k:        123456,
			expected: -1,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := smallestRepunitDivByK(c.k)
			assert.Equal(t, c.expected, result)
		})
	}
}
