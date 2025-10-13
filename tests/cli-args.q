import cli
import io

main() {
	args := cli.args()

	loop i := 1..args.len {
		io.writeLine(args[i])
	}
}