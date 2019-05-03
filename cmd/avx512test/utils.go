package main

import (
	"strings"

	"github.com/quasilyte/avx512test/internal/x86encode"
	"golang.org/x/arch/x86/x86csv"
)

func goAsmString(inst *x86csv.Inst, args []instArg) string {
	op := inst.GoOpcode()
	if len(args) == 0 {
		return op
	}

	// Collect args in reverse.
	goArgs := make([]string, 0, len(args))
	for i := len(args) - 1; i >= 0; i-- {
		goArgs = append(goArgs, args[i].goSyntax)
	}

	return op + " " + strings.Join(goArgs, ", ")
}

func instREXW(inst *x86csv.Inst) []x86encode.InstParam {
	switch {
	case strings.Contains(inst.Encoding, ".WIG"):
		return []x86encode.InstParam{x86encode.ParamRexW0, x86encode.ParamRexW1}
	case strings.Contains(inst.Encoding, ".W1"):
		return []x86encode.InstParam{x86encode.ParamRexW1}
	default:
		return []x86encode.InstParam{x86encode.ParamRexW0}
	}
}

func instVL(inst *x86csv.Inst) []x86encode.InstParam {
	switch {
	case strings.Contains(inst.Encoding, ".LIG"):
		if evexEncoded(inst) {
			return []x86encode.InstParam{
				x86encode.ParamVexL128,
				x86encode.ParamVexL256,
				x86encode.ParamVexL512,
			}
		}
		return []x86encode.InstParam{x86encode.ParamVexL128, x86encode.ParamVexL256}
	case strings.Contains(inst.Encoding, ".512"):
		return []x86encode.InstParam{x86encode.ParamVexL512}
	case strings.Contains(inst.Encoding, ".256"):
		return []x86encode.InstParam{x86encode.ParamVexL256}
	default:
		return []x86encode.InstParam{x86encode.ParamVexL128}
	}
}

func evexEncoded(inst *x86csv.Inst) bool {
	return strings.HasPrefix(inst.Encoding, "EVEX")
}

func normalizeCPUID(cpuid string) string {
	cpuid = strings.Replace(cpuid, "+AVX512VL", "", 1)
	cpuid = strings.Replace(cpuid, "+AVX512F", "", 1)
	return cpuid
}
