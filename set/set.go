package set

type Set struct {
	m map[interface{}]bool
}

func NewSet() *Set {
	return &Set{
		m: make(map[interface{}]bool, 0),
	}
}

func (s *Set) Add(v interface{}) {
	s.m[v] = true
}

func (s *Set) Contains(v interface{}) bool {
	_, ok := s.m[v]
	return ok
}

func (s *Set) Delete(v interface{}) {
	if s.Contains(v) {
		delete(s.m, v)
	}
}

func (s *Set) Size(v interface{}) int {
	return len(s.m)
}

func (s *Set) ToArray() []interface{} {
	keys := make([]interface{}, 0)
	for k := range s.m {
		keys = append(keys, k)
	}
	return keys
}
