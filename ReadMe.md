# RISC-V Emulator in Go

A simple RV32I RISC-V emulator written in Go for learning computer architecture, instruction decoding, and emulator development.

## Current Features

### CPU Core

* Program Counter (PC)
* 32 General Purpose Registers (`x0-x31`)
* Byte-addressable memory
* Fetch → Decode → Execute pipeline

### Memory Operations

* Read8
* Write8
* Read32
* Write32
* Little-endian memory access

### Implemented Instructions

#### Arithmetic

* ADD
* SUB
* ADDI

#### Logical

* AND
* OR
* XOR

#### Memory

* LW (Load Word)
* SW (Store Word)

### Instruction Decoding

Implemented decoding for:

* R-Type instructions
* I-Type instructions
* Load instructions
* Store instructions

### Immediate Handling

* 12-bit immediate extraction
* S-Type immediate reconstruction
* Sign extension support

## Project Structure

```text
riscv-emulator/
├── cmd/
│   └── emulator/
│       └── main.go
├── internal/
│   └── cpu/
│       ├── cpu.go
│       ├── fetch.go
│       ├── decode.go
│       ├── execute.go
│       ├── instructions.go
│       ├── opcodes.go
│       ├── memory.go
│       ├── registers.go
│       └── step.go
└── go.mod
```

## Execution Flow

```text
Memory
  ↓
Fetch
  ↓
Decode
  ↓
Instruction Struct
  ↓
Execute
  ↓
Registers / Memory
```

## Example

```assembly
ADDI x1, x0, 10
ADDI x2, x0, 20
ADD  x3, x1, x2
```

Result:

```text
x3 = 30
```

## Roadmap

### Completed

* [x] Register file
* [x] Memory subsystem
* [x] Fetch stage
* [x] Decode stage
* [x] Execute stage
* [x] ADD
* [x] SUB
* [x] ADDI
* [x] AND
* [x] OR
* [x] XOR
* [x] LW
* [x] SW
* [x] Sign extension

### Next

* [ ] BEQ
* [ ] BNE
* [ ] Branch decoding
* [ ] JAL
* [ ] JALR
* [ ] ELF loading
* [ ] More RV32I instructions


