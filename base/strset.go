package base

type void struct{}

type innerSet map[string]void

type StrSet struct {
	inner innerSet
	// performance reasons
	keys []string
}

func MakeStrSet(data []string) StrSet {
	sets := make(innerSet, len(data))
	keys := make([]string, len(data))
	counter := 0

	for _, value := range data {
		// should be unique
		if _, ok := sets[value]; !ok {
			sets[value] = void{}
			keys[counter] = value
			counter += 1
		}
	}

	return StrSet{
		sets, keys,
	}
}

func EmptyStrSet() StrSet {
	return StrSet{
		map[string]void{},
		[]string{},
	}
}

func (s *StrSet) Add(value string) {
	if _, ok := s.inner[value]; !ok {
		s.inner[value] = void{}
		s.keys = append(s.keys, value)
	}
}

func (s *StrSet) Contains(value string) bool {
	_, ok := s.inner[value]
	return ok
}

func (s *StrSet) Delete(value string) {
	delete(s.inner, value)
}

func (s *StrSet) Values() []string {
	return s.keys
}
