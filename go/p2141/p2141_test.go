package p2141

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		n         int
		batteries []int
		expected  int64
	}{
		{
			n:         2,
			batteries: []int{3, 3, 3},
			expected:  4,
		},
		{
			n:         2,
			batteries: []int{1, 1, 1, 1},
			expected:  2,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := maxRunTime(c.n, c.batteries)
			assert.Equal(t, c.expected, result)
		})
	}
}
