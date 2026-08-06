# Design

A programming language is a data format that stores temporal and spatial data.
Instructions are temporal and data is spatial.
These are the two most important constructs in software.

Example for spatial data:

```q
Point {
	x int
	y int
}
```

Example for temporal instructions:

```q
sum(a int, b int) -> int {
	return a + b
}
```

Designing a data format that is durable and withstands the test of time is a challenge that has proven to be difficult.
One essential factor that helps achieve this property is to reduce the features the language supports to the absolute minimum needed.
With fewer features, there is a smaller surface of areas that can potentially become obsolete in the future.
It also has the added benefit of the data format becoming easier to read.

The following sections will be dedicated to exploring the design options in detail.

## Functions

Machines execute a list of instructions that operate on data.
In theory it is possible to write programs for these machines in assembly code directly.
Even assembly code lets you reuse common functionality with the use of labels to which one can jump.
However, they lack any agreement on where the data they operate on must exist.
This agreement is called an Application Binary Interface (ABI).

Now, we could in theory just define a specification that developers have to follow.
This is problematic because such a regulation does not **enforce** the rules of the specification.
In other words, "code can work together" is a weaker property than "code is guaranteed to work together".
This brings us to the first axiom in the philosophy of designing Q:

> Enforced is better than optional.

We need an abstraction to describe reusable code in such a way that conformance to the ABI is enforced.
A sum function in assembly code could look like this:

```
.sum:
	add r0, r0, r1
	ret
```

The `add` instruction here adds register 1 to register 0 and saves the result in register 0.

This is not a good data format to store our program code.
It assumes that both operands are in registers 0 and 1 and that the return value must reside in register 0 which might not be true in another developer's program who decided to use a different ABI.

The first step to make our code reusable across different programs is to leave out the specific registers we're using.
Here the `in` and `out` registers are abstract and can be filled in by the compiler:

```
.sum:
	add out0, in0, in1
	ret
```

This is already a significant upgrade, but how does a developer know that this function operates on 2 and not 3 operands?
This is where function signatures come into play:

```
.sum(in0, in1) -> (out0):
	add out0, in0, in1
	ret
```

The added information allows the compiler to check that the function is called with the correct number of arguments and that there is a single result.
However, the `add` instruction operates on integers only and would fail if we tried to get the sum of an integer and a float.

We need an enforced restriction which guarantees that the operand registers have contents that make sense for the instructions we're using.
We'll call these restrictions "types" and `int` will be one such type, representing integers.
They make it clear for the compiler that the function can only be called with integers and that the result register contains an integer.

```
.sum(in0 int, in1 int) -> (out0 int):
	add out0, in0, in1
	ret
```

This leaves us with one final problem: The instruction names and the format of these instructions are specific to a certain architecture.
If we want the code to be reusable on multiple architectures, we need to abstract these.
Here is a possible abstraction using the first keyword `return` and a symbolic expression that abstracts away the exact `add` instruction used on the machine.

```
.sum(in0 int, in1 int) -> (out0 int):
	return (+ in0 in1)
```

While symbolic expressions are perfect for computers, code should be easily understood by humans.
Code that is hard to understand often leads to software bugs.
We want to prevent these, so we'll stick to the proven math notation that humans have been using for centuries:

```
.sum(in0 int, in1 int) -> (out0 int):
	return in0 + in1
```

Another thing that will help humans understand the purpose of each parameter is a name.
Note that due to the commutative nature of the sum operation, names don't really add any extra information in this specific example but in general it helps the reader understand the intended purpose of each parameter.
We no longer need to number the registers because they are implied by the order of the parameters and the ABI of our compiler:

```
.sum(a int, b int) -> (sum int):
	return a + b
```

This is almost the final design of how functions are written in Q.
Notice that the end of the function is not clearly visible,
so we'll add symbols for the start and end of the function block to make it immediately clear from just glancing at the function where it ends:

```
.sum(a int, b int) -> (sum int) {
	return a + b
}
```

The `.` prefix (or `func` in Go, or `fn` in Rust and Zig) is no longer required because Q expects function definitions by default:

```
sum(a int, b int) -> (sum int) {
	return a + b
}
```

You can optionally leave out the names for the return values if the function name already conveys the same information:

```
sum(a int, b int) -> int {
	return a + b
}
```

This is how functions are defined in Q.