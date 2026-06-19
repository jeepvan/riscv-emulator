package main

import (
	"fmt"

	"github.com/jeepvan/riscv-emulator/internal/cpu"
)

func main() {
	c := cpu.NewCPU()

	fmt.Println("RISC-V Emulator")
	fmt.Printf("Memory Size: %d bytes\n", len(c.Mem))

	// TODO:
	// Load ELF
	// Set PC
	// Run CPU
}
