package p1437

import (
	"fmt"
	"testing"
)

func TestKLengthApart(t *testing.T) {
	cases := []struct {
		n        []int
		k        int
		expected bool
	}{
		{
			n:        []int{1, 0, 0, 0, 1, 0, 0, 1},
			k:        2,
			expected: true,
		},
		{
			n:        []int{1, 0, 0, 1, 0, 1},
			k:        2,
			expected: false,
		},
		{
			n:        []int{0, 1, 0, 1},
			k:        1,
			expected: true,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := kLengthApart(c.n, c.k)
			if result != c.expected {
				t.Errorf("expected %v; got %v", c.expected, result)
			}
		})
	}
}
