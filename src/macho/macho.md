# Mach-O

## Notes

- The start of the file must be loaded in some segment.
- The start of the file must be marked as readable + executable.
- Load command size must be divisible by 8.
- Segments must be page-aligned in the file.

## Links

- https://github.com/apple-oss-distributions/xnu/blob/main/bsd/kern/mach_loader.c
- https://github.com/apple-oss-distributions/xnu/blob/main/bsd/kern/mach_loader.c#L2021-L2027
- https://github.com/apple-oss-distributions/xnu/blob/main/EXTERNAL_HEADERS/mach-o/loader.h
- https://en.wikipedia.org/wiki/Mach-O
- https://github.com/aidansteele/osx-abi-macho-file-format-reference
- https://stackoverflow.com/questions/39863112/what-is-required-for-a-mach-o-executable-to-load