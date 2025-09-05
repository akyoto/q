import strings

main() {
	assert strings.trimLeft(" Hello").len == 5
	assert strings.trimLeft("  Hello").len == 5
	assert strings.trimRight("Hello ").len == 5
	assert strings.trimRight("Hello  ").len == 5
	assert strings.trim(" Hello ").len == 5
	assert strings.trim("  Hello  ").len == 5
}