package metadata

import (
	"testing"

	"gotest.tools/assert"
)

func TestSimpleVersionParsing(t *testing.T) {
	var version Version
	var expected = Version{[]int{1, 1, 0}}

	if version, err := ParseVersion("1.1.0"); err != nil {
		t.Errorf("error occured when parsing, expected no error")
	} else {
		assert.DeepEqual(t, version, expected)
	}
}
