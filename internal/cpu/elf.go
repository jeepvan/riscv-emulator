package cpu

import (
	"debug/elf"
	"fmt"
)

func InspectELF(path string) ([]byte, uint64, error) {
	f, err := elf.Open(path)
	if err != nil {
		return nil, 0, err
	}
	defer f.Close()
	text := f.Section(".text")
	if text == nil {
		return nil, 0, fmt.Errorf(".text section not found")
	}
	data, err := text.Data()
	if err != nil {
		return nil, 0, err
	} else {
		return data, f.Entry, nil
	}
}

func LoadELF(cpu *CPU, path string) error {
	data,_, err := InspectELF(path)
	if err != nil {
		return err
	}
	copy(cpu.Mem, data)
	cpu.PC = 0   // TODO: Load sections at ELF virtual addresses.
	return nil
}
