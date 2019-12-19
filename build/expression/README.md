# expression

This package parses an expression in Q. It supports mathematical operators as well as function calls like `sum(3, 4) + 5`.

An expression consists of a token (either an operator or the leaf token like numbers and identifiers). Operators and function calls have child expressions. Operators use the child expressions as operands while function calls use them as parameters for the call.

Function call nodes have the boolean `IsFunctionCall` set to true.

## Examples

### `1 + 2`

```text
  +
 / \
1   2
```

### `sum(1, 2)`

```text
  sum
 /   \
1     2
```

### `1 + 2 * 3`

```text
  +
 / \
1   *
   / \
  2   3
```
