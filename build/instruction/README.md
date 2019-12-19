# instruction

This package parses instructions on a very high level.

## Example

The following code...

```q
if x > 5 {
	doSomething()
}
```

...would result in 3 instructions:

```text
IfStart
Call
IfEnd
```
