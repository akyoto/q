package scanner

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// scanner is used to scan files before the actual compilation step.
type scanner struct {
	files  chan *fs.File
	errors chan error
	build  *build.Build
	queued sync.Map
	group  sync.WaitGroup
}

// queue scans the list of files.
func (s *scanner) queue(files ...string) {
	for _, file := range files {
		stat, err := os.Stat(file)

		if err != nil {
			s.errors <- err
			return
		}

		if stat.IsDir() {
			s.queueDirectory(file, "main")
		} else {
			s.queueFile(file, "main")
		}
	}
}

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
				if s.build.OS != build.Linux {
					return
				}

			case "mac":
				if s.build.OS != build.Mac {
					return
				}

			case "unix":
				if s.build.OS != build.Linux && s.build.OS != build.Mac {
					return
				}

			case "windows":
				if s.build.OS != build.Windows {
					return
				}

			case "x86":
				if s.build.Arch != build.X86 {
					return
				}

			case "arm":
				if s.build.Arch != build.ARM {
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

// queueFile queues a single file to be scanned.
func (s *scanner) queueFile(file string, pkg string) {
	s.group.Add(1)

	go func() {
		defer s.group.Done()
		err := s.scanFile(file, pkg)

		if err != nil {
			s.errors <- err
		}
	}()
}

// scanFile scans a single file.
func (s *scanner) scanFile(path string, pkg string) error {
	contents, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	tokens := token.Tokenize(contents)

	file := &fs.File{
		Path:    path,
		Package: pkg,
		Bytes:   contents,
		Tokens:  tokens,
	}

	s.files <- file

	for i := 0; i < len(tokens); i++ {
		switch tokens[i].Kind {
		case token.NewLine:
		case token.Comment:
		case token.Identifier:
			i, err = s.scanFunction(file, tokens, i)
		case token.Import:
			i, err = s.scanImport(file, tokens, i)
		case token.EOF:
			return nil
		case token.Invalid:
			return errors.New(&invalidCharacter{Character: tokens[i].String(file.Bytes)}, file, tokens[i].Position)
		default:
			return errors.New(&invalidTopLevel{Instruction: tokens[i].String(file.Bytes)}, file, tokens[i].Position)
		}

		if err != nil {
			return err
		}
	}

	return nil
}