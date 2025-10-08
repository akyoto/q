package core

import "git.urbach.dev/cli/q/src/fs"

// ReceiveFile receives a file from the scanner.
func (env *Environment) ReceiveFile(file *fs.File) {
	env.Files = append(env.Files, file)
}