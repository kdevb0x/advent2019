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

func interpretSource(fpath string) (code string, noun int64, verb int64) {
	src, err := ic.ParseSourceCodeFile(fpath)
	if err != nil {
		panic(err)
	}
	var j, k int64
out:
	for j = 0; j <= 99; j++ {
		for k = 99; k >= 0; k-- {
			src.Data[1] = j
			src.Data[2] = k
			for i := 0; i < len(src.Data); i += 4 {
				if src.Data[i] == int64(ic.Halt) {
					break
				}
				op, ok := ic.Ops[ic.Opcode(src.Data[i])]
				if ok {
					sum := op(src.Data[src.Data[i+1]], src.Data[src.Data[i+2]])
					idx := src.Data[i+3]
					src.Data[idx] = sum
				}
			}
		}
		if src.Data[0] == 19690720 {
			noun = j
			verb = k
			break out
		}
		if src.Data[0] != 19690720 {
			src.Reset()
			continue
		}
	}
	code = formatCode(src.Data)
	return

}

func formatCode(ops []int64) string {
	var opstr []string
	for _, op := range ops {
		opstr = append(opstr, strconv.Itoa(int(op)))

	}
	return strings.Join(opstr, ",")

}

func main() {
	code, noun, verb := interpretSource("../../input")
	fmt.Printf("noun: %d, verb: %d, result: %v\n", noun, verb, code[0])
}
