package cpu

type Instruction struct {
	Opcode uint32

	Rd  uint32
	Rs1 uint32
	Rs2 uint32

	Funct3 uint32
	Funct7 uint32

	Imm int32
}
