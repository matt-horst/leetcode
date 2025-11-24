package p1018

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		nums     []int
		expected []bool
	}{
		{
			nums: []int{0, 1, 1},
			expected: []bool{true, false, false},
		},
		{
			nums: []int{1, 1, 1},
			expected: []bool{false, false, false},
		},
		{
			nums: []int{1,0,1,1,1,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,1,1,1,1,0,0,0,0,1,1,1,0,0,0,0,0,1,0,0,0,1,0,0,1,1,1,1,1,1,0,1,1,0,1,0,0,0,0,0,0,1,0,1,1,1,0,0,1,0},
			expected: []bool{false,false,true,false,false,false,false,false,false,false,true,true,true,true,true,true,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,true,false,false,false,true,false,false,true,false,false,true,true,true,true,true,true,true,false,false,true,false,false,false,false,true,true},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			result := prefixesDivBy5(c.nums)
			assert.ElementsMatch(t, c.expected, result)
		})
	}
}
