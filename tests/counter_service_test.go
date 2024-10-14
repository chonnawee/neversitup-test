package tests

import (
	"neversitup-test/internal/core/implements/services/counter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSmileys(t *testing.T) {
	type testCase struct {
		name    string
		smileys []string
		expeced int
	}
	cases := []testCase{
		{name: "case1", smileys: []string{":)", ";(", ";}", ":-D"}, expeced: 2},
		{name: "case2", smileys: []string{";D", ":-(", ":-)", ";~)"}, expeced: 3},
		{name: "case3", smileys: []string{";]", ":[", ";*", ":$", ";-D"}, expeced: 1},
	}
	service := counter.NewCounterServiceImplement()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			smileys := service.CountSmileyInSlice(c.smileys)
			assert.Equal(t, c.expeced, smileys)
		})
	}
}
