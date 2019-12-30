// Copyright (C) 2019-2020 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func parseInput(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var text []string

	s := bufio.NewScanner(f)
	for s.Scan() {
		text = append(text, s.Text())
		if s.Err() == io.EOF {
			break
		}

	}
	var masses []int
	for _, n := range text {
		m, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		masses = append(masses, m)
	}
	return masses, nil
}

func fuelPerModule(mass int) int {
	thrd := mass / 3
	f := math.Floor(float64(thrd))
	return int(f) - 2
}

func totalFuelNeeded(mods []int) int {
	var total int

	for _, n := range mods {
		var t int
		for (n/3)-2 > 0 {
			t = (n / 3) - 2
			n = t
			total += n
		}
	}
	return total
}

func main() {
	input := os.Args[1]
	mods, err := parseInput(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("total fuel needed: %d\n", totalFuelNeeded(mods))
}
