// Copyright (C) 2019-2020 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package intcode

// opcode in hexidecimal format
type OpcodeHex int64

const (
	_      OpcodeHex = iota
	AddHex           // 0x01
	MulHex           // 0x02

	HaltHex OpcodeHex = 0x63 // 99
)

// opcode in decimal format
type Opcode int64

const (
	_    Opcode = iota
	Add         // 1
	Mul         // 2
	Halt Opcode = 99
)

var Ops = map[Opcode]func(int64, int64) int64{
	Add: func(op1, op2 int64) int64 {
		return AsmAdd(op1, op2)
	},
	Mul: func(op1, op2 int64) int64 {
		return AsmMul(op1, op2)
	},
	Halt: func(op1, op2 int64) int64 {
		return int64(Halt)
	},
}

func ValidOpCode(n Opcode) bool {
	if n == Add || n == Mul || n == Halt {
		return true
	}
	return false
}
