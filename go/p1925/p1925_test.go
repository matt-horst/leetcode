package p1925

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		n        int
		expected int
	}{
		{
			n:        5,
			expected: 2,
		},
		{
			n:        10,
			expected: 4,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := countTriples(c.n)
			assert.Equal(t, c.expected, result)
		})
	}
}
