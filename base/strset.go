package base

// StrSet : Simple Set implementation that holds its keys explicitly
// backed by map[string]void.
//
// By default, you get key information from map[string]void by using
// `reflect` but since it uses reflection (contact go runtime directly),
// for performance reasons we manage our keys by our self.
//
// Since mostly we never actually call StrSet.Delete in our code.
//
type StrSet struct {
	inner map[string]int
	// performance reasons
	keys []string
}

// MakeStrSet : Create StrSet by giving string slices.
//
// Returned StrSet will have unique elements & ignore
// second time duplicated element being added.
//
func MakeStrSet(data []string) StrSet {
	var sets = make(map[string]int)
	var keys = make([]string, 0)
	var counter = 0

	for _, value := range data {
		// should be unique
		if _, ok := sets[value]; !ok {
			sets[value] = counter
			keys = append(keys, value)
			counter += 1
		}
	}

	return StrSet{
		sets, keys,
	}
}

// EmptyStrSet : Create empty StrSet (empty key & inner).
//
func EmptyStrSet() StrSet {
	return StrSet{
		map[string]int{},
		[]string{},
	}
}

// Add : Adding new value to current StrSet (uniquely).
//
func (s *StrSet) Add(value string) {
	if _, ok := s.inner[value]; !ok {
		s.inner[value] = len(s.keys) - 1
		s.keys = append(s.keys, value)
	}
}

// Contains : Check whether current StrSet has value value.
//
func (s *StrSet) Contains(value string) bool {
	_, ok := s.inner[value]
	return ok
}

// Delete : Delete value from StrSet.
//
// O(N) + O(delete(s.inner, value)) + re-slice(s.keys)
//
func (s *StrSet) Delete(value string) {
	var idx = s.inner[value]
	var last_idx = len(s.keys) - 1

	for ii := idx; ii < last_idx; ii += 1 {
		next := ii + 1
		old := s.keys[ii]
		s.inner[old] = next
		s.keys[ii] = s.keys[next]
	}

	delete(s.inner, value)

	s.keys = s.keys[:last_idx]
}

// Values : Get current elements of StrSet.
//
func (s *StrSet) Values() []string {
	return s.keys
}

// Size : Get current element size of StrSet.
//
func (s *StrSet) Size() int {
	return len(s.keys)
}
