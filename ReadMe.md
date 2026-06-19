# RVEMU

A lightweight RISC-V RV32I emulator written in Go.

RVEMU implements a fetch-decode-execute pipeline, memory subsystem, ELF loading, and support for executing compiled RISC-V binaries.

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
git clone <repository-url>
cd riscv-emulator

go build -o rvemu ./cmd/emulator
```

## Running

```bash
./rvemu hello.elf
```

Or install globally:

```bash
sudo cp rvemu /usr/local/bin/
```

Then:

```bash
rvemu hello.elf
```

## Example

Compile a simple RISC-V program:

```c
int main() {
    return 42;
}
```

Run it:

```bash
rvemu hello.elf
```

## Current Status

### Implemented

* RV32I core instruction support
* Memory subsystem
* Program loader
* ELF loader
* Execution of compiled RISC-V binaries

### Future Improvements

* Additional RV32I instructions
* Proper ELF virtual address mapping
* System call support
* Debugger and instruction tracing

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
