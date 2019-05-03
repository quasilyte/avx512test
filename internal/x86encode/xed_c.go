package x86encode

// #cgo LDFLAGS: -lxed
// #include "xed.h"
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

var xedState C.xed_state_t

func init() {
	C.xed_state_zero(&xedState)
	xedState.stack_addr_width = C.XED_ADDRESS_WIDTH_64b
	xedState.mmode = C.XED_MACHINE_MODE_LONG_64
}

func xedTablesInit() { C.xed_tables_init() }

func xedEncode(inst *Inst) ([]byte, error) {
	const bufCapacity = C.XED_MAX_INSTRUCTION_BYTES
	buf := make([]byte, bufCapacity)

	var enc C.xed_encoder_instruction_t
	eosz := C.xed_uint_t(32) // Default

	iclass := C.xed_iclass_enum_t(iclassByOpcode[inst.Opcode])
	if iclass == 0 {
		panic(fmt.Sprintf("no iclass found for %q", inst.Opcode))
	}

	for _, param := range inst.Params {
		switch param {
		case ParamEOSZ8:
			eosz = 8
		case ParamEOSZ16:
			eosz = 16
		case ParamEOSZ32:
			eosz = 32
		case ParamEOSZ64:
			eosz = 64
		}
	}

	operands, err := xedConvertOperands(inst)
	if err != nil {
		return nil, err
	}

	switch len(inst.Args) {
	case 0:
		C.xed_inst0(&enc, xedState, iclass, eosz)
	case 1:
		C.xed_inst1(&enc, xedState, iclass, eosz,
			operands[0])
	case 2:
		C.xed_inst2(&enc, xedState, iclass, eosz,
			operands[0],
			operands[1])
	case 3:
		C.xed_inst3(&enc, xedState, iclass, eosz,
			operands[0],
			operands[1],
			operands[2])
	case 4:
		C.xed_inst4(&enc, xedState, iclass, eosz,
			operands[0],
			operands[1],
			operands[2],
			operands[3])
	case 5:
		C.xed_inst5(&enc, xedState, iclass, eosz,
			operands[0],
			operands[1],
			operands[2],
			operands[3],
			operands[4])
	default:
		return nil, fmt.Errorf("unexpected number of args: %d", len(inst.Args))
	}

	var req C.xed_encoder_request_t
	C.xed_encoder_request_zero_set_mode(&req, &enc.mode)
	ok := C.xed_convert_to_encoder_request(&req, &enc)
	if ok == 0 {
		return nil, errors.New("encoder request conversion failed")
	}

	for _, param := range inst.Params {
		switch param {
		case ParamRexW0:
			C.xed3_operand_set_rexw(&req, C.xed_bits_t(0))
		case ParamRexW1:
			C.xed3_operand_set_rexw(&req, C.xed_bits_t(1))

		case ParamVexL128:
			C.xed3_operand_set_vl(&req, C.xed_bits_t(0))
		case ParamVexL256:
			C.xed3_operand_set_vl(&req, C.xed_bits_t(1))
		case ParamVexL512:
			C.xed3_operand_set_vl(&req, C.xed_bits_t(2))
		}
	}

	codeLen := C.uint(0)
	errCode := C.xed_encode(
		&req,
		(*C.uint8_t)(unsafe.Pointer(&buf[0])),
		C.uint(cap(buf)),
		&codeLen,
	)
	if errCode != C.XED_ERROR_NONE {
		return nil, fmt.Errorf("xed error: %s", xedErrCodeToString(errCode))
	}

	return buf[:codeLen], nil
}

func xedConvertOperands(inst *Inst) ([]C.xed_encoder_operand_t, error) {
	operands := make([]C.xed_encoder_operand_t, len(inst.Args))
	for i, arg := range inst.Args {
		op, err := xedOperand(arg)
		if err != nil {
			return nil, fmt.Errorf("error in Args[%d]: %v", i, err)
		}
		operands[i] = op
	}
	return operands, nil
}

func xedOperand(arg Argument) (C.xed_encoder_operand_t, error) {
	var invalid C.xed_encoder_operand_t

	switch arg := arg.(type) {
	case *RegArgument:
		return C.xed_reg(C.xed_reg_enum_t(registerByName[arg.Name])), nil

	case *ImmArgument:
		switch {
		case arg.Width == 8 && arg.Unsigned:
			return C.xed_imm0(C.xed_uint64_t(arg.Value), 8), nil
		case arg.Width == 16 && arg.Unsigned:
			return C.xed_imm0(C.xed_uint64_t(arg.Value), 16), nil
		case arg.Width == 32 && arg.Unsigned:
			return C.xed_imm0(C.xed_uint64_t(arg.Value), 32), nil
		case arg.Width == 64 && arg.Unsigned:
			return C.xed_imm0(C.xed_uint64_t(arg.Value), 64), nil
		case arg.Width == 8 && !arg.Unsigned:
			return C.xed_simm0(C.xed_int32_t(arg.Value), 8), nil
		case arg.Width == 16 && !arg.Unsigned:
			return C.xed_simm0(C.xed_int32_t(arg.Value), 16), nil
		case arg.Width == 32 && !arg.Unsigned:
			return C.xed_simm0(C.xed_int32_t(arg.Value), 32), nil
		default:
			return invalid, errors.New("bad width/signedness combination for immediate")
		}

	case *MemArgument:
		var disp C.xed_enc_displacement_t
		disp.displacement = C.xed_uint64_t(arg.Disp)
		switch arg.DispWidth {
		case Disp8:
			disp.displacement_bits = 8
		case Disp32:
			disp.displacement_bits = 32
		case DispSmallest:
			switch {
			case arg.Disp == 0:
				disp.displacement_bits = 0
			case arg.Disp >= -128 && arg.Disp <= 127:
				disp.displacement_bits = 8
			default:
				disp.displacement_bits = 32
			}
		default:
			return invalid, fmt.Errorf("invalid memory argument disp width: %d", arg.DispWidth)
		}

		scale := C.xed_uint_t(arg.Scale)
		switch arg.Scale {
		case 0:
			scale = 1 // Default
		case 1, 2, 4, 8:
			// OK.
		default:
			return invalid, fmt.Errorf("invalid memory argument scale: %d", arg.Scale)
		}

		base := C.xed_reg_enum_t(registerByName[arg.Base])
		index := C.xed_reg_enum_t(registerByName[arg.Index])
		bitSize := C.xed_uint_t(arg.Width)
		return C.xed_mem_bisd(base, index, scale, disp, bitSize), nil

	default:
		return invalid, fmt.Errorf("invalid argument type: %T", arg)
	}
}

func xedErrCodeToString(errCode C.xed_error_enum_t) string {
	return C.GoString(C.xed_error_enum_t2str(errCode))
}
