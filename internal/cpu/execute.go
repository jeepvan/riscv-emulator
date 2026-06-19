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
			int32(inst.Imm),
		)
	case OpLoad:
		cpu.ExecuteLW(
			inst.Rd,
			inst.Rs1,
			int32(inst.Imm),
		)
	case OpStore:
		cpu.ExecuteSW(
			inst.Rs1,
			inst.Rs2,
			int32(inst.Imm),
		)
	case OpBranch:
		switch inst.Funct3 {
		case 0b000:
			cpu.ExecuteBEQ(
				inst.Rs1,
				inst.Rs2,
				int32(inst.Imm),
			)
		case 0b001:
			cpu.ExecuteBNE(
				inst.Rs1,
				inst.Rs2,
				int32(inst.Imm),
			)
		}

	}
	cpu.Regs[0] = 0
}
func (cpu *CPU) ExecuteADDI(rd, rs1 uint32, imm int32) {
	cpu.Regs[rd] = uint32(int32(cpu.Regs[rs1]) + imm)
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
func (cpu *CPU) ExecuteLW(rd, rs1 uint32, imm int32) {
	addr := uint32(int32(cpu.Regs[rs1]) + imm)
	cpu.Regs[rd] = cpu.Read32(addr)
}

func (cpu *CPU) ExecuteSW(rs1, rs2 uint32, imm int32) {
	addr := uint32(int32(cpu.Regs[rs1]) + imm)
	value := cpu.Regs[rs2]
	cpu.Write32(addr, value)
}

func (cpu *CPU) ExecuteBEQ(rs1, rs2 uint32, imm int32) {
	if cpu.Regs[rs1] == cpu.Regs[rs2] {
		cpu.PC = uint32(int32(cpu.PC) + imm)
	}
}
func (cpu *CPU) ExecuteBNE(rs1, rs2 uint32, imm int32) {
	if cpu.Regs[rs1] != cpu.Regs[rs2] {
		cpu.PC = uint32(int32(cpu.PC) + imm)
	}
}
