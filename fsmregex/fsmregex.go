package fsmregex

import "fmt"

const FSMENDLINE = '$'
const ASCIISIZE = 128

type FSM struct {
	state int
	c     [][ASCIISIZE]int // store all ascii characters as int
	// and track state of each char
}

func (f *FSM) Dump() {
	for i := 0; i <= 128; i++ {
		fmt.Printf("%s  ", string(rune(i)))
		for j := 0; j < len(f.c); j++ {
			fmt.Printf("%d ", f.c[j][i])
		}
		fmt.Println()
	}
}

func (f *FSM) addCol(c [ASCIISIZE]int) {
	if len(f.c) == 0 {
		f.c = [][ASCIISIZE]int{}
	}
	f.c = append(f.c, c)
}

func (f *FSM) Compile(s string) {
	// reset f
	f.c = [][ASCIISIZE]int{}
	f.state = 0

	f.addCol([ASCIISIZE]int{}) // default state
	for _, c := range []rune(s) {
		switch c {
		case FSMENDLINE:
			// also handles FSMENDLINE
			t := [ASCIISIZE]int{}
			t[FSMENDLINE] = len(f.c) + 1
			f.addCol(t)
		case '.':
			// . means check for match any printable character
			t := [ASCIISIZE]int{}
			l := len(f.c)
			for i := 32; i <= 127; i++ {
				t[i] = l + 1
			}
			f.addCol(t)
		default:
			t := [ASCIISIZE]int{}
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
