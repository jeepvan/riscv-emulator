package main

import "fmt"

type CPU struct {
	PC   uint32
	Regs [32]uint32
	Mem  []byte
}

func main() {
	cpu := CPU{
		Mem: []byte{
			0x93, 0x00, 0xA0, 0x00,
			0x13, 0x01, 0x40, 0x01,
			0xB3, 0x81, 0x20, 0x00,
		},
	}
	cpu.PC = 0
	for cpu.PC < uint32(len(cpu.Mem)) {
		cpu.Step()
	}
	cpu.DumpRegisters()
}

func (cpu *CPU) DumpRegisters() {
	for i := 0; i < 8; i++ {
		fmt.Printf("x%d=%d\n", i, cpu.Regs[i])
	}
	fmt.Println()
}

func (cpu *CPU) DecodeandExecute(intr uint32) {
	opcode := intr & 0x7F
	rd := (intr >> 7) & 0x1F
	funct3 := (intr >> 12) & 0x7
	rs1 := (intr >> 15) & 0x1F

	switch opcode {

	case 0x33:

		rs2 := (intr >> 20) & 0x1F
		funct7 := (intr >> 25) & 0x7F

		switch funct3 {
		case 0:
			switch funct7 {
			case 0:
				cpu.ExecuteADD(rd, rs1, rs2)
			case 0x20:
				cpu.ExecuteSUB(rd, rs1, rs2)
			}
		case 0b111:
			cpu.ExecuteAND(rd, rs1, rs2)
		case 0b110:
			cpu.ExecuteOR(rd, rs1, rs2)
		case 0b100:
			cpu.ExecuteXOR(rd, rs1, rs2)
		}

	case 0x13:
		imm := (intr >> 20)

		cpu.ExecuteADDI(rd, rs1, imm)

	}
	cpu.Regs[0] = 0

}
func (cpu *CPU) Step() {
	intr := cpu.Fetch()

	fmt.Printf(
		"PC=0x%08X INSTR=0x%08X\n",
		cpu.PC,
		intr,
	)

	cpu.DecodeandExecute(intr)
	cpu.PC += 4

}
func (cpu *CPU) Fetch() uint32 {
	b0 := cpu.Mem[cpu.PC]
	b1 := cpu.Mem[cpu.PC+1]
	b2 := cpu.Mem[cpu.PC+2]
	b3 := cpu.Mem[cpu.PC+3]

	instr := uint32(b0) | uint32(b1)<<8 | uint32(b2)<<16 | uint32(b3)<<24
	return instr
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
