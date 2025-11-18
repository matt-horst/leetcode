package p717

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsOneBitCharacter(t *testing.T) {
	cases := []struct {
		bits     []int
		expected bool
	}{
		{
			bits: []int{1, 0, 0},
			expected: true,
		},
		{
			bits: []int{1, 1, 1, 0},
			expected: false,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := isOneBitCharacter(c.bits)
			assert.Equal(t, result, c.expected)
		})
	}
}
