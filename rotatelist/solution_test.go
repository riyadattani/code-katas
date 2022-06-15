package rotatelist_test

import (
	"fmt"
	"testing"

	"code-katas/rotatelist"

	"github.com/alecthomas/assert/v2"
)

//Write a function that rotates a list by n elements. For example [1,2,3,4,5,6] rotated by 2 becomes [3,4,5,6,1,2].
//Bonus: Try solving this without creating a copy of the list.

func TestRotate(t *testing.T) {
	var tcs = []struct {
		givenRotateNumber int
		givenList         []int
		want              []int
	}{
		{givenRotateNumber: 2, givenList: []int{1, 2, 3, 4, 5, 6}, want: []int{3, 4, 5, 6, 1, 2}},
		{givenRotateNumber: 7, givenList: []int{1, 2, 3, 4, 5, 6}, want: []int{}},
		{givenRotateNumber: -7, givenList: []int{1, 2, 3, 4, 5, 6}, want: []int{}},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("Given a list: %v, rotate the list given a number: %q", tc.givenList, tc.givenRotateNumber), func(t *testing.T) {
			assert.Equal(t, tc.want, rotatelist.Rotate(tc.givenList, tc.givenRotateNumber))
		})
	}
}
