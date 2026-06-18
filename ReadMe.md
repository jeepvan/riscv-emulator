# RISC-V Emulator

A RISC-V RV32I emulator written in Go.

## Features

* RV32I instruction decoding
* Fetch → Decode → Execute pipeline
* Register file (x0-x31)
* Program Counter (PC)
* R-Type instructions

  * ADD
  * SUB
  * AND
  * OR
  * XOR
* I-Type instructions

  * ADDI

## Architecture

```text
Memory
  ↓
Fetch
  ↓
Decode
  ↓
Instruction
  ↓
Executegit remote add origin git@github.com:jeepvan/riscv-emulator.git
  ↓
Registers
```

## Current Status

The emulator can execute simple RV32I programs loaded into memory.

## Roadmap

* [ ] Sign extension for immediates
* [ ] LW
* [ ] SW
* [ ] BEQ
* [ ] BNE
* [ ] JAL
* [ ] ELF loading
* [ ] Test suite
* [ ] RV32I compliance
