package p1523
import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	cases := []struct {
		low int
		high int
		expected int
	}{
		{
			low: 3,
			high: 7,
			expected: 3,
		},
		{
			low: 8,
			high: 10,
			expected: 1,
		},
		{
			low: 8,
			high: 11,
			expected: 2,
		},
		{
			low: 7,
			high: 12,
			expected: 3,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := countOdds(c.low, c.high)
			assert.Equal(t, c.expected, result)
		})
	}
}
