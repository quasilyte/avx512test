package main

import (
	"fmt"

	"github.com/quasilyte/avx512test/internal/x86encode"
)

var intelRegToGoRegMap = map[string]string{
	// Note that the list is incomplete.

	"EBP":   "BP",
	"RBP":   "BP",
	"ESI":   "SI",
	"RSI":   "SI",
	"EDI":   "DI",
	"RDI":   "DI",
	"EAX":   "AX",
	"RAX":   "AX",
	"R8D":   "R8",
	"R8":    "R8",
	"R9D":   "R9",
	"R9":    "R9",
	"R10D":  "R10",
	"R10":   "R10",
	"R14D":  "R14",
	"R14":   "R14",
	"R15D":  "R15",
	"R15":   "R15",
	"EDX":   "DX",
	"RDX":   "DX",
	"EBX":   "BX",
	"RBX":   "BX",
	"ECX":   "CX",
	"RCX":   "CX",
	"ESP":   "SP",
	"RSP":   "SP",
	"XMM4":  "X4",
	"XMM10": "X10",
	"XMM29": "X29",
	"YMM3":  "Y3",
	"YMM4":  "Y4",
	"YMM5":  "Y5",
	"YMM9":  "Y9",
	"YMM11": "Y11",
	"YMM28": "Y28",
	"YMM30": "Y30",
	"ZMM4":  "Z4",
	"ZMM9":  "Z9",
	"ZMM12": "Z12",
	"ZMM31": "Z31",
}

func intelRegToGoReg(intelName string) string {
	goName := intelRegToGoRegMap[intelName]
	if goName == "" {
		panic(fmt.Sprintf("empty Intel->Go reg mapping for %q", intelName))
	}
	return goName
}

func memoryExpression(mem *x86encode.MemArgument) string {
	base := intelRegToGoReg(mem.Base)
	expr := fmt.Sprintf("(%s)", base)

	scale := 1 // Default
	if mem.Scale != 0 {
		scale = mem.Scale
	}

	// Append index expression, if any.
	if mem.Index != "" {
		index := intelRegToGoReg(mem.Index)
		expr += fmt.Sprintf("(%s*%d)", index, scale)
	}

	// Prepend displacement, if any.
	if mem.Disp != 0 {
		expr = fmt.Sprint(mem.Disp) + expr
	}

	return expr
}
