main() {
	foo("foo")
	bar("bar")
}

foo(s string) {
	_(s.ptr)
}

bar(s string) {
	_(s.len)
}

_(_ any) {}