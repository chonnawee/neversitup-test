package odd_test

import (
	"neversitup-test/internal/core/implements/services/odd"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOddInt(t *testing.T) {
	type testCase struct {
		name    string
		numbers []int
		expeced int
	}
	cases := []testCase{
		{name: "numbers [7]", numbers: []int{7}, expeced: 7},
		{name: "numbers [0]", numbers: []int{0}, expeced: 0},
		{name: "numbers [1,1,2]", numbers: []int{1, 1, 2}, expeced: 2},
		{name: "numbers [1,2,2,3,3,3,4,3,3,3,2,2,1]", numbers: []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}, expeced: 4},
	}
	service := odd.NewOddServiceImplement()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			permutation := service.FindIntInSlice(c.numbers)
			assert.Equal(t, c.expeced, permutation)
		})
	}
}
