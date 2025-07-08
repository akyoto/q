# Intermediate Representation

## Static Single Assignment

The SSA IR follows a simple rule: every value is assigned exactly once.

### Basic Blocks

Every function has a list of basic blocks. Basic blocks store values in the order they appear in the original code.

### Instructions

Instructions are no different from values. The calculation `1 + 2` is represented as 3 values using 2 int constants and 1 binary operation:

```
t0 = 1
t1 = 2
t2 = t0 + t1
```

All of these are considered values even in cases where the type of the value is `void` such as in procedural function calls with side effects.

### Pointers

There are no IDs or indices, therefore values don't know their position in a basic block.

Values reference other values by including a pointer to them.