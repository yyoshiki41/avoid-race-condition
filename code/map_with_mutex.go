func (s *Storage) Get(key string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	if v, ok := s.data[key]; ok {
		return v
	}
	return ""
}

func (s *Storage) Set(key string, val string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = val
}

func (s *Storage) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, key)
}
