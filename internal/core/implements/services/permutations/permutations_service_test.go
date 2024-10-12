package permutations_test

import (
	"neversitup-test/internal/core/implements/services/permutations"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutations(t *testing.T) {
	type testCase struct {
		name    string
		str     string
		expeced []string
	}
	cases := []testCase{
		{name: "a", str: "a", expeced: []string{"a"}},
		{name: "ab", str: "ab", expeced: []string{"ab", "ba"}},
		{name: "abc", str: "abc", expeced: []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		// {name: "aabb", str: "aabb", expeced: []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}},
	}
	service := permutations.NewPermutationsServiceImplement()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			permutation := service.Generate(c.str)
			assert.Equal(t, c.expeced, permutation)
		})
	}
}
