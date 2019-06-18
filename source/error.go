package source

import (
	"errors"
	"fmt"
)

func UnsupportedLookupPath(path string) error {
	return errors.New(fmt.Sprintf("path %s is unsupported path", path))
}

func LookupPathIncomplete(path string) error {
	return errors.New(fmt.Sprintf("path %s is incomplete path", path))
}
