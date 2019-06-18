package source

type DynamicFileSource struct {
	path    string
	inner   *StaticSource
	watcher *fsnotify.Watcher
	cancel  chan<- cancel
}

func (s *DynamicFileSource) Close() {
	s.cancel <- cancel{}
	defer s.watcher.Close()
}
