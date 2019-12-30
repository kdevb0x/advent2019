// Copyright (C) 2019-2020 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package intcode

// opcode in hexidecimal format
type opcodeHex int

const (
	_      opcodeHex = iota
	addHex           // 0x01
	mulHex           // 0x02

	hltHex opcodeHex = 0x63 // 99
)

// opcode in decimal format
type opcode int

const (
	_   opcode = iota
	add        // 1
	mul        // 2
	hlt opcode = 99
)

var ops = map[opcode]func(...int) int{
	add: func(oprnd ...int) int {
		if len(oprnd) != 2 {
			return 0
		}
		return oprnd[0] + oprnd[1]
	},
	mul: func(oprnd ...int) int {
		if len(oprnd) != 2 {
			return 0
		}
		return oprnd[0] * oprnd[1]
	},
	hlt: func(oprnd ...int) int {
		return int(hlt)
	},
}

func validOpCode(n opcode) bool {
	if n == add || n == mul || n == hlt {
		return true
	}
	return false
}
