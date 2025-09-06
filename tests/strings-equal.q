import strings

main() {
	assert !strings.equal("world", "universe")
	assert strings.equal("world", "world")
	assert !strings.equal("世界", "宇宙")
	assert strings.equal("世界", "世界")
}