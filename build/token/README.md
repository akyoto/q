# token

This package takes an input file and turns the code into tokens that are easier to work with.

## Example

The following code...

```q
main() {
	print("Hello World")
}
```

...would produce the following tokens:

* Identifier `main`
* GroupStart `(`
* GroupEnd `)`
* BlockStart `{`
* Newline
* Identifier `print`
* GroupStart `(`
* Text `Hello World`
* GroupEnd `)`
* Newline
* BlockEnd `}`
* Newline
