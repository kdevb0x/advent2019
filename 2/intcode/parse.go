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
	"strings"
)

var (
	ErrNegIdxErr        = errors.New("error: index cannot be negative")
	ErrInvalidOpCodeErr = errors.New("error: invalid or unrecognized opcode")
	// Not an error per se...
	ErrEncounteredHalt = errors.New("encountered halt opcode; ending program execution")
)

type SourceCode struct {
	FilePath   string
	Data       []int64 // instructions and arguments
	StartState []int64 // the initial memory state of the program.

}

func ParseSourceCodeFile(path string) (*SourceCode, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var code SourceCode
	code.FilePath = path
	s := bufio.NewScanner(f)
	// s.Split(bufio.ScanRunes)
	for s.Scan() {
		if s.Err() == io.EOF {
			break
		}
		tx := s.Text()
		txs := strings.Split(tx, ",")
		for _, n := range txs {

			// remove commas so index matched position in code.Data
			switch n {
			case ",", " ", "\n", "":
				continue
			}
			tok, err := strconv.Atoi(n)
			if err != nil {
				println(err.Error())
				return nil, err
			}

			code.Data = append(code.Data, int64(tok))
		}
	}
	code.StartState = code.Data
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
	c.Data[off] = m
	return count, nil
}

func (c *SourceCode) SetValue(idx int, val int64) error {
	if idx < 0 {
		return ErrNegIdxErr
	}

	str := strconv.FormatInt(val, 10)
	_, err := c.WriterAt([]byte(str), int64(idx))
	if err != nil {
		return err
	}
	return nil

}

// Perform executes the operation op with parameters c.Data[p1] and c.Data[p2].
func (c *SourceCode) Perform(op int64, p1, p2 int) error {
	if !ValidOpCode((Opcode)(op)) {
		return ErrInvalidOpCodeErr
	}
	p1idx := c.Data[p1]

	p2idx := c.Data[p2]
	p1val := c.Data[p1idx]
	p2val := c.Data[p2idx]
	setidx := c.Data[p2+1]
	if oper, ok := Ops[Opcode(op)]; !ok {
		return ErrInvalidOpCodeErr
	} else {
		fval := oper(p1val, p2val)
		if Opcode(fval) == Halt {
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

func (c *SourceCode) Reset() {
	c.Data = c.StartState
}
