package main

import (
	"fmt"

	"github.com/dropdevrahul/fmsregex/fmsregex"
)

func main() {
	f := fmsregex.FSM{}
	inputs := []string{"bc", "bcd", "abc", "abcd", "hello, world"}
	f.Compile("bc$")
	f.Dump()
	for _, i := range inputs {
		fmt.Printf(i)
		fmt.Println(f.Match(i))
	}
}
