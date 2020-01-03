// Copyright (C) 2019-2020 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package main

import (
	"fmt"
	"strconv"
	"strings"

	ic "github.com/kdevb0x/advent2019/2/intcode"
)

func p2(prog []int64, output int64) int {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			if output == run(setState(prog, noun, verb))[0] {
				return 100*noun + verb
			}

		}
	}
	return -1

}

func setState(prog []int64, noun int, verb int) []int64 {
	prog[1] = int64(noun)
	prog[2] = int64(verb)
	return prog

}

func interpretSource(fpath string) (code []int64) {
	src, err := ic.ParseSourceCodeFile(fpath)
	if err != nil {
		panic(err)
	}
	return src.Data[:]

}

func run(prog []int64) []int64 {
	p := make([]int64, len(prog))
	copy(p, prog)
	var pars = 0
	for i := 0; i < len(p); i += pars {
		switch p[i] {
		case 1:
			prog[prog[i+3]] = prog[prog[i+1]] + prog[prog[i+2]]
			pars = 4
		case 2:
			prog[prog[i+3]] = prog[prog[i+1]] * prog[prog[i+2]]
			pars = 4
		case 99:
			return prog
		}
	}
	return prog

}

func formatCode(ops []int64) string {
	var opstr = make([]string, len(ops))
	for _, op := range ops {
		opstr = append(opstr, strconv.Itoa(int(op)))

	}
	return strings.Join(opstr, ",")

}

func main() {
	p := interpretSource("../../input")
	ans := p2(p, 19690720)
	fmt.Printf("%s\n", ans)
}
