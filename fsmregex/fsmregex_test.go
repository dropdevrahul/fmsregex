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
		"group-[range]-check": {
			input: "abc1d",
			regex: "abc[12345]d$",
			exp:   true,
		},
		"group-[range]-check-last-char-]": {
			input: "abc1",
			regex: "abc[12345]$",
			exp:   true,
		},
		"group-[0-9]-check-1": {
			input: "abc1",
			regex: "abc[0-9]$",
			exp:   true,
		},
		"group-[0-9]-check-2": {
			input: "abc8",
			regex: "abc[0-9]$",
			exp:   true,
		},
		"group-[0-9]-check-3": {
			input: "abc9",
			regex: "abc[0-9]$",
			exp:   true,
		},
		"group-[a-b]-check-3": {
			input: "abcbabc",
			regex: "abc[a-b]",
			exp:   true,
		},
		"group-[a-b]-check-4": {
			input: "hello a 8 abc",
			regex: "hello.[a-z].[8-8]",
			exp:   true,
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
