// Copyright (C) 2019-2020 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package main

import (
	ic "github.com/kdevb0x/advent2019/2/intcode"
)

func main() {
	src, err := ic.ParseSourceCodeFile("../../input")
	if err != nil {
		panic(err)
	}

}
