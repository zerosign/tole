package template

import (
	"reflect"
)

// Function : internal function being used for tracking function
//
// TODO: enable return size of 2 (since having error in function also nice (mostly doing I/O))
//
type Function struct {
	name       string        // name of the function
	inner      reflect.Value // original pointer that holds instance of the function
	kind       reflect.Type  // pointer that holds type of the function
	paramSize  int           // param size of the function
	returnSize int           // return size of the function
}

var (
	emptyFunction = Function{}
)

func EmptyFunc() Function {
	return emptyFunction
}

func FuncFromReflect(name string, raw reflect.Value) (Function, error) {

	if raw.Kind() != reflect.Func {
		return EmptyFunc(), FunctionTypeError(name)
	}

	kind := raw.Type()

	// single return only
	// TODO: we could provide double return for function in here
	if kind.NumOut() != 1 {
		return EmptyFunc(), ReturnTypeError(name)
	}

	return Function{
		name,
		raw,
		kind,
		kind.NumIn(),
		kind.NumOut(),
	}, nil
}

func (f Function) Name() string {
	return f.name
}

func (f Function) ParamSize() int {
	return f.paramSize
}

func (f Function) InputType(idx int) (reflect.Type, error) {
	if idx < f.ParamSize() {
		return f.kind.In(idx), nil
	} else {
		return nil, ArgumentOutOfBound(f.name, idx)
	}
}

// HasOptions : check whether last argument of the function
//              are option types or not
//
func (f Function) HasOptions() bool {
	// TODO: @zerosign
	// var argType = f.InputType(f.ParamSize() - 1)
	// return OptionType.AssignableTo(argType)

	return false
}
