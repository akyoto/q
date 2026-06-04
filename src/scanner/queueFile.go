package scanner

// queueFile queues a single file to be scanned.
func (s *scanner) queueFile(file string, pkg string) {
	s.group.Go(func() {
		err := s.scanFile(file, pkg)

		if err != nil {
			s.errors <- err
		}
	})
}