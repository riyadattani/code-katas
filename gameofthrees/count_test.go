package gameofthrees_test

import (
	"fmt"
	"testing"

	"code-katas/gameofthrees"

	"github.com/alecthomas/assert/v2"
)

//Given any number, repeatedly divide the number by 3 till you reach 1. Add or subtract 1 whenever division by 3 is not possible. At each stage output the number.
//e.g. if you start with 100, it’s not divisible by 3, so subtract 1.
//100, 99
//Then keep dividing by 3
//100, 99, 33, 11
//11 is not divisible by 3, so add 1.
//100, 99, 33, 11, 12
//Keep going till you get to 1.
//100, 99, 33, 11, 12, 4, 3, 1

func TestCount(t *testing.T) {
	var tcs = []struct {
		given int
		want  []int
	}{
		{given: 100, want: []int{100, 99, 33, 11, 12, 4, 3, 1}},
		{given: 6, want: []int{6, 2, 3, 1}},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("Given the number %v, repeatedly divide the number by 3 till you reach 1", tc.given), func(t *testing.T) {
			gotNums := gameofthrees.Count(tc.given)
			assert.Equal(t, tc.want, gotNums)
		})
	}
}
