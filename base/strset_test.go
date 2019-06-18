package base

import (
	"reflect"
	"testing"
)

func TestStrSet_MakeStrSet(t *testing.T) {
	expected := StrSet{
		map[string]void{"test": void{}},
		[]string{"test"},
	}

	result := MakeStrSet([]string{"test"})

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("strset not equals\nexpected: %#v\nreality: %#v\n", expected, result)
	}
}
