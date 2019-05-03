package main

import (
	"fmt"
	"strings"

	"github.com/quasilyte/avx512test/internal/x86encode"
	"golang.org/x/arch/x86/x86csv"
)

// This file acts as a configuration.
// The output depends on this file contents directly.

// TODO(quasilyte): less hardcoded options. Generate programmatically.
// What we trying to do here looks like property-based testing,
// but without an appropriate framework.

type instArg struct {
	goSyntax string
	data     x86encode.Argument
}

// instArgsBySyntax maps x86csv operand syntax string to a list
// of appropriate arguments that can be used to cover it.
//
// Initialized inside init().
var instArgsBySyntax map[string][]instArg

func init() {
	type mem = x86encode.MemArgument
	type imm = x86encode.ImmArgument
	type reg = x86encode.RegArgument

	makeRegArgs := func(name, goFmt, xedFmt string, ids ...int) []instArg {
		args := make([]instArg, len(ids))
		for i, id := range ids {
			goSyntax := fmt.Sprintf(goFmt, name, id)
			data := &reg{Name: fmt.Sprintf(xedFmt, name, id)}
			args[i] = instArg{goSyntax, data}
		}
		return args
	}

	makeMaskRegArgs := func(ids ...int) []instArg {
		return makeRegArgs("K", "%s%d", "%s%d", ids...)
	}

	makeVecRegArgs := func(name string, ids ...int) []instArg {
		return makeRegArgs(name, "%s%d", "%sMM%d", ids...)
	}

	makeUint8Args := func(values ...uint64) []instArg {
		args := make([]instArg, len(values))
		for i, v := range values {
			goSyntax := fmt.Sprintf("$%d", v)
			data := &imm{Width: 8, Value: v, Unsigned: true}
			args[i] = instArg{goSyntax, data}
		}
		return args
	}

	memoryListToArgs := func(width uint, list []*mem) []instArg {
		args := make([]instArg, len(list))
		for i, mem := range list {
			mem.Width = width
			args[i].goSyntax = memoryExpression(mem)
			args[i].data = mem
		}
		return args
	}

	makeMemArgs := func(width uint) []instArg {
		return memoryListToArgs(width, []*mem{
			{Base: "RSP", Disp: 17},
			{Base: "RBP", Index: "RSI", Scale: 4, Disp: -17},
			{Base: "RAX", Disp: 7},
			{Base: "RDI"},
			{Base: "R15", Index: "R15", Disp: 99},
			{Base: "RDX"},
			{Base: "RBP", Index: "RSI", Scale: 8, Disp: -17},
			{Base: "R15"},
			{Base: "RSI", Index: "RDI", Scale: 8, Disp: 7},
			{Base: "R14", Disp: -15},
			{Base: "RSI", Index: "RDI", Disp: 7},
			{Base: "RDX", Index: "RBX", Scale: 8, Disp: 15},
			{Base: "RDI", Index: "R8", Disp: -7},
			{Base: "RSP"},
			{Base: "RCX", Disp: -7},
			{Base: "RDX", Index: "RBX", Scale: 4, Disp: 15},
			{Base: "R15", Index: "R15", Scale: 8, Disp: 99},
			{Base: "RAX", Index: "RCX", Scale: 8, Disp: 7},
			{Base: "RAX"},
			{Base: "RSI", Disp: 7},
			{Base: "RBX"},
			{Base: "RBP", Index: "RSI", Disp: -17},
			{Base: "R8", Index: "R14", Scale: 4, Disp: 15},
			{Base: "RCX", Index: "RDX", Scale: 4, Disp: -7},
			{Base: "R8"},
			{Base: "RDX", Index: "RBX", Scale: 2, Disp: 15},
			{Base: "RSP", Index: "RBP", Disp: 17},
			{Base: "RCX", Index: "RDX", Scale: 8, Disp: -7},
			{Base: "RBP", Index: "RSI", Scale: 2, Disp: -17},
			{Base: "RAX", Index: "RCX", Scale: 2, Disp: 7},
			{Base: "R8", Index: "R14", Disp: 15},
			{Base: "R8", Index: "R14", Scale: 2, Disp: 15},
			{Base: "R14"},
			{Base: "RDI", Index: "R8", Scale: 8, Disp: -7},
			{Base: "R15", Index: "R15", Scale: 4, Disp: 99},
			{Base: "RDX", Disp: 15},
			{Base: "RCX"},
			{Base: "R15", Disp: 99},
			{Base: "R15", Index: "R15", Scale: 2, Disp: 99},
			{Base: "RDI", Disp: -7},
			{Base: "RCX", Index: "RDX", Disp: -7},
			{Base: "R14", Index: "R15", Scale: 4, Disp: -15},
			{Base: "RDX", Index: "RBX", Disp: 15},
			{Base: "RCX", Index: "RDX", Scale: 2, Disp: -7},
			{Base: "RBP", Disp: -17},
			{Base: "R14", Index: "R15", Scale: 8, Disp: -15},
			{Base: "RSP", Index: "RBP", Scale: 2, Disp: 17},
			{Base: "RDI", Index: "R8", Scale: 4, Disp: -7},
			{Base: "R8", Disp: 15},
			{Base: "RBP"},
			{Base: "R8", Index: "R14", Scale: 8, Disp: 15},
			{Base: "R14", Index: "R15", Scale: 2, Disp: -15},
			{Base: "R14", Index: "R15", Disp: -15},
			{Base: "RBX", Disp: -15},
			{Base: "RAX", Index: "RCX", Scale: 4, Disp: 7},
			{Base: "RAX", Index: "RCX", Disp: 7},
			{Base: "RSI"},
			{Base: "RSI", Index: "RDI", Scale: 2, Disp: 7},
			{Base: "RSP", Index: "RBP", Scale: 8, Disp: 17},
			{Base: "RSP", Index: "RBP", Scale: 4, Disp: 17},
			{Base: "RSI", Index: "RDI", Scale: 4, Disp: 7},
			{Base: "RDI", Index: "R8", Scale: 2, Disp: -7},
		})
	}

	makeVMemXArgs := func(width uint) []instArg {
		return memoryListToArgs(width, []*mem{
			{Base: "RAX", Index: "XMM4"},
			{Base: "RBP", Index: "XMM10", Scale: 2},
			{Base: "R10", Index: "XMM29", Scale: 8},
			{Base: "RDX", Index: "XMM10", Scale: 4},
			{Base: "RSP", Index: "XMM4", Scale: 2},
			{Base: "R14", Index: "XMM29", Scale: 8},
		})
	}
	makeVMemYArgs := func(width uint) []instArg {
		return memoryListToArgs(width, []*mem{
			{Base: "RAX", Index: "YMM3"},
			{Base: "RBP", Index: "YMM9", Scale: 2},
			{Base: "R10", Index: "YMM28", Scale: 8},
			{Base: "RDX", Index: "YMM9", Scale: 4},
			{Base: "RSP", Index: "YMM3", Scale: 2},
			{Base: "R14", Index: "YMM28", Scale: 8},
		})
	}
	makeVMemZArgs := func(width uint) []instArg {
		return memoryListToArgs(width, []*mem{
			{Base: "RAX", Index: "ZMM9"},
			{Base: "RBP", Index: "ZMM12", Scale: 2},
			{Base: "R10", Index: "ZMM31", Scale: 8},
			{Base: "RDX", Index: "ZMM12", Scale: 4},
			{Base: "RSP", Index: "ZMM9", Scale: 2},
			{Base: "R14", Index: "ZMM31", Scale: 8},
		})
	}

	instArgsBySyntax = map[string][]instArg{
		// Skip broadcasts.
		// They are tested separately (as all other suffixes).
		"m32bcst": {},
		"m64bcst": {},

		// GPR args.
		"r32": {
			{"AX", &reg{Name: "EAX"}},
			{"R9", &reg{Name: "R9D"}},
			{"CX", &reg{Name: "ECX"}},
			{"SP", &reg{Name: "ESP"}},
			{"R14", &reg{Name: "R14D"}},
		},
		"r64": {
			{"DX", &reg{Name: "RDX"}},
			{"BP", &reg{Name: "RBP"}},
			{"R10", &reg{Name: "R10"}},
			{"CX", &reg{Name: "RCX"}},
			{"R9", &reg{Name: "R9"}},
			{"R13", &reg{Name: "R13"}},
		},

		// Vector registar range (block) args.
		"zmm+3": {
			{"[Z0-Z3]", &reg{Name: "ZMM0"}},
			{"[Z10-Z13]", &reg{Name: "ZMM10"}},
			{"[Z20-Z23]", &reg{Name: "ZMM20"}},
			{"[Z1-Z4]", &reg{Name: "ZMM1"}},
			{"[Z11-Z14]", &reg{Name: "ZMM11"}},
			{"[Z21-Z24]", &reg{Name: "ZMM21"}},
			{"[Z2-Z5]", &reg{Name: "ZMM2"}},
			{"[Z12-Z15]", &reg{Name: "ZMM12"}},
			{"[Z22-Z25]", &reg{Name: "ZMM22"}},
			{"[Z4-Z7]", &reg{Name: "ZMM4"}},
			{"[Z14-Z17]", &reg{Name: "ZMM14"}},
			{"[Z24-Z27]", &reg{Name: "ZMM24"}},
		},
		"xmm+3": {
			{"[X0-X3]", &reg{"XMM0"}},
			{"[X10-X13]", &reg{"XMM10"}},
			{"[X20-X23]", &reg{"XMM20"}},
			{"[X1-X4]", &reg{"XMM1"}},
			{"[X11-X14]", &reg{"XMM11"}},
			{"[X21-X24]", &reg{"XMM21"}},
			{"[X2-X5]", &reg{"XMM2"}},
			{"[X12-X15]", &reg{"XMM12"}},
			{"[X22-X25]", &reg{"XMM22"}},
			{"[X4-X7]", &reg{"XMM4"}},
			{"[X14-X17]", &reg{"XMM14"}},
			{"[X24-X27]", &reg{"XMM24"}},
		},

		// K operand for KOP instructions.
		"k": makeMaskRegArgs(7, 0, 1, 2, 3, 4, 5, 6),
		// K operand for write masks. Can't be K0.
		"{k}": makeMaskRegArgs(7, 1, 2, 3, 4, 5, 6),

		// Immediate args.
		"imm8u:1": makeUint8Args(1, 0),
		"imm8u:2": makeUint8Args(3, 0, 1, 2),
		"imm8u:4": makeUint8Args(
			15,
			0,
			1,
			2,
			3,
			4,
			5,
			6,
			7,
			8,
			9,
			10,
			11,
			12,
			13,
			14,
		),
		"imm8u": makeUint8Args(
			255,
			0,
			97,
			81,
			42,
			79,
			64,
			27,
			47,
			82,
			126,
			94,
			121,
			13,
			65,
			67,
		),

		// Memory args.
		"m8":   makeMemArgs(8),
		"m16":  makeMemArgs(16),
		"m32":  makeMemArgs(32),
		"m64":  makeMemArgs(64),
		"m128": makeMemArgs(128),
		"m256": makeMemArgs(256),
		"m512": makeMemArgs(512),

		// VMem args.
		"vmx:32": makeVMemXArgs(32),
		"vmx:64": makeVMemXArgs(64),
		"vmy:8":  makeVMemYArgs(8),
		"vmy:32": makeVMemYArgs(32),
		"vmy:64": makeVMemYArgs(64),
		"vmz:8":  makeVMemZArgs(8),
		"vmz:32": makeVMemZArgs(32),
		"vmz:64": makeVMemZArgs(64),

		// Vector register args.
		"xmm": makeVecRegArgs("X",
			22, 30, 3,
			11, 15, 30,
			13, 6, 12,
			23, 30, 8,
			20, 2, 9,
			26, 19, 0,
			31, 16, 7,
			8, 1, 0,
			15, 0, 16,
			21, 0, 28,
			22, 7, 19,
			7, 16, 31,
			1, 7, 9,
			15, 12, 0,
			12, 14, 5,
			17, 15, 8,
			3, 26, 23,
			13, 28, 24,
			9, 15, 26,
			18, 21, 1,
			11, 31, 3,
			7, 0, 0,
			24, 20, 7,
			9, 7, 14,
			5, 31, 3,
			21, 1, 11,
			13, 0, 30,
			16, 14, 11,
			14, 19, 8,
			8, 26, 23,
			12, 16, 23,
			23, 11, 31,
			24, 14, 0,
			11, 23, 2,
			20, 5, 25,
			0, 9, 13,
			2, 8, 9,
			2, 31, 11,
			22, 5, 14,
			0, 17, 7,
			15, 11, 0,
			18, 8, 27,
			25, 3, 18,
			15, 28, 15,
			7, 13, 8,
			24, 7, 0,
			22, 1, 11,
			6, 7, 8,
			31, 3, 28,
			20, 24, 7,
			20, 16, 12,
			6, 17, 28,
			6, 1, 8,
			8, 6, 0,
			11, 16, 6,
			6, 22, 12,
			16, 28, 8,
			15, 11, 1,
			19, 13, 2,
			14, 0, 0,
			25, 11, 17,
			18, 11, 9,
			2, 24, 2,
			2, 27, 26,
		),
		"ymm": makeVecRegArgs("Y",
			14, 31, 25,
			2, 22, 27,
			8, 9, 22,
			9, 14, 1,
			6, 1, 9,
			0, 19, 31,
			22, 9, 23,
			31, 5, 0,
			5, 19, 31,
			28, 2, 24,
			27, 0, 11,
			31, 3, 14,
			2, 13, 27,
			15, 22, 20,
			18, 24, 9,
			3, 19, 23,
			19, 14, 21,
			5, 16, 2,
			21, 20, 6,
			31, 6, 11,
			19, 7, 6,
			0, 3, 5,
			20, 12, 3,
			5, 28, 7,
			0, 22, 13,
			12, 1, 14,
			17, 7, 9,
			31, 8, 1,
			28, 13, 7,
			2, 21, 12,
			9, 1, 9,
			3, 2, 9,
			12, 21, 14,
			30, 26, 7,
			16, 1, 30,
			31, 22, 6,
			21, 7, 0,
			28, 20, 14,
			24, 13, 20,
			14, 21, 1,
			26, 30, 12,
			22, 3, 15,
			1, 27, 19,
			5, 17, 13,
			21, 7, 30,
			13, 18, 24,
			8, 11, 24,
			5, 24, 21,
			16, 9, 13,
			9, 6, 3,
			7, 6, 26,
			11, 26, 12,
			14, 18, 31,
			18, 3, 24,
			2, 7, 21,
			14, 8, 20,
			11, 24, 1,
			5, 18, 20,
			20, 9, 28,
			28, 1, 8,
			11, 27, 17,
			16, 12, 6,
			26, 3, 8,
			28, 1, 23,
		),
		"zmm": makeVecRegArgs("Z",
			0, 8,
			15, 12,
			14, 27,
			11, 5,
			13, 14,
			5, 23,
			2, 2,
			6, 14,
			26, 14,
			28, 6,
			13, 21,
			26, 3,
			3, 0,
			21, 13,
			11, 25,
			3, 12,
			27, 15,
			23, 5,
			23, 6,
			8, 28,
			21, 5,
			16, 13,
			12, 27,
			22, 11,
			6, 8,
			25, 12,
			12, 17,
			9, 12,
			6, 25,
			3, 21,
			8, 2,
			3, 27,
			7, 9,
			0, 6,
			20, 28,
			3, 30,
			9, 19,
			12, 22,
			11, 5,
			18, 24,
			2, 21,
			7, 13,
			6, 16,
			6, 22,
			1, 15,
			13, 13,
			18, 8,
			22, 7,
			2, 31,
			20, 9,
			1, 3,
			12, 16,
			28, 13,
			14, 28,
			3, 12,
			15, 30,
			19, 15,
			5, 1,
			3, 5,
			14, 15,
			21, 8,
			16, 9,
			20, 0,
			23, 19,
			0, 11,
			0, 25,
			24, 12,
			0, 26,
			9, 3,
			9, 25,
			9, 28,
			20, 0,
			17, 0,
			17, 23,
			31, 0,
			21, 9,
			6, 9,
			1, 9,
			20, 9,
			30, 5,
			26, 22,
			7, 21,
			16, 25,
			14, 13,
			12, 13,
			21, 9,
			2, 7,
			27, 25,
			23, 9,
			27, 14,
			3, 0,
			14, 7,
			8, 24,
			22, 25,
			1, 16,
			6, 2,
		),

		// End of instArgsBySyntax literal.
	}
}

