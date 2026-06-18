package cpu

func (cpu *CPU) Fetch() uint32 {
	b0 := cpu.Mem[cpu.PC]
	b1 := cpu.Mem[cpu.PC+1]
	b2 := cpu.Mem[cpu.PC+2]
	b3 := cpu.Mem[cpu.PC+3]

	instr := uint32(b0) | uint32(b1)<<8 | uint32(b2)<<16 | uint32(b3)<<24
	return instr
}
