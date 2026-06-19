package cpu

type CPU struct {
	PC   uint32
	Regs [32]uint32
	Mem  []byte
}

func NewCPU() *CPU {
	cpu := &CPU{
		Mem: make([]byte, 1024*1024),
	}

	cpu.Regs[2] = uint32(len(cpu.Mem))

	return cpu
}
