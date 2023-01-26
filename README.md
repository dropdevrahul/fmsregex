### Simple project to demonstrate usage of Finite State Machines
Uses finite state machine to implement regex in Golang.

Currently supports '$' and '.' in usual sens of regex that is $ means regex match should end and . means any printable character can be present (only 1 character)

```
package main

import (
	"github.com/dropdevrahul/fmsregex/fmsregex"
)

func main() {
  f := fmsregex.FSM{}
  f.Compile("abc.g$")
  res := f.Match("abceg") // true
  res = f.Match("abcfg") // true
  res = f.Match("abcega") // false

  f.Compile("abcg")
  res = f.Match("abcgagagdsajd") // true

  f.Compile("abcg$")
  res = f.Match("abcgagagdsajd") // false
}

```
