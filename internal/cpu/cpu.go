package cpu

type CPU struct {
    PC   uint32
    Regs [32]uint32
    Mem  []byte
}

func New() *CPU{
	return &CPU{}
}