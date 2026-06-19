# RVEMU

A lightweight RISC-V RV32I emulator written in Go.

RVEMU implements a fetch-decode-execute pipeline, memory subsystem, ELF loading, and execution of compiled RISC-V binaries.

## Features

### CPU Core

* 32 General Purpose Registers (x0-x31)
* Program Counter (PC)
* Fetch → Decode → Execute pipeline
* RV32I instruction decoding

### Supported Instructions

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

#### Control Flow

* BEQ (Branch if Equal)
* BNE (Branch if Not Equal)
* JAL (Jump and Link)
* JALR (Jump and Link Register)

### ELF Support

* ELF file parsing
* Entry point extraction
* `.text` section loading
* Execution of compiled RV32I binaries

## Architecture

```text
ELF File
    ↓
ELF Loader
    ↓
Memory
    ↓
Fetch
    ↓
Decode
    ↓
Execute
    ↓
Registers / Memory Updates
```

## Project Structure

```text
cmd/
└── emulator/
    └── main.go

examples/
├── hello.c
└── hello.elf

internal/cpu/
├── cpu.go
├── memory.go
├── fetch.go
├── decode.go
├── execute.go
├── loader.go
├── elf.go
├── step.go
└── opcodes.go
```

## Building

```bash
git clone https://github.com/jeepvan/riscv-emulator.git
cd riscv-emulator

go build -o rvemu ./cmd/emulator
```

## Running

Run the included example:

```bash
./rvemu examples/hello.elf
```

Or install globally:

```bash
sudo cp rvemu /usr/local/bin/
rvemu examples/hello.elf
```

## Example Program

`examples/hello.c`

```c
int main() {
    return 42;
}
```

Run it with:

```bash
rvemu examples/hello.elf
```

## Current Status

### Implemented

* RV32I core instruction support
* Memory subsystem
* Program loader
* ELF loader
* Execution of compiled RV32I binaries

### Known Limitations

* Partial RV32I instruction coverage
* ELF sections are currently loaded at address 0
* No system call support
* No operating system support

### Future Improvements

* Complete RV32I instruction set
* Proper ELF virtual address mapping
* System call support
* Debugger and instruction tracing
* Linux userspace support

## Learning Outcomes

This project explores:

* Computer Architecture
* Instruction Set Architectures (ISA)
* RISC-V RV32I
* Memory Management
* ELF Executable Format
* Emulator Design
* Systems Programming in Go

## License

MIT License
