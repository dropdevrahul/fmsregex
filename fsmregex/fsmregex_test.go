package fsmregex_test

import (
	"testing"

	"github.com/dropdevrahul/fsmregex/fsmregex"
	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	cases := map[string]struct {
		regex string
		input string
		exp   bool
	}{
		"basic-failed": {
			input: "abchel",
			regex: "hel",
			exp:   false,
		},
		"basic": {
			input: "hello",
			regex: "hel",
			exp:   true,
		},
		"exact": {
			input: "hel",
			regex: "hel",
			exp:   true,
		},
		"terminated": {
			input: "hello",
			regex: "hel$",
			exp:   false,
		},
		"terminated-exact": {
			input: "hel$",
			regex: "hel",
			exp:   true,
		},
		"dot-1": {
			input: "abceg",
			regex: "abc.g",
			exp:   true,
		},
		"dot-2": {
			input: "abcfg",
			regex: "abc.g",
			exp:   true,
		},
		"dot-no-char-after": {
			input: "abc",
			regex: "abc.g",
			exp:   false,
		},
		"dot-failed": {
			input: "abce",
			regex: "abc.g",
			exp:   false,
		},
	}

	for tl, tc := range cases {
		t.Run(tl, func(t *testing.T) {
			f := fsmregex.FSM{}
			f.Compile(tc.regex)
			res := f.Match(tc.input)
			assert.Equal(t, tc.exp, res)
		})
	}
}
