package metadata

import (
	"testing"
)

func TestParseFullVersion(t *testing.T) {
	version, err := ParseVersion("0.0.1")

	if err != nil {
		t.Error(err)
	}

	expected := []int{0, 0, 1}
	expectedVersion := versionFromIntSlice(expected)

	for ii := 0; ii < len(version.inner); ii += 1 {
		if version.inner[ii] != expected[ii] {
			t.Errorf(
				"expected value for idx %d should be %d but got %d",
				ii,
				expected[ii],
				version.inner[ii],
			)
		}
	}

	if !version.Equal(expectedVersion) {
		t.Errorf(
			"expected version should be %v but got %v",
			expectedVersion,
			version,
		)
	}
}

func TestParsePartialVersion(t *testing.T) {
	version, err := ParseVersion("0.1")

	if err != nil {
		t.Error(err)
	}

	expected := []int{0, 1, 0}
	expectedVersion := versionFromIntSlice(expected)

	for ii := 0; ii < len(version.inner); ii += 1 {
		if version.inner[ii] != expected[ii] {
			t.Errorf(
				"expected value for idx %d should be %d but got %d",
				ii,
				expected[ii],
				version.inner[ii],
			)
		}
	}

	if !version.Equal(expectedVersion) {
		t.Errorf(
			"expected version should be %v but got %v",
			expectedVersion,
			version,
		)
	}

}
