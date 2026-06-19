package main

import (
	"fmt"

	"github.com/jeepvan/riscv-emulator/internal/cpu"
)

func main() {
	c := cpu.CPU{
		Mem: make([]byte, 1024),
	}

	program := []uint32{
		cpu.EncodeADDI(1, 0, 10),
		cpu.EncodeADDI(2, 0, 20),
		0x002081B3, // ADD x3, x1, x2
	}

	cpu.LoadProgram(&c, program)

	c.PC = 0

	c.Run(3)

	fmt.Println("x3 =", c.Regs[3])
}
