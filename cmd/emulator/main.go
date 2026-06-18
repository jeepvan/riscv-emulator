package main

import (
	"fmt"

	"github.com/jeepvan/riscv-emulator/internal/cpu"
)

func main() {
	cpu := cpu.CPU{
		Mem: make([]byte, 256),
	}

	cpu.Regs[1] = 100
	cpu.Regs[5] = 42

	cpu.ExecuteSW(1, 5, 0)
	cpu.ExecuteLW(6, 1, 0)

	fmt.Println(cpu.Regs[6])
}
