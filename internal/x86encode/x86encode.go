// Package x86encode implements simple x86 instructions encoder.
//
// Ignores existence of 32-bit CPU mode.
// Ignores existence of multi-immediate operands instructions.
// Ignores many other avx512test-irrelevant things, like rel-operands and so on.
// This package exists solely to satisfy avx512test needs.
//
// Uses Intel XED under the hood.
package x86encode

import (
	"fmt"
)

// ToHexString returns hex string of octets that describe machine code
// that can be used to encode requested instruction.
//
// Please note that sometimes there are more than one way to
// encode the same instruction. There are no guarantees
// regarding which form will be used. It can also vary between
// different XED versions.
func ToHexString(inst *Inst) (string, error) {
	return encodeToHexString(inst)
}

// Inst describes a single instruction to be encoded.
type Inst struct {
	// Opcode in Intel syntax.
	Opcode string

	Params []InstParam

	Args []Argument
}

type InstParam int

//go:generate stringer -type=InstParam
const (
	ParamBad InstParam = iota
	ParamRexW0
	ParamRexW1
	ParamVexL128
	ParamVexL256
	ParamVexL512
	ParamEOSZ8
	ParamEOSZ16
	ParamEOSZ32
	ParamEOSZ64
)

type DisplacementKind int

const (
	DispSmallest DisplacementKind = iota
	Disp8
	Disp32
)

// Argument carries arbitrary instruction operand.
// Can be memory, immediate or register.
type Argument interface {
	argument()
}

// RegArgument describes register operand.
type RegArgument struct {
	Name string
}

// ImmArgument describes immediate (const) operand.
// There can be only one such operand (if explicit operands are considered).
type ImmArgument struct {
	// Unsigned is true for uint-like arguments.
	// False for int-like arguments.
	Unsigned bool

	// Width is immediate constant width in bits.
	//
	// Common values are:
	//	8  | uint8  or int8
	//	16 | uint16 or int16
	//	32 | uint32 or int32
	//	64 | uint64
	Width uint

	// Value holds uninterpreted integer value of immediate operand.
	Value uint64
}

// MemArgument describes memory operand.
// There can be only one such operand (if explicit operands are considered).
type MemArgument struct {
	// Base register name. (SIB.B)
	//
	// Empty string means "no base".
	Base string

	// Index register name. (SIB.I)
	//
	// Empty string means "no scaled index".
	Index string

	// Scaling factor. (SIB.S)
	//
	// Zero value means "no explicit scaling factor",
	// which usually implies scaling factor of 1.
	Scale int

	// Width is a pointer size in bits.
	//
	// Common values are:
	//	8   | BYTE PTR
	//	16  | WORD PTR
	//	32  | DWORD PTR
	//	64  | QWORD PTR
	//	80  | TBYTE PTR (x87)
	//	128 | XMMWORD PTR
	//	256 | YMMWORD PTR
	//	512 | ZMMWORD PTR
	Width uint

	// Disp is a 32-but pointer.
	// Displacement is encoded as 8 or 32 bit immediate value.
	// Exceptions like MOVABS are ignored.
	Disp int32

	// DispWidth determines displacement encoding strategy.
	// Especially important for EVEX instructions.
	DispWidth DisplacementKind
}

func (*RegArgument) argument() {}
func (*ImmArgument) argument() {}
func (*MemArgument) argument() {}

func encodeToHexString(inst *Inst) (string, error) {
	encoding, err := encodeToBytes(inst)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", encoding), nil
}

func encodeToBytes(inst *Inst) ([]byte, error) {
	xedTablesInit() // Safe to be called multiple times
	return xedEncode(inst)
}
