package template

import (
	"reflect"
)

// FunctionTables : a table to register a function with a name.
//
//
type FunctionTables struct {
	inner map[string]Function
}

var (
	emptyTables = FunctionTables{}
	zero        reflect.Value
)

// EmptyHelpers : Create empty Helpers
//
func EmptyTables() FunctionTables {
	return emptyTables
}

// MakeHelpers : create helpers statically by using map of functions (reflect.Value)
//
func MakeFunctionTables(rawFunctions map[string]reflect.Value) (FunctionTables, []error) {
	var errors []error = make([]error, 0)
	var tables = FunctionTables{make(map[string]Function)}

	for name, raw := range rawFunctions {

		if _, ok := tables.inner[name]; ok {
			errors = append(errors, DuplicateFunctionError(name))
			continue
		}

		function, err := FuncFromReflect(name, raw)

		if err != nil {
			errors = append(errors, err)
		}

		tables.inner[name] = function

	}

	return tables, errors
}

// Register : register a function to a name
//
// this method will return error if there is already function that being registered
// to that name
//
func (tables FunctionTables) Register(name string, raw interface{}) error {

	if _, ok := tables.inner[name]; ok {
		return DuplicateFunctionError(name)
	}

	function, err := FuncFromReflect(name, reflect.ValueOf(raw))

	if err != nil {
		return err
	}

	tables.inner[name] = function

	return nil
}

func (tables FunctionTables) Get(name string) (Function, bool) {
	value, ok := tables.inner[name]
	return value, ok
}

func (tables FunctionTables) Call(name string, params Params) (reflect.Value, error) {
	var function Function
	var ok bool = false

	if function, ok = tables.Get(name); !ok {
		return zero, FunctionNotExist(name)
	}

	// TODO: @zerosign
	// mean there is an options args in function arguments
	// propagating an options in this new runtime
	// are still questionable (it might not need to separates options & params)
	if function.ParamSize() == params.ParamSize()+1 {

	}

	return zero, nil
}

func CallFunc(name string, params Params) reflect.Value {
	// TODO: @zerosign
	// I need to figure out why raymond.Options also include ast.Node
	// probably for local variables ? but ya need to know that first
	// however all call stacks are being evaluated at evalProgram
	// and most all mutation happens when stack being popped (hidden from user API)
	//
	return zero
}
