package template

import (
	"fmt"
)

func DuplicateFunctionError(name string) error {
	return fmt.Errorf("helper %s already being declared", name)
}

func FunctionTypeError(name string) error {
	return fmt.Errorf("helper %s type are not function", name)
}

func ReturnTypeError(name string) error {
	return fmt.Errorf("helper %s should only return a string or SafeString", name)
}

func FunctionNotExist(name string) error {
	return fmt.Errorf("function %s didn't exists or not being declared", name)
}

func ArgumentOutOfBound(name string, idx int) error {
	return fmt.Errorf("function %s doesn't have argument with index %d", name, idx)
}
