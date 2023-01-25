package fmsregex

import "fmt"

const FSMENDLINE = '$'

type FSM struct {
	state int
	c     [][128]int // store all ascii characters as int
	// and track state of each char
}

func (f *FSM) Dump() {
	for i, j := range f.c {
		for k, _ := range j {
			fmt.Printf("%d %d => %d \n", i, k, f.c[i][k])
		}
	}
}

func (f *FSM) addCol(c [128]int) {
	if len(f.c) == 0 {
		f.c = [][128]int{}
	}
	f.c = append(f.c, c)
}

func (f *FSM) Compile(s string) {
	f.addCol([128]int{}) // default state
	for _, c := range []rune(s) {
		switch c {
		default:
			t := [128]int{}
			t[int(c)] = len(f.c) + 1
			f.addCol(t)
		}
	}
}

func (f *FSM) Match(s string) bool {
	state := 1 // starting state
	for _, i := range []rune(s) {
		if state == 0 || state >= len(f.c) {
			break
		}

		state = f.c[state][int(i)]
	}

	if state == 0 {
		return false
	}

	if state < len(f.c) {
		state = f.c[state][FSMENDLINE]
	}

	return state >= len(f.c)
}
