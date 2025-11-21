package p1930

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	cases := []struct {
		s        string
		expected int
	}{
		{
			s: "aabca",
			expected: 3,
		}, 
		{
			s: "adc",
			expected: 0,
		},
		{
			s: "bbcbaba",
			expected: 4,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := countPalindromicSubsequence(c.s)
			assert.Equal(t, c.expected, result)
		})
	}
}
