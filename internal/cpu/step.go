package cpu

import "fmt"

func (cpu *CPU) Step() {
	instr := cpu.Fetch()

	fmt.Printf("instr=%08X\n", instr)

	inst := Decode(instr)
	cpu.PcChanged = false
	cpu.Execute(inst)
	if !cpu.PcChanged {
		cpu.PC += 4
	}
}
func (cpu *CPU) Run(maxSteps int) {
	for i := 0; i < maxSteps; i++ {

		fmt.Printf(
			"PC=%d x1=%d x2=%d x3=%d\n",
			cpu.PC,
			cpu.Regs[1],
			cpu.Regs[2],
			cpu.Regs[3],
		)

		cpu.Step()
	}
}
