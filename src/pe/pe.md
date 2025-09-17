# Portable Executable

## Notes

Unlike Linux, Windows does not ignore zero-length sections at the end of a file and will
fail loading them because they don't exist within the file. Adding a single byte to the
section can fix this problem, but it's easier to just remove the section header entirely.
The solution used here is to guarantee that the data section is never empty by always
importing a few core functions from "kernel32.dll".

## DLL function pointers

The section where the DLL function pointers are stored does not need to be marked as writable.
The Windows executable loader resolves the pointers before they are loaded into memory.

The stack must be 16 byte aligned before a DLL function is called.

## Links

- https://learn.microsoft.com/en-us/windows/win32/debug/pe-format
- https://learn.microsoft.com/en-us/previous-versions/ms809762(v=msdn.10)
- https://learn.microsoft.com/en-us/archive/msdn-magazine/2002/february/inside-windows-win32-portable-executable-file-format-in-detail
- https://learn.microsoft.com/en-us/archive/msdn-magazine/2002/march/inside-windows-an-in-depth-look-into-the-win32-portable-executable-file-format-part-2
- https://blog.kowalczyk.info/articles/pefileformat.html
- https://keyj.emphy.de/win32-pe/
- https://corkamiwiki.github.io/PE
- https://github.com/ayaka14732/TinyPE-on-Win10