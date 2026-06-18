package cpu

func (cpu *CPU) Execute(inst Instruction) {
	switch inst.Opcode {

	case OpRType:

		switch inst.Funct3 {

		case 0:

			switch inst.Funct7 {

			case 0:
				cpu.ExecuteADD(
					inst.Rd,
					inst.Rs1,
					inst.Rs2,
				)

			case 0x20:
				cpu.ExecuteSUB(
					inst.Rd,
					inst.Rs1,
					inst.Rs2,
				)
			}
		}
	case OpIType:
		cpu.ExecuteADDI(
			inst.Rd,
			inst.Rs1,
			uint32(inst.Imm),
		)

	}
	cpu.Regs[0] = 0
}
func (cpu *CPU) ExecuteADDI(rd, rs1, imm uint32) {
	cpu.Regs[rd] = cpu.Regs[rs1] + imm

}
func (cpu *CPU) ExecuteADD(rd, rs1, rs2 uint32) {
	cpu.Regs[rd] = cpu.Regs[rs1] + cpu.Regs[rs2]

}
func (cpu *CPU) ExecuteSUB(rd, rs1, rs2 uint32) {
	cpu.Regs[rd] = cpu.Regs[rs1] - cpu.Regs[rs2]

}
func (cpu *CPU) ExecuteAND(rd, rs1, rs2 uint32) {
	cpu.Regs[rd] = cpu.Regs[rs1] & cpu.Regs[rs2]

}
func (cpu *CPU) ExecuteOR(rd, rs1, rs2 uint32) {
	cpu.Regs[rd] = cpu.Regs[rs1] | cpu.Regs[rs2]

}
func (cpu *CPU) ExecuteXOR(rd, rs1, rs2 uint32) {
	cpu.Regs[rd] = cpu.Regs[rs1] ^ cpu.Regs[rs2]

}
