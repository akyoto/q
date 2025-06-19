package scanner

import "os"

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