// Snippet extracted from: arith_386.s
// This file provides fast assembly versions for the elementary
// arithmetic operations on vectors implemented in arith.go.

// func mulWW(x, y Word) (z1, z0 Word)
TEXT ·mulWW(SB),NOSPLIT,$0
    MOVL x+0(FP), AX
    MULL y+4(FP)
    MOVL DX, z1+8(FP)
    MOVL AX, z0+12(FP)
    RET