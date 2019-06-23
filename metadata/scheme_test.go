package metadata

import (
	"testing"
)

func TestParseScheme(t *testing.T) {
	var expected = []string{"git", "https"}

	scheme, err := ParseScheme("git+https")

	if err != nil {
		t.Error(err)
	}

	if scheme.Base() != "git" {
		t.Errorf(
			"base scheme for this test should be `%s` but got %s",
			expected[0],
			scheme.Base(),
		)
	}

	if scheme.Extension() != "https" {
		t.Errorf(
			"base scheme for this test should be `%s` but got %s",
			expected[1],
			scheme.Extension(),
		)
	}
}
