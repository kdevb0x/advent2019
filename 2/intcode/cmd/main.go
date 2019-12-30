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

func interpretSource(fpath string) string {
	src, err := ic.ParseSourceCodeFile(fpath)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(src.Data); i += 4 {
		if src.Data[i] == int64(ic.Halt) {
			break
		}
		op, ok := ic.Ops[ic.Opcode(src.Data[i])]
		if ok {
			sum := op(src.Data[i+1], src.Data[i+2])
			idx := src.Data[i+3]
			src.Data[idx] = sum
		}
	}
	return formatCode(src.Data)

}

func formatCode(ops []int64) string {
	var opstr []string
	for _, op := range ops {
		opstr = append(opstr, strconv.Itoa(int(op)))

	}
	return strings.Join(opstr, ",")

}

func main() {
	code := interpretSource("../../input")
	fmt.Println(code)
}
