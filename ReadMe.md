# RVEMU

A lightweight RISC-V RV32I emulator written in Go.

RVEMU implements a complete fetch-decode-execute pipeline, memory subsystem, ELF loading, and execution of compiled RISC-V binaries.

## Features

### CPU Core

* 32 general-purpose registers (x0-x31)
* Program Counter (PC)
* Fetch → Decode → Execute pipeline
* RV32I instruction decoding and execution

### Supported Instructions

#### Arithmetic

* ADD
* SUB
* ADDI

#### Logical

* AND
* OR
* XOR

#### Memory Access

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
Register / Memory Updates
```

## Repository Structure

```text
cmd/
└── emulator/
    └── main.go

examples/
├── hello.c
└── hello.elf

internal/cpu/
├── cpu.go
├── decode.go
├── elf.go
├── execute.go
├── fetch.go
├── loader.go
├── memory.go
├── opcodes.go
└── step.go
```

## Build

```bash
git clone https://github.com/jeepvan/riscv-emulator.git
cd riscv-emulator

go build -o rvemu ./cmd/emulator
```

## Usage

Run a compiled RISC-V ELF binary:

```bash
./rvemu examples/hello.elf
```

Or install globally:

```bash
sudo cp rvemu /usr/local/bin/
rvemu examples/hello.elf
```

## Example

Source:

```c
int main() {
    return 42;
}
```

Compile:

```bash
riscv64-unknown-elf-gcc \
  -march=rv32i \
  -mabi=ilp32 \
  -nostdlib \
  -nostartfiles \
  hello.c \
  -o hello.elf
```

Run:

```bash
rvemu hello.elf
```

## Current Status

Implemented:

* RV32I fetch-decode-execute pipeline
* Register file and memory subsystem
* ELF loader
* Control flow instructions
* Execution of compiled RV32I binaries

## Roadmap

* Complete RV32I instruction coverage
* Proper ELF virtual address mapping
* Debugger and instruction tracing
* System call support

## License

MIT License
