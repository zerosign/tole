package metadata

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// ConstSeparatorDot : Separator for version string
	//
	ConstSeparatorDot = "."
)

var (
	emptyVersion = Version{}
)

// Version: Version struct
type Version struct {
	inner []int
}

// newVersion : create version from separate major, minor, patch int
//
func newVersion(major, minor, patch int) Version {
	return Version{[]int{major, minor, patch}}
}

// versionFromIntSlice : create version from int slice
//
// this method doesn't check whether slice size correct or not.
//
func versionFromIntSlice(slices []int) Version {
	return Version{[]int{slices[0], slices[1], slices[2]}}
}

// FullVersion : create (major,minor,patch) Version format
//
func FullVersion(major, minor, patch int) Version {
	return Version{[]int{major, minor, patch}}
}

// PartialVersion : create (major,minor,*) Version format
//
func PartialVersion(major, minor int) Version {
	return Version{[]int{major, minor, 0}}
}

// Major : Get current major version
//
func (v Version) Major() int {
	return v.inner[0]
}

// Minor : Get current minor version
//
func (v Version) Minor() int {
	return v.inner[1]
}

// Patch : Get current patch version
//
func (v Version) Patch() int {
	return v.inner[2]
}

// Equal : check for Version equality
//
func (v Version) Equal(other Version) bool {
	var result = true

	for ii := 0; ii < len(v.inner); ii += 1 {
		result = result && (v.inner[ii] == other.inner[ii])
	}

	return result
}

// Format : format version into major.minor.patch formats
//
func (v Version) Format() string {
	if len(v.inner) == 1 {
		return fmt.Sprintf("%d", v.Major())
	} else if len(v.inner) == 2 {
		return fmt.Sprintf("%d.%d", v.Major(), v.Minor())
	} else if len(v.inner) == 3 {
		return fmt.Sprintf("%d.%d.%d", v.Major(), v.Minor(), v.Patch())
	} else {
		return "0.0.0"
	}
}

// ParseVersion : Parse version string into Version
//
// supported string format (%d.%d) or (%d.%d.%d)
//
func ParseVersion(v string) (Version, error) {
	var err error

	data := make([]int, 3)

	for idx, v := range strings.Split(v, ConstSeparatorDot) {
		if idx == 3 {
			break
		}

		value, err := strconv.Atoi(v)
		if err != nil {
			return emptyVersion, err
		}
		data[idx] = value
	}

	return Version{data}, err
}
