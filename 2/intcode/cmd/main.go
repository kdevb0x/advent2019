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
	for ; noun <= 99; noun++ {
		for ; verb <= 99; verb++ {
			src.Data[1] = noun
			src.Data[2] = verb
		halt:
			for i := 0; i < len(src.Data); i += 4 {
				switch ic.Opcode(src.Data[i]) {
				case ic.Halt:
					break halt
				case ic.Add:
					sum := ic.AsmAdd(src.Data[src.Data[i+1]], src.Data[src.Data[i+2]])
					idx := src.Data[i+3]
					src.Data[idx] = sum
				case ic.Mul:
					prod := ic.AsmMul(src.Data[src.Data[i+1]], src.Data[src.Data[i+2]])
					idx := src.Data[i+3]
					src.Data[idx] = prod

				}
			}
			fmt.Println(src.Data[0])
			if src.Data[0] == 19690720 {
				return
			}
			src.Reset()
		}

	}

	code = formatCode(src.Data)
	return

}

func formatCode(ops []int64) string {
	var opstr = make([]string, len(ops))
	for _, op := range ops {
		opstr = append(opstr, strconv.Itoa(int(op)))

	}
	return strings.Join(opstr, ",")

}

func main() {
	code, noun, verb := interpretSource("../../input")
	fmt.Printf("noun: %d, verb: %d, result: %v\n", noun, verb, code)
}
