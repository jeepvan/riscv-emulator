package cpu

func Decode(instr uint32) Instruction {
	return Instruction{
		Opcode: instr & 0x7F,
		Rd:     (instr >> 7) & 0x1F,
		Funct3: (instr >> 12) & 0x7,
		Rs1:    (instr >> 15) & 0x1F,
		Rs2:    (instr >> 20) & 0x1F,
		Funct7: (instr >> 25) & 0x7F,
		Imm:    int32(instr >> 20),
	}
}

// func (cpu *CPU) DecodeandExecute(intr uint32) {
// 	opcode := intr & 0x7F
// 	rd := (intr >> 7) & 0x1F
// 	funct3 := (intr >> 12) & 0x7
// 	rs1 := (intr >> 15) & 0x1F

// 	switch opcode {

// 	case 0x33:

// 		rs2 := (intr >> 20) & 0x1F
// 		funct7 := (intr >> 25) & 0x7F

// 		switch funct3 {
// 		case 0:
// 			switch funct7 {
// 			case 0:
// 				cpu.ExecuteADD(rd, rs1, rs2)
// 			case 0x20:
// 				cpu.ExecuteSUB(rd, rs1, rs2)
// 			}
// 		case 0b111:
// 			cpu.ExecuteAND(rd, rs1, rs2)
// 		case 0b110:
// 			cpu.ExecuteOR(rd, rs1, rs2)
// 		case 0b100:
// 			cpu.ExecuteXOR(rd, rs1, rs2)
// 		}

// 	case 0x13:
// 		imm := (intr >> 20)

// 		cpu.ExecuteADDI(rd, rs1, imm)

// 	}
// 	cpu.Regs[0] = 0

// }
