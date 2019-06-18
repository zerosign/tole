package manifest

import (
	"errors"
	"fmt"
)

func UnsupportedManifest(kind string) error {
	return errors.New(fmt.Sprintf("unsupported manifest type: %s", kind))
}

func FieldNotExists(root, field string) error {
	return errors.New(fmt.Sprintf("required field %s not exists for %s", field, root))
}

func DuplicatedField(field string, oldValue, newValue interface{}) error {
	return errors.New(
		fmt.Sprintf(
			"duplicate value detected in field %s, old: %s, new: %s",
			field, oldValue, newValue,
		),
	)
}
