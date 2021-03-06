// Copyright (C) 2019-2020 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package intcode

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type MemState []int64

func LoadProgram(path string) (MemState, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	n := strings.ReplaceAll(string(f), "\n", "")
	r := strings.Split(n, ",")
	var mem []int64
	for _, n := range r {
		q, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		mem = append(mem, int64(q))
	}

	return mem, nil

}

func RunProgram(m MemState) MemState {
	// var pc int
out:
	for i := 0; i < len(m); i += 4 {
		switch m[i] {
		case 1:
			ansidx, arg1idx, arg2idx := m[i+3], m[i+1], m[i+2]
			m[ansidx] = m[arg1idx] + m[arg2idx]
		// 	pc = 4
		case 2:
			ansidx, arg1idx, arg2idx := m[i+3], m[i+1], m[i+2]
			m[ansidx] = m[arg1idx] * m[arg2idx]
			// pc = 4
		case 99:
			// pc = 1
			break out
		}
	}
	return m
}

//go:noescape
// AsmAdd is an ADD instruction implemented in assembly.
func AsmAdd(a, b int64) int64

//go:noescape
// AsmMul is an MUL instruction implemented in assembly.
func AsmMul(a, b int64) int64
