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
		inst.Imm = int32(instr >> 20)

	case OpIType:
		inst.Rd = (instr >> 7) & 0x1F
		inst.Funct3 = (instr >> 12) & 0x7
		inst.Rs1 = (instr >> 15) & 0x1F
		inst.Imm = int32(instr >> 20)

	case OpLoad:
		inst.Rd = (instr >> 7) & 0x1F    //destination register
		inst.Funct3 = (instr >> 12) & 0x7 //load type
		inst.Rs1 = (instr >> 15) & 0x1F  //base address registter
		inst.Imm = int32(instr >> 20) //address offset
	case OpStore:
		imm115 := (instr >> 25) & 0x7F  // upper 7 bits of immediate
		imm40 := (instr >> 7) & 0x1F    //lower 5 bit of immdediate
		inst.Funct3 = (instr >> 12) & 0x7
		inst.Rs1 = (instr >> 15) & 0x1F  //base address register
		inst.Rs2 = (instr >> 20) & 0x1F  //register carrying value to store
		inst.Imm = int32((imm115 << 5) | imm40) //offset
	}
	return inst
}
