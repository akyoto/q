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

			switch condition {
			case "linux":
				if s.build.OS != config.Linux {
					return
				}

			case "mac":
				if s.build.OS != config.Mac {
					return
				}

			case "unix":
				if s.build.OS != config.Linux && s.build.OS != config.Mac {
					return
				}

			case "windows":
				if s.build.OS != config.Windows {
					return
				}

			case "x86":
				if s.build.Arch != config.X86 {
					return
				}

			case "arm":
				if s.build.Arch != config.ARM {
					return
				}

			default:
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