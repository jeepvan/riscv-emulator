package cpu

type CPU struct {
	PC   uint32
	Regs [32]uint32
	Mem  []byte
}

func NewCPU() *CPU {
	return &CPU{
		Mem: make([]byte, 1024*1024),
	}
}
