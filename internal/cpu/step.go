package cpu

func (cpu *CPU) Step() {
	instr := cpu.Fetch()
	oldPC := cpu.PC
	inst := Decode(instr)
	cpu.Execute(inst)
	if oldPC == cpu.PC {
		cpu.PC += 4
	}

}