// peeksPerArgBySyntax define how many forms from available args list we take
// for a single instruction test.
//
// Higher numbers increase output test suite size significantly.
var peeksPerArgBySyntax = map[string]int{
	"m8":      2,
	"m16":     2,
	"m32":     2,
	"m64":     2,
	"m128":    2,
	"m256":    2,
	"m512":    2,
	"xmm":     1,
	"ymm":     1,
	"zmm":     2,
	"k":       2,
	"{k}":     1,
	"imm8u":   1,
	"imm8u:1": 1,
	"imm8u:2": 1,
	"imm8u:4": 1,
	"r32":     2,
	"r64":     2,

	"vmx:32": 3,
	"vmx:64": 3,
	"vmy:8":  3,
	"vmy:32": 3,
	"vmy:64": 3,
	"vmz:8":  3,
	"vmz:32": 3,
	"vmz:64": 3,

	"zmm+3": 3,
	"xmm+3": 3,

	"m32bcst": 0,
	"m64bcst": 0,
}

// argReplacer is used to erase/replace arguments before parsing them.
var argReplacer = strings.NewReplacer(
	"{sae}", "",
	"{er}", "",
)

// argNormalizeMap replaces x86csv-style args to a form that can be used
// to access args table.
var argNormalizeMap = map[string]string{
	"rmr32": "r32",
	"rmr64": "r64",

	"xmm1":   "xmm",
	"xmm2":   "xmm",
	"xmmV":   "xmm",
	"xmmV+3": "xmm+3",
	"xmmIH":  "xmm",

	"ymm1":  "ymm",
	"ymm2":  "ymm",
	"ymmV":  "ymm",
	"ymmIH": "ymm",

	"zmm1":   "zmm",
	"zmm2":   "zmm",
	"zmmV":   "zmm",
	"zmmV+3": "zmm+3",

	"k1": "k",
	"kV": "k",
	"k2": "k",

	"vm32x": "vmx",
	"vm64x": "vmx",
	"vm32y": "vmy",
	"vm64y": "vmy",
	"vm64z": "vmz",
	"vm32z": "vmz",
}

