# ELF

## Basic structure

1. ELF header [0x00 : 0x40]
2. Program headers [0x40 : 0xB0]
3. String table
4. Section headers
5. Padding
6. Executable code [0x1000 : ...]
7. Padding
8. Read-only data [0x2000 : ...]

## Entry point

The executables are compiled as position-independent executables (PIE). Therefore the entry point is defined as a file offset instead of a static virtual address.

## Padding

To ensure that execution permissions are properly applied,
the code section and the data section are aligned on page boundaries. Permissions like read, write and execute can only be applied to an entire page in memory.

## Links

- https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/fs/binfmt_elf.c
- https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/arch/x86/include/asm/elf.h
- https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/include/uapi/linux/elf.h
- https://lwn.net/Articles/631631/
- https://en.wikipedia.org/wiki/Executable_and_Linkable_Format
- https://www.muppetlabs.com/~breadbox/software/tiny/teensy.html
- https://nathanotterness.com/2021/10/tiny_elf_modernized.html