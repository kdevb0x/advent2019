// Copyright (C) 2019-2020 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package main

import (
	"github.com/kdevb0x/advent2019/2/intcode"
)

func ex2() (noun, verb int64) {
	var match int64 = 19690720
	m, err := intcode.LoadProgram("../../input")
	if err != nil {
		panic(err)
	}
	for n := 0; n <= 99; n++ {
		for v := 0; v <= 99; v++ {
			m[1] = int64(n)
			m[2] = int64(v)
			final := intcode.RunProgram(m)
			if final[0] == match {
				return int64(n), int64(v)
			}
		}
	}
	return -1, -1
}
func main() {
	noun, verb := ex2()
	println((100 * noun) + verb)
}
