package fsmregex

import (
	"errors"
	"fmt"
)

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

func (f *FSM) Compile(s string) error {
	// reset f
	f.c = [][ASCIISIZE]int{}
	f.state = 0

	f.addCol([ASCIISIZE]int{}) // default state
	chars := []rune(s)
	for i := 0; i < len(chars); i++ {
		c := chars[i]

		switch c {
		// support for $ end of regex string to be matched
		// should not have any characters after char before
		// $ in regex
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
		case '[':
			// match range
			// find closing ] or raise error
			t := [ASCIISIZE]int{}
			i += 1
			g := []rune{}

			// in case we have range like [16789]
			for k := 1; i < len(chars); i++ {
				if chars[i] == ']' {
					break
				}
				if k == 2 && chars[i] == '-' {
					if chars[i+2] != ']' {
						return errors.New("invalid range")
					}
					// check for ranges a-z A-Z 1-n
					if (chars[i-1] < 'a' && chars[i-1] > 'z') &&
						(chars[i+1] < 'a' || chars[i+1] > 'z') {
						return errors.New("Invalid range")
					}
					if (chars[i-1] < 'A' && chars[i-1] > 'Z') &&
						(chars[i+1] < 'A' || chars[i+1] > 'Z') {
						return errors.New("Invalid range")
					}
					if chars[i-1] < '0' && chars[i+1] > '9' {
						return errors.New("Invalid range")
					}

					// valid range
					for r := chars[i-1] + 1; r <= chars[i+1]; r++ {
						g = append(g, r)
						fmt.Printf("g: %d len t %d \n", int(r), int(len(t)))
						t[r] = len(f.c) + 1
					}

					i += 2
					//break from for loop
					break
				}

				g = append(g, chars[i])
				t[chars[i]] = len(f.c) + 1
				k += 1
			}

			//if g is still empty till end of string
			// error
			if len(g) == 0 && i == len(chars) {
				return errors.New("[ has no end ]")
			}

			f.addCol(t)
		default:
			t := [ASCIISIZE]int{}

			t[int(c)] = len(f.c) + 1
			f.addCol(t)
		}
	}

	return nil
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
