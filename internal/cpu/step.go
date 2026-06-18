package cpu

func (cpu *CPU) Step() {
	instr := cpu.Fetch()

	inst := Decode(instr)
	cpu.Execute(inst)
	cpu.PC += 4

}
