package x86encode

import (
	"testing"
)

func TestEncode(t *testing.T) {
	type reg = RegArgument
	type imm = ImmArgument
	type mem = MemArgument

	tests := []struct {
		inst Inst
		want string
	}{
		{Inst{Opcode: "NOP"}, "90"},

		{
			Inst{
				Opcode: "INC",
				Args: []Argument{
					&reg{Name: "EAX"},
				},
			},
			"ffc0",
		},

		{
			Inst{
				Opcode: "ADD",
				Args: []Argument{
					&reg{Name: "EAX"},
					&imm{Value: 0x10, Width: 8},
				},
			},
			"83c010",
		},

		{
			Inst{
				Opcode: "VAESDEC",
				Args: []Argument{
					&reg{Name: "XMM11"},
					&reg{Name: "XMM12"},
					&mem{Base: "EDX", Width: 128},
				},
			},
			"67c46219de1a",
		},

		{
			Inst{
				Opcode: "VADDPD",
				Params: []InstParam{ParamVexL256},
				Args: []Argument{
					&reg{Name: "YMM0"},
					&reg{Name: "K3"},
					&reg{Name: "YMM5"},
					&reg{Name: "YMM22"},
				},
			},
			"62b1d52b58c6",
		},

		{
			Inst{
				Opcode: "VANDPD",
				Params: []InstParam{ParamRexW1},
				Args: []Argument{
					&reg{Name: "XMM0"},
					&reg{Name: "K3"},
					&reg{Name: "XMM5"},
					&reg{Name: "XMM22"},
				},
			},
			"62b1d50b54c6",
		},

		{
			Inst{
				Opcode: "KNOTQ",
				Params: []InstParam{ParamRexW1},
				Args: []Argument{
					&reg{Name: "K1"},
					&reg{Name: "K1"},
				},
			},
			"c4e1f844c9",
		},

		{
			Inst{
				Opcode: "KORD",
				Params: []InstParam{ParamRexW1, ParamVexL256},
				Args: []Argument{
					&reg{Name: "K6"},
					&reg{Name: "K1"},
					&reg{Name: "K3"},
				},
			},
			"c4e1f545f3",
		},
	}

	for _, test := range tests {
		have, err := ToHexString(&test.inst)
		if err != nil {
			t.Errorf("encoding failed: %v", err)
			continue
		}
		if have != test.want {
			t.Errorf("encoding result mismatch:\nhave: %q\nwant: %q",
				have, test.want)
		}
	}
}
