package cpu

func (cpu *CPU) Read8(addr uint32) byte {
	return cpu.Mem[addr]
}

func (cpu *CPU) Write8(addr uint32, value byte) {
	cpu.Mem[addr] = value
}
func (cpu *CPU) Read32(addr uint32) uint32 {
	b0 := cpu.Read8(addr)
	b1 := cpu.Read8(addr + 1)
	b2 := cpu.Read8(addr + 2)
	b3 := cpu.Read8(addr + 3)

	instr := uint32(b0) | uint32(b1)<<8 | uint32(b2)<<16 | uint32(b3)<<24
	return instr
}

func (cpu *CPU) Write32(addr uint32, value uint32) {
	b0 := byte(value)
	b1 := byte(value >> 8)
	b2 := byte(value >> 16)
	b3 := byte(value >> 24)

	cpu.Write8(addr, b0)
	cpu.Write8(addr+1, b1)
	cpu.Write8(addr+2, b2)
	cpu.Write8(addr+3, b3)

}
