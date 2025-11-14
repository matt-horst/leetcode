package p3228

import (
	"fmt"
	"testing"
)

func TestMaxOperations(t *testing.T) {
	cases := []struct {
		input string
		expects int
	} {
		{
			input: "1001101",
			expects: 4,
		},
		{
			input: "00111",
			expects: 0,
		},
		{
			input: "110",
			expects: 2,
		},
		{
			input: "0011010101100",
			expects: 15,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := maxOperations(c.input)
			if result != c.expects {
				t.Errorf("Expected %v, but got %v", c.expects, result)
			}
		})
	}
}
