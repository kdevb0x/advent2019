// Copyright (C) 2019-2020 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

#include "textflag.h"
#include "go_asm.h"

TEXT ·AsmAdd(SB), NOSPLIT, $0-24
	MOVQ    r0+0(FP), BX
	MOVQ    r1+8(FP), AX
	ADDQ 	BX, AX
	MOVQ    AX, ret+16(FP)
	RET

TEXT ·AsmMul(SB), NOSPLIT, $0-24
	MOVQ    r0+0(FP), BX
	MOVQ    r1+8(FP), AX
	IMULL 	BX, AX
	MOVQ    AX, ret+16(FP)
	RET
