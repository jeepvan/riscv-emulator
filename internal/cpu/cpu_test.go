package cpu

import "testing"

// ---------------------------------------------------------------------------
// Reference encoders — one per instruction format.
// These are spec bit layouts (RISC-V Unprivileged ISA, ch. 2), independent of
// decode.go, so a bug in Decode cannot hide inside the tests.
// ---------------------------------------------------------------------------

func encR(funct7, rs2, rs1, funct3, rd uint32) uint32 {
	return funct7<<25 | rs2<<20 | rs1<<15 | funct3<<12 | rd<<7 | 0x33
}

func encI(imm int32, rs1, funct3, rd, opcode uint32) uint32 {
	return (uint32(imm)&0xFFF)<<20 | rs1<<15 | funct3<<12 | rd<<7 | opcode
}

func encS(imm int32, rs2, rs1, funct3 uint32) uint32 {
	u := uint32(imm)
	return ((u>>5)&0x7F)<<25 | rs2<<20 | rs1<<15 | funct3<<12 | (u&0x1F)<<7 | 0x23
}

func encB(imm int32, rs2, rs1, funct3 uint32) uint32 {
	u := uint32(imm)
	return ((u>>12)&1)<<31 | ((u>>5)&0x3F)<<25 | rs2<<20 | rs1<<15 |
		funct3<<12 | ((u>>1)&0xF)<<8 | ((u>>11)&1)<<7 | 0x63
}

func encJ(imm int32, rd uint32) uint32 {
	u := uint32(imm)
	return ((u>>20)&1)<<31 | ((u>>1)&0x3FF)<<21 | ((u>>11)&1)<<20 |
		((u>>12)&0xFF)<<12 | rd<<7 | 0x6F
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func TestSignExtend(t *testing.T) {
	cases := []struct {
		value uint32
		bits  int
		want  int32
	}{
		{0x000, 12, 0},
		{0x001, 12, 1},
		{0x7FF, 12, 2047},
		{0x800, 12, -2048},
		{0xFFF, 12, -1},
		{0x1FFF, 13, -1},
		{0x0FFF, 13, 4095},
		{0x1FFFE0, 21, -32},
	}
	for _, c := range cases {
		if got := SignExtend(c.value, c.bits); got != c.want {
			t.Errorf("SignExtend(%#x, %d) = %d, want %d", c.value, c.bits, got, c.want)
		}
	}
}

func TestMemoryLittleEndian(t *testing.T) {
	c := NewCPU()
	c.Write32(0x100, 0x11223344)
	for i, want := range []byte{0x44, 0x33, 0x22, 0x11} {
		if got := c.Read8(0x100 + uint32(i)); got != want {
			t.Errorf("byte %d = %#x, want %#x (RISC-V is little-endian)", i, got, want)
		}
	}
	if got := c.Read32(0x100); got != 0x11223344 {
		t.Errorf("Read32 round-trip = %#x, want 0x11223344", got)
	}
}

func TestEncodeADDIRoundTrip(t *testing.T) {
	inst := Decode(EncodeADDI(1, 2, 5)) // your encoder through your decoder
	if inst.Rd != 1 || inst.Rs1 != 2 || inst.Imm != 5 {
		t.Errorf("EncodeADDI(1,2,5) decoded to rd=%d rs1=%d imm=%d, want 1/2/5",
			inst.Rd, inst.Rs1, inst.Imm)
	}
}

// ---------------------------------------------------------------------------
// Architectural state rules
// ---------------------------------------------------------------------------

func TestX0Hardwired(t *testing.T) {
	c := NewCPU()
	c.Write32(0, encI(5, 0, 0, 0, 0x13)) // addi x0, x0, 5
	c.Step()
	if c.Regs[0] != 0 {
		t.Errorf("x0 = %d after write, must always read 0", c.Regs[0])
	}
}

func TestStepAdvancesPC(t *testing.T) {
	c := NewCPU()
	c.Write32(0, encI(1, 0, 0, 5, 0x13)) // addi x5, x0, 1
	c.Step()
	if c.PC != 4 {
		t.Errorf("PC = %d after non-branch, want 4", c.PC)
	}
}

// ---------------------------------------------------------------------------
// ALU
// ---------------------------------------------------------------------------

func TestADDI(t *testing.T) {
	cases := []struct {
		name   string
		rs1Val uint32
		imm    int32
		want   uint32
	}{
		{"positive", 10, 5, 15},
		{"negative imm", 10, -1, 9},
		{"sign-extends to full width", 0, -1, 0xFFFFFFFF},
		{"wraps modulo 2^32", 0xFFFFFFFF, 1, 0},
		{"min imm", 0, -2048, 0xFFFFF800},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewCPU()
			c.Regs[6] = tc.rs1Val
			c.Write32(0, encI(tc.imm, 6, 0, 5, 0x13)) // addi x5, x6, imm
			c.Step()
			if c.Regs[5] != tc.want {
				t.Errorf("addi: x5 = %#x, want %#x", c.Regs[5], tc.want)
			}
		})
	}
}

