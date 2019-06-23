package base

import (
	"fmt"
)

func ArrayConvError(v Value) error {
	return fmt.Errorf("can't convert to array from %#v", v)
}

func MapConvError(v Value) error {
	return fmt.Errorf("can't convert to map from %#v", v)
}

func UnknownTypeError(v Value) error {
	return fmt.Errorf("unknown type for %#v", v)
}

func PrimitiveQueryError(v Value) error {
	return fmt.Errorf("can't query primitive value %#v", v)
}
