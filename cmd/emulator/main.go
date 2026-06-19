package main

import (
	"fmt"
	"os"

	"github.com/jeepvan/riscv-emulator/internal/cpu"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: rvemu <program.elf>")
		os.Exit(1)
	}

	c := cpu.NewCPU()

	err := cpu.LoadELF(c, os.Args[1])
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1000; i++ {
		instr := c.Fetch()

		if instr == 0 {
			fmt.Println("Program finished")
			break
		}

		c.Step()
	}
}