// NOTE: AND / OR / XOR go through the full Decode→Execute path here.
// The Execute*() helpers exist in execute.go but the R-type funct3 switch
// never routes to them — these subtests fail until that wiring exists.
func TestRType(t *testing.T) {
	cases := []struct {
		name   string
		funct7 uint32
		funct3 uint32
		a, b   uint32
		want   uint32
	}{
		{"add", 0x00, 0b000, 7, 3, 10},
		{"add wraps", 0x00, 0b000, 0xFFFFFFFF, 1, 0},
		{"sub", 0x20, 0b000, 7, 3, 4},
		{"sub borrows", 0x20, 0b000, 0, 1, 0xFFFFFFFF},
		{"xor", 0x00, 0b100, 0b1100, 0b1010, 0b0110},
		{"or", 0x00, 0b110, 0b1100, 0b1010, 0b1110},
		{"and", 0x00, 0b111, 0b1100, 0b1010, 0b1000},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewCPU()
			c.Regs[6], c.Regs[7] = tc.a, tc.b
			c.Write32(0, encR(tc.funct7, 7, 6, tc.funct3, 5)) // op x5, x6, x7
			c.Step()
			if c.Regs[5] != tc.want {
				t.Errorf("%s: x5 = %#x, want %#x", tc.name, c.Regs[5], tc.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// Loads / stores
// ---------------------------------------------------------------------------

func TestLoadStoreRoundTrip(t *testing.T) {
	c := NewCPU()
	c.Regs[6] = 0x200
	c.Regs[7] = 0xDEADBEEF
	c.Write32(0, encS(16, 7, 6, 0b010))       // sw x7, 16(x6)   -> mem[0x210]
	c.Write32(4, encI(16, 6, 0b010, 5, 0x03)) // lw x5, 16(x6)
	c.Step()
	c.Step()
	if got := c.Read32(0x210); got != 0xDEADBEEF {
		t.Errorf("sw: mem[0x210] = %#x, want 0xDEADBEEF", got)
	}
	if c.Regs[5] != 0xDEADBEEF {
		t.Errorf("lw: x5 = %#x, want 0xDEADBEEF", c.Regs[5])
	}
}

func TestLoadStoreNegativeOffset(t *testing.T) {
	c := NewCPU()
	c.Regs[6] = 0x200
	c.Regs[7] = 0xCAFEBABE
	c.Write32(0, encS(-4, 7, 6, 0b010))       // sw x7, -4(x6)  -> mem[0x1FC]
	c.Write32(4, encI(-4, 6, 0b010, 5, 0x03)) // lw x5, -4(x6)
	c.Step()
	c.Step()
	if c.Regs[5] != 0xCAFEBABE {
		t.Errorf("negative offset: x5 = %#x, want 0xCAFEBABE (addr math with S-type imm sign)", c.Regs[5])
	}
}

// ---------------------------------------------------------------------------
// Branches
// ---------------------------------------------------------------------------

func TestBranches(t *testing.T) {
	cases := []struct {
		name            string
		funct3          uint32
		a, b            uint32
		imm             int32
		startPC, wantPC uint32
	}{
		{"beq taken forward", 0b000, 5, 5, 12, 0, 12},
		{"beq not taken", 0b000, 5, 6, 12, 0, 4},
		{"bne taken forward", 0b001, 5, 6, 12, 0, 12},
		{"bne not taken", 0b001, 5, 5, 12, 0, 4},
		{"beq taken backward", 0b000, 1, 1, -16, 0x40, 0x30},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewCPU()
			c.PC = tc.startPC
			c.Regs[6], c.Regs[7] = tc.a, tc.b
			c.Write32(tc.startPC, encB(tc.imm, 7, 6, tc.funct3))
			c.Step()
			if c.PC != tc.wantPC {
				t.Errorf("%s: PC = %#x, want %#x", tc.name, c.PC, tc.wantPC)
			}
		})
	}
}

// A jump-to-self (`j .`) must loop forever — it's the standard bare-metal
// halt/spin idiom. Fails until Step stops inferring "did we jump?" from
// oldPC == PC and Execute reports the next PC explicitly (Phase 0, item 4).
func TestSelfJumpSpins(t *testing.T) {
	c := NewCPU()
	c.Write32(0, encJ(0, 0)) // jal x0, 0
	c.Step()
	if c.PC != 0 {
		t.Errorf("`j .`: PC = %d, want 0 (self-jump must not fall through)", c.PC)
	}
}

// ---------------------------------------------------------------------------
// JAL / JALR — the call/return pair
// ---------------------------------------------------------------------------

func TestJAL(t *testing.T) {
	t.Run("forward, saves link", func(t *testing.T) {
		c := NewCPU()
		c.PC = 8
		c.Write32(8, encJ(16, 1)) // jal x1, +16
		c.Step()
		if c.PC != 24 {
			t.Errorf("jal: PC = %d, want 24", c.PC)
		}
		if c.Regs[1] != 12 {
			t.Errorf("jal: x1 = %d, want 12 (PC+4 — decode must extract rd)", c.Regs[1])
		}
	})
	t.Run("backward", func(t *testing.T) {
		c := NewCPU()
		c.PC = 0x40
		c.Write32(0x40, encJ(-0x20, 2)) // jal x2, -32
		c.Step()
		if c.PC != 0x20 {
			t.Errorf("jal backward: PC = %#x, want 0x20", c.PC)
		}
		if c.Regs[2] != 0x44 {
			t.Errorf("jal backward: x2 = %#x, want 0x44", c.Regs[2])
		}
	})
}

func TestJALR(t *testing.T) {
	t.Run("jumps to rs1+imm, saves link", func(t *testing.T) {
		c := NewCPU()
		c.Regs[5] = 0x300
		c.Write32(0, encI(0, 5, 0, 1, 0x67)) // jalr x1, 0(x5)
		c.Step()
		if c.PC != 0x300 {
			t.Errorf("jalr: PC = %#x, want 0x300", c.PC)
		}
		if c.Regs[1] != 4 {
			t.Errorf("jalr: x1 = %d, want 4", c.Regs[1])
		}
	})
	// Spec (Unprivileged ISA §2.5.1): target = (rs1 + imm) & ~1.
	// The naive one-line implementation misses this.
	t.Run("clears target bit 0", func(t *testing.T) {
		c := NewCPU()
		c.Regs[5] = 0x301
		c.Write32(0, encI(0, 5, 0, 1, 0x67))
		c.Step()
		if c.PC != 0x300 {
			t.Errorf("jalr odd target: PC = %#x, want 0x300 (LSB must be cleared)", c.PC)
		}
	})
	// rd == rs1 is legal and common. rs1 must be read BEFORE rd is written,
	// or the jump target is computed from the freshly-written link value.
	t.Run("rd == rs1 reads before writing", func(t *testing.T) {
		c := NewCPU()
		c.Regs[1] = 0x200
		c.Write32(0, encI(0, 1, 0, 1, 0x67)) // jalr x1, 0(x1)
		c.Step()
		if c.PC != 0x200 {
			t.Errorf("jalr rd==rs1: PC = %#x, want 0x200 (read rs1 before writing rd)", c.PC)
		}
		if c.Regs[1] != 4 {
			t.Errorf("jalr rd==rs1: x1 = %d, want 4", c.Regs[1])
		}
	})
}
