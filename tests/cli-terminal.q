import cli
import io

main() {
	assert !cli.isTerminal(io.stdin)
}