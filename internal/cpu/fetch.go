package cpu

func (cpu *CPU) Fetch() uint32 {
    return cpu.Read32(cpu.PC)
}
