# build

This package contains the source code for the compiler invoked by the `build` command.

## Organization

I organized the files in such a way that a single file always represents a particular aspect of the language (e.g. `if` conditions in `If.go`) even if the methods in the file are spread among different types.

## State

The core type for the compiler is the `State` class which captures the entire compiler state for a single function compilation. We are using a parallel function compiler which compiles every function in a separate goroutine, therefore we'll create a new state object for every function. The result of a compilation is a list of assembler instructions which are then optimized and fed to the final linker.
