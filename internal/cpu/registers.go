package cpu

import "fmt"

func (cpu *CPU) DumpRegisters() {
	for i := 0; i < 8; i++ {
		fmt.Printf("x%d=%d\n", i, cpu.Regs[i])
	}
	fmt.Println()
}
