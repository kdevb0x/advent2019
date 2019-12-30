// Copyright (C) 2019-2020 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package intcode

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"os"
	"strconv"
)

var (
	ErrNegIdxErr        = errors.New("error: index cannot be negative")
	ErrInvalidOpCodeErr = errors.New("error: invalid or unrecognized opcode")
	// Not an error per se...
	ErrEncounteredHalt = errors.New("encountered halt opcode; ending program execution")
)

type SourceCode struct {
	FilePath string
	Data     []int // instructions and arguments

}

func ParseSourceCodeFile(path string) (*SourceCode, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var code SourceCode
	code.FilePath = path
	s := bufio.NewScanner(f)
	for s.Scan() {
		tx := s.Text()
		// remove commas so index matched position in code.Data
		if tx == "," || tx == " " {
			continue
		}
		tok, err := strconv.Atoi(tx)
		if err != nil {
			println(err.Error())
			return nil, err
		}

		code.Data = append(code.Data, tok)
	}
	return &code, nil

}

func (c *SourceCode) ReadAt(p []byte, off int64) (n int, err error) {
	for i := 0; i < len(p); i++ {
		if i >= len(c.Data) {
			return n, io.EOF
		}
		if c.Data[i] == ',' {
			continue
		}
		n += binary.PutVarint(p, int64(c.Data[i]))
		p = p[n:]
	}
	return n, nil
}

func (c *SourceCode) WriterAt(p []byte, off int64) (n int, err error) {
	if off < 0 {
		return 0, ErrNegIdxErr
	}
	m, count := binary.Varint(p)
	c.Data[off] = int(m)
	return count, nil
}

func (c *SourceCode) SetValue(idx, val int) error {
	if idx < 0 {
		return ErrNegIdxErr
	}

	str := strconv.FormatInt(int64(val), 10)
	_, err := c.WriterAt([]byte(str), int64(idx))
	if err != nil {
		return err
	}
	return nil

}

// Perform executes the operation op with parameters c.Data[p1] and c.Data[p2].
func (c *SourceCode) Perform(op int, p1, p2 int) error {
	if !validOpCode((opcode)(op)) {
		return ErrInvalidOpCodeErr
	}
	p1idx := c.Data[p1]

	p2idx := c.Data[p2]
	p1val := c.Data[p1idx]
	p2val := c.Data[p2idx]
	setidx := c.Data[p2+1]
	if oper, ok := ops[opcode(op)]; !ok {
		return ErrInvalidOpCodeErr
	} else {
		fval := oper(p1val, p2val)
		if opcode(fval) == hlt {
			return ErrEncounteredHalt
		}
		c.Data[setidx] = fval

	}
	return nil

}

// Implemented in assembly
func AsmAdd(r0, r1 int64) int64

// Implemented in asm
func AsmMul(r0, r1 int64) int64