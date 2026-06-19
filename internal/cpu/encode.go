package cpu

func EncodeADDI(rd, rs1 uint32, imm uint32) uint32 {
	return (imm << 20) |
		(rs1 << 15) |
		(0 << 12) |
		(rd << 7) |
		0x13
}

func LoadProgram(cpu *CPU, program []uint32) {
	for i, instr := range program {
		cpu.Write32(uint32(i*4), instr)
	}
}
