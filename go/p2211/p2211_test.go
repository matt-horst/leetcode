package p2211

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		directions string
		expected   int
	}{
		{
			directions: "RLRSLL",
			expected: 5,
		},
		{
			directions: "LLRR",
			expected: 0,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := countCollisions(c.directions)
			assert.Equal(t, c.expected, result)
		})
	}
}