// vmemWidths maps Intel opcode to VMem width.
// VMem memory size can't be inferred from vm32/vm64 alone.
var vmemWidths = map[string]uint16{
	"VSCATTERQPD":    64,
	"VGATHERDPD":     64,
	"VGATHERQPD":     64,
	"VPGATHERDQ":     64,
	"VPGATHERQQ":     64,
	"VPSCATTERDQ":    64,
	"VPSCATTERQQ":    64,
	"VSCATTERDPD":    64,
	"VGATHERDPS":     32,
	"VGATHERQPS":     32,
	"VPGATHERDD":     32,
	"VPGATHERQD":     32,
	"VPSCATTERDD":    32,
	"VPSCATTERQD":    32,
	"VSCATTERDPS":    32,
	"VSCATTERQPS":    32,
	"VGATHERPF0DPD":  8,
	"VGATHERPF0DPS":  8,
	"VGATHERPF0QPD":  8,
	"VGATHERPF0QPS":  8,
	"VGATHERPF1DPD":  8,
	"VGATHERPF1DPS":  8,
	"VGATHERPF1QPD":  8,
	"VGATHERPF1QPS":  8,
	"VSCATTERPF0DPD": 8,
	"VSCATTERPF0DPS": 8,
	"VSCATTERPF0QPD": 8,
	"VSCATTERPF0QPS": 8,
	"VSCATTERPF1DPD": 8,
	"VSCATTERPF1DPS": 8,
	"VSCATTERPF1QPD": 8,
	"VSCATTERPF1QPS": 8,
}

func normalizeArg(inst *x86csv.Inst, arg string) string {
	arg = argReplacer.Replace(arg)
	if normalized := argNormalizeMap[arg]; normalized != "" {
		switch normalized {
		case "vmx", "vmy", "vmz":
			if width := vmemWidths[inst.IntelOpcode()]; width != 0 {
				return fmt.Sprintf("%s:%d", normalized, width)
			}
		}
		return normalized
	}

	return arg
}

func argsCartesianProd(args [][]instArg) (c [][]instArg) {
	if len(args) == 0 {
		return [][]instArg{nil}
	}
	c2 := argsCartesianProd(args[1:])
	for _, arg := range args[0] {
		for _, rest := range c2 {
			c = append(c, append([]instArg{arg}, rest...))
		}
	}
	return c
}
