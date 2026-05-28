package main

type Store struct {
	stations map[string]Station
}

func NewStore() *Store {
	return &Store{
		stations: make(map[string]Station),
	}
}

func (s *Store) Put(st Station) {
	s.stations[st.ID] = st
}

func (s *Store) Has(id string) bool {
	_, ok := s.stations[id]
	return ok
}

func (s *Store) Get(id string) (Station, bool) {
	st, ok := s.stations[id]
	return st, ok
}

func (s *Store) Delete(id string) bool {
	if !s.Has(id) {
		return false
	}

	delete(s.stations, id)
	return true
}

func (s *Store) All() []Station {
	result := make([]Station, 0, len(s.stations))

	for _, st := range s.stations {
		result = append(result, st)
	}

	return result
}
