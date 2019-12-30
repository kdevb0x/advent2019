// Copyright (C) 2019-2020 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package intcode

// opcode in hexidecimal format
type OpcodeHex int

const (
	_      OpcodeHex = iota
	AddHex           // 0x01
	MulHex           // 0x02

	HaltHex OpcodeHex = 0x63 // 99
)

// opcode in decimal format
type Opcode int

const (
	_    Opcode = iota
	Add         // 1
	Mul         // 2
	Halt Opcode = 99
)

var Ops = map[Opcode]func(...int) int{
	Add: func(oprnd ...int) int {
		if len(oprnd) != 2 {
			return 0
		}
		return oprnd[0] + oprnd[1]
	},
	Mul: func(oprnd ...int) int {
		if len(oprnd) != 2 {
			return 0
		}
		return oprnd[0] * oprnd[1]
	},
	Halt: func(oprnd ...int) int {
		return int(Halt)
	},
}

func ValidOpCode(n Opcode) bool {
	if n == Add || n == Mul || n == Halt {
		return true
	}
	return false
}
