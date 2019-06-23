package base

import (
	"testing"
)

func TestMakeStrSetByEmptyStringSlice(t *testing.T) {
	sets := MakeStrSet([]string{})
	if sets.Size() != 0 {
		t.Errorf("Can't create StrSet with empty string slice")
	}
}

func TestMakeStrSetByStringSlice(t *testing.T) {
	sets := MakeStrSet([]string{"", "", ""})

	if sets.Size() != 1 {
		t.Errorf("Can't create StrSet with string slice, size should be 1")
	}
}

func TestMakeStrSetByUniqueStringSlice(t *testing.T) {
	sets := MakeStrSet([]string{"", "1", "2"})

	if sets.Size() != 3 {
		t.Errorf("Can't create StrSet with unique string slice")
	}
}

func TestStrSetAddString(t *testing.T) {
	sets := EmptyStrSet()
	sets.Add("1")
	sets.Add("1")
	sets.Add("2")

	if sets.Size() != 2 {
		t.Errorf("StrSet add string, size should be 2")
	}
}

func TestStrSetDeleteString(t *testing.T) {
	sets := MakeStrSet([]string{"test", "test", "test2", "test3", "test"})

	sets.Delete("test2")

	if sets.Size() != 2 {
		t.Errorf("StrSet delete string, size should be 2")
	}

	results := sets.Values()
	expected := []string{"test", "test3"}

	// Somehow can't use reflect.DeepEqual
	// for comparing expected & results in here
	for ii := 0; ii < len(expected); ii += 1 {
		if results[ii] != expected[ii] {
			t.Errorf("StrSet index %d should be %s but %s", ii, expected[ii], results[ii])
		}
	}

}
