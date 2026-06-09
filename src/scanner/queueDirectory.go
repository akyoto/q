package scanner

import (
	"path/filepath"
	"strings"

	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/fs"
)

// queueDirectory queues an entire directory to be scanned.
func (s *scanner) queueDirectory(directory string, pkg string) {
	_, loaded := s.queued.LoadOrStore(directory, nil)

	if loaded {
		return
	}

	err := fs.Walk(directory, func(name string) {
		if !strings.HasSuffix(name, ".q") {
			return
		}

		tmp := name[:len(name)-2]

		for {
			underscore := strings.LastIndexByte(tmp, '_')

			if underscore == -1 {
				break
			}

			condition := tmp[underscore+1:]
			pass := false

			switch condition {
			case "linux":
				pass = s.build.OS == config.Linux
			case "mac":
				pass = s.build.OS == config.Mac
			case "unix":
				pass = s.build.OS == config.Linux || s.build.OS == config.Mac
			case "windows":
				pass = s.build.OS == config.Windows
			case "winux":
				pass = s.build.OS == config.Linux || s.build.OS == config.Windows
			case "x86":
				pass = s.build.Arch == config.X86
			case "arm":
				pass = s.build.Arch == config.ARM
			}

			if !pass {
				return
			}

			tmp = tmp[:underscore]
		}

		fullPath := filepath.Join(directory, name)
		s.queueFile(fullPath, pkg)
	})

	if err != nil {
		s.errors <- err
	}
}