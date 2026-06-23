import cli
import io

main() {
	args := cli.args()

	loop i := 0..args.len {
		io.writeLine(args[i])
	}
}