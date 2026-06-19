package cpu

func Decode(instr uint32) Instruction {
	inst := Instruction{}
	inst.Opcode = instr & 0x7F

	switch inst.Opcode {
	case OpRType:
		inst.Rd = (instr >> 7) & 0x1F
		inst.Funct3 = (instr >> 12) & 0x7
		inst.Rs1 = (instr >> 15) & 0x1F
		inst.Rs2 = (instr >> 20) & 0x1F
		inst.Funct7 = (instr >> 25) & 0x7F

	case OpIType:
		inst.Rd = (instr >> 7) & 0x1F
		inst.Funct3 = (instr >> 12) & 0x7
		inst.Rs1 = (instr >> 15) & 0x1F
		imm := (instr >> 20) & 0xFFF
		inst.Imm = SignExtend(imm, 12)

	case OpLoad:
		inst.Rd = (instr >> 7) & 0x1F     //destination register
		inst.Funct3 = (instr >> 12) & 0x7 //load type
		inst.Rs1 = (instr >> 15) & 0x1F   //base address registter
		imm := (instr >> 20) & 0xFFF
		inst.Imm = SignExtend(imm, 12) //address offset
	case OpStore:
		imm115 := (instr >> 25) & 0x7F // upper 7 bits of immediate
		imm40 := (instr >> 7) & 0x1F   //lower 5 bit of immdediate
		inst.Funct3 = (instr >> 12) & 0x7
		inst.Rs1 = (instr >> 15) & 0x1F //base address register
		inst.Rs2 = (instr >> 20) & 0x1F //register carrying value to store
		imm := (imm115 << 5) | imm40
		inst.Imm = SignExtend(imm, 12) //offset
	case OpBranch:
		inst.Funct3 = (instr >> 12) & 0x7
		inst.Rs1 = (instr >> 15) & 0x1F
		inst.Rs2 = (instr >> 20) & 0x1F
		imm12 := (instr >> 31) & 0x1
		imm105 := (instr >> 25) & 0x3F
		imm41 := (instr >> 8) & 0xF
		imm11 := (instr >> 7) & 0x1
		imm := (imm12 << 12) | (imm105 << 5) | (imm41 << 1) | (imm11 << 11)
		inst.Imm = SignExtend(imm, 13)
	case OPJAL:
		imm20 := (instr >> 31) & 0x1
		imm101 := (instr >> 21) & 0x3FFF
		imm11 := (instr >> 11) & 0x1
		imm1912 := (instr >> 7) & 0xFF

		imm := (imm20 << 20) | (imm1912 << 12) | (imm11 << 11) | (imm101 << 1)
		inst.Imm = SignExtend(imm, 21)
	}

	return inst
}

func SignExtend(value uint32, bits int) int32 {
	shift := 32 - bits
	return int32(value<<shift) >> shift
}
