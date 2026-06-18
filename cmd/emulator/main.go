package main

import "github.com/jeepvan/riscv-emulator/internal/cpu"

func main() {
	cpu := cpu.New()
	cpu.Mem = []byte{
		0x93, 0x00, 0xA0, 0x00,
		0x13, 0x01, 0x40, 0x01,
		0xB3, 0x81, 0x20, 0x00,
	}
	for cpu.PC < uint32(len(cpu.Mem)) {
		cpu.Step()
	}
	cpu.DumpRegisters()
}
