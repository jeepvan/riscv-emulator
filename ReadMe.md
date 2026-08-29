# riscv-emulator

![ci](https://github.com/jeepvan/riscv-emulator/actions/workflows/go.yml/badge.svg)

a risc-v emulator written in go from scratch, standard library only.

## what works right now

rv32i subset, 12 instructions:

    add sub addi and or xor lw sw beq bne jal jalr

- fetch/decode/execute loop with explicit next-pc handling
- 32 registers + pc, x0 hardwired to zero
- 1 MiB flat little-endian memory at address 0
- minimal elf loading: copies .text to address 0 and starts there.
  entry point and other segments are ignored for now
- instruction encoders for all five formats, usable to hand-assemble
  test programs from go

## what does not exist yet

no csrs, no traps, no interrupts, no devices, no clean way for a
program to halt. programs run until they fetch a zero word or hit the
1000 step limit.

## build and run

requires go 1.25+.

    go build ./...
    go run ./cmd/emulator examples/hello.elf

## test

    go test ./...

table driven, spec level tests for every implemented instruction,
including the cases that were once real bugs here: jal's link
register, jalr's lsb clearing, rd == rs1 ordering.