package metadata

import (
	"errors"
	"strings"
)

// Scheme : a extensible scheme representation that allows extension
//
type Scheme []string

// Base : Get base scheme
//
// example: (git+https) means base will be git
//
func (s Scheme) Base() string {
	return s[0]
}

// Extension : Get extension scheme
//
// example: (git+https) means extension will be https
//
func (s Scheme) Extension() string {
	return s[1]
}

// ParseScheme : parse scheme for given string scheme
//
// returns actual scheme if size is between 0 to 2
//         empty scheme & error if size is more than 2 or 0
//
func ParseScheme(scheme string) (Scheme, error) {
	split := strings.Split(scheme, "+")
	size := len(split)

	if size == 0 || size > 2 {
		return Scheme{}, errors.New("")
	} else {
		return split, nil
	}
}
