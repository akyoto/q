import cli

main() {
	path, err := cli.env("PATH")
	assert err == 0
	assert path.len > 0
}