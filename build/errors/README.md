# errors

This package lists all the possible compiler errors.

## Usage

To capture the call stack, the `errors.New(error)` constructor should be used. It will let the compiler developers know which particular line in the source code generated the error.

## Examples

### `MissingFunctionName`

```go
errors.New(errors.MissingFunctionName)
```

### `UnknownFunction`

```go
errors.New(&errors.UnknownFunction{Name: "doesNotExist"})
```

```go
errors.New(&errors.UnknownFunction{Name: "prin", CorrectName: "print"})
```
