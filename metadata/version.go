package metadata

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	SEP_DOT = "."
)

var (
	emptyVersion = Version{}
)

type Version struct {
	inner []int
}

func FullVersion(major, minor, patch int) Version {
	return Version{[]int{major, minor, patch}}
}

func PartialVersion(major, minor int) Version {
	return Version{[]int{major, minor, 0}}
}

func (v Version) Major() int {
	return v.inner[0]
}

func (v Version) Minor() int {
	return v.inner[1]
}

func (v Version) Patch() int {
	return v.inner[2]
}

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

// 1.0.1
func ParseVersion(v string) (Version, error) {
	var err error
	data := make([]int, 0)
	for _, v := range strings.Split(v, SEP_DOT) {
		value, err := strconv.Atoi(v)
		if err != nil {
			return emptyVersion, err
		}
		data = append(data, value)
	}

	return Version{data}, err
}
